package party

import (
	"io"

	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
)

// ID represents a unique identifier for a participant in our scheme.
//
// You should think of this as a 32 byte slice. We represent it as a string
// to have a comparable type, but using more than 32 bytes will lead to inconsistencies
// because of how we use this ID numerically later.
//
// This ID is used as an interpolation point of a polynomial sharing of the secret key.
type ID string

// Scalar converts this ID into a scalar.
//
// All of the IDs of our participants form a polynomial sharing of the secret
// scalar value used for ECDSA.
func (id ID) Scalar() *curve.Scalar {
	out := new(curve.Scalar)
	out.SetBytes([]byte(id))
	return out
}

// WriteTo makes ID implement the io.WriterTo interface.
//
// This writes out the content of this ID, in a domain separated way.
func (id ID) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(id))
	return int64(n), err
}

// Domain implements WriterToWithDomain, and separates this type within hash.Hash.
func (ID) Domain() string {
	return "ID"
}
