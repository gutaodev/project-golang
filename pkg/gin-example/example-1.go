package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {
	// 1. 创建路由
	r := gin.Default()
	// 2. 绑定路由规则
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	// 3. 设置参数
	// 3.1 路径参数，参数设置在url中
	r.GET("user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		fmt.Println("name ", name, "action: ", action)
		context.String(http.StatusOK, "name: "+name+", action: ", action)
	})
	// 3.2 请求参数
	r.GET("queryparam", func(context *gin.Context) {
		name := context.DefaultQuery("name", "default user")
		fmt.Println("Query Parma name:", name)
		context.String(http.StatusOK, "param name:"+name)
	})
	// 3.3 表单参数
	r.POST("posthttp", func(context *gin.Context) {
		userType := context.DefaultPostForm("userType", "post")
		user := context.PostForm("user")
		fmt.Println("user: ", user, "type: ", userType)
		context.String(http.StatusOK, fmt.Sprintf("user: %s, type: %s", user, userType))
	})
	// 4. 服务启动
	r.Run(":8080")
}
