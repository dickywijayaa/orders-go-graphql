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
	// to do remove to loadermiddleware
	return od.OrderRepo.GetOrderById(obj.OrderId)
}