//go:build relic
// +build relic

package verification

import (
	"fmt"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/crypto/hash"
	"github.com/onflow/flow-go/model/encoding"
	"github.com/onflow/flow-go/model/flow"
)

// StakingVerifier is a verifier capable of verifying staking signature for each
// verifying operation. It's used primarily with collection cluster where hotstuff without beacon signers is used.
type StakingVerifier struct {
	stakingHasher hash.Hasher
	// TODO: to be replaced by module/signature.PublicKeyAggregator in V2
	keysAggregator *stakingKeysAggregator
}

// NewStakingVerifier creates a new single verifier with the given dependencies.
func NewStakingVerifier() *StakingVerifier {
	return &StakingVerifier{
		stakingHasher:  crypto.NewBLSKMAC(encoding.CollectorVoteTag),
		keysAggregator: newStakingKeysAggregator(),
	}
}

// VerifyVote verifies the validity of a single signature from a vote.
// Usually this method is only used to verify the proposer's vote, which is
// the vote included in a block proposal.
// TODO: return error only, because when the sig is invalid, the returned bool
func (v *StakingVerifier) VerifyVote(signer *flow.Identity, sigData []byte, block *model.Block) (bool, error) {

	// create the to-be-signed message
	msg := MakeVoteMessage(block.View, block.BlockID)

	// verify each signature against the message
	stakingValid, err := signer.StakingPubKey.Verify(sigData, msg, v.stakingHasher)
	if err != nil {
		return false, fmt.Errorf("internal error while verifying staking signature: %w", err)
	}
	if !stakingValid {
		return false, fmt.Errorf("invalid staking sig")
	}

	return true, nil
}

// VerifyQC verifies the validity of a single signature on a quorum certificate.
//
// In the single verification case, `sigData` represents a single signature (`crypto.Signature`).
func (v *StakingVerifier) VerifyQC(signers flow.IdentityList, sigData []byte, block *model.Block) (bool, error) {
	// verify the aggregated staking signature
	msg := MakeVoteMessage(block.View, block.BlockID)

	aggregatedKey, err := v.keysAggregator.aggregatedStakingKey(signers)
	if err != nil {
		return false, fmt.Errorf("could not compute aggregated key: %w", err)
	}
	stakingValid, err := aggregatedKey.Verify(sigData, msg, v.stakingHasher)
	if err != nil {
		return false, fmt.Errorf("internal error while verifying staking signature: %w", err)
	}
	return stakingValid, nil
}
