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

package native

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ledgerwatch/erigon/common"
	"github.com/ledgerwatch/erigon/core/vm"
)

// FIXME: call stack hierarchy has not been properly handled

const (
	LabelTransfer         = "Transfer"
	LabelInternalTransfer = "Internal-Transfer"
)

type opsCallFrame struct {
	Type    string          `json:"type"`
	Label   string          `json:"label"`
	From    string          `json:"from"`
	To      string          `json:"to,omitempty"`
	Value   string          `json:"value,omitempty"`
	GasIn   string          `json:"gasIn"`
	GasCost string          `json:"gasUsed"`
	Input   string          `json:"input"`
	Output  string          `json:"output,omitempty"`
	Error   string          `json:"error,omitempty"`
	Calls   []*opsCallFrame `json:"calls,omitempty"`
	parent  *opsCallFrame   `json:"-"`
}

type opsTracer struct {
	callstack    opsCallFrame
	currentDepth int
	currentFrame *opsCallFrame
	interrupt    uint32 // Atomic flag to signal execution interruption
	reason       error  // Textual reason for the interruption
}

// newOpsTracer returns a native go tracer which tracks
// call frames of a tx, and implements vm.EVMLogger.
func NewOpsTracer() vm.Tracer {
	// First callframe contains tx context info
	// and is populated on start and end.
	return &opsTracer{}
}

// CaptureStart implements the EVMLogger interface to initialize the tracing operation.
func (t *opsTracer) CaptureStart(env *vm.EVM, depth int, from, to common.Address,
	precompile, create bool, callType vm.CallType, input []byte, gas uint64,
	value *big.Int, code []byte) {
	fmt.Println("CaptureStart", env, from, to, create, input, gas, value)
	t.callstack = opsCallFrame{
		Type:  "CALL",
		From:  addrToHex(from),
		To:    addrToHex(to),
		Input: bytesToHex(input),
		GasIn: uintToHex(gas),
		Value: bigToHex(value),
	}
	if create {
		t.callstack.Type = "CREATE"
	}
	t.currentDepth = depth
	t.currentFrame = &t.callstack
}

// CaptureEnd is called after the call finishes to finalize the tracing.
func (t *opsTracer) CaptureEnd(depth int, output []byte, startGas, endGas uint64, duration time.Duration, err error) {
	fmt.Println("CaptureEnd", output, startGas, endGas, err)
	t.callstack.GasCost = uintToHex(startGas - endGas)
	if err != nil {
		t.callstack.Error = err.Error()
		if err.Error() == "execution reverted" && len(output) > 0 {
			t.callstack.Output = bytesToHex(output)
		}
	} else {
		t.callstack.Output = bytesToHex(output)
	}
}

// Note the result has no "0x" prefix
func getLogValueHex(scope *vm.ScopeContext) string {
	offset := scope.Stack.Back(0).Uint64()
	length := scope.Stack.Back(1).Uint64()
	return hex.EncodeToString(scope.Memory.Data()[offset : offset+length])
}

// code modified from `4byte.go`
func (t *opsTracer) isPrecompiled(env *vm.EVM, addr common.Address) bool {
	activePrecompiles := vm.ActivePrecompiles(env.ChainRules())
	for _, p := range activePrecompiles {
		if p == addr {
			return true
		}
	}
	return false
}

// CaptureState implements the EVMLogger interface to trace a single step of VM execution.
func (t *opsTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
	if err != nil {
		t.reason = err
		return
	}
	if depth < t.currentDepth {
		t.currentFrame = t.currentFrame.parent
		t.currentDepth -= 1
	}
	if op == vm.LOG0 || op == vm.LOG1 || op == vm.LOG2 || op == vm.LOG3 || op == vm.LOG4 {
		var topic0, topic1, topic2, topic3, logInput string
		switch op {
		case vm.LOG1:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			logInput = topic0
		case vm.LOG2:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = scope.Stack.Back(3).String()[2:] // remove "0x" prefix
			logInput = topic0 + topic1
		case vm.LOG3:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = scope.Stack.Back(3).String()[2:] // remove "0x" prefix
			topic2 = scope.Stack.Back(4).String()[2:] // remove "0x" prefix
			logInput = topic0 + topic1 + topic2
		case vm.LOG4:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = scope.Stack.Back(3).String()[2:] // remove "0x" prefix
			topic2 = scope.Stack.Back(4).String()[2:] // remove "0x" prefix
			topic3 = scope.Stack.Back(5).String()[2:] // remove "0x" prefix
			logInput = topic0 + topic1 + topic2 + topic3
		}
		var label string
		// FIXME: add docs about the magic number
		if op != vm.LOG0 && topic0 == "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
			label = LabelTransfer
		}
		frame := opsCallFrame{
			Type:    op.String(),
			Label:   label,
			From:    scope.Contract.Address().String(),
			Input:   logInput,
			Value:   getLogValueHex(scope),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		return
	}

	switch op {
	case vm.CREATE, vm.CREATE2:
		value := scope.Stack.Back(0)
		offset := scope.Stack.Back(1).Uint64()
		length := scope.Stack.Back(2).Uint64()
		frame := opsCallFrame{
			Type:    op.String(),
			From:    scope.Contract.Address().String(),
			Input:   hex.EncodeToString(scope.Memory.Data()[offset : offset+length]),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			Value:   value.String(),
			parent:  t.currentFrame,
		}
		if !value.IsZero() {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
		return
	case vm.SELFDESTRUCT:
		// TODO: capture in `CaptureSelfDestruct`?
		value := env.IntraBlockState().GetBalance(scope.Contract.Address())
		frame := opsCallFrame{
			Type:    op.String(),
			From:    scope.Contract.Address().String(),
			To:      scope.Stack.Back(0).String(),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			Value:   value.String(),
			parent:  t.currentFrame,
		}
		if value.Uint64() != 0 {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		return
	case vm.CALL, vm.CALLCODE:
		var to common.Address = scope.Stack.Back(1).Bytes20()
		if t.isPrecompiled(env, to) {
			return
		}
		argOffset := scope.Stack.Back(3).Uint64()
		argLength := scope.Stack.Back(4).Uint64()
		value := scope.Stack.Back(2)
		frame := opsCallFrame{
			Type:    op.String(),
			From:    scope.Contract.Address().String(),
			To:      to.String(),
			Value:   value.String(),
			Input:   hex.EncodeToString(scope.Memory.Data()[argOffset : argOffset+argLength]),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		if value.IsZero() {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		return
	case vm.DELEGATECALL, vm.STATICCALL:
		var to common.Address = scope.Stack.Back(1).Bytes20()
		if t.isPrecompiled(env, to) {
			return
		}
		argOffset := scope.Stack.Back(2).Uint64()
		argLength := scope.Stack.Back(3).Uint64()
		frame := opsCallFrame{
			Type:    op.String(),
			From:    scope.Contract.Address().String(),
			To:      to.String(),
			Input:   hex.EncodeToString(scope.Memory.Data()[argOffset : argOffset+argLength]),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		return
	}
}

// CaptureFault implements the EVMLogger interface to trace an execution fault.
func (t *opsTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64,
	scope *vm.ScopeContext, depth int, err error) {
	fmt.Println("CaptureFault", pc, op, gas, cost, depth, err)
}

func (t *opsTracer) CaptureSelfDestruct(from common.Address, to common.Address, value *big.Int) {
}
func (t *opsTracer) CaptureAccountRead(account common.Address) error {
	return nil
}
func (t *opsTracer) CaptureAccountWrite(account common.Address) error {
	return nil
}

// GetResult returns the json-encoded nested list of call traces, and any
// error arising from the encoding or forceful termination (via `Stop`).
func (t *opsTracer) GetResult() (json.RawMessage, error) {
	res, err := json.Marshal(t.callstack)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(res), t.reason
}

func addrToHex(a common.Address) string {
	return strings.ToLower(a.Hex())
}

func bytesToHex(s []byte) string {
	return "0x" + common.Bytes2Hex(s)
}

func uintToHex(n uint64) string {
	return "0x" + strconv.FormatUint(n, 16)
}

func bigToHex(n *big.Int) string {
	if n == nil {
		return ""
	}
	return "0x" + n.Text(16)
}
