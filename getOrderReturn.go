package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderReturn(c *gin.Context) {

	data := returnOrder()

	res := gin.H{
		"Orders": data,
	}

	c.JSON(http.StatusBadRequest, res)
	return

}

func returnOrder() []Orders {
	Data := []Orders{}
	//user_id := CurrentUser["currentUser"].Id
	SQL := "SELECT orders.id, orders.student_id,orders.book_id AS book_book_id ,orders.order_request_id,orders.issue_date,orders.return_date, books.book_name,books.book_author,books.book_cover_image,books.book_id AS order_book_id from orders  LEFT join books  ON  orders.book_id=books.book_id WHERE orders.return_status=$1 "

	rows, err := DB.Query(SQL, 0)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	order := Orders{}
	var issueDate time.Time
	var returnDate time.Time

	for rows.Next() {
		rows.Scan(&order.ID, &order.Student_id, &order.Book_id, &order.Order_request_id, &issueDate, &returnDate, &order.Book_name, &order.Book_author, &order.Book_cover_image, &order.Book_id)

		order.Issue_date = issueDate.Format("2006-02-01")
		order.Return_date = returnDate.Format("2006-02-01")

		Data = append(Data, order)

	}

	fmt.Println(Data)
	return Data

}
