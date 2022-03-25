package reporters

import (
	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/cmd/util/ledger/migrations"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
)

// EpochContractReporter reports the current epoch counter from the FlowEpoch smart contract.
type EpochContractReporter struct {
	Chain flow.Chain
	Log   zerolog.Logger
}

func (e *EpochContractReporter) Name() string {
	return "EpochContractReporter"
}

func (e *EpochContractReporter) Report(payload []ledger.Payload) error {
	l := migrations.NewView(payload)
	prog := programs.NewEmptyPrograms()
	vm := fvm.NewVirtualMachine(fvm.NewInterpreterRuntime())
	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(e.Chain))

	scriptCode := `
	pub fun main(account: Address, contract: String): String {
		return String.encodeHex(getAccount(account).contracts.get(name: contract)!.code)
	}
	`
	script := fvm.Script([]byte(scriptCode))
	script.WithArguments(jsoncdc.MustEncode(
		cadence.BytesToAddress(flow.HexToAddress("0x9eca2b38b18b5dfe").Bytes())),
		jsoncdc.MustEncode(cadence.String("FlowIDTableStaking")),
	)

	err := vm.Run(ctx, script, l, prog)
	if err != nil {
		e.Log.
			Error().
			Err(err).
			Msg("error running get FlowIDTableStaking contract script")
		// don't fail the rest of the reporters
		return nil
	}

	if script.Err == nil && script.Value != nil {
		epochCounter := script.Value.ToGoValue().(string)
		e.Log.
			Info().
			Str("FlowIDTableStaking", epochCounter).
			Msg("Fetched FlowIDTableStaking from 0x9eca2b38b18b5dfe")
	} else {
		e.Log.
			Error().
			Err(script.Err).
			Msg("Failed to get FlowIDTableStaking")
	}
	return nil
}

var _ ledger.Reporter = &EpochContractReporter{}
