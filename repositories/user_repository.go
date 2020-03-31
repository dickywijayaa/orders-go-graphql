package repositories

import (
	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type UserRepository struct {
	DB *pg.DB
}

func (u *UserRepository) GetUser() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Model(&users).Select()
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