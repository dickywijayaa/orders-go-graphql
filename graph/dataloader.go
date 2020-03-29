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
	UserRepo repositories.UserRepository
}

const userLoaderKey = "userloader"

func (dl *Dataloader) DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 10,
			wait:	  1 * time.Millisecond,
			fetch:	func(ids []string) ([]*models.User, []error) {
				users, err := dl.UserRepo.GetUserByIds(ids)
				
				if err != nil {
					return nil, err
				}

				u := make(map[string]*models.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*models.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}