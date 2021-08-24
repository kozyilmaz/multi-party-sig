package presign

import (
	"errors"
	"fmt"

	"github.com/cronokirby/safenum"
	"github.com/taurusgroup/multi-party-sig/internal/elgamal"
	"github.com/taurusgroup/multi-party-sig/internal/hash"
	"github.com/taurusgroup/multi-party-sig/internal/round"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/paillier"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	zkaffg "github.com/taurusgroup/multi-party-sig/pkg/zk/affg"
	zkaffp "github.com/taurusgroup/multi-party-sig/pkg/zk/affp"
)

var _ round.Round = (*presign3)(nil)

type presign3 struct {
	*presign2

	// DeltaShareAlpha[j] = αᵢⱼ
	DeltaShareAlpha map[party.ID]*safenum.Int
	// DeltaShareBeta[j] = βᵢⱼ
	DeltaShareBeta map[party.ID]*safenum.Int
	// ChiShareAlpha[j] = α̂ᵢⱼ
	ChiShareAlpha map[party.ID]*safenum.Int
	// ChiShareBeta[j] = β̂ᵢⱼ
	ChiShareBeta map[party.ID]*safenum.Int

	// DeltaCiphertext[j][k] = Dₖⱼ
	DeltaCiphertext map[party.ID]map[party.ID]*paillier.Ciphertext
	// ChiCiphertext[j][k] = D̂ₖⱼ
	ChiCiphertext map[party.ID]map[party.ID]*paillier.Ciphertext
}

type broadcast3 struct {
	// DeltaCiphertext[k] = Dₖⱼ
	DeltaCiphertext map[party.ID]*paillier.Ciphertext
	// ChiCiphertext[k] = D̂ₖⱼ
	ChiCiphertext map[party.ID]*paillier.Ciphertext
}

type message3 struct {
	broadcast3
	DeltaF     *paillier.Ciphertext // DeltaF = Fᵢⱼ
	DeltaProof *zkaffp.Proof
	ChiF       *paillier.Ciphertext // ChiF = F̂ᵢⱼ
	ChiProof   *zkaffg.Proof
}

// VerifyMessage implements round.Round.
//
// - verify zkaffg, zkaffp.
func (r *presign3) VerifyMessage(msg round.Message) error {
	from, to := msg.From, msg.To
	body, ok := msg.Content.(*message3)
	if !ok || body == nil {
		return round.ErrInvalidContent
	}

	if !body.DeltaProof.Verify(r.HashForID(from), zkaffp.Public{
		Kv:       r.K[to],
		Dv:       body.broadcast3.DeltaCiphertext[to],
		Fp:       body.DeltaF,
		Xp:       r.G[from],
		Prover:   r.Paillier[from],
		Verifier: r.Paillier[to],
		Aux:      r.Pedersen[to],
	}) {
		return errors.New("failed to validate affp proof for Delta MtA")
	}

	if !body.ChiProof.Verify(r.HashForID(from), zkaffg.Public{
		Kv:       r.K[to],
		Dv:       body.broadcast3.ChiCiphertext[to],
		Fp:       body.ChiF,
		Xp:       r.ECDSA[from],
		Prover:   r.Paillier[from],
		Verifier: r.Paillier[to],
		Aux:      r.Pedersen[to],
	}) {
		return errors.New("failed to validate affg proof for Chi MtA")
	}

	return nil
}

// StoreMessage implements round.Round.
//
// - Decrypt MtA shares,
// - save αᵢⱼ, α̂ᵢⱼ.
func (r *presign3) StoreMessage(msg round.Message) error {
	from, body := msg.From, msg.Content.(*message3)

	// αᵢⱼ
	DeltaShareAlpha, err := r.SecretPaillier.Dec(body.broadcast3.DeltaCiphertext[r.SelfID()])
	if err != nil {
		return fmt.Errorf("failed to decrypt alpha share for delta: %w", err)
	}
	// α̂ᵢⱼ
	ChiShareAlpha, err := r.SecretPaillier.Dec(body.broadcast3.ChiCiphertext[r.SelfID()])
	if err != nil {
		return fmt.Errorf("failed to decrypt alpha share for chi: %w", err)
	}

	r.DeltaShareAlpha[from] = DeltaShareAlpha
	r.ChiShareAlpha[from] = ChiShareAlpha

	r.DeltaCiphertext[from] = body.broadcast3.DeltaCiphertext
	r.ChiCiphertext[from] = body.broadcast3.ChiCiphertext
	return nil
}

// Finalize implements round.Round
//
// - Γ = ∑ⱼ Γⱼ
// - Δᵢ = [kᵢ]Γ
// - δᵢ = γᵢ kᵢ + ∑ⱼ αᵢⱼ + βᵢⱼ
// - χᵢ = xᵢ kᵢ + ∑ⱼ α̂ᵢⱼ + β̂ᵢⱼ
// - Ẑⱼ, b̂ⱼ
func (r *presign3) Finalize(out chan<- *round.Message) (round.Session, error) {
	// δᵢ = γᵢ kᵢ
	KShareInt := curve.MakeInt(r.KShare)
	DeltaShare := new(safenum.Int).Mul(r.GammaShare, KShareInt, -1)

	// χᵢ = xᵢ kᵢ
	ChiShare := new(safenum.Int).Mul(curve.MakeInt(r.SecretECDSA), KShareInt, -1)

	for _, j := range r.OtherPartyIDs() {
		//δᵢ += αᵢⱼ + βᵢⱼ
		DeltaShare.Add(DeltaShare, r.DeltaShareAlpha[j], -1)
		DeltaShare.Add(DeltaShare, r.DeltaShareBeta[j], -1)

		// χᵢ += α̂ᵢⱼ + β̂ᵢⱼ
		ChiShare.Add(ChiShare, r.ChiShareAlpha[j], -1)
		ChiShare.Add(ChiShare, r.ChiShareBeta[j], -1)
	}

	// ElGamalChi = Ẑⱼ = (b̂ⱼ⋅G, χᵢ+b̂ⱼ⋅Yᵢ)
	// ElGamalChiNonce = b̂ⱼ
	ElGamalChi, ElGamalChiNonce := elgamal.Encrypt(r.ElGamal[r.SelfID()], r.Group().NewScalar().SetNat(ChiShare.Mod(r.Group().Order())))

	DeltaShareScalar := r.Group().NewScalar().SetNat(DeltaShare.Mod(r.Group().Order()))
	err := r.SendMessage(out, &message4{
		DeltaShare: DeltaShareScalar,
		ElGamalChi: ElGamalChi,
	}, "")
	if err != nil {
		return r, err
	}

	return &presign4{
		presign3:        r,
		ElGamalChiNonce: ElGamalChiNonce,
		ElGamalChi:      map[party.ID]*elgamal.Ciphertext{r.SelfID(): ElGamalChi},
		DeltaShares:     map[party.ID]curve.Scalar{r.SelfID(): DeltaShareScalar},
		ChiShare:        r.Group().NewScalar().SetNat(ChiShare.Mod(r.Group().Order())),
	}, nil
}

// MessageContent implements round.Round.
func (presign3) MessageContent() round.Content { return &message3{} }

// Number implements round.Round.
func (presign3) Number() round.Number { return 3 }

// Init implements round.Content.
func (m *message3) Init(group curve.Curve) {
	m.ChiProof = zkaffg.Empty(group)
}

// BroadcastData implements broadcast.Broadcaster.
func (m broadcast3) BroadcastData() []byte {
	h := hash.New()
	ids := make([]party.ID, 0, len(m.DeltaCiphertext))
	for id := range m.DeltaCiphertext {
		ids = append(ids, id)
	}
	sortedIDs := party.NewIDSlice(ids)
	for _, id := range sortedIDs {
		_ = h.WriteAny(id, m.DeltaCiphertext[id], m.ChiCiphertext[id])
	}
	return h.Sum()
}
