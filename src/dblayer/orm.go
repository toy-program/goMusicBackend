package dblayer

import (
	"goMusicBackend/src/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBORM struct {
	*gorm.DB
}

// NewORM create new ORM
func NewORM(dbname, con string) (*DBOR, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

// GetAllProducts get all products
func (db *DBORM) GetAllProducts() (products []models.Product, err Error) {
	return products, db.Find(&products).Error
}

// GetPromos get promotions
func (db *DBORM) GetPromos() (products []models.Product, err Error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

// GetCustomerByName get customer by name
func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err Error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}
