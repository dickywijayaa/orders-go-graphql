package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/repositories"
)

type Dataloader struct {
	UserRepo 	repositories.UserRepository
	OrderRepo	repositories.OrderRepository
}

type Loaders struct {
	getUserByIds 		*UserLoader
	getOrderByBuyerIds	*OrderLoader
}

const ctxLoaderKey = "ctxloader"

func (dl *Dataloader) LoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := Loaders{}

		loaders.getUserByIds = &UserLoader{
			maxBatch: 100,
			wait:	  2 * time.Millisecond,
			fetch:	func(user_ids []string) ([]*models.User, []error) {
				users, err := dl.UserRepo.GetUserByIds(user_ids)
				
				if err != nil {
					return nil, err
				}

				u := make(map[string]*models.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*models.User, len(user_ids))

				for i, id := range user_ids {
					result[i] = u[id]
				}

				return result, nil
			},
		}

		loaders.getOrderByBuyerIds = &OrderLoader{
			maxBatch: 100,
			wait:	  2 * time.Millisecond,
			fetch: 	func(buyer_ids []string) ([]*models.Order, []error)	{
				// need to work on this and implement in resolver
				return dl.OrderRepo.GetOrderByBuyerIds(buyer_ids)
			},
		}

		ctx := context.WithValue(r.Context(), ctxLoaderKey, loaders)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ctxLoaders(ctx context.Context) Loaders {
	return ctx.Value(ctxLoaderKey).(Loaders)
}