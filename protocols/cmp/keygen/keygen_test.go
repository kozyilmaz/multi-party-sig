package keygen

import (
	mrand "math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taurusgroup/multi-party-sig/internal/round"
	"github.com/taurusgroup/multi-party-sig/internal/test"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp/config"
)

var group = curve.Secp256k1{}

func checkOutput(t *testing.T, rounds []round.Session) {
	N := len(rounds)
	newConfigs := make([]*config.Config, 0, N)
	for _, r := range rounds {
		require.IsType(t, &round.Output{}, r)
		resultRound := r.(*round.Output)
		require.IsType(t, &config.Config{}, resultRound.Result)
		c := resultRound.Result.(*config.Config)
		newConfigs = append(newConfigs, c)
	}

	firstConfig := newConfigs[0]
	pk := firstConfig.PublicPoint()
	for _, c := range newConfigs {
		assert.True(t, pk.Equal(c.PublicPoint()), "RID is different")
		assert.Equal(t, firstConfig.RID, c.RID, "RID is different")
		assert.EqualValues(t, firstConfig.ChainKey, c.ChainKey, "ChainKey is different")
		for id, p := range firstConfig.Public {
			assert.True(t, p.Equal(c.Public[id]), "public is not equal")
		}
		assert.NoError(t, c.Validate(), "failed to validate new config")
	}
}

func TestKeygen(t *testing.T) {
	pl := pool.NewPool(0)
	defer pl.TearDown()

	N := 2
	partyIDs := test.PartyIDs(N)

	rounds := make([]round.Session, 0, N)
	for _, partyID := range partyIDs {
		r, err := StartKeygen(group, partyIDs, N-1, partyID, pl)(nil)
		require.NoError(t, err, "round creation should not result in an error")
		rounds = append(rounds, r)
	}

	for {
		err, done := test.Rounds(group, rounds, nil)
		require.NoError(t, err, "failed to process round")
		if done {
			break
		}
	}
	checkOutput(t, rounds)
}

func TestRefresh(t *testing.T) {
	pl := pool.NewPool(0)
	defer pl.TearDown()

	N := 4
	T := N - 1
	configs, _ := test.GenerateConfig(group, N, T, mrand.New(mrand.NewSource(1)), pl)

	rounds := make([]round.Session, 0, N)
	for _, c := range configs {
		r, err := StartRefresh(c, pl)(nil)
		require.NoError(t, err, "round creation should not result in an error")
		rounds = append(rounds, r)

	}

	for {
		err, done := test.Rounds(group, rounds, nil)
		require.NoError(t, err, "failed to process round")
		if done {
			break
		}
	}
	checkOutput(t, rounds)
}
