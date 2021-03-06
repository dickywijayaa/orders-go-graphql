package repositories

import (
	"fmt"
	"errors"

	"github.com/dickywijayaa/orders-go-graphql/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-pg/pg/v9"
)

type UserRepository struct {
	DB *pg.DB
}

func (u *UserRepository) GetUser(input *models.FilterUser, limit *int, offset *int) ([]*models.User, error) {
	var users []*models.User
	query := u.DB.Model(&users)
	
	if input != nil {
		if input.Name != nil && *input.Name != "" {
			query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *input.Name))
		}

		if input.Email != nil && *input.Email != "" {
			query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", *input.Email))
		}
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Limit(*offset)
	}

	err := query.Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("email = ?", email).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("name = ?", name).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByIds(ids []string) ([]*models.User, []error) {
	var users []*models.User
	err := u.DB.Model(&users).Where("id in (?)", pg.In(ids)).Select()
	
	if err != nil {
		return nil, []error{err}
	}

	return users, nil
}

func (u *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserRepository) DeleteUser(id string) (string, error) {
	user := &models.User{
		ID: id,
	}
	err := u.DB.Delete(user)

	if err != nil {
		return "", err
	}

	return id, err
}

func (u *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	_, err := u.DB.Model(user).Where("id = ?", user.ID).Update()
	return user, err
}

func (u *UserRepository) Login(input models.LoginUserInput) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("email = ?", input.Email).First()
	if err != nil {
		return nil, errors.New("user not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}