package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Authentication struct{}

// Create user with password
func (Authentication) Create(c *gin.Context) {
	fmt.Println("hello")
}
