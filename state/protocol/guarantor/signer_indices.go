package guarantor

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/packer"
	"github.com/onflow/flow-go/state/protocol"
)

func FindGuarantors(state protocol.State, guarantee *flow.CollectionGuarantee) ([]flow.Identifier, error) {
	snapshot := state.AtBlockID(guarantee.ReferenceBlockID)
	cluster, err := snapshot.Epochs().Current().ClusterByChainID(guarantee.ChainID)

	if err != nil {
		// protocol state must have validated the block that contains the guarantee, so the cluster
		// must be found, otherwise, it's an internal error
		return nil, fmt.Errorf(
			"internal error retrieving collector clusters for guarantee (ReferenceBlockID: %v, ChainID: %v): %w",
			guarantee.ReferenceBlockID, guarantee.ChainID, err)
	}

	guarantorIDs, err := packer.DecodeSignerIdentifiersFromIndices(cluster.Members().NodeIDs(), guarantee.SignerIndices)
	if err != nil {
		return nil, fmt.Errorf("could not decode guarantor indices: %v", err)
	}

	return guarantorIDs, nil
}
