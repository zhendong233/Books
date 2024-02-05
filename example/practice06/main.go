package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func router1() http.Handler {
	// gin.Default()默认使用了Logger和Recovery中间件，其中：
	// Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	// Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	// 如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "use server1"})
	})
	return e
}

func router2() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "use server2"})
	})
	return e
}

func main() {
	// 运行多个服务
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server2 := &http.Server{
		Addr:         ":8081",
		Handler:      router2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	var eg errgroup.Group
	eg.Go(func() error {
		return server1.ListenAndServe()
	})
	eg.Go(func() error {
		return server2.ListenAndServe()
	})
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
