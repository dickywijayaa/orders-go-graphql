// generate new loader ? here is an example :
// go run github.com/vektah/dataloaden ShippingMethodLoader string *github.com/dickywijayaa/orders-go-graphql/models.ShippingMethod
// go run github.com/vektah/dataloaden OrderSliceLoader string []*github.com/dickywijayaa/orders-go-graphql/models.Order

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
	UserRepo 			repositories.UserRepository
	OrderRepo			repositories.OrderRepository
	OrderDetailRepo		repositories.OrderDetailRepository
	ProvinceRepo		repositories.ProvinceRepository
	ShippingMethodRepo	repositories.ShippingMethodRepository
}

type Loaders struct {
	getUserByIds 			*UserLoader
	getOrderByBuyerIds		*OrderSliceLoader
	getOrderByIds 			*OrderLoader
	getOrderDetails			*OrderDetailLoader
	getProvinceByIds		*ProvinceLoader
	getShippingMethodByIds	*ShippingMethodLoader
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

		loaders.getOrderByBuyerIds = &OrderSliceLoader{
			maxBatch: 100,
			wait: 2 * time.Millisecond,
			fetch: func(buyer_ids []string) ([][]*models.Order, []error) {
				// need to work on this and implement in resolver
				orders, err := dl.OrderRepo.GetOrderByBuyerIds(buyer_ids)
				if err != nil {
					return nil, err
				}

				result := make([][]*models.Order, len(buyer_ids))
				for i, id := range buyer_ids {
					for _, order := range orders {
						if order.BuyerId == id {
							result[i] = append(result[i], order)
						}
					}
				}

				return result, nil
			},
		}

		loaders.getOrderByIds = &OrderLoader{
			maxBatch: 100,
			wait: 2 * time.Millisecond,
			fetch: func(ids []string) ([]*models.Order, []error) {
				orders, err := dl.OrderRepo.GetOrderByIds(ids)
				if err != nil {
					return nil, []error{err}
				}

				o := make(map[string]*models.Order, len(orders))

				for _, order := range orders {
					o[order.ID] = order
				}

				result := make([]*models.Order, len(ids))

				for i, id := range ids {
					result[i] = o[id]
				}

				return result, nil
			},
		}

		loaders.getOrderDetails = &OrderDetailLoader{
			maxBatch: 100,
			wait:	  2 * time.Millisecond,
			fetch: 	func(order_ids []string) ([][]*models.OrderDetail, []error)	{
				order_details, err := dl.OrderDetailRepo.GetOrderDetails(order_ids)
				if err != nil {
					return nil, []error{err}
				}


				result := make([][]*models.OrderDetail, len(order_ids))
				for i, id := range order_ids {
					for _, order_detail := range order_details {
						if order_detail.OrderId == id {
							result[i] = append(result[i], order_detail)
						}
					}
				}

				return result, nil
			},
		}

		loaders.getProvinceByIds = &ProvinceLoader{
			maxBatch: 100,
			wait:	  2 * time.Millisecond,
			fetch:	func(ids []string) ([]*models.Province, []error) {
				provinces, err := dl.ProvinceRepo.GetProvinceByIds(ids)
				
				if err != nil {
					return nil, []error{err}
				}

				p := make(map[string]*models.Province, len(provinces))

				for _, province := range provinces {
					p[province.ID] = province
				}

				result := make([]*models.Province, len(ids))

				for i, id := range ids {
					result[i] = p[id]
				}

				return result, nil
			},
		}

		loaders.getShippingMethodByIds = &ShippingMethodLoader{
			maxBatch: 100,
			wait:	  2 * time.Millisecond,
			fetch:	func(ids []string) ([]*models.ShippingMethod, []error) {
				shipping_methods, err := dl.ShippingMethodRepo.GetShippingMethodByIds(ids)
				
				if err != nil {
					return nil, []error{err}
				}

				sp := make(map[string]*models.ShippingMethod, len(shipping_methods))

				for _, shipping_method := range shipping_methods {
					sp[shipping_method.ID] = shipping_method
				}

				result := make([]*models.ShippingMethod, len(ids))

				for i, id := range ids {
					result[i] = sp[id]
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), ctxLoaderKey, loaders)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ctxLoaders(ctx context.Context) Loaders {
	return ctx.Value(ctxLoaderKey).(Loaders)
}