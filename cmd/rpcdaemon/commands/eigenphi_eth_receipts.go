package commands

import (
	"context"
	"fmt"
	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/common/hexutil"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/rpc"
)

// GetTransactionReceiptsByBlockNumber implements eth_getTransactionReceiptsByBlockNumber. Returns the receipts of transactions given the block number.
func (api *APIImpl) GetTransactionReceiptsByBlockNumber(ctx context.Context, blockNumber rpc.BlockNumber) ([]map[string]interface{}, error) {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Extract transactions from block
	block, bErr := api.blockByRPCNumber(blockNumber, tx)
	if bErr != nil {
		return nil, bErr
	}
	if block == nil {
		return nil, fmt.Errorf("could not find block  %d", blockNumber.Int64())
	}

	cc, err := api.chainConfig(tx)
	if err != nil {
		return nil, err
	}
	txs := block.Transactions()

	receipts, err := api.getReceipts(ctx, tx, cc, block, block.Body().SendersFromTxs())
	if err != nil {
		return nil, fmt.Errorf("getReceipts error: %w", err)
	}
	txReceipts := make([]map[string]interface{}, 0, len(txs))

	for idx, receipt := range receipts {
		tx := txs[idx]

		from, _ := tx.GetSender()

		fields := map[string]interface{}{
			"blockHash":         block.Hash(),
			"blockNumber":       hexutil.Uint64(blockNumber),
			"transactionHash":   tx.Hash(),
			"transactionIndex":  hexutil.Uint64(idx),
			"from":              from,
			"to":                tx.GetTo(),
			"gasUsed":           hexutil.Uint64(receipt.GasUsed),
			"cumulativeGasUsed": hexutil.Uint64(receipt.CumulativeGasUsed),
			"contractAddress":   nil,
			"logs":              receipt.Logs,
			"logsBloom":         receipt.Bloom,
		}

		// Assign receipt status or post state.
		if len(receipt.PostState) > 0 {
			fields["root"] = hexutil.Bytes(receipt.PostState)
		} else {
			fields["status"] = hexutil.Uint(receipt.Status)
		}
		if receipt.Logs == nil {
			fields["logs"] = [][]*types.Log{}
		}
		// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
		if receipt.ContractAddress != (common.Address{}) {
			fields["contractAddress"] = receipt.ContractAddress
		}

		txReceipts = append(txReceipts, fields)
	}

	return txReceipts, nil
}
