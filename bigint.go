package gql_bigint

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// BigInt represents a 64-bit signed integer (int64)
// which is usable as a GraphQL scalar type.
type BigInt int64

func (b BigInt) Ptr() *BigInt {
	return &b
}

func (b BigInt) Int64() int64 {
	return int64(b)
}

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

func (b BigInt) MarshalGQL(w io.Writer) {
	if err := json.NewEncoder(w).Encode(b.Int64()); err != nil {
		panic(errors.Wrapf(err, "error marshaling BigInt %[1]v", b))
	}
}
