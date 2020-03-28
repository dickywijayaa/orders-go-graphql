package graph

import (
	"context"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/model"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type mutationResolver struct { *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}


func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	//do validate

	user := models.User{
		Name: input.Name,
		Email: input.Email,
	}
	return r.UserRepo.CreateUser(&user)
}