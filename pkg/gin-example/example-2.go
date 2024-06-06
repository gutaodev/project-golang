package main

import (
	"fmt"
	"gutaooo.dev/project-golang/pkg/gin-example/routers"
	"gutaooo.dev/project-golang/pkg/gin-example/routers/blog"
	"gutaooo.dev/project-golang/pkg/gin-example/routers/shop"
)

func main() {

	routers.Include(shop.Routers, blog.Routers)

	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("Server setup failed, ", err)
	}
}
