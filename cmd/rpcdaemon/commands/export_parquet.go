package commands

import (
	"fmt"
	"github.com/apache/arrow/go/v8/arrow"
	"github.com/apache/arrow/go/v8/arrow/array"
	"github.com/apache/arrow/go/v8/arrow/memory"
	"github.com/apache/arrow/go/v8/parquet"
	"github.com/apache/arrow/go/v8/parquet/compress"
	"github.com/apache/arrow/go/v8/parquet/pqarrow"
	"github.com/apache/arrow/go/v8/parquet/schema"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/pb/go/protobuf"
	"os"
)

type PlainStackFrame struct {
	FrameId         string
	Type            string
	Label           string
	From            string
	To              string
	ContractCreated string
	Value           string
	Input           string
	Error           string
}

type ExportTraceParquet struct {
	BlockNumber      int64
	TransactionHash  string
	TransactionIndex int32
	FromAddress      string
	ToAddress        string
	GasPrice         int64
	Input            string
	Nonce            int64
	TransactionValue string
	Stack            []PlainStackFrame
}

func dfs(node *protobuf.StackFrame, prefix string, sks []PlainStackFrame) {
	if node == nil {
		return
	}
	sks = append(sks, PlainStackFrame{
		FrameId:         prefix,
		Type:            node.Type,
		Label:           node.Label,
		From:            node.From,
		To:              node.To,
		ContractCreated: node.ContractCreated,
		Value:           node.Value,
		Input:           node.Input,
		Error:           node.Error,
	})
	for i, call := range node.GetCalls() {
		cPrefix := fmt.Sprintf("%s_%d", prefix, i)
		dfs(call, cPrefix, sks)
	}
}

func (e *ExportTraceParquet) setFromPb(tx *protobuf.TraceTransaction) {
	if e == nil {
		panic("receiver: ExportTraceParquet is nil")
	}
	if tx == nil {
		return
	}
	e.BlockNumber = tx.BlockNumber
	e.TransactionHash = tx.TransactionHash
	e.TransactionIndex = tx.TransactionIndex
	e.FromAddress = tx.FromAddress
	e.ToAddress = tx.ToAddress
	e.GasPrice = tx.GasPrice
	e.Input = tx.Input
	e.Nonce = tx.Nonce
	e.TransactionValue = tx.TransactionValue
	e.Stack = make([]PlainStackFrame, 0)

	dfs(tx.Stack, "0", e.Stack)
}

func exportParquet(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	psc, err := schema.NewSchemaFromStruct(&ExportTraceParquet{})
	if err != nil {
		return fmt.Errorf("failed to create schema from struct: %v", err)
	}

	sc, err := pqarrow.FromParquet(psc, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create array schema from schema: %v", err)
	}

	wr, err := pqarrow.NewFileWriter(sc, file,
		parquet.NewWriterProperties(parquet.WithCompression(compress.Codecs.Zstd),
			parquet.WithDictionaryDefault(true),
			parquet.WithDataPageSize(100*1024)),
		pqarrow.DefaultWriterProps())
	if err != nil {
		return fmt.Errorf("failed to create file writer: %v", err)
	}
	defer wr.Close()
	var data = make([]ExportTraceParquet, 2)
	return saveParquet(wr, sc, data)

}
func saveParquet(wr *pqarrow.FileWriter, sc *arrow.Schema, data []ExportTraceParquet) error {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	b := array.NewRecordBuilder(mem, sc)
	defer b.Release()
	for _, v := range data {

		//BlockNumber      int64
		b.Field(0).(*array.Int64Builder).Append(v.BlockNumber)
		//TransactionHash  string
		b.Field(1).(*array.StringBuilder).Append(v.TransactionHash)
		//TransactionIndex int32
		b.Field(2).(*array.Int32Builder).Append(v.TransactionIndex)
		//FromAddress      string
		b.Field(3).(*array.StringBuilder).Append(v.FromAddress)
		//ToAddress        string
		b.Field(4).(*array.StringBuilder).Append(v.ToAddress)
		//GasPrice         int64
		b.Field(5).(*array.Int64Builder).Append(v.GasPrice)
		//Input            string
		b.Field(6).(*array.StringBuilder).Append(v.Input)
		//Nonce            int64
		b.Field(7).(*array.Int64Builder).Append(v.Nonce)
		//TransactionValue string
		b.Field(8).(*array.StringBuilder).Append(v.TransactionValue)
		//Stack            []PlainStackFrame
		stacksBuilder := b.Field(9).(*array.ListBuilder)
		stacksBuilder.Append(true)
		for _, stack := range v.Stack {
			stacksBuilder.Append(true)

			stackBuilder := stacksBuilder.ValueBuilder().(*array.StructBuilder)

			//FrameId         string
			stackBuilder.FieldBuilder(0).(*array.StringBuilder).Append(stack.FrameId)
			//Type            string
			stackBuilder.FieldBuilder(1).(*array.StringBuilder).Append(stack.Type)
			//Label           string
			stackBuilder.FieldBuilder(2).(*array.StringBuilder).Append(stack.Label)
			//From            string
			stackBuilder.FieldBuilder(3).(*array.StringBuilder).Append(stack.From)
			//To              string
			stackBuilder.FieldBuilder(4).(*array.StringBuilder).Append(stack.To)
			//ContractCreated string
			stackBuilder.FieldBuilder(5).(*array.StringBuilder).Append(stack.ContractCreated)
			//Value           string
			stackBuilder.FieldBuilder(6).(*array.StringBuilder).Append(stack.Value)
			//Input           string
			stackBuilder.FieldBuilder(7).(*array.StringBuilder).Append(stack.Input)
			//Error           string
			stackBuilder.FieldBuilder(8).(*array.StringBuilder).Append(stack.Error)
		}

	}

	record := b.NewRecord()
	return wr.Write(record)
}
