package commands

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/consensus/ethash"
	"github.com/ledgerwatch/erigon/core/rawdb"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"github.com/ledgerwatch/erigon/ethdb"
	"github.com/ledgerwatch/erigon/rpc"
	"github.com/ledgerwatch/erigon/turbo/transactions"
)

// TraceBlockByStep implements debug_traceBlockByStep. Returns Geth style blocks traces with startNumber and step size.
func (api *PrivateDebugAPIImpl) TraceBlockByStep(ctx context.Context, startBlock, step rpc.BlockNumber, config *tracers.TraceConfig, stream *jsoniter.Stream) error {
	// TODO check endBlock less than latest height
	if step <= 0 {
		return fmt.Errorf("step must be positive")
	}

	for block := startBlock; block <= startBlock+step; block++ {
		err := api.traceSingleBlock(ctx, block, config, stream)
		if err != nil {
			stream.Write(nil)
			return err
		}
	}
	return nil
}

// TraceBlockByRange implements debug_traceBlockByRange. Returns Geth style blocks traces with startNumber and endNumber.
func (api *PrivateDebugAPIImpl) TraceBlockByRange(ctx context.Context, startBlockNumber, endBlockNumber rpc.BlockNumber, config *tracers.TraceConfig, stream *jsoniter.Stream) error {
	if startBlockNumber >= endBlockNumber {
		return fmt.Errorf("startBlock >= endBlock")
	}

	// TODO check endBlock less or equal than latest height

	for block := startBlockNumber; block <= endBlockNumber; block++ {
		err := api.traceSingleBlock(ctx, block, config, stream)
		if err != nil {
			stream.Write(nil)
			return err
		}
	}
	return nil
}

func (api *PrivateDebugAPIImpl) traceSingleBlock(ctx context.Context, blockNr rpc.BlockNumber, config *tracers.TraceConfig, stream *jsoniter.Stream) error {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		stream.WriteNil()
		return err
	}
	defer tx.Rollback()
	var block *types.Block
	block, err = api.blockByRPCNumber(blockNr, tx)
	if err != nil {
		stream.WriteNil()
		return err
	}

	chainConfig, err := api.chainConfig(tx)
	if err != nil {
		stream.WriteNil()
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
		stream.WriteNil()
		return err
	}

	signer := types.MakeSigner(chainConfig, block.NumberU64())
	stream.WriteArrayStart()
	for idx, tx := range block.Transactions() {
		select {
		default:
		case <-ctx.Done():
			stream.WriteNil()
			return ctx.Err()
		}
		ibs.Prepare(tx.Hash(), block.Hash(), idx)
		msg, _ := tx.AsMessage(*signer, block.BaseFee())

		transactions.TraceTx(ctx, msg, blockCtx, txCtx, ibs, config, chainConfig, stream)
		_ = ibs.FinalizeTx(chainConfig.Rules(blockCtx.BlockNumber), reader)
		if idx != len(block.Transactions())-1 {
			stream.WriteMore()
		}
		stream.Flush()
	}
	stream.WriteArrayEnd()
	stream.Flush()
	return nil
}
