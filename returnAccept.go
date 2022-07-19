package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnAccept(c *gin.Context) {
	reqBody := Order_Return{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}
	data := getOrderDetails(reqBody.OrderId)

	if data.Fine > 0 {
		res := gin.H{
			"success": false,
			"message": "Please Clear Your Fine",
		}
		c.JSON(http.StatusBadRequest, res)
		return

	}

	result := UpdateReturnStaus(data)

	if result == false {
		res := gin.H{
			"success": false,
			"message": "Something Went Wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return

	}

	res := gin.H{
		"success": true,
		"message": "Order Return Request Accepted",
	}
	c.JSON(http.StatusOK, res)
}

func UpdateReturnStaus(Data Order_Return) bool {
	email := CurrentUser["currentUser"].Email
	userSQL := `UPDATE orders SET  return_status=$1,accepted_by=$2 WHERE id=$3`

	_, err2 := DB.Exec(userSQL, 1, email, Data.OrderId)

	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}
