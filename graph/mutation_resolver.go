package graph

import (
	"context"
	"errors"
	"regexp"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/model"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type mutationResolver struct { *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}


func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	if (len(input.Name) < 5) {
		return nil, errors.New("name is not long enough.")
	}

	rgx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	result := rgx.MatchString(input.Email)
	
	if !result {
		return nil, errors.New("invalid email format")
	}

	user := models.User{
		Name: input.Name,
		Email: input.Email,
	}
	return r.UserRepo.CreateUser(&user)
}