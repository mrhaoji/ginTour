package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", startPage)
	r.Run(":8085")
}

func startPage(c *gin.Context)  {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")
}