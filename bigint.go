package gql_bigint

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// BigInt represents a 64-bit signed integer (int64)
// which is usable as a GraphQL scalar type.
type BigInt int64

// Ptr returns a pointer to b.
func (b BigInt) Ptr() *BigInt {
	return &b
}

// Int64 returns the underlying inbuilt int64 value of b.
func (b BigInt) Int64() int64 {
	return int64(b)
}

// Float64 returns the float64 representation of b
func (b BigInt) Float64() float64 {
	return float64(b.Int64())
}

// MarshalGQL implements the graphql.Marshaler interface found in
// gqlgen, allowing the type to be marshaled by gqlgen and sent over
// the wire.
func (b BigInt) MarshalGQL(w io.Writer) {
	if err := json.NewEncoder(w).Encode(b.Int64()); err != nil {
		panic(errors.Wrapf(err, "error marshaling BigInt %[1]v", b))
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in
// gqlgen, allowing the type to be received by a graphql client and unmarshaled.
func (b *BigInt) UnmarshalGQL(v interface{}) error {
	check, ok := v.(json.Number)
	if !ok {
		return errors.New("BigInt must be a valid integer value")
	}

	conv, err := check.Int64()
	if err != nil {
		return err
	}

	*b = BigInt(conv)

	return nil
}
