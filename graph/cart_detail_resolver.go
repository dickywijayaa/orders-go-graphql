package graph

import (
	"context"
	"fmt"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type cartDetailResolver struct{ *Resolver }

func (r *Resolver) CartDetail() generated.CartDetailResolver { 
	return &cartDetailResolver{r}
}

func (r *cartDetailResolver) Cart(ctx context.Context, obj *models.CartDetail) (*models.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *cartDetailResolver) Product(ctx context.Context, obj *models.CartDetail) (*models.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *cartDetailResolver) Quantity(ctx context.Context, obj *models.CartDetail) (int, error) {
	panic(fmt.Errorf("not implemented"))
}