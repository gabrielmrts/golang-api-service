package controllers

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/gabrielmrts/golang-api-service/instances"
	"github.com/gabrielmrts/golang-api-service/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (s UserController) FindAll(c *gin.Context) {
	db := instances.GetDatabaseInstance()

	rows, err := db.Query("SELECT username, email, password FROM users")

	if err != nil {
		panic(err)
	}

	users := make([]string, 0, 100)

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Username, &user.Email, &user.Password)

		if err != nil {
			panic(err)
		}

		users = append(users, user.Username, user.Email, user.Password)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}

func (s UserController) FindOne(c *gin.Context) {
	db := instances.GetDatabaseInstance()
	var user models.User

	err := db.QueryRow("SELECT username, email, password FROM users WHERE id = $1", c.Param("id")).Scan(&user.Username, &user.Email, &user.Password)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func (s UserController) Add(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	hasher := sha512.New()
	hasher.Write([]byte(user.Password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	user.Password = hash

	user.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}

func (s UserController) Auth(c *gin.Context) {
	var user models.User

	db := instances.GetDatabaseInstance()

	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	hasher := sha512.New()
	hasher.Write([]byte(user.Password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	err := db.QueryRow(`
			SELECT email FROM users 
			WHERE username = $1 AND password = $2
		`, user.Username, encryptedPassword).Scan(&user.Email)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Incorrect credentials",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
	})
}
