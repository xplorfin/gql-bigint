package gql_test_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"

	gqltest "github.com/xplorfin/gql-bigint/gql-test"

	bigint "github.com/xplorfin/gql-bigint"
	"github.com/xplorfin/gql-bigint/gql-test/generated"
)

func TestBigInt(t *testing.T) {
	var resp struct {
		BigIntQuery struct {
			Value bigint.BigInt
		}
	}

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &gqltest.Resolver{}})
	srv := handler.NewDefaultServer(schema)
	c := client.New(srv)

	c.MustPost(`query { 
		bigIntQuery { 
			value
		} 
	}
	`, &resp)

	assert.NotZero(t, resp.BigIntQuery.Value)
	assert.IsType(t, bigint.BigInt(1), resp.BigIntQuery.Value)
}

func TestBigUInt(t *testing.T) {
	var resp struct {
		BigUIntQuery struct {
			Value bigint.BigUInt
		}
	}

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &gqltest.Resolver{}})
	srv := handler.NewDefaultServer(schema)
	c := client.New(srv)

	c.MustPost(`query { 
		bigUIntQuery { 
			value
		} 
	}
	`, &resp)

	assert.NotZero(t, resp.BigUIntQuery.Value)
	assert.IsType(t, bigint.BigUInt(1), resp.BigUIntQuery.Value)
}
