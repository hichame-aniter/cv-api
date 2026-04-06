package handlers

import (
	"cv-api/internal/models"
	"cv-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBasics(c *gin.Context) {
	// TODO
	// Remove Sensitive informations
	// email
	// phone
	// url
	// location

	data, err := services.GetBasics()
	if err != nil {
		c.Error(err) // let middleware handle it
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}
func UpdateBasics(c *gin.Context) {
	var input models.Basics

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err) // let middleware handle it
		return
	}

	err := services.UpdateBasics(input)
	if err != nil {
		c.Error(err) // let middleware handle it
		return
	}

	c.Status(http.StatusOK)
}
