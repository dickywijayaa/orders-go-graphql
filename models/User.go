package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 			string		`json:"id"`
	Name		string 		`json:"name"`
	Email 		string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"-" pg:",softdelete"`
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) GenerateToken() (*AuthToken, error) {
	response := &AuthToken{
		Token: "random",
		ExpiredAt: time.Now(),
	}
	return response, nil
}