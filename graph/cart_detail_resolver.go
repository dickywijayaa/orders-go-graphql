package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type cartDetailResolver struct{ *Resolver }

func (r *Resolver) CartDetail() generated.CartDetailResolver { 
	return &cartDetailResolver{r}
}

func (r *cartDetailResolver) Cart(ctx context.Context, obj *models.CartDetail) (*models.Cart, error) {
	return r.CartRepo.GetCartById(obj.CartId)
}

func (r *cartDetailResolver) Product(ctx context.Context, obj *models.CartDetail) (*models.Product, error) {
	return r.ProductRepo.GetProductById(obj.ProductId)
}