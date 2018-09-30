package controllers

import (
	"TestImport/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetDateTimeNow(c *gin.Context)  {

	str:=time.Now().String()
	fmt.Println(str)
	c.JSON( http.StatusOK,common.TimFormat( str,"yyyy-MM-dd HH:mm"))
}
