package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Product 상품
type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImageAlt    string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` //sql.NullFloat64
	ProductName string  `gorm:"column:productname" json:"productname"`
	Description string
}

// TableName Product table name
func (Product) TableName() string {
	return "products"
}

// Customer 고객
type Customer struct {
	gorm.Model
	Name      string `json:"name"`
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email" json:"email"`
	LoggedIn  bool   `gorm:"column:loggedin" json:"loggedin"`
}

// TableName Customer table name
func (Customer) TableName() string {
	return "customers"
}

// Order 주문
type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id"`
	ProductID    int       `gorm:"column:product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

// TableName Order Tablename
func (Order) TableName() string {
	return "orders"
}
