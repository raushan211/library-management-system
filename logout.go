package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {

	c.SetCookie("id", "", -1, "", "", false, false)
	res := gin.H{
		"success": true,
		"message": "logout successfull",
	}
	c.JSON(http.StatusOK, res)
	return
}
