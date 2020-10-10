package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func setupRouter() *gin.Engine  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	//r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
	//r.getCurrentPath
	// gin.H 是map[string]interface{}的缩写
	//r.GET("/someJSON", func(c *gin.Context) {
	//	// 方式一：自己拼接JSON
	//	c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	//})
	//r.GET("/moreJSON", func(c *gin.Context) {
	//	// 方法二：使用结构体
	//	var msg struct {
	//		Name    string `json:"user"`
	//		Message string
	//		Age     int
	//	}
	//	msg.Name = "小王子"
	//	msg.Message = "Hello world!"
	//	msg.Age = 18
	//	c.JSON(http.StatusOK, msg)
	//})
	//r.GET("/someYAML", func(c *gin.Context) {
	//	c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	//})
	//r.GET("/someProtoBuf", func(c *gin.Context) {
	//	reps := []int64{int64(1), int64(2)}
	//	label := "test"
	//	// protobuf 的具体定义写在 testdata/protoexample 文件中。
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  reps,
	//	}
	//	// 请注意，数据在响应中变为二进制数据
	//	// 将输出被 protoexample.Test protobuf 序列化了的数据
	//	c.ProtoBuf(http.StatusOK, data)
	//})
	//r.GET("/user/search", func(c *gin.Context) {
	//	username := c.DefaultQuery("username", "小王子")
	//	//username := c.Query("username")
	//	address := c.Query("address")
	//	//输出json结果给调用方
	//	c.JSON(http.StatusOK, gin.H{
	//		"message":  "ok",
	//		"username": username,
	//		"address":  address,
	//	})
	//})
	//r.POST("/user/search", func(c *gin.Context) {
	//	// DefaultPostForm取不到值时会返回指定的默认值
	//	//username := c.DefaultPostForm("username", "小王子")
	//	username := c.PostForm("username")
	//	address := c.PostForm("address")
	//	//输出json结果给调用方
	//	c.JSON(http.StatusOK, gin.H{
	//		"message":  "ok",
	//		"username": username,
	//		"address":  address,
	//	})
	//})
	//r.GET("/user/search/:username/:address", func(c *gin.Context) {
	//	username := c.Param("username")
	//	address := c.Param("address")
	//	//输出json结果给调用方
	//	c.JSON(http.StatusOK, gin.H{
	//		"message":  "ok",
	//		"username": username,
	//		"address":  address,
	//	})
	//})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}
