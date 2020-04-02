package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
	"github.com/dickywijayaa/orders-go-graphql/graph/model"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Users(ctx context.Context, input *model.FilterUser, limit *int, offset *int) ([]*models.User, error) {
	return r.UserRepo.GetUser(input, limit, offset)
}

func (r *queryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	return r.OrderRepo.GetOrder()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepo.GetUserById(id)
}

func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	return r.OrderRepo.GetOrderById(id)
}