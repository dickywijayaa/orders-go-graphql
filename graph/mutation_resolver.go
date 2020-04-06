package graph

import (
	"fmt"
	"log"
	"context"
	"errors"
	"regexp"

	"github.com/dickywijayaa/orders-go-graphql/middleware"
	"github.com/dickywijayaa/orders-go-graphql/models"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
)

const VALIDATION_NAME_LENGTH_ERROR_MESSAGE = "name is not long enough."
const VALIDATION_NAME_EXISTS_ERROR_MESSAGE = "name is already exists."
const VALIDATION_PASSWORD_LENGTH_ERROR_MESSAGE = "password is not long enough."
const VALIDATION_EMAIL_FORMAT_ERROR_MESSAGE = "invalid email format."
const VALIDATION_EMAIL_EXISTS_ERROR_MESSAGE = "email is already exists."
const VALIDATION_USER_NOT_EXISTS_ERROR_MESSAGE = "user not exists."
const PAYLOAD_UPDATE_EMPTY_ERROR_MESSAGE = "nothing value to be updated."
const VALIDATION_NOT_EXISTS_PRODUCT_ERROR_MESSAGE = "product is not exists."
const VALIDATION_NULL_QUANTITY_ERROR_MESSAGE = "quantity must be more than 0."

type mutationResolver struct { *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.AuthResponse, error) {
	if (len(input.Name) < 5) {
		return nil, errors.New(VALIDATION_NAME_LENGTH_ERROR_MESSAGE)
	}

	result := checkEmailRegex(input.Email)
	
	if !result {
		return nil, errors.New(VALIDATION_EMAIL_FORMAT_ERROR_MESSAGE)
	}

	check_email_exists, err := r.UserRepo.GetUserByEmail(input.Email)
	if check_email_exists != nil && err == nil {
		return nil, errors.New(VALIDATION_EMAIL_EXISTS_ERROR_MESSAGE)
	}

	if (len(input.Password) < 5) {
		return nil, errors.New(VALIDATION_PASSWORD_LENGTH_ERROR_MESSAGE)
	}

	user_data := models.User{
		Name: input.Name,
		Email: input.Email,
		Role: "default", // need to be updated
	}

	err = user_data.HashPassword(input.Password)
	if err != nil {
		log.Printf("error when hash password : %v", err)
		return nil, errors.New("something went wrong")
	}
	
	tx, err := r.UserRepo.DB.Begin()
	if err != nil {
		log.Printf("error when begin transaction : %v", err)
		return nil, errors.New("something went wrong")
	}

	defer tx.Rollback()
	
	user, err := r.UserRepo.CreateUser(&user_data)
	if err != nil {
		log.Printf("error when insert user : %v", err)
		return nil, errors.New("something went wrong")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error when commit : %v", err)
		return nil, errors.New("something went wrong")
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.Printf("error when generate token : %v", err)
		return nil, errors.New("something went wrong")
	}

	response := &models.AuthResponse{
		Auth: token,
		User: user,
	}

	return response, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	return r.UserRepo.DeleteUser(id)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UpdateUser) (*models.User, error) {
	if input.Name == nil && input.Email == nil {
		// make sure there is something to update
		return nil, errors.New(PAYLOAD_UPDATE_EMPTY_ERROR_MESSAGE)
	}

	user, err := r.UserRepo.GetUserById(id)
	if err != nil || user == nil {
		return nil, errors.New(VALIDATION_USER_NOT_EXISTS_ERROR_MESSAGE)
	}

	if input.Name != nil {
		if (len(*input.Name) < 5) {
			return nil, errors.New(VALIDATION_NAME_LENGTH_ERROR_MESSAGE)
		}

		user.Name = *input.Name
	}

	if input.Email != nil {
		check_email := checkEmailRegex(*input.Email)
		if !check_email {
			return nil, errors.New(VALIDATION_EMAIL_FORMAT_ERROR_MESSAGE)
		}

		user.Email = *input.Email
	}

	user, err = r.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginUserInput) (*models.AuthToken, error) {
	user, err := r.UserRepo.Login(input)
	if err != nil {
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func checkEmailRegex(email string) bool {
	rgx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	result := rgx.MatchString(email)

	if !result {
		return false
	}
	return true
}

func (r *mutationResolver) AddToCart(ctx context.Context, input *models.AddCartInput) (*models.Cart, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if input.ProductID == "" {
		return nil, errors.New(VALIDATION_NOT_EXISTS_PRODUCT_ERROR_MESSAGE)
	}

	if input.Quantity < 1 {
		return nil, errors.New(VALIDATION_NULL_QUANTITY_ERROR_MESSAGE)
	}

	_, err = r.ProductRepo.GetProductById(input.ProductID)
	if err != nil {
		return nil, errors.New("product not exists")
	}

	tx, err := r.CartRepo.DB.Begin()
	if err != nil {
		log.Printf("error when begin transaction : %v", err)
		return nil, errors.New("something went wrong")
	}

	defer tx.Rollback()
	
	result, err := r.CartRepo.AddToCart(tx, user.ID, input)
	if err != nil {
		log.Printf("error when insert user : %v", err)
		return nil, errors.New("something went wrong")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error when commit : %v", err)
		return nil, errors.New("something went wrong")
	}
	
	return result, err
}

func (r *mutationResolver) RemoveFromCart(ctx context.Context, productID string) (*models.Cart, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, err
	}

	if productID == "" {
		return nil, errors.New(VALIDATION_NOT_EXISTS_PRODUCT_ERROR_MESSAGE)
	}

	return r.CartRepo.DeleteCart(user.ID, productID)
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input models.CreateOrderInput) (*models.Order, error) {
	panic(fmt.Errorf("not implemented"))
}