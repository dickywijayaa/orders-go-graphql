package graph

import (
	"context"
	"fmt"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type productResolver struct{ *Resolver }

func (r *Resolver) Product() generated.ProductResolver { 
	return &productResolver{r} 
}

func (r *productResolver) Seller(ctx context.Context, obj *models.Product) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}