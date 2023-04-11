// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trace

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/holiman/uint256"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/core/vm/stack"
	"github.com/ledgerwatch/erigon/eigenphi/utils/fourbyte"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"strconv"
	"strings"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/core/vm"
)

const (
	LabelTransfer         = "Transfer"
	LabelInternalTransfer = "Internal-Transfer"
)

type OpsCallFrame struct {
	Type            string          `json:"type"`
	Label           string          `json:"label"`
	From            string          `json:"from"`
	To              string          `json:"to,omitempty"`
	ContractCreated string          `json:"contract_created,omitempty"`
	Value           string          `json:"value,omitempty"`
	GasIn           string          `json:"gasIn"`
	GasCost         string          `json:"gasCost"`
	Input           string          `json:"input,omitempty"`
	FourBytes       string          `json:"four_bytes"`
	Error           string          `json:"error,omitempty"`
	Calls           []*OpsCallFrame `json:"calls,omitempty"`
	parent          *OpsCallFrame   `json:"-"`
	code            []byte          `json:"-"` // for calculating CREATE2 contract address
	salt            *uint256.Int    `json:"-"` // for calculating CREATE2 contract address
}

var _ tracers.Tracer = (*OpsTracer)(nil)

type OpsTracer struct {
	callstack    OpsCallFrame
	currentDepth int
	currentFrame *OpsCallFrame
	interrupt    uint32 // Atomic flag to signal execution interruption
	initialized  bool
}

func (t *OpsTracer) CaptureTxStart(gasLimit uint64) {
	fmt.Println("CaptureTxStart", gasLimit)
	return
}

func (t *OpsTracer) CaptureTxEnd(restGas uint64) {
	fmt.Println("CaptureTxEnd", restGas)
	return
}

func (t *OpsTracer) CaptureEnter(op vm.OpCode, from libcommon.Address, to libcommon.Address, precompile bool,
	create bool, input []byte, gas uint64, value *uint256.Int, code []byte) {
	fmt.Println("CaptureEnter", op, from, to, precompile, create,
		hex.EncodeToString(input), gas, bigToHex(value))
	if precompile {
		return
	}
}

func (t *OpsTracer) CaptureExit(output []byte, usedGas uint64, err error) {
	fmt.Println("CaptureExit", hex.EncodeToString(output), usedGas, err)
}

func (t *OpsTracer) Stop(err error) {
	fmt.Println("Stop", err)
	return
}

// newOpsTracer returns a native go tracer which tracks
// call frames of a tx, and implements vm.EVMLogger.
func NewOpsTracer() tracers.Tracer {
	// First callframe contains tx context info
	// and is populated on start and end.
	return &OpsTracer{}
}

var labelDb *fourbyte.Database

func init() {
	labelDb, _ = fourbyte.New()
}

// CaptureStart implements the EVMLogger interface to initialize the tracing operation.
func (t *OpsTracer) CaptureStart(env vm.VMInterface, from, to common.Address,
	precompile bool, create bool, input []byte, gas uint64, value *uint256.Int, code []byte) {
	fmt.Println("CaptureStart", from.String(), to.String(), precompile,
		create, hex.EncodeToString(input), gas, bigToHex(value))
	if precompile {
		return
	}

	if t.initialized {
		frame := OpsCallFrame{
			parent: t.currentFrame,
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
	} else {
		t.callstack = OpsCallFrame{}
		t.currentFrame = &t.callstack
		t.initialized = true
	}

	t.currentFrame.Type = "CALL"
	t.currentFrame.From = addrToHex(from)
	t.currentFrame.To = addrToHex(to)
	t.currentFrame.GasIn = uintToHex(gas)
	t.currentFrame.Value = bigToHex(value)
	t.currentFrame.FourBytes = getInputFourBytes(input)
	if create {
		t.currentFrame.Type = "CREATE"
		t.currentFrame.ContractCreated = addrToHex(to)
	}
}

// CaptureEnd is called after the call finishes to finalize the tracing.
// func (t *OpsTracer) CaptureEnd(depth int, output []byte, startGas, endGas uint64, duration time.Duration, err error) error {
// fmt.Println("CaptureEnd", depth, t.currentDepth, err)
// precompiled calls don't have a callframe
func (t *OpsTracer) CaptureEnd(output []byte, usedGas uint64, err error) {
	fmt.Println("CaptureEnd", hex.EncodeToString(output), usedGas, err)
	t.currentFrame.GasCost = uintToHex(usedGas)
	if err != nil {
		t.currentFrame.Error = err.Error()
		t.currentFrame.Calls = []*OpsCallFrame{}
	}

	t.currentFrame = t.currentFrame.parent
	t.currentDepth -= 1
	return
}

// Note the result has no "0x" prefix
func getLogValueHex(stack *stack.Stack, memory *vm.Memory) string {
	offset := stack.Back(0).Uint64()
	length := stack.Back(1).Uint64()
	if memory.Len() < int(offset+length) {
		memory.Resize(offset + length)
	}
	return hex.EncodeToString(memory.Data()[offset : offset+length])
}

// code modified from `4byte.go`
func (t *OpsTracer) isPrecompiled(env *vm.EVM, addr common.Address) bool {
	activePrecompiles := vm.ActivePrecompiles(env.ChainRules())
	for _, p := range activePrecompiles {
		if p == addr {
			return true
		}
	}
	return false
}

func (t *OpsTracer) getLabel(topic0 string) string {
	//if op != vm.LOG0 {
	topic0Bs, _ := hex.DecodeString(topic0)
	label, _ := labelDb.Selector(topic0Bs)
	//}
	return label
}

func getInputFourBytes(input []byte) string {
	if len(input) < 4 {
		return ""
	}
	return hex.EncodeToString(input[:4])
}

// CaptureState implements the EVMLogger interface to trace a single step of VM execution.
//func (t *OpsTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *stack.Stack, rData []byte, contract *vm.Contract, depth int, err error) error {

func (t *OpsTracer) CaptureState(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
	fmt.Println("CaptureState", pc, op.String(), gas, cost, depth, err)
	stack := scope.Stack
	contract := scope.Contract
	memory := scope.Memory
	if err != nil {
		if t.currentFrame != nil {
			t.currentFrame.Error = err.Error()
			t.currentFrame.Calls = []*OpsCallFrame{}
			t.currentFrame = t.currentFrame.parent
			t.currentDepth -= 1
		}
		return
	}
	// Fix txs like 0x3494b6a2f62a558c46660691f68e4e2a47694e0b02fad1969e1f0dc725fc9ee5,
	// where a sub-CALL is failed but the whole tx is not reverted.
	if t.currentDepth == depth+1 && (t.currentFrame.Type == vm.CALL.String() ||
		t.currentFrame.Type == vm.CALLCODE.String() ||
		t.currentFrame.Type == vm.DELEGATECALL.String() ||
		t.currentFrame.Type == vm.STATICCALL.String() ||
		t.currentFrame.Type == vm.CREATE.String() ||
		t.currentFrame.Type == vm.CREATE2.String()) {
		t.currentFrame.Error = "Subcall reverted"
		t.currentFrame = t.currentFrame.parent
		t.currentDepth -= 1
	}

	if op == vm.LOG0 || op == vm.LOG1 || op == vm.LOG2 || op == vm.LOG3 || op == vm.LOG4 {
		var topic0, topic1, topic2, topic3, logInput string
		switch op {
		case vm.LOG1:
			topic0 = stack.Back(2).String()[2:] // remove "0x" prefix
			logInput = topic0
		case vm.LOG2:
			topic0 = stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = stack.Back(3).String()[2:] // remove "0x" prefix
			logInput = strings.Join([]string{topic0, topic1}, " ")
		case vm.LOG3:
			topic0 = stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = stack.Back(3).String()[2:] // remove "0x" prefix
			topic2 = stack.Back(4).String()[2:] // remove "0x" prefix
			logInput = strings.Join([]string{topic0, topic1, topic2}, " ")
		case vm.LOG4:
			topic0 = stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = stack.Back(3).String()[2:] // remove "0x" prefix
			topic2 = stack.Back(4).String()[2:] // remove "0x" prefix
			topic3 = stack.Back(5).String()[2:] // remove "0x" prefix
			logInput = strings.Join([]string{topic0, topic1, topic2, topic3}, " ")
		}
		var label = t.getLabel(topic0)
		frame := OpsCallFrame{
			Type:    op.String(),
			Label:   label,
			From:    strings.ToLower(contract.Address().String()),
			Input:   logInput,
			Value:   getLogValueHex(stack, memory),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		return
	}

	switch op {
	case vm.CREATE, vm.CREATE2:
		value := stack.Back(0)
		from := contract.Address()
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(from.String()),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			Value:   value.String(),
			parent:  t.currentFrame,
		}
		if op == vm.CREATE2 {
			frame.salt = stack.Back(3)
		}
		if !value.IsZero() {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
	case vm.SELFDESTRUCT:
		var to common.Address = stack.Back(0).Bytes20()
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(contract.Address().String()),
			To:      strings.ToLower(to.String()),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			//Value:   value.String(),
			parent: t.currentFrame,
		}
		//if value.Uint64() != 0 {
		//	frame.Label = LabelInternalTransfer
		//}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
	case vm.CALL, vm.CALLCODE:
		var to common.Address = stack.Back(1).Bytes20()
		//if t.isPrecompiled(env, to) {
		//	return
		//}
		value := stack.Back(2)
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(contract.Address().String()),
			To:      strings.ToLower(to.String()),
			Value:   value.String(),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		if !value.IsZero() {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
	case vm.DELEGATECALL, vm.STATICCALL:
		var to common.Address = stack.Back(1).Bytes20()
		//if t.isPrecompiled(env, to) {
		//	return
		//}

		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(contract.Address().String()),
			To:      strings.ToLower(to.String()),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}

		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
	}
	return
}

// CaptureFault implements the EVMLogger interface to trace an execution fault.
// func (t *OpsTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *stack.Stack, contract *vm.Contract, depth int, err error) error {
// fmt.Println("CaptureFault", pc, op, gas, cost, depth, err)
func (t *OpsTracer) CaptureFault(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {
	fmt.Println("CaptureFault", pc, op.String(), gas, cost, depth, err)
	return
}

// GetResult returns the json-encoded nested list of call traces, and any
// error arising from the encoding or forceful termination (via `Stop`).
func (t *OpsTracer) GetResult() (json.RawMessage, error) {
	if len(t.callstack.Error) != 0 {
		t.callstack.Calls = []*OpsCallFrame{}
	}
	errString := t.callstack.Error
	var traceErr error
	if len(errString) > 0 {
		t.callstack.Calls = []*OpsCallFrame{}
		traceErr = errors.New(errString)
	}
	res, err := json.Marshal(t.callstack)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(res), traceErr
}

func (t *OpsTracer) GetCallStack() *OpsCallFrame {
	if len(t.callstack.Error) != 0 {
		t.callstack.Calls = []*OpsCallFrame{}
	}
	errString := t.callstack.Error
	if len(errString) > 0 {
		t.callstack.Calls = []*OpsCallFrame{}
	}
	return &t.callstack
}

func addrToHex(a common.Address) string {
	return strings.ToLower(a.Hex())
}

func uintToHex(n uint64) string {
	return "0x" + strconv.FormatUint(n, 16)
}

func bigToHex(n *uint256.Int) string {
	if n == nil {
		return ""
	}
	return n.Hex()
}
