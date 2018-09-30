package main

import (
	"TestImport/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.Hello)
	router.GET("/user", controllers.GetUser)

	router.POST("/addperson", controllers.AddUser)

	//router.GET("/persons", GetPersonsApi)
	//
	//router.GET("/person/:id", GetPersonApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//
	//router.DELETE("/person/:id", DelPersonApi)

	router.GET("/GetTimeNow",controllers.GetDateTimeNow)

	return router
}
