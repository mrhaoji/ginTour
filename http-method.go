package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	r := gin.Default()

	r.GET("/someGet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "GET",
		})
	})
	r.POST("/somePost", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "POST",
		})
	})
	r.PUT("/somePut", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "PUT",
		})
	})
	r.DELETE("/someDelete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "DELETE",
		})
	})
	r.PATCH("/somePatch", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "PATCH",
		})
	})
	r.HEAD("/someHead", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "HEAD",
		})
	})
	r.OPTIONS("/someOptions", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "OPTIONS",
		})
	})

	r.Run()
}
