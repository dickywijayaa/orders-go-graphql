package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type productResolver struct{ *Resolver }

func (r *Resolver) Product() generated.ProductResolver { 
	return &productResolver{r} 
}

func (r *productResolver) Seller(ctx context.Context, obj *models.Product) (*models.User, error) {
	return ctxLoaders(ctx).getUserByIds.Load(obj.SellerId)
}