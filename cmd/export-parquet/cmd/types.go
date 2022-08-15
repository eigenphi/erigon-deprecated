package cmd

import "github.com/ledgerwatch/erigon/rpc"

type JsonrpcParam struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
}

type jsonrpcResult struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  rpc.BlockNumber `json:"result"`
}
