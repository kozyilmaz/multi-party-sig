package refresh

import (
	"fmt"

	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/polynomial"
	"github.com/taurusgroup/cmp-ecdsa/pkg/party"
	"github.com/taurusgroup/cmp-ecdsa/pkg/round"
	"github.com/taurusgroup/cmp-ecdsa/pkg/session"
	zkmod "github.com/taurusgroup/cmp-ecdsa/pkg/zk/mod"
	zkprm "github.com/taurusgroup/cmp-ecdsa/pkg/zk/prm"
	zksch "github.com/taurusgroup/cmp-ecdsa/pkg/zk/sch"
)

type output struct {
	*round4
}

// ProcessMessage implements round.Round
//
// - verify all Schnorr proofs
//   - if refresh, skip first coefficient
// - decrypt share
// - verify VSS
func (r *output) ProcessMessage(msg round.Message) error {
	j := msg.GetHeader().From
	partyJ := r.LocalParties[j]
	body := msg.(*Message).GetRefresh4()

	// verify Schnorr proofs
	if len(body.VSSSchnorrResponse) != len(r.Self.VSSCommitments) {
		return fmt.Errorf("refresh.output.ProcessMessage(): party %s: wrong number of Schnorr proofs", j)
	}

	coefficient := partyJ.VSSPolynomial.Coefficients
	if !zksch.VerifyMulti(r.Hash.CloneWithID(j), partyJ.VSSCommitments, coefficient, body.VSSSchnorrResponse) {
		return fmt.Errorf("refresh.output.ProcessMessage(): party %s: failed to validate sch proof", j)
	}

	// decrypt share
	xJdec := r.PaillierSecret.Dec(body.Share)
	xJ := curve.NewScalarBigInt(xJdec)
	if xJdec.Cmp(xJ.BigInt()) != 0 {
		return fmt.Errorf("refresh.output.ProcessMessage(): party %s: decrypted share is not in correct range", j)
	}

	// verify share with VSS
	index := r.SelfID.Scalar()
	vss := partyJ.VSSPolynomial.Evaluate(index) // Fⱼ(idJ)
	X := curve.NewIdentityPoint().ScalarBaseMult(xJ)
	if !X.Equal(vss) {
		return fmt.Errorf("refresh.output.ProcessMessage(): party %s: failed to validate share from VSS", j)
	}

	// verify zkmod
	if !body.Mod.Verify(r.Hash.CloneWithID(j), zkmod.Public{N: partyJ.Public.Pedersen.N}) {
		return fmt.Errorf("refresh.output.ProcessMessage(): party %s: mod proof failed to verify", j)
	}

	// verify zkprm
	if !body.Prm.Verify(r.Hash.CloneWithID(j), zkprm.Public{Pedersen: partyJ.Public.Pedersen}) {
		return fmt.Errorf("refresh.output.ProcessMessage(): party %s: prm proof failed to verify", j)
	}

	partyJ.ShareReceived = xJ

	return partyJ.AddMessage(msg)
}

// GenerateMessages implements round.Round
//
// - sum of all received shares
// - compute group public key and individual public keys
// - recompute session SSID
// - validate Session
func (r *output) GenerateMessages() ([]round.Message, error) {
	// add all shares to our secret
	newSecretShareECDSA := curve.NewScalar()
	if r.isRefresh() {
		newSecretShareECDSA.Add(newSecretShareECDSA, r.S.Secret().ECDSA)
	}
	for _, partyJ := range r.LocalParties {
		newSecretShareECDSA.Add(newSecretShareECDSA, partyJ.ShareReceived)
	}
	newSecret := &party.Secret{
		ID:       r.SelfID,
		ECDSA:    newSecretShareECDSA,
		Paillier: r.PaillierSecret,
	}
	// set RID if we are in keygen
	if r.isKeygen() {
		newSecret.RID = r.rho
	} else {
		newSecret.RID = r.S.Secret().RID
	}

	// [F₁(X), ..., Fₙ(X)]
	allPolyExps := make([]*polynomial.Exponent, r.S.N())
	for i, partyIDJ := range r.S.PartyIDs() {
		partyJ := r.LocalParties[partyIDJ]
		allPolyExps[i] = partyJ.VSSPolynomial
	}
	// summedPoly = F(X) = ∑Fⱼ(X)
	summedPoly, err := polynomial.Sum(allPolyExps)
	if err != nil {
		return nil, fmt.Errorf("refresh.output.GenerateMessages(): sum polynomial exponent: %w", err)
	}

	// compute public key Xⱼ = F(j)
	for _, idJ := range r.S.PartyIDs() {
		index := idJ.Scalar()
		newPublicShareECDSA := summedPoly.Evaluate(index)
		if r.isRefresh() {
			oldPublicShareECDSA := r.S.Public(idJ).ECDSA
			newPublicShareECDSA.Add(newPublicShareECDSA, oldPublicShareECDSA)
		}

		r.LocalParties[idJ].Public.ECDSA = newPublicShareECDSA
	}

	newPublicECDSA := summedPoly.Constant()
	if r.isRefresh() {
		oldPublicECDSA := curve.FromPublicKey(r.S.PublicKey())
		newPublicECDSA.Add(newPublicECDSA, oldPublicECDSA)
	}
	publicKey := newPublicECDSA.ToPublicKey()

	public := make(map[party.ID]*party.Public, r.S.N())
	for idJ, partyJ := range r.LocalParties {
		public[idJ] = partyJ.Public.Clone()
	}

	newSession, err := session.NewSession(r.S.Threshold(), public, publicKey, newSecret.Clone(), nil)
	if err != nil {
		return nil, fmt.Errorf("refresh.output.GenerateMessages(): compute SSID: %w", err)
	}
	r.S = newSession

	return nil, nil
}

// Finalize implements round.Round
func (r *output) Finalize() (round.Round, error) {
	return nil, nil
}

func (r *output) MessageType() round.MessageType {
	return MessageTypeRefresh4
}
