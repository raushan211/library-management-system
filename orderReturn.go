package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Order_Return struct {
	OrderId    int `json:"order_id" binding:"required"`
	Student_id int `json:"student_id"`
	Fine       int `json:"fine"`
}

func order_return(c *gin.Context) {
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
	if CurrentUser["currentUser"].Id != data.Student_id {

		res := gin.H{
			"error": "invalid Order ID",
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := UpdateFine(data)

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
		"message": "Order Return Request Successfully Submited",
	}
	c.JSON(http.StatusOK, res)
}

func getOrderDetails(OrderId int) Order_Return {
	Data := Order_Return{}
	var returnDate time.Time

	userSQL := "SELECT id,student_id, return_date FROM orders WHERE id=$1"

	rows := DB.QueryRow(userSQL, OrderId)

	err := rows.Scan(&Data.OrderId, &Data.Student_id, &returnDate)
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now()

	today := currentTime.Format("2006-01-02")
	layout := "2006-01-02"
	t22, _ := time.Parse(layout, today)

	day_diff := (t22).Sub(returnDate).Hours() / 24

	if day_diff < 0 {
		Data.Fine = CalculateFine(int(day_diff))
	} else {
		Data.Fine = 0
	}

	//fmt.Println(day_diff)
	fmt.Println(Data)
	return Data

}

func CalculateFine(day int) int {

	fine := 10 * day
	return fine
}

func UpdateFine(Data Order_Return) bool {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	layout := "2006-01-02"
	t22, _ := time.Parse(layout, today)
	userSQL := `UPDATE orders SET  actual_return_date=$1, fine=$2, return_status=$3 WHERE id=$4`

	_, err2 := DB.Exec(userSQL, t22, Data.Fine, 0, Data.OrderId)

	if err2 != nil {
		log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}
