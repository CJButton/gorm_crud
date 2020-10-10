package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Ping)
		v1.GET("/pong", Pong)

		// auth := new(controllers.Create)
		// auth := new(controller.Authentication)
		// auth.POST("/", controllers.auth)

	}

	r.Run(":8080")
}

// Ping is a test controller
func Ping(c *gin.Context) {
	body := map[string]string{"message": "11"}

	SuccessComposer(body, c, http.StatusOK)
	return
}

// Pong blah blah
func Pong(c *gin.Context) {
	errors := "oh no!"
	FailureComposer(c, http.StatusBadRequest, errors)
	return
}

// SuccessComposer formats successful requests
func SuccessComposer(body interface{}, c *gin.Context, status int) {
	c.JSON(status, gin.H{
		"body":        body,
		"success":     true,
		"request_url": c.Request.URL.Path,
		"method":      c.Request.Method,
		"status":      status,
	})
}

// FailureComposer balh blah
func FailureComposer(c *gin.Context, status int, errors interface{}) {
	c.AbortWithStatusJSON(status, gin.H{
		"body":        nil,
		"success":     false,
		"request_url": c.Request.URL.Path,
		"method":      c.Request.Method,
		"errors":      errors,
	})
}
