package rest

import (
	"github.com/gin-gonic/gin"
)

// RunAPIWithHandler run api server with handler
func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	r.GET("/products", h.GetProducts)

	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/users")
	{
		userGroup.POST("", h.AddUser)
		userGroup.POST("/charge", h.Charge)
		userGroup.POST("/signin", h.SignIn)
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)

	}

	return r.Run(address)
}

// RunAPI run server
func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, h)
}
