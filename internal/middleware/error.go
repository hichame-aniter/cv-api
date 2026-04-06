package middleware

import (
	appErrors "cv-api/internal/errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Log real error
			log.Println("ERROR:", err)
			fmt.Println("MIDDLWARE CALLED:", err)
			// Map error → response
			switch err {
			case appErrors.ErrNotFound:
				fmt.Println("MIDDLWARE CHOOSE NOTFOUND:", err)
				c.JSON(http.StatusNotFound, APIError{
					Message: "Resource not found",
					Code:    "NOT_FOUND",
				})
			case appErrors.ErrInvalid:
				fmt.Println("MIDDLWARE CHOOSE INVALID:", err)
				c.JSON(http.StatusBadRequest, APIError{
					Message: "Invalid request",
					Code:    "INVALID_REQUEST",
				})
			default:
				fmt.Println("MIDDLWARE CHOOSE DEFAULT:", err)
				c.JSON(http.StatusInternalServerError, APIError{
					Message: "Internal server error",
					Code:    "INTERNAL_ERROR",
				})
			}
		}
	}
}
