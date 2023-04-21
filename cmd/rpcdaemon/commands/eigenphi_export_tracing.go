package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	jsoniter "github.com/json-iterator/go"
	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon/core"
	"github.com/ledgerwatch/erigon/core/rawdb"
	"github.com/ledgerwatch/erigon/core/state"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/core/vm/evmtypes"
	"github.com/ledgerwatch/erigon/eigenphi/opstrace"
	"github.com/ledgerwatch/erigon/eigenphi/pb/go/protobuf"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"github.com/ledgerwatch/erigon/rpc"
	"github.com/ledgerwatch/erigon/turbo/transactions"
	"go.uber.org/zap"
	"io"
	"math/big"
	"os"
	"strings"
)

var _ PrivateDebugAPI = (*PrivateDebugAPIImpl)(nil)

// modified from api.TraceTransaction
func (api *PrivateDebugAPIImpl) traceTx(ctx context.Context, hash common.Hash) (traceResult *trace.OpsCallFrame,
	blockNumber uint64, txIndex uint64, err error) {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return
	}
	defer tx.Rollback()
	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		return
	}
	// Retrieve the transaction and assemble its EVM context
	blockNumber, ok, err := api.txnLookup(ctx, tx, hash)
	if err != nil {
		return
	}
	if !ok {
		err = fmt.Errorf("txn not found")
		return
	}
	// check pruning to ensure we have history at this block level
	err = api.BaseAPI.checkPruneHistory(tx, blockNumber)
	if err != nil {
		return
	}
	// Private API returns 0 if transaction is not found.
	if blockNumber == 0 && chainConfig.Bor != nil {
		var blockNumPtr *uint64
		blockNumPtr, err = rawdb.ReadBorTxLookupEntry(tx, hash)
		if err != nil {
			return
		}
		if blockNumPtr == nil {
			err = fmt.Errorf("block not found")
			return
		}
		blockNumber = *blockNumPtr
	}
	block, err := api.blockByNumberWithSenders(tx, blockNumber)
	if err != nil {
		return
	}
	if block == nil {
		err = fmt.Errorf("block not found")
		return
	}
	var txn types.Transaction
	for i, transaction := range block.Transactions() {
		if transaction.Hash() == hash {
			txIndex = uint64(i)
			txn = transaction
			break
		}
	}
	if txn == nil {
		var borTx types.Transaction
		borTx, _, _, _, err = rawdb.ReadBorTransaction(tx, hash)
		if err != nil {
			return
		}

		if borTx != nil {
			err = fmt.Errorf("borTx is nil")
			return
		}
		err = fmt.Errorf("transaction %#x not found", hash)
		return
	}
	engine := api.engine()

	msg, blockCtx, txCtx, ibs, _, err := transactions.ComputeTxEnv(ctx, engine, block,
		chainConfig, api._blockReader, tx, int(txIndex), api.historyV3(tx))
	if err != nil {
		return
	}
	config := &tracers.TraceConfig{}
	traceResult, err = trace.TraceTxByOpsTracer(ctx, msg, blockCtx, txCtx, ibs, config, chainConfig)
	return traceResult, blockNumber, txIndex, err
}

func (api *PrivateDebugAPIImpl) EigenphiTraceByTxHash(ctx context.Context, hash common.Hash, stream *jsoniter.Stream) error {
	tracerResult, _, _, err := api.traceTx(ctx, hash)
	if err != nil {
		stream.WriteNil()
		return err
	}
	out := jsonpb.Marshaler{EmitDefaults: true}
	pbTrace := toPbCallTrace(tracerResult)
	if err := out.Marshal(stream, pbTrace); err != nil {
		return fmt.Errorf("proto marshal %s", err)
	}
	return nil
}

func (api *PrivateDebugAPIImpl) getActualTxMessage(ctx context.Context, tx kv.Tx,
	hash common.Hash) (core.Message, error) {

	blockNum, ok, err := api.txnLookup(ctx, tx, hash)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("txn not found")
	}
	block, err := api.blockByNumberWithSenders(tx, blockNum)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block not found")
	}
	var txnIndex uint64
	var txn types.Transaction
	for i, transaction := range block.Transactions() {
		if transaction.Hash() == hash {
			txnIndex = uint64(i)
			txn = transaction
			break
		}
	}
	if txn == nil {
		var borTx types.Transaction
		borTx, _, _, _, err = rawdb.ReadBorTransaction(tx, hash)
		if err != nil {
			return nil, err
		}

		if borTx != nil {
			return nil, fmt.Errorf("borTx is nil")
		}
		return nil, fmt.Errorf("transaction %#x not found", hash)
	}

	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		return nil, err
	}
	engine := api.engine()

	msg, _, _, _, _, err := transactions.ComputeTxEnv(ctx, engine, block, chainConfig,
		api._blockReader, tx, int(txnIndex), api.historyV3(tx))
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (api *PrivateDebugAPIImpl) SimulateTxAtIndex(ctx context.Context, hash common.Hash,
	blockNumber rpc.BlockNumber, txIndex uint64) (*trace.OpsCallFrame, error) {

	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get actual tx info
	msg, err := api.getActualTxMessage(ctx, tx, hash)
	if err != nil {
		return nil, err
	}

	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		return nil, err
	}

	block, err := api.blockByNumberWithSenders(tx, uint64(blockNumber.Int64()))
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("got empty block")
	}
	engine := api.engine()
	_, blockCtx, txCtx, ibs, _, err := transactions.ComputeTxEnv(
		ctx, engine, block, chainConfig, api._blockReader, tx,
		int(txIndex), api.historyV3(tx))
	if err != nil {
		return nil, err
	}
	// Trace the transaction and return
	config := &tracers.TraceConfig{}
	return trace.TraceTxByOpsTracer(ctx, msg, blockCtx, txCtx, ibs, config, chainConfig)
}

func (api *PrivateDebugAPIImpl) EigenphiSimulateTxTraceByHash(ctx context.Context, hash common.Hash,
	blockNumber rpc.BlockNumber, txIndex uint64, stream *jsoniter.Stream) error {

	tracerResult, err := api.SimulateTxAtIndex(ctx, hash, blockNumber, txIndex)
	if err != nil {
		stream.WriteNil()
		return err
	}
	pbTrace := toPbCallTrace(tracerResult)
	ptxs := make([]PlainStackFrame, 0)
	dfs(pbTrace, "0", &ptxs)
	return json.NewEncoder(stream).Encode(PlainTraceByTx{
		BlockNumber:      uint64(blockNumber),
		TransactionHash:  hash.String(),
		TransactionIndex: txIndex,
		PlainTraces:      ptxs,
	})
}

func (api *PrivateDebugAPIImpl) EigenphiTraceByNumber(ctx context.Context, height rpc.BlockNumber, stream *jsoniter.Stream) error {
	defer stream.Flush()

	config := &tracers.TraceConfig{}
	if _, err := stream.Write([]byte("[")); err != nil {
		return fmt.Errorf("write array begin err: %s", err)
	}
	if err := api.TraceSingleBlock(ctx, height, config, stream, []byte(",")); err != nil {
		stream.WriteNil()
		return err
	}
	stream.Write([]byte("]"))
	return nil
}

func (api *PrivateDebugAPIImpl) EigenphiPlainTraceByTxHash(ctx context.Context, hash common.Hash, stream *jsoniter.Stream) error {
	tracerResult, blockNumber, txIndex, err := api.traceTx(ctx, hash)
	if err != nil {
		stream.WriteNil()
		return err
	}
	pbTrace := toPbCallTrace(tracerResult)
	ptxs := make([]PlainStackFrame, 0)
	dfs(pbTrace, "0", &ptxs)
	return json.NewEncoder(stream).Encode(PlainTraceByTx{
		BlockNumber:      blockNumber,
		TransactionHash:  hash.String(),
		TransactionIndex: txIndex,
		PlainTraces:      ptxs,
	})
}

type PlainTraceByTx struct {
	BlockNumber      uint64            `json:"blockNumber"`
	TransactionHash  string            `json:"transactionHash"`
	TransactionIndex uint64            `json:"transactionIndex"`
	PlainTraces      []PlainStackFrame `json:"plainTraces"`
}

func (api *PrivateDebugAPIImpl) EigenphiPlainTraceByNumber(ctx context.Context, height rpc.BlockNumber, stream *jsoniter.Stream) error {
	defer stream.Flush()

	tracerName := "opsTracer"
	traces, err := api.TraceSingleBlockRaw(ctx, height, &tracers.TraceConfig{Tracer: &tracerName})
	if err != nil {
		stream.WriteNil()
		return err
	}

	rets := make([]PlainTraceByTx, 0)

	for i := range traces {
		ptxs := make([]PlainStackFrame, 0)
		dfs(traces[i].Stack, "0", &ptxs)
		rets = append(rets, PlainTraceByTx{
			BlockNumber:      uint64(height.Int64()),
			TransactionHash:  traces[i].TransactionHash,
			TransactionIndex: uint64(traces[i].TransactionIndex),
			PlainTraces:      ptxs,
		})
	}

	return json.NewEncoder(stream).Encode(rets)
}

func (api *PrivateDebugAPIImpl) GetSimpleBlock(ctx context.Context, blockNr rpc.BlockNumber, output *os.File) error {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var block *types.Block
	block, err = api.blockByRPCNumber(blockNr, tx)
	if err != nil {
		return err
	}
	simpleBlock := toPbSimpleBlock(block)

	out := jsonpb.Marshaler{EmitDefaults: true}
	if err := out.Marshal(output, &simpleBlock); err != nil {
		return fmt.Errorf("failed to marshal block: %s", err)
	}

	if _, err := output.WriteString("\n"); err != nil {
		return fmt.Errorf("failed to write newline: %s", err)
	}
	return nil
}

func (api *PrivateDebugAPIImpl) GetSimpleTx(ctx context.Context, blockNr rpc.BlockNumber, output *os.File) error {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var block *types.Block
	block, err = api.blockByRPCNumber(blockNr, tx)
	if err != nil {
		return err
	}

	out := jsonpb.Marshaler{EmitDefaults: true}
	for idx, tx := range block.Transactions() {
		select {
		default:
		case <-ctx.Done():
			return ctx.Err()
		}

		rtx := newRPCTransaction(tx, block.Hash(), block.NumberU64(), uint64(idx), big.NewInt(0))
		pbTx := toPbSimpleTransaction(rtx, tx, block.Time())
		if err := out.Marshal(output, &pbTx); err != nil {
			return fmt.Errorf("failed to encode transaction: %s", err)
		}

		if _, err := output.WriteString("\n"); err != nil {
			return fmt.Errorf("failed to write newline: %s", err)
		}
	}
	return nil
}

func (api *PrivateDebugAPIImpl) TraceSingleBlockRaw(ctx context.Context, blockNr rpc.BlockNumber, config *tracers.TraceConfig) ([]protobuf.TraceTransaction, error) {

	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var block *types.Block
	block, err = api.blockByRPCNumber(blockNr, tx)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block %d not found", blockNr)
	}

	// if we've pruned this history away for this block then just return early
	// to save any red herring errors
	err = api.BaseAPI.checkPruneHistory(tx, block.NumberU64())
	if err != nil {
		return nil, err
	}

	if config.BorTraceEnabled == nil {
		config.BorTraceEnabled = newBoolPtr(false)
	}

	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		return nil, err
	}
	engine := api.engine()

	_, blockCtx, _, ibs, _, err := transactions.ComputeTxEnv(ctx, engine, block,
		chainConfig, api._blockReader, tx, 0, api.historyV3(tx))
	if err != nil {
		return nil, err
	}

	signer := types.MakeSigner(chainConfig, block.NumberU64())
	rules := chainConfig.Rules(block.NumberU64(), block.Time())
	rets := make([]protobuf.TraceTransaction, len(block.Transactions()))

	for idx, txn := range block.Transactions() {
		select {
		default:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
		ibs.Prepare(txn.Hash(), block.Hash(), idx)

		msg, _ := txn.AsMessage(*signer, block.BaseFee(), rules)
		if msg.FeeCap().IsZero() && engine != nil {
			syscall := func(contract common.Address, data []byte) ([]byte, error) {
				return core.SysCallContract(contract, data, *chainConfig, ibs, block.Header(), engine, true /* constCall */)
			}
			msg.SetIsFree(engine.IsServiceTransaction(msg.From(), syscall))
		}

		txCtx := evmtypes.TxContext{
			TxHash:   txn.Hash(),
			Origin:   msg.From(),
			GasPrice: msg.GasPrice(),
		}

		tracerResult, err := trace.TraceTxByOpsTracer(ctx, msg, blockCtx, txCtx, ibs, config, chainConfig)
		if err == nil {
			err = ibs.FinalizeTx(rules, state.NewNoopWriter())
		}
		if err != nil {
			// TODO handle trace transaction error
			zap.L().Sugar().Errorf("TraceTxByOpsTracer error: %s %s", txn.Hash(), err)
		}

		var baseFee *big.Int
		if chainConfig.IsLondon(block.Number().Uint64()) && block.Hash() != (common.Hash{}) {
			baseFee = block.BaseFee()
		}
		rtx := newRPCTransaction(txn, block.Hash(), block.NumberU64(), uint64(idx), baseFee)
		pbTx := toPbTraceTransaction(rtx, txn, tracerResult, int64(block.Time()))
		rets[idx] = pbTx
	}
	return rets, nil
}

func (api *PrivateDebugAPIImpl) TraceSingleBlock(ctx context.Context, blockNr rpc.BlockNumber, config *tracers.TraceConfig, output io.Writer, separator []byte) error {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var block *types.Block
	block, err = api.blockByRPCNumber(blockNr, tx)
	if err != nil {
		return err
	}
	if block == nil {
		return fmt.Errorf("block %d not found", blockNr)
	}

	// if we've pruned this history away for this block then just return early
	// to save any red herring errors
	err = api.BaseAPI.checkPruneHistory(tx, block.NumberU64())
	if err != nil {
		return err
	}

	if config.BorTraceEnabled == nil {
		config.BorTraceEnabled = newBoolPtr(false)
	}

	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		return err
	}
	engine := api.engine()

	_, blockCtx, _, ibs, _, err := transactions.ComputeTxEnv(ctx, engine, block,
		chainConfig, api._blockReader, tx, 0, api.historyV3(tx))
	if err != nil {
		return err
	}

	signer := types.MakeSigner(chainConfig, block.NumberU64())
	rules := chainConfig.Rules(block.NumberU64(), block.Time())

	out := jsonpb.Marshaler{EmitDefaults: true}
	for idx, txn := range block.Transactions() {
		select {
		default:
		case <-ctx.Done():
			return ctx.Err()
		}
		ibs.Prepare(txn.Hash(), block.Hash(), idx)

		msg, _ := txn.AsMessage(*signer, block.BaseFee(), rules)
		if msg.FeeCap().IsZero() && engine != nil {
			syscall := func(contract common.Address, data []byte) ([]byte, error) {
				return core.SysCallContract(contract, data, *chainConfig, ibs, block.Header(), engine, true /* constCall */)
			}
			msg.SetIsFree(engine.IsServiceTransaction(msg.From(), syscall))
		}

		txCtx := evmtypes.TxContext{
			TxHash:   txn.Hash(),
			Origin:   msg.From(),
			GasPrice: msg.GasPrice(),
		}

		tracerResult, err := trace.TraceTxByOpsTracer(ctx, msg, blockCtx, txCtx, ibs, config, chainConfig)
		if err == nil {
			err = ibs.FinalizeTx(rules, state.NewNoopWriter())
		}
		if err != nil {
			// TODO handle trace transaction error
			zap.L().Sugar().Errorf("TraceTxByOpsTracer error: %s %s", txn.Hash(), err)
		}

		var baseFee *big.Int
		if chainConfig.IsLondon(block.Number().Uint64()) && block.Hash() != (common.Hash{}) {
			baseFee = block.BaseFee()
		}
		rtx := newRPCTransaction(txn, block.Hash(), block.NumberU64(), uint64(idx), baseFee)
		pbTx := toPbTraceTransaction(rtx, txn, tracerResult, int64(block.Time()))
		if err := out.Marshal(output, &pbTx); err != nil {
			return fmt.Errorf("failed to encode transaction: %s", err)
		}
		if idx+1 != len(block.Transactions()) {
			if _, err := output.Write(separator); err != nil {
				return fmt.Errorf("failed to write newline: %s", err)
			}
		}
	}
	return nil
}

func toPbSimpleBlock(block *types.Block) protobuf.Block {
	return protobuf.Block{
		BlockNumber:    int64(block.NumberU64()),
		BlockHash:      block.Hash().String(),
		ParentHash:     block.ParentHash().String(),
		Miner:          block.Header().Coinbase.String(),
		BlockSize:      int32(block.Size()),
		GasLimit:       int64(block.GasLimit()),
		GasUsed:        int64(block.GasUsed()),
		BlockTimestamp: int64(block.Time()),
	}
}
func toPbSimpleTransaction(rtx *RPCTransaction, tx types.Transaction, blkTs uint64) protobuf.Transaction {
	var to string
	if rtx.To != nil {
		to = strings.ToLower(rtx.To.String())
	}
	return protobuf.Transaction{
		TransactionHash:  rtx.Hash.String(),
		TransactionIndex: int32(*rtx.TransactionIndex),
		BlockNumber:      rtx.BlockNumber.ToInt().Int64(),
		FromAddress:      strings.ToLower(rtx.From.String()),
		ToAddress:        to,
		Nonce:            int64(tx.GetNonce()),
		TransactionValue: tx.GetValue().String(),
		BlockTimestamp:   int64(blkTs),
	}
}

func toPbTraceTransaction(rtx *RPCTransaction, tx types.Transaction, tc *trace.OpsCallFrame, blockTimestamp int64) protobuf.TraceTransaction {
	var to string
	if rtx.To != nil {
		to = strings.ToLower(rtx.To.String())
	}
	return protobuf.TraceTransaction{
		BlockNumber:      rtx.BlockNumber.ToInt().Int64(),
		TransactionHash:  rtx.Hash.String(),
		TransactionIndex: int32(*rtx.TransactionIndex),
		FromAddress:      strings.ToLower(rtx.From.String()),
		ToAddress:        to,
		GasPrice:         rtx.GasPrice.ToInt().Int64(),
		Input:            hexutility.Encode(tx.GetData()),
		Nonce:            int64(tx.GetNonce()),
		TransactionValue: tx.GetValue().String(),
		Stack:            toPbCallTrace(tc),
		BlockTimestamp:   blockTimestamp,
	}
}

func toPbCallTrace(in *trace.OpsCallFrame) *protobuf.StackFrame {
	if in == nil {
		return &protobuf.StackFrame{}
	}

	var calls []*protobuf.StackFrame
	if in.Calls != nil {
		calls = make([]*protobuf.StackFrame, len(in.Calls))
		for i, c := range in.Calls {
			calls[i] = toPbCallTrace(c)
		}
	}
	return &protobuf.StackFrame{
		Type:            in.Type,
		Label:           in.Label,
		From:            in.From,
		To:              in.To,
		ContractCreated: in.ContractCreated,
		Value:           in.Value,
		Input:           in.Input,
		Error:           in.Error,
		Calls:           calls,
		FourBytes:       in.FourBytes,
	}

}
