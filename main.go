package main

import (
	"github.com/gin-gonic/gin" // Gin框架
	"github.com/CYZH1307/tiktok/config"
	"github.com/CYZH1307/tiktok/controller" // controller
)

func main() {
	router := gin.Default() // 创建一个默认的 Gin 路由器对象
	initRouter(router) // 在router.go中定义
	_ = router.Run(config.Port)
}