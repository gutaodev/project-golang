package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func postHandler(c *gin.Context) {
	c.String(http.StatusOK, "bolg post handler")
}

func commentHandler(c *gin.Context) {
	c.String(http.StatusOK, "bolg comment handler")
}

func likeHandler(c *gin.Context) {
	c.String(http.StatusOK, "bolg like handler")
}
