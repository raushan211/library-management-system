package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Order_approve(c *gin.Context) {
	reqBody := Order{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := approve_action(reqBody.OrderRequestId, reqBody.Approved)

	if result == false {
		res := gin.H{
			"success": false,
			"message": "Something went wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}
	msg := ""
	if reqBody.Approved == "true" {
		msg = "Order Approved"
	} else if reqBody.Approved == "false" {
		msg = "Order Approved"
	} else {
		res := gin.H{
			"success": false,
			"message": "Invalid Request",
		}
		c.JSON(http.StatusBadRequest, res)

	}

	res := gin.H{
		"success": true,
		"message": msg,
	}
	c.JSON(http.StatusOK, res)
}

func approve_action(request_id int, approve string) bool {
	fmt.Println("request id", request_id)
	OrderData := Order{}
	userSQL := "SELECT id, book_id, student_id,return_date, issue_date FROM order_request WHERE id=$1"

	row := DB.QueryRow(userSQL, request_id)

	err := row.Scan(&OrderData.ID, &OrderData.Book_id, &OrderData.Student_id, &OrderData.Return_date, &OrderData.Issue_date)
	fmt.Println(err, OrderData)

	if err != nil {
		// log the error

		return false
	}

	sqlStatement := `
	INSERT INTO orders(book_id, order_request_id,student_id,issue_date, return_date, approved)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err2 := DB.Exec(sqlStatement, OrderData.Book_id, request_id, OrderData.Student_id, OrderData.Issue_date, OrderData.Return_date, approve)
	fmt.Println(err2)

	if err2 != nil {
		//log.Fatal(err2)
		//return false
	}

	return true
}
