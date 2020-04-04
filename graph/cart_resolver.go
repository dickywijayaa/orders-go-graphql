package graph

import (
	"context"
	"fmt"

	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

type cartResolver struct{ *Resolver }

func (r *Resolver) Cart() generated.CartResolver { 
	return &cartResolver{r}
}

func (r *cartResolver) Buyer(ctx context.Context, obj *models.Cart) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}