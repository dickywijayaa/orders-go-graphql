package graph

import (
	"context"
	
	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type userAddressResolver struct { *Resolver }

func (r *Resolver) UserAddress() generated.UserAddressResolver { 
	return &userAddressResolver{r} 
}

func (r *userAddressResolver) User(ctx context.Context, obj *models.UserAddress) (*models.User, error) {
	return ctxLoaders(ctx).getUserByIds.Load(obj.UserId)
}

func (r *userAddressResolver) Province(ctx context.Context, obj *models.UserAddress) (*models.Province, error) {
	return ctxLoaders(ctx).getProvinceByIds.Load(obj.ProvinceId)
}