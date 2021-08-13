package signature

import (
	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/model/flow"
)

type ConsensusSigPackerImpl struct {
	committees hotstuff.Committee
}

func (p *ConsensusSigPackerImpl) Combine(sig *hotstuff.AggregatedSignatureData) ([]flow.Identifier, []byte, error) {
	panic("to be implemented")
}

func (p *ConsensusSigPackerImpl) Split(signerIDs []flow.Identifier, sigData []byte) (*hotstuff.AggregatedSignatureData, error) {
	panic("to be implemented")
}
