package gql_bigint

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// BigUInt represents a 64-bit unsigned integer (uint64)
// which is usable as a GraphQL scalar type.
type BigUInt uint64

func (b BigUInt) Ptr() *BigUInt {
	return &b
}

func (b BigUInt) UInt64() uint64 {
	return uint64(b)
}

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

func (b BigUInt) MarshalGQL(w io.Writer) {
	if err := json.NewEncoder(w).Encode(b.UInt64()); err != nil {
		panic(errors.Wrapf(err, "error marshaling BigUInt %[1]v", b))
	}
}
