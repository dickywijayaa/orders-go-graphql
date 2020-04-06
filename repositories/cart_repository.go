package repositories

import (
	"log"
	"errors"

	"github.com/dickywijayaa/orders-go-graphql/models"

	"github.com/go-pg/pg/v9"
)

type CartRepository struct {
	DB *pg.DB
}

func (c *CartRepository) GetCartById(id string) (*models.Cart, error) {
	var cart models.Cart
	err := c.DB.Model(&cart).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *CartRepository) GetCartByBuyerId(buyer_id string) (*models.Cart, error) {
	var cart models.Cart
	err := c.DB.Model(&cart).Where("buyer_id = ?", buyer_id).First()
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *CartRepository) AddToCart(tx *pg.Tx, buyer_id string, input *models.AddCartInput) (*models.Cart, error) {
	// check existing cart
	var cart models.Cart
	count_cart, err := tx.Model(&cart).Where("buyer_id = ?", buyer_id).SelectAndCount()
	if err != nil && count_cart != 0 {
		log.Printf("failed when check cart : %v", err)
		return nil, errors.New("something went wrong")
	}

	cart_id := ""

	if count_cart == 0 {
		// create new cart
		new_cart := &models.Cart{
			BuyerId: buyer_id,
		}

		_, err := tx.Model(new_cart).Returning("*").Insert()
		if err != nil {
			log.Printf("failed when insert new cart : %v", err)
			return nil, errors.New("something went wrong")
		}
		cart_id = new_cart.Id
		cart = *new_cart
	} else {
		cart_id = cart.Id
	}

	// check existing cart details, if exists just update the quantity
	var cart_detail models.CartDetail

	check_cart_detail, err := tx.Model(&cart_detail).Where("cart_id = ?", cart_id).Where("product_id = ?", input.ProductID).SelectAndCount()
	if err != nil && check_cart_detail != 0 {
		log.Printf("failed when check cart detail : %v", err)
		return nil, errors.New("something went wrong")
	}
	
	if check_cart_detail != 0 {
		cart_detail.Quantity += input.Quantity
		_, err := tx.Model(&cart_detail).Where("id = ?", cart_detail.Id).Update()
		if err != nil {
			log.Printf("failed when update cart detail : %v", err)
			return nil, errors.New("something went wrong")
		}
	} else {
		new_cart_detail := &models.CartDetail{
			CartId: cart_id,
			ProductId: input.ProductID,
			Quantity: input.Quantity,
		}
	
		// query insert
		_, err = tx.Model(new_cart_detail).Insert()
		if err != nil {
			log.Printf("failed when insert cart detail : %v", err)
			return nil, errors.New("something went wrong")
		}
	}

	return &cart, nil
}