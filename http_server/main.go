package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

var users []User

type account struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", ping)

	r.POST("/login", login)

	r.GET("/users", all)
	r.GET("/user/:id", get)
	r.POST("/user", creator)
	r.PUT("/user/:id", edit)
	r.DELETE("/user/:id", delete)

	r.Run(":7777")
}

func login(c *gin.Context) {
	fmt.Println(c.Request.Header.Get("Content-Type"))
	fmt.Println(c.Request.Header.Get("token"))

	a := new(account)
	c.BindJSON(&a)

	c.JSON(http.StatusOK, gin.H{
		"id":    1,
		"token": "112222",
	})
}

func creator(c *gin.Context) {
	u := new(User)
	c.BindJSON(&u)

	users = append(users, *u)
}

func all(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func get(c *gin.Context) {
	id := getUserID(c)
	fmt.Println(id)
}

func edit(c *gin.Context) {
	id := getUserID(c)
	fmt.Println(id)
}

func delete(c *gin.Context) {
	id := getUserID(c)
	fmt.Println(id)
}

// heath check
func ping(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}

func getUserID(c *gin.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))

	return id
}
