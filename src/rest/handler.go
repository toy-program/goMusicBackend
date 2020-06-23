package rest

import (
	"goMusicBackend/src/dblayer"
	"goMusicBackend/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandlerInterface 핸들러 인터페이스
type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	GetOrders(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	Charge(c *gin.Context)
}

// Handler Handler
type Handler struct {
	db dblayer.DBLayer
}

// NewHandler New Handler
func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

// GetProducts get products
func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}

	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetPromos get promotions
func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		return
	}

	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promos)
}

// SignIn sign in
func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	customer, err = h.db.SignInUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, customer)
}

// AddUser Add user
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// SignOut sign out
func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		return
	}

	p := c.Param("id")

	id, err := strconv.Atoi(p) // 정수형으로 변환
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.SignOutUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetOrders get orders by id
func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}

	p := c.Param("id")

	// 문자열  => 정수형

	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Charge charge fee
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}
}
