package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

// 原生 SQL ORM 层
var db *sql.DB

type Person struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

// 返回全部信息
func (p Person) getAll() (persons []Person, err error) {
	rows, err := db.Query("SELECT id, first_name, last_name FROM person")
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		// 遍历所有行
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		// 将 person 添加到 persons
		persons = append(persons, person)
	}

	defer rows.Close()
	return
}

// 通过 id 查询
func (p Person) get() (person Person, err error) {
	row := db.QueryRow("SELECT id, first_name, last_name FROM person WHERE id = ?", p.Id)
	err = row.Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		return
	}
	return
}

// POST 新增数据
func (p Person) add() (Id int, err error) {
	stmt, err := db.Prepare("INSERT INTO person(first_name, last_name) VALUES (?, ?)")
	if err != nil {
		return
	}
	// 执行插入操作
	rs, err := stmt.Exec(p.FirstName, p.LastName)
	if err != nil {
		return
	}
	// 返回插入的 id
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	// 将 id 类型转换
	Id = int(id)
	defer stmt.Close()
	return
}

// DELETE 通过 id 删除
func (p Person) del() (rows int, err error) {
	stmt, err := db.Prepare("DELETE FROM person WHERE id = ?")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	// 删除的行数
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
	return
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()
	
	r.GET("/person", func(c *gin.Context) {
		p := Person{}
		persons, err := p.getAll()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count": len(persons),
		})
	})

	r.GET("/person/:id", func(c *gin.Context) {
		var result gin.H
		id := c.Param("id")
		Id, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		p := Person{
			Id: Id,
		}
		person, err := p.get()
		fmt.Println(err)
		if err != nil {
			result = gin.H{
				"result": nil,
				"count": 0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count": 1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	r.POST("/person", func(c *gin.Context) {
		var p Person
		err := c.Bind(&p)
		if err != nil {
			log.Fatal(err)
		}
		Id, err := p.add()
		fmt.Print("id=", Id)
		name := p.FirstName + " " + p.LastName
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s 插入成功", name),
		})
	})
	
	r.DELETE("/person/:id", func(c *gin.Context) {
		id := c.Param("id")

		Id, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			log.Fatal(err)
		}
		p := Person{Id: int(Id)}
		rows, err := p.del()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("delete rows ", rows)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})

	r.Run(":8080")
}
