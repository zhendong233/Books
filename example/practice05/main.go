package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// nolint
func main() {
	r := gin.Default()
	// hello world
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	// query使用练习
	r.GET("/user", func(c *gin.Context) {
		username := c.DefaultQuery("username", "jack")
		// age := c.Query("age")
		age, err := strconv.Atoi(c.Query("age"))
		if err != nil {
			c.JSON(http.StatusBadRequest, c.Error(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "hello",
			"username": username,
			"age":      age,
		})
	})
	// form使用练习
	r.POST("/user/form", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})
	// request body获取练习
	r.POST("/new-user", func(c *gin.Context) {
		b, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, c.Error(err))
			return
		}
		user := struct {
			User string `json:"user"`
			Age  int    `json:"age"`
		}{}
		if err := json.Unmarshal(b, &user); err != nil {
			c.JSON(http.StatusBadRequest, c.Error(err))
			return
		}
		c.JSON(http.StatusOK, user)
	})
	// form 文件上传
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("a")
		if err != nil {
			c.JSON(http.StatusInternalServerError, c.Error(err))
			return
		}
		dst := fmt.Sprintf("./example/file/%s", file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, c.Error(err))
			return
		}
		c.Status(http.StatusAccepted)
	})
	// form 多文件上传
	r.POST("/upload/multi", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, c.Error(err))
			return
		}
		files := form.File["f"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, c.Error(errors.New("no files")))
			return
		}

		for _, file := range files {
			dst := fmt.Sprintf("./example/file/%s", file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, c.Error(err))
				return
			}
		}
		c.Status(http.StatusAccepted)
	})
	// 重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})
	// 路由重定向
	r.GET("/test/router", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	// 路由组：我们可以将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别。
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/name", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"name": "nike"})
		})
		shopGroup.GET("/address", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"address": "japan"})
		})
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
