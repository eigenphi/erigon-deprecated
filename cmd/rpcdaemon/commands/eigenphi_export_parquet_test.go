package commands

import (
	"os"
	"testing"
)

func Test_exportParquetWithData(t *testing.T) {

	file, err := os.OpenFile("erigon.tmp.parquet", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	t.Log("Created parquet file:", file.Name())
	data := []ExportTraceParquet{
		{
			BlockNumber:      1,
			TransactionHash:  "hash",
			TransactionIndex: 0,
			FromAddress:      "from",
			ToAddress:        "toaddr",
			GasPrice:         1,
			Input:            "x",
			Nonce:            121,
			TransactionValue: "va",
			Stack: []PlainStackFrame{
				{
					FrameId:         "ccc",
					Type:            "ccc",
					Label:           "ccc",
					From:            "ccc",
					To:              "ccc",
					ContractCreated: "ccc",
					Value:           "ccc",
					Input:           "ccc",
					Error:           "ccc",
					ChildrenCount:   0,
				},
				{
					FrameId:         "bbb",
					Type:            "bbb",
					Label:           "bbb",
					From:            "bbb",
					To:              "bbb",
					ContractCreated: "bbb",
					Value:           "bbb",
					Input:           "bbb",
					Error:           "bbb",
					ChildrenCount:   0,
				},
				{
					FrameId:         "aaa",
					Type:            "aaa",
					Label:           "aaa",
					From:            "aaa",
					To:              "aaa",
					ContractCreated: "aaa",
					Value:           "aaa",
					Input:           "aaa",
					Error:           "aaa",
					ChildrenCount:   0,
				},
			},
		},
		{
			BlockNumber:      2,
			TransactionHash:  "hash",
			TransactionIndex: 2,
			FromAddress:      "from",
			ToAddress:        "toaddr",
			GasPrice:         2,
			Input:            "x",
			Nonce:            222,
			TransactionValue: "va",
			Stack: []PlainStackFrame{
				{
					FrameId:         "222aaa",
					Type:            "222aaa",
					Label:           "222aaa",
					From:            "222aaa",
					To:              "222aaa",
					ContractCreated: "222aaa",
					Value:           "222aaa",
					Input:           "222aaa",
					Error:           "222aaa",
					ChildrenCount:   0,
				},
			},
		},
	}

	if err := ExportParquetWithData(file, data); err != nil {
		t.Fatal(err)
	}
}
