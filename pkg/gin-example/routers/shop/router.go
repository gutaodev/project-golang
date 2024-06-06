package shop

import "github.com/gin-gonic/gin"

func Routers(c *gin.Engine) {
	c.GET("/goods", goodsHandler)
	c.GET("/checkout", checkoutHandler)
}
