package zkenc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taurusgroup/cmp-ecdsa/pkg/hash"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/sample"
	"github.com/taurusgroup/cmp-ecdsa/pkg/zk"
)

func TestEnc(t *testing.T) {
	verifier := zk.Pedersen
	prover := zk.ProverPaillierPublic

	k := sample.IntervalL()
	K, rho := prover.Enc(k)
	public := Public{
		K:      K,
		Prover: prover,
		Aux:    verifier,
	}

	proof := NewProof(hash.New(), public, Private{
		K:   k,
		Rho: rho,
	})
	out, err := proof.Marshal()
	require.NoError(t, err, "failed to marshal proof")
	proof2 := &Proof{}
	require.NoError(t, proof2.Unmarshal(out), "failed to unmarshal proof")
	assert.Equal(t, proof, proof2)
	out2, err := proof2.Marshal()
	assert.Equal(t, out, out2)
	proof3 := &Proof{}
	require.NoError(t, proof3.Unmarshal(out2), "failed to marshal 2nd proof")
	assert.Equal(t, proof, proof3)

	assert.True(t, proof2.Verify(hash.New(), public))
}
