package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql/graph/generated"
	"go-graphql/graph/model"
	"go-graphql/pkg/mutations"
	"go-graphql/pkg/queries"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	mutation := &mutations.OrderMutation{DB: r.DB}
	return mutation.CreateOrder(ctx, input)
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input model.OrderInput) (*model.Order, error) {
	mutation := &mutations.OrderMutation{DB: r.DB}
	return mutation.UpdateOrder(ctx, orderID, input)
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
	mutation := &mutations.OrderMutation{DB: r.DB}
	return mutation.DeleteOrder(ctx, orderID)
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	query := &queries.OrderQuery{DB: r.DB}
	return query.Orders(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
