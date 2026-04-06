package handlers

import (
	"cv-api/internal/models"
	"cv-api/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExperiences(c *gin.Context) {
	data, err := services.GetExperiences()
	if err != nil {
		c.Error(err) // let middleware handle it
		return
	}

	c.JSON(http.StatusOK, data)
}
func GetExperience(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("HANDLER CALLED WITH:", id)
	data, err := services.GetExperience(id)
	if err != nil {
		fmt.Println("ERROR RETURNED:", err)
		c.Error(err) // let middleware handle it
		return
	}

	c.JSON(http.StatusOK, data)
}
func UpdateExperience(c *gin.Context) {
	id := c.Param("id") // assuming URL param like /experiences/:id

	var updatedExp models.Experience
	if err := c.ShouldBindJSON(&updatedExp); err != nil {
		c.Error(err)
		return
	}

	// Ensure the ID from URL matches the one in body (optional but recommended)
	updatedExp.ID = id

	if err := services.UpdateExperience(id, updatedExp); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, updatedExp)
}
