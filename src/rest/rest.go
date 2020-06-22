package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		// 클라이언트에게 상품 목록 반환
	})

	r.GET("/promos", func(c *gin.Context) {
		// 클라이언트에게 프로모션 목록 반환
	})

	r.POST("/users/signin", func(c *gin.Context) {
		// 사용자 로그인
	})

	r.POST("/users", func(c *gin.Context) {
		// 사용자 추가
	})

	r.POST("/users/:id/signout", func(c *gin.Context) {
		// 해당 ID의 사용자의 주문 내역 조회
	})

	r.GET("/users/:id/orders", func(c *gin.Context) {
		// 주문 목록 조회
	})

	r.POST("/users/charge", func(c *gin.Context) {
		// 신용카드 결제 처리
	})
}
