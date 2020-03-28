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
	return o.UserRepo.GetUserById(obj.BuyerId)
}

func (o *orderResolver) TotalPrice(ctx context.Context, obj *models.Order) (float64, error) {
	return o.OrderRepo.GetOrderTotalPrice(obj.ID)
}

func (o *orderResolver) Details(ctx context.Context, obj *models.Order) ([]*models.OrderDetail, error) {
	return o.OrderDetailRepo.GetOrderDetailsByOrderId(obj.ID)
}