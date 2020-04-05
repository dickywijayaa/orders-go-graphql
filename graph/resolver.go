//go run github.com/99designs/gqlgen

package graph

import (
	"github.com/dickywijayaa/orders-go-graphql/repositories"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserRepo			repositories.UserRepository
	OrderRepo			repositories.OrderRepository
	OrderDetailRepo		repositories.OrderDetailRepository
	UserAddressRepo		repositories.UserAddressRepository
	ProvinceRepo		repositories.ProvinceRepository
	ProductRepo 		repositories.ProductRepository
	ShippingMethodRepo	repositories.ShippingMethodRepository
	CartRepo			repositories.CartRepository
	CartDetailRepo		repositories.CartDetailRepository
}