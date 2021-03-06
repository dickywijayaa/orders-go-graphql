// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type AddCartInput struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type AuthResponse struct {
	Auth *AuthToken `json:"auth"`
	User *User      `json:"user"`
}

type AuthToken struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type CreateOrderInput struct {
	ShippingCost     float64 `json:"shipping_cost"`
	ShippingMethodID string  `json:"shipping_method_id"`
}

type FilterUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
