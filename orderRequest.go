package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func orderRequest(c *gin.Context) {
	reqBody := Order_request{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if CurrentUser["currentUser"].Id != reqBody.Student_id {

		res := gin.H{
			"error": "student id is invalid",
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	ress := day_diff(reqBody.Issue_date, reqBody.Return_date)

	if ress > 10 {

		res := gin.H{
			"error": "Maximum difference date not more than 10 days",
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if !isBookExist(reqBody.Book_id) {

		res := gin.H{
			"error": "Invalid Book ID",
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if isOutOfStock(reqBody.Book_id) {

		res := gin.H{
			"error": "Book Out Of Stock",
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	_, err_result := createOrderRequest(reqBody)

	if err_result != "" {
		res := gin.H{
			"error": err_result,
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"success": true,
		"message": "Order Request Successfully Submited",
	}
	c.JSON(http.StatusOK, res)
}

func createOrderRequest(reqbody Order_request) (bool, string) {
	var result = true
	var err_responce = ""
	var approved = "pending"

	sqlStatement := `
INSERT INTO order_request(book_id, student_id,issue_date, return_date, approved)
VALUES ($1, $2, $3, $4, $5)`
	_, err2 := DB.Exec(sqlStatement, reqbody.Book_id, reqbody.Student_id, reqbody.Issue_date, reqbody.Return_date, approved)
	fmt.Println(err2)

	if err2 != nil {
		err_responce = "Something went wrong"
		return false, err_responce
	}
	//log.Fatal("ERror in insert: ", err2)

	result = false
	return result, err_responce

}

func day_diff(t1, t2 string) float64 {
	layout := "2006-01-02"
	// str := "2014-11-12T11:45:26.371Z"

	t11, _ := time.Parse(layout, t1)
	t22, _ := time.Parse(layout, t2)
	fmt.Println("date:", (t22).Sub(t11).Hours()/24)
	day_diff := (t22).Sub(t11).Hours() / 24

	return day_diff

}
