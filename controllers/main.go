package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hakisolos/waitlist/models"
	"github.com/kamva/mgm/v3"
)

//var err error

func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "waitlist api running",
	})
}

func JoinController(c *gin.Context) {
	var users models.User
	err := c.ShouldBindJSON(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	err = mgm.Coll(&users).Create(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully added to waitlist",
		"data":    users,
	})
}

func GetUsersController(c *gin.Context) {
	var users []models.User

	err := mgm.Coll(&models.User{}).SimpleFind(&users, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
