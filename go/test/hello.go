package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个默认的路由引擎
    router := gin.Default()

    // 定义一个路由处理函数
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Gin!")
    })

    // 启动HTTP服务，默认监听在:8080端口
    router.Run()
}