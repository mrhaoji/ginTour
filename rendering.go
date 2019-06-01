package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "hey JSON",
		})
	})
	
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "hey XML",
		})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "hey YAML",
		})
	})
	
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"

		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
