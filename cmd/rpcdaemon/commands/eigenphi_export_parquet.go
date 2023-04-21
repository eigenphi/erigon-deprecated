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
	"github.com/ledgerwatch/erigon/eigenphi/pb/go/protobuf"
	"io"
	"os"
)

type PlainStackFrame struct {
	FrameId         string `parquet:"fieldid=0,logical=String" json:"frameId"`
	Type            string `parquet:"fieldid=1,logical=String" json:"type"`
	Label           string `parquet:"fieldid=2,logical=String" json:"label"`
	From            string `parquet:"fieldid=3,logical=String" json:"from"`
	To              string `parquet:"fieldid=4,logical=String" json:"to"`
	ContractCreated string `parquet:"fieldid=5,logical=String" json:"contractCreated"`
	Value           string `parquet:"fieldid=6,logical=String" json:"value"`
	Input           string `parquet:"fieldid=7,logical=String" json:"input"`
	Error           string `parquet:"fieldid=8,logical=String" json:"error"`
	ChildrenCount   int32  `parquet:"fieldid=9" json:"childrenCount"`
	FourBytes       string `parquet:"fieldid=10,logical=String" json:"fourBytes"`
}

type ExportTraceParquet struct {
	BlockNumber      int64             `parquet:"fieldid=0" json:"blockNumber"`
	TransactionHash  string            `parquet:"fieldid=1,logical=String" json:"transactionHash"`
	TransactionIndex int32             `parquet:"fieldid=2" json:"transactionIndex"`
	FromAddress      string            `parquet:"fieldid=3,logical=String" json:"fromAddress"`
	ToAddress        string            `parquet:"fieldid=4,logical=String" json:"toAddress"`
	GasPrice         int64             `parquet:"fieldid=5" json:"gasPrice"`
	Input            string            `parquet:"fieldid=6,logical=String" json:"input"`
	Nonce            int64             `parquet:"fieldid=7" json:"nonce"`
	TransactionValue string            `parquet:"fieldid=8,logical=String" json:"transactionValue"`
	Stack            []PlainStackFrame `parquet:"fieldid=9" json:"stack"`
	BlockTimestamp   int64             `parquet:"fieldid=10" json:"blockTimestamp"`
}

func dfs(node *protobuf.StackFrame, prefix string, sks *[]PlainStackFrame) {
	if node == nil {
		return
	}
	*sks = append(*sks, PlainStackFrame{
		FrameId:         prefix,
		Type:            node.Type,
		Label:           node.Label,
		From:            node.From,
		To:              node.To,
		ContractCreated: node.ContractCreated,
		Value:           node.Value,
		Input:           node.Input,
		Error:           node.Error,
		ChildrenCount:   int32(len(node.GetCalls())),
		FourBytes:       node.FourBytes,
	})
	for i, call := range node.GetCalls() {
		cPrefix := fmt.Sprintf("%s_%d", prefix, i)
		dfs(call, cPrefix, sks)
	}
}

func (e *ExportTraceParquet) SetFromPb(tx *protobuf.TraceTransaction) {
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
	e.BlockTimestamp = tx.BlockTimestamp
	dfs(tx.Stack, "0", &e.Stack)
}

func ExportParquet(filename string, traces []protobuf.TraceTransaction) error {
	tmpfile := filename + ".tmp"
	tmpf, err := os.OpenFile(tmpfile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer tmpf.Close()

	var data = make([]ExportTraceParquet, len(traces))
	for i := range traces {
		data[i].SetFromPb(&traces[i])
	}

	if err := ExportParquetWithData(tmpf, data); err != nil {
		return fmt.Errorf("export parquet file to %s: %w", tmpfile, err)
	}
	return os.Rename(tmpfile, filename)
}

func ExportParquetWithData(writer io.Writer, data []ExportTraceParquet) error {

	psc, err := schema.NewSchemaFromStruct(&ExportTraceParquet{})
	if err != nil {
		return fmt.Errorf("failed to create schema from struct: %v", err)
	}

	sc, err := pqarrow.FromParquet(psc, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create array schema from schema: %v", err)
	}

	wr, err := pqarrow.NewFileWriter(sc, writer, parquet.NewWriterProperties(parquet.WithCompression(compress.Codecs.Zstd),
		parquet.WithDictionaryDefault(true),
		parquet.WithDataPageSize(100*1024),
	),
		pqarrow.DefaultWriterProps(),
	)
	if err != nil {
		return fmt.Errorf("failed to create file writer: %v", err)
	}
	defer wr.Close()

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
		lb := b.Field(9).(*array.ListBuilder)
		lb.Append(true)
		for _, stack := range v.Stack {
			lvb := lb.ValueBuilder().(*array.StructBuilder)
			lvb.Append(true)

			//FrameId         string
			lvb.FieldBuilder(0).(*array.StringBuilder).Append(stack.FrameId)
			//Type            string
			lvb.FieldBuilder(1).(*array.StringBuilder).Append(stack.Type)
			//Label           string
			lvb.FieldBuilder(2).(*array.StringBuilder).Append(stack.Label)
			//From            string
			lvb.FieldBuilder(3).(*array.StringBuilder).Append(stack.From)
			//To              string
			lvb.FieldBuilder(4).(*array.StringBuilder).Append(stack.To)
			//ContractCreated string
			lvb.FieldBuilder(5).(*array.StringBuilder).Append(stack.ContractCreated)
			//Value           string
			lvb.FieldBuilder(6).(*array.StringBuilder).Append(stack.Value)
			//Input           string
			lvb.FieldBuilder(7).(*array.StringBuilder).Append(stack.Input)
			//Error           string
			lvb.FieldBuilder(8).(*array.StringBuilder).Append(stack.Error)
			//ChildrenCount   int32  `parquet:"fieldid=9"`
			lvb.FieldBuilder(9).(*array.Int32Builder).Append(stack.ChildrenCount)
			//FourBytes
			lvb.FieldBuilder(10).(*array.StringBuilder).Append(stack.FourBytes)
		}

		//BlockTimestamp
		b.Field(10).(*array.Int64Builder).Append(v.BlockTimestamp)
	}

	record := b.NewRecord()
	return wr.Write(record)
}