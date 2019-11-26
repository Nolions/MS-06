package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type User struct {
	Uid string `json:"uid"`
	InputUser
}

type InputUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func newUser() *User {
	return &User{Uid: uuid.Must(uuid.NewV4()).String()}
}

func main() {
	seed()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("GAE_PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/health", healthHandler)
	r.GET("/user/:uid", findHandler)
	r.POST("/user", creatorHandler)
	r.PUT("/user/:uid", editHandler)
	r.DELETE("/user/:uid", removeHandler)
	r.GET("/users", allHandler)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World")
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}

func creatorHandler(c *gin.Context) {
	ru := new(InputUser)
	err := c.BindJSON(&ru)
	if err != nil {
		log.Fatal("error:", err)
	}

	u := newUser()
	u.Name = ru.Name
	u.Email = ru.Email
	users = append(users, *u)
	c.JSON(http.StatusNoContent, gin.H{})
}

func editHandler(c *gin.Context) {
	i := findUserIndex(getUserID(c))
	u := &users[i]

	ru := new(InputUser)
	err := c.BindJSON(&ru)
	if err != nil {
		log.Fatal("error:", err)
	}

	u.Email = ru.Email
	u.Name = ru.Name
}

func removeHandler(c *gin.Context) {
	i := findUserIndex(getUserID(c))

	if i != -1 {
		users = append(users[:i], users[i+1:]...)
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func allHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func findHandler(c *gin.Context) {
	i := findUserIndex(getUserID(c))

	c.JSON(http.StatusOK, users[i])
}

func seed() {
	users = append(users, User{
		Uid: uuid.Must(uuid.NewV4()).String(),
		InputUser: InputUser{
		Name:  "John",
		Email: "John@gmail.com",
	},
	})
}

func getUserID(c *gin.Context) string {
	return c.Param("uid")
}

func findUserIndex(uid string) int {
	for i := 0; i < len(users); i++ {
		if uid == users[i].Uid {
			return i
		}
	}
	return -1
}
