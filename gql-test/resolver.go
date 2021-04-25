package gql_test

import (
	"math/rand"
	"time"

	bigint "github.com/xplorfin/gql-bigint"
)

//go:generate go run github.com/99designs/gqlgen

func init() {
	rand.Seed(time.Now().UnixNano())
}

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r Resolver) BigInt() bigint.BigInt {
	return bigint.BigInt(rand.Int63())
}

func (r Resolver) BigUInt() bigint.BigUInt {
	return bigint.BigUInt(rand.Uint64())
}
