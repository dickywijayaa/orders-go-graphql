package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/middleware"
	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Users(ctx context.Context, input *models.FilterUser, limit *int, offset *int) ([]*models.User, error) {
	return r.UserRepo.GetUser(input, limit, offset)
}

func (r *queryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	return r.OrderRepo.GetOrder()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepo.GetUserById(id)
}

func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	return r.OrderRepo.GetOrderById(id)
}

func (r *queryResolver) UserAddress(ctx context.Context) ([]*models.UserAddress, error) {
	return r.UserAddressRepo.GetUserAddress()
}

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	return r.ProductRepo.GetProducts()
}

func (r *queryResolver) Cart(ctx context.Context) (*models.Cart, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	return r.CartRepo.GetCartByBuyerId(user.ID)
}