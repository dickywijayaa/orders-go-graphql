package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type cartResolver struct{ *Resolver }

func (r *Resolver) Cart() generated.CartResolver { 
	return &cartResolver{r}
}

func (r *cartResolver) Buyer(ctx context.Context, obj *models.Cart) (*models.User, error) {
	return ctxLoaders(ctx).getUserByIds.Load(obj.BuyerId)
}

func (r *cartResolver) Details(ctx context.Context, obj *models.Cart) (*models.CartDetail, error) {
	return r.CartDetailRepo.GetCartDetailByCartId(obj.Id)
}