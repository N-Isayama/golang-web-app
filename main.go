package main

import (
	"api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("vendor/html/template/*.html")
	router.Static("/assets", "./assets")

	user := router.Group("/user")
	{
		user.POST("/signup", api.UserSignUp)
		user.POST("/login", api.UserLogIn)
	}

	router.GET("/", api.Home)
	router.GET("/login", api.LogIn)
	router.GET("/signup", api.SignUp)
	router.NoRoute(api.NoRoute)

	router.Run(":8080")
}
