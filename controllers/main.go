package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hakisolos/waitlist/config"
	"github.com/hakisolos/waitlist/models"
)

//var err error

func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "waitlist api running",
	})
}

func JoinController(c *gin.Context) {
	var users models.Waiter
	err := c.ShouldBindJSON(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	var email string
	err = config.DB.QueryRow("SELECT name FROM waiters WHERE email = $1", users.Email).Scan(&email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "you have already joined",
		})
		return
	}
	_, err = config.DB.Exec("INSERT INTO waiters(name,email) VALUES ($1, $2)", users.Name, users.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	if err != sql.ErrNoRows && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "joined successfully",
	})
}

func GetUsersController(c *gin.Context) {

	rows, err := config.DB.Query("SELECT id,name,email FROM waiters")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	defer rows.Close()

	var usr []models.Waiter

	for rows.Next() {
		var u models.Waiter
		rows.Scan(&u.ID, &u.Name, &u.Email)
		usr = append(usr, u)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": usr,
	})
}
