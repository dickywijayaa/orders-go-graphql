package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type orderDetailResolver struct { *Resolver }

func (r *Resolver) OrderDetail() generated.OrderDetailResolver {
	return &orderDetailResolver{r}
}

func (od *orderDetailResolver) Order(ctx context.Context, obj *models.OrderDetail) (*models.Order, error) {
	return ctxLoaders(ctx).getOrderByIds.Load(obj.OrderId)
}