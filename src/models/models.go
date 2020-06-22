package models

import "time"

// Product 상품
type Product struct {
	Image       string  `json: "img"`
	ImageAlt    string  `json: "imgalt"`
	Price       float64 `json: "price"`
	Promotion   float64 `json: "promotion"`
	ProductName string  `json: "productname`
	Description string  `json: "desc"`
}

// Customer 고객
type Customer struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
	Email     string `json: "email"`
	LoggedIn  bool   `json: "loggedin"`
}

// Order 주문
type Order struct {
	Product
	Customer
	CustomerID   int       `json: "customer_id"`
	ProductID    int       `json: "product_id`
	Price        float64   `json: "sell_price"`
	PurchaseDate time.Time `json: "purchase_date"`
}