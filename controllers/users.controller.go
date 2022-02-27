package controllers

import (
	"net/http"

	"github.com/gabrielmrts/golang-api-service/factories"
	"github.com/gabrielmrts/golang-api-service/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (s UserController) FindAll(c *gin.Context) {
	db := factories.GetDatabaseInstance()

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
	db := factories.GetDatabaseInstance()
	var user models.User

	err := db.QueryRow("SELECT username, email, password FROM users WHERE id = $1", c.Param("id")).Scan(&user.Username, &user.Email, &user.Password)

	if err != nil {
		panic(err)
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

	user.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
