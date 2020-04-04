package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
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

const JWT_SECRET = "jwtsecret"

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
	expired_at := time.Now().Add(time.Hour * 24 * 7)
	claims := &jwt.StandardClaims{
		ExpiresAt: expired_at.Unix(),
		Issuer: "orders_go_graphql",
		Id: u.ID,
		IssuedAt: time.Now().Unix(),
	}

	secret := []byte(JWT_SECRET)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	response := &AuthToken{
		Token: fmt.Sprintf("%v", result),
		ExpiredAt: expired_at,
	}

	return response, nil
}