package reporters

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-core-contracts/lib/go/templates"

	"github.com/onflow/flow-go/cmd/util/ledger/migrations"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/systemcontracts"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
)

// EpochCounterReporter reports the current epoch counter from the FlowEpoch smart contract.
type EpochCounterReporter struct {
	Chain                   flow.Chain
	Log                     zerolog.Logger
	PreviousStateCommitment flow.StateCommitment
}

func (e *EpochCounterReporter) Name() string {
	return "EpochCounterReporter"
}

func (e *EpochCounterReporter) Report(payload []ledger.Payload) error {
	script, address, err := ExecuteCurrentEpochScript(e.Chain, payload)
	if err != nil {
		e.Log.
			Error().
			Err(err).
			Msg("error running GetCurrentEpochCounter script")
		// don't fail the rest of the reporters
		return nil
	}

	if script.Err == nil && script.Value != nil {
		epochCounter := script.Value.ToGoValue().(uint64)
		e.Log.
			Info().
			Uint64("epochCounter", epochCounter).
			Str("flowEpochAddress", address.HexWithPrefix()).
			Msg("Fetched epoch counter from the FlowEpoch smart contract")
	} else {
		e.Log.
			Error().
			Err(script.Err).
			Msg("Failed to get epoch counter")
	}
	return nil
}

func ExecuteCurrentEpochScript(c flow.Chain, payload []ledger.Payload) (*fvm.ScriptProcedure, flow.Address, error) {
	l := migrations.NewView(payload)
	prog := programs.NewEmptyPrograms()
	vm := fvm.NewVirtualMachine(fvm.NewInterpreterRuntime())
	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(c))

	sc, err := systemcontracts.SystemContractsForChain(c.ChainID())
	if err != nil {
		return nil, flow.Address{}, fmt.Errorf("error getting SystemContracts for chain %s: %w", c.String(), err)
	}
	address := sc.Epoch.Address
	scriptCode := templates.GenerateGetCurrentEpochCounterScript(templates.Environment{
		EpochAddress: address.Hex(),
	})
	script := fvm.Script(scriptCode)
	return script, address, vm.Run(ctx, script, l, prog)
}

var _ ledger.Reporter = &EpochCounterReporter{}
