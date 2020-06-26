package dblayer

import (
	"goMusicBackend/src/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBORM database orm interface
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

func (db *DBORM) GetCustomerByID(id int) (product models.Product, error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) SignInUser (email, pass string) (customer models.Customer, error) {
	if !checkPassword(pass) {
		return customer, errors.New("Invalid password")
	}

	result := db.Table("Customers").Where(&models.Customer{Email: email})
	err = result.Update("loggedIn", true).Error


	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer.Error)
}