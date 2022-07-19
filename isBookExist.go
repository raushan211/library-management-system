package main

func isBookExist(bookId string) bool {
	var count int

	userSQL := "SELECT COUNT(*) FROM books WHERE book_id=$1"

	row := DB.QueryRow(userSQL, bookId)

	err := row.Scan(&count)

	if err != nil {
		//log.Fatal(err)
	}

	if count == 1 {
		return true
	} else {
		return false
	}

}
