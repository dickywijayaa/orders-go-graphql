package graph

import (
	"context"
	"fmt"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type orderResolver struct { *Resolver }

func (r *Resolver) Order() generated.OrderResolver {
	return &orderResolver{r}
}

func (o *orderResolver) Buyer(ctx context.Context, obj *models.Order) (*models.User, error) {
	return ctxLoaders(ctx).getUserByIds.Load(obj.BuyerId)
}

func (o *orderResolver) Details(ctx context.Context, obj *models.Order) ([]*models.OrderDetail, error) {
	return ctxLoaders(ctx).getOrderDetails.Load(obj.ID)
}

func (r *orderResolver) TotalShippingCost(ctx context.Context, obj *models.Order) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}