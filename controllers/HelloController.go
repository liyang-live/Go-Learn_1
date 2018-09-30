package controllers

import (
	"TestImport/models"
	"TestImport/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello word!")
}

func GetUser(c *gin.Context) {
	p := models.Person{Id: 111, Name: "User 1", Age: 18, Sex: 1}
	c.JSON(http.StatusOK, p)
}

func AddUser(c *gin.Context) {
	//p := models.Person{Name: "User 1", Age: 18, Sex: 1}
	var p  models.Person
	c.Bind(&p)

	people:=models.Person{Name:p.Name,Age:p.Age,Sex:p.Sex}

	service.AddPerson(people)
	//c.JSON(http.StatusOK, p)
}

