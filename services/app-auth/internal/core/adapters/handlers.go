package adapters

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Register() gin.HandlerFunc
	Logout() gin.HandlerFunc
	Login() gin.HandlerFunc
	GetUserByID() gin.HandlerFunc
}
