package gql_test

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/xplorfin/gql-bigint/gql-test/generated"
	"github.com/xplorfin/gql-bigint/gql-test/model"
)

func (r *queryResolver) BigIntQuery(ctx context.Context) (*model.BigIntObject, error) {
	return &model.BigIntObject{
		Value: r.BigInt(),
	}, nil
}

func (r *queryResolver) BigUIntQuery(ctx context.Context) (*model.BigUIntObject, error) {
	return &model.BigUIntObject{
		Value: r.BigUInt(),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
