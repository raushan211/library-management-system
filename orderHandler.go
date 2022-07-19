package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func orderHandler(c *gin.Context) {

	data := getLiveOrders()

	res := gin.H{
		"Orders": data,
	}

	c.JSON(http.StatusBadRequest, res)
	return

}
