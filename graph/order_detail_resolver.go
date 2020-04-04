package graph

import (
	"context"
	"fmt"

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

func (r *orderDetailResolver) ShippingCost(ctx context.Context, obj *models.OrderDetail) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *orderDetailResolver) Seller(ctx context.Context, obj *models.OrderDetail) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *orderDetailResolver) ShippingMethod(ctx context.Context, obj *models.OrderDetail) (*models.ShippingMethod, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *orderDetailResolver) ItemPrice(ctx context.Context, obj *models.OrderDetail) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *orderDetailResolver) ItemWeight(ctx context.Context, obj *models.OrderDetail) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}