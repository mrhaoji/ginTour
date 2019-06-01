package main

import (
	"fmt"
	//"html/template"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	r := gin.Default()
	fmt.Println("===env===", os.Args)


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.LoadHTMLFiles("templates/index.html")
	r.GET("index", func(c *gin.Context) {
		fmt.Println("test")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	//r.LoadHTMLGlob("templates/**/*")
	//r.GET("posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.html", gin.H{
	//		"title": "Posts",
	//	})
	//})
	//r.GET("users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.html", gin.H{
	//		"title": "User",
	//	})
	//})

	//r.Delims("{[{", "}]}")
	//r.SetFuncMap(template.FuncMap{
	//	"formatAsDate": formatAsDate,
	//})
	//r.LoadHTMLFiles("templates/raw.tmpl")
	//
	//r.GET("/raw", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
	//		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	//	})
	//})

	r.Run(":8080")
}