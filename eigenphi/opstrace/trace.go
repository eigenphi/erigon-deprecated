package trace

import (
	"context"
	"errors"
	"fmt"
	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/core"
	"github.com/ledgerwatch/erigon/core/vm"
	"github.com/ledgerwatch/erigon/core/vm/evmtypes"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"time"
)

// TraceTx configures a new tracer according to the provided configuration, and
// executes the given message in the provided environment. The return value will
// be tracer dependent.
func TraceTxByOpsTracer(
	ctx context.Context,
	message core.Message,
	blockCtx evmtypes.BlockContext,
	txCtx evmtypes.TxContext,
	ibs evmtypes.IntraBlockState,
	config *tracers.TraceConfig,
	chainConfig *chain.Config,
) (*OpsCallFrame, error) {
	// Assemble the structured logger or the JavaScript tracer
	var (
		tracer = NewOpsTracer()
		err    error
	)

	timeout := 5 * time.Minute
	if config.Timeout != nil {
		if timeout, err = time.ParseDuration(*config.Timeout); err != nil {
			return nil, err
		}
	}

	deadlineCtx, cancel := context.WithTimeout(ctx, timeout)
	go func() {
		<-deadlineCtx.Done()
		if t, ok := tracer.(tracers.Tracer); ok {
			t.Stop(errors.New("execution timeout"))
		}
	}()
	defer cancel()

	vmenv := vm.NewEVM(blockCtx, txCtx, ibs, chainConfig, vm.Config{Debug: true, Tracer: tracer})
	var refunds = true
	if config != nil && config.NoRefunds != nil && *config.NoRefunds {
		refunds = false
	}
	core.ApplyMessage(vmenv, message, new(core.GasPool).AddGas(message.Gas()), refunds, false)

	t, ok := tracer.(*OpsTracer)
	if !ok {
		return nil, fmt.Errorf("tracer is not OpsTracer")
	}
	return t.GetCallStack(), nil
}
