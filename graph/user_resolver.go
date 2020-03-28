package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type userResolver struct { *Resolver }

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

func (u *userResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	return u.OrderRepo.GetBuyerOrders(obj.ID)
}