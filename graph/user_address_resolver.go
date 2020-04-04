package graph

import (
	"context"
	"fmt"
	
	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type userAddressResolver struct { *Resolver }

func (r *Resolver) UserAddress() generated.UserAddressResolver { 
	return &userAddressResolver{r} 
}

func (r *userAddressResolver) User(ctx context.Context, obj *models.UserAddress) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userAddressResolver) Province(ctx context.Context, obj *models.UserAddress) (*models.Province, error) {
	panic(fmt.Errorf("not implemented"))
}