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
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon/crypto"
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

type OpsCallFrame struct {
	Type    string           `json:"type"`
	Label   string           `json:"label"`
	From    string           `json:"from"`
	To      string           `json:"to,omitempty"`
	Value   string           `json:"value,omitempty"`
	GasIn   string           `json:"gasIn"`
	GasCost string           `json:"gasCost"`
	Input   string           `json:"input,omitempty"`
	Error   string           `json:"error,omitempty"`
	Calls   []*OpsCallFrame  `json:"calls,omitempty"`
	parent  *OpsCallFrame    `json:"-"`
	scope   *vm.ScopeContext `json:"-"`
	code    []byte           `json:"-"` // for calculating CREATE2 contract address
	salt    *uint256.Int     `json:"-"` // for calculating CREATE2 contract address
}

type OpsTracer struct {
	callstack    OpsCallFrame
	currentDepth int
	currentFrame *OpsCallFrame
	interrupt    uint32 // Atomic flag to signal execution interruption
	reason       error  // Textual reason for the interruption
	initialized  bool
}

// newOpsTracer returns a native go tracer which tracks
// call frames of a tx, and implements vm.EVMLogger.
func NewOpsTracer() vm.Tracer {
	// First callframe contains tx context info
	// and is populated on start and end.
	return &OpsTracer{}
}

// CaptureStart implements the EVMLogger interface to initialize the tracing operation.
func (t *OpsTracer) CaptureStart(env *vm.EVM, depth int, from, to common.Address,
	precompile, create bool, callType vm.CallType, input []byte, gas uint64,
	value *big.Int, code []byte) {

	fmt.Println("CaptureStart", depth, callType)
	if callType == vm.CREATE2T {
		create2Frame := t.currentFrame
		codeHash := crypto.Keccak256Hash(input)
		contractAddr := crypto.CreateAddress2(
			common.HexToAddress(create2Frame.From),
			common.Hash(create2Frame.salt.Bytes32()),
			codeHash.Bytes())
		create2Frame.To = contractAddr.String()
	}
	if t.initialized {
		return
	}
	t.callstack = OpsCallFrame{
		Type:  "CALL",
		From:  addrToHex(from),
		To:    addrToHex(to),
		GasIn: uintToHex(gas),
		Value: bigToHex(value),
	}
	if create {
		t.callstack.Type = "CREATE"
	}
	t.currentDepth = depth + 1 // depth is the value before "CALL" or "CREATE"
	t.currentFrame = &t.callstack
	t.initialized = true
}

// CaptureEnd is called after the call finishes to finalize the tracing.
func (t *OpsTracer) CaptureEnd(depth int, output []byte, startGas, endGas uint64, duration time.Duration, err error) {
	fmt.Println("CaptureEnd", depth, t.currentDepth, err)
	// precompiled calls don't have a callframe
	if depth == t.currentDepth {
		return
	}
	t.currentFrame.GasCost = uintToHex(startGas - endGas)
	if err != nil {
		t.currentFrame.Error = err.Error()
	}

	t.currentFrame = t.currentFrame.parent
	t.currentDepth -= 1
}

// Note the result has no "0x" prefix
func getLogValueHex(scope *vm.ScopeContext) string {
	offset := scope.Stack.Back(0).Uint64()
	length := scope.Stack.Back(1).Uint64()
	if scope.Memory.Len() < int(offset+length) {
		scope.Memory.Resize(offset + length)
	}
	return hex.EncodeToString(scope.Memory.Data()[offset : offset+length])
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

// CaptureState implements the EVMLogger interface to trace a single step of VM execution.
func (t *OpsTracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
	fmt.Println("CaptureState", depth, t.currentDepth, op.String())
	if err != nil {
		t.reason = err
		if t.currentFrame != nil {
			t.currentFrame.Error = err.Error()
		}
		return
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
		frame := OpsCallFrame{
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
		from := scope.Contract.Address()
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    from.String(),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			Value:   value.String(),
			parent:  t.currentFrame,
			scope:   scope,
		}
		if op == vm.CREATE {
			nonce := env.IntraBlockState().GetNonce(from)
			frame.To = crypto.CreateAddress(from, nonce).String()
		}
		if op == vm.CREATE2 {
			frame.salt = scope.Stack.Back(3)
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
		frame := OpsCallFrame{
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
		value := scope.Stack.Back(2)
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    scope.Contract.Address().String(),
			To:      to.String(),
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
		return
	case vm.DELEGATECALL, vm.STATICCALL:
		var to common.Address = scope.Stack.Back(1).Bytes20()
		if t.isPrecompiled(env, to) {
			return
		}
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    scope.Contract.Address().String(),
			To:      to.String(),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
		return
	}
}

// CaptureFault implements the EVMLogger interface to trace an execution fault.
func (t *OpsTracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64,
	scope *vm.ScopeContext, depth int, err error) {
	fmt.Println("CaptureFault", pc, op, gas, cost, depth, err)
}

func (t *OpsTracer) CaptureSelfDestruct(from common.Address, to common.Address, value *big.Int) {
}
func (t *OpsTracer) CaptureAccountRead(account common.Address) error {
	return nil
}
func (t *OpsTracer) CaptureAccountWrite(account common.Address) error {
	return nil
}

// GetResult returns the json-encoded nested list of call traces, and any
// error arising from the encoding or forceful termination (via `Stop`).
func (t *OpsTracer) GetResult() (json.RawMessage, error) {
	res, err := json.Marshal(t.callstack)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(res), t.reason
}

func (t *OpsTracer) GetCallStack() *OpsCallFrame {
	return &t.callstack
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
