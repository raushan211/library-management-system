package main

import (
	"fmt"
	"log"
)

func isOutOfStock(bookId string) bool {
	//var count int
	data := Book{}

	userSQL := "SELECT book_id, currrent_stock FROM books WHERE book_id=$1"

	err := DB.QueryRow(userSQL, bookId).Scan(&data.Book_id, &data.Current_stock)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("count: ", data)

	if data.Current_stock > 0 {
		return false
	} else {
		return true
	}

}
