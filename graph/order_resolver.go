package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type orderResolver struct { *Resolver }

func (r *Resolver) Order() generated.OrderResolver {
	return &orderResolver{r}
}

func (o *orderResolver) Buyer(ctx context.Context, obj *models.Order) (*models.User, error) {
	return ctxLoaders(ctx).getUserByIds.Load(obj.BuyerId) // single query to fetch many rows
}

func (o *orderResolver) Details(ctx context.Context, obj *models.Order) ([]*models.OrderDetail, error) {
	return ctxLoaders(ctx).getOrderDetails.Load(obj.ID) // single query to fetch many rows
}