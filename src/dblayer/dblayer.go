package dblayer

import (
	"goMusicBackend/src/models"
)

// DBLayer DB 레이어
type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(models.Customer) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
