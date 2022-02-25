package commands

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	jsoniter "github.com/json-iterator/go"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/pb/go/protobuf"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/common/hexutil"
	"github.com/ledgerwatch/erigon/consensus/ethash"
	"github.com/ledgerwatch/erigon/core/rawdb"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"github.com/ledgerwatch/erigon/eth/tracers/native"
	"github.com/ledgerwatch/erigon/ethdb"
	"github.com/ledgerwatch/erigon/rpc"
	"github.com/ledgerwatch/erigon/turbo/transactions"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"math/big"
	"os"
	"strings"
)

//// TraceBlockByStep implements debug_traceBlockByStep. Returns Geth style blocks traces with startNumber and step size.
//func (api *PrivateDebugAPIImpl) TraceBlockByStep(ctx context.Context, startBlock, step rpc.BlockNumber, config *tracers.TraceConfig, stream *jsoniter.Stream) error {
//	// TODO check endBlock less than latest height
//	if step <= 0 {
//		return fmt.Errorf("step must be positive")
//	}
//
//	for block := startBlock; block <= startBlock+step; block++ {
//		err := api.TraceSingleBlock(ctx, block, config, stream)
//		if err != nil {
//			stream.Write(nil)
//			return err
//		}
//	}
//	return nil
//}
//
// TraceBlockByRange implements debug_traceBlockByRange. Returns Geth style blocks traces with startNumber and endNumber.
func (api *PrivateDebugAPIImpl) TraceBlockByRange(ctx context.Context, startBlockNumber, endBlockNumber rpc.BlockNumber, config *tracers.TraceConfig, stream *jsoniter.Stream) error {
	if startBlockNumber >= endBlockNumber {
		return fmt.Errorf("startBlock >= endBlock")
	}

	// TODO check endBlock less or equal than latest height

	//for block := startBlockNumber; block <= endBlockNumber; block++ {
	//	err := api.TraceSingleBlock(ctx, block, config, stream)
	//	if err != nil {
	//		stream.Write(nil)
	//		return err
	//	}
	//}
	return nil
}

type TxOpsTracer struct {
	Hash  common.Hash
	trace native.OpsCallFrame
}

func (api *PrivateDebugAPIImpl) getSimpleBlock(ctx context.Context, blockNr rpc.BlockNumber, output *os.File) error {
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

func (api *PrivateDebugAPIImpl) getSimpleTx(ctx context.Context, blockNr rpc.BlockNumber, output *os.File) error {
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

func (api *PrivateDebugAPIImpl) TraceSingleBlock(ctx context.Context, blockNr rpc.BlockNumber, config *tracers.TraceConfig, output *os.File, outProto bool) error {
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

	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		return err
	}

	contractHasTEVM := func(contractHash common.Hash) (bool, error) { return false, nil }
	if api.TevmEnabled {
		contractHasTEVM = ethdb.GetHasTEVM(tx)
	}

	getHeader := func(hash common.Hash, number uint64) *types.Header {
		return rawdb.ReadHeader(tx, hash, number)
	}

	_, blockCtx, txCtx, ibs, reader, err := transactions.ComputeTxEnv(ctx, block, chainConfig, getHeader, contractHasTEVM, ethash.NewFaker(), tx, block.Hash(), 0)
	if err != nil {
		return err
	}

	signer := types.MakeSigner(chainConfig, block.NumberU64())

	out := jsonpb.Marshaler{EmitDefaults: true}
	for idx, tx := range block.Transactions() {
		select {
		default:
		case <-ctx.Done():
			return ctx.Err()
		}
		ibs.Prepare(tx.Hash(), block.Hash(), idx)
		msg, _ := tx.AsMessage(*signer, block.BaseFee())

		tracerResult, err := transactions.TraceTxByOpsTracer(ctx, msg, blockCtx, txCtx, ibs, config, chainConfig)
		_ = ibs.FinalizeTx(chainConfig.Rules(blockCtx.BlockNumber), reader)
		if err != nil {
			// TODO handle trace transaction error
			zap.L().Sugar().Errorf("TraceTxByOpsTracer error: %s %s", tx.Hash(), err)
		}

		var baseFee *big.Int
		if chainConfig.IsLondon(block.Number().Uint64()) && block.Hash() != (common.Hash{}) {
			baseFee = block.BaseFee()
		}
		rtx := newRPCTransaction(tx, block.Hash(), block.NumberU64(), uint64(idx), baseFee)
		pbTx := toPbTraceTransaction(rtx, tx, tracerResult)
		if outProto {
			marshal, err := proto.Marshal(&pbTx)
			if err != nil {
				return fmt.Errorf("proto marshal %s", err)
			}
			if _, err := output.Write(marshal); err != nil {
				return fmt.Errorf("write protobuf data: %s", err)
			}
		} else {
			if err := out.Marshal(output, &pbTx); err != nil {
				return fmt.Errorf("failed to encode transaction: %s", err)
			}
			if _, err := output.WriteString("\n"); err != nil {
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

func toPbTraceTransaction(rtx *RPCTransaction, tx types.Transaction, tc *native.OpsCallFrame) protobuf.TraceTransaction {
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
		Input:            hexutil.Encode(tx.GetData()),
		Nonce:            int64(tx.GetNonce()),
		TransactionValue: tx.GetValue().String(),
		Stack:            toPbCallTrace(tc),
	}
}

func toPbCallTrace(in *native.OpsCallFrame) *protobuf.StackFrame {
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
	}

}
