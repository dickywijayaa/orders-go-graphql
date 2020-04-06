package models

import (
	"time"
)

type UserAddress struct {
	Id			string		`json:"id"`
	UserId		string		`json:"user_id"`
	Address 	string 		`json:"address"`
	ProvinceId	string		`json:"province_id"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"-" pg:",soft_delete"`
}