package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func goodsHandler(c *gin.Context) {
	c.String(http.StatusOK, "shop goods handler")
}

func checkoutHandler(c *gin.Context) {
	c.String(http.StatusOK, "shop checkout handler")
}
