package gql_bigint

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// BigUInt represents a 64-bit unsigned integer (uint64)
// which is usable as a GraphQL scalar type.
type BigUInt uint64

// Ptr returns a pointer to b.
func (b BigUInt) Ptr() *BigUInt {
	return &b
}

// UInt64 returns the underlying inbuilt uint64 value of b.
func (b BigUInt) UInt64() uint64 {
	return uint64(b)
}

// MarshalGQL implements the graphql.Marshaler interface found in
// gqlgen, allowing the type to be marshaled by gqlgen and sent over
// the wire.
func (b BigUInt) MarshalGQL(w io.Writer) {
	if err := json.NewEncoder(w).Encode(b.UInt64()); err != nil {
		panic(errors.Wrapf(err, "error marshaling BigUInt %[1]v", b))
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in
// gqlgen, allowing the type to be received by a graphql client and unmarshaled.
func (b *BigUInt) UnmarshalGQL(v interface{}) error {
	check, ok := v.(json.Number)
	if !ok {
		return errors.New("BigUInt must be a valid integer value")
	}

	conv, err := check.Int64()
	if err != nil {
		return err
	}

	*b = BigUInt(uint64(conv))

	return nil
}
