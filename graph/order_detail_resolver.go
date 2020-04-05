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

func (r *orderDetailResolver) Seller(ctx context.Context, obj *models.OrderDetail) (*models.User, error) {
	return ctxLoaders(ctx).getUserByIds.Load(obj.SellerId)
}

func (r *orderDetailResolver) ShippingMethod(ctx context.Context, obj *models.OrderDetail) (*models.ShippingMethod, error) {
	return ctxLoaders(ctx).getShippingMethodByIds.Load(obj.ShippingMethodId)
}