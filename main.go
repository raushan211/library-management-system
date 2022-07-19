package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DB_DSN = "postgres://localhost:5432/lms?sslmode=disable"
)

var (
	DB *sql.DB
)

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email" binding:"email"`
	Password   string `json:"password" binding:"required,min=8,alphanum"`
	Type       int    `json:"type"`
}

func setupRoutes(r *gin.Engine) {

	r.POST("signup/admin", Signup)
	r.POST("signup/user", Signup)
	r.POST("login", login)
	r.POST("logout", logout)

	// r.GET("/orders", isLogin(), orderHandler)
	// r.GET("/orders/history", isLogin(), orderHistoryHandler)

	// r.POST("/order", isLogin(), orderRequest)

	// r.POST("/order/return", isLogin(), order_return)

	// r.GET("/search/:name", isLogin(), searchHandler)

	// //ADMIN
	// r.POST("/return/approve", isAdmin(), ReturnAccept)

	// r.POST("/order/approve", isAdmin(), Order_approve)

	// r.GET("/order/return", isAdmin(), OrderReturn)

}

type Book struct {
	Book_id          string `json:"book_id"`
	Current_stock    int    `json:"current_stock"`
	Book_name        string `json:"book_name"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
}
type Order_request struct {
	ID                 int    `json:"id"`
	Book_id            string `json:"book_id"`
	Student_id         int    `json:"student_id"`
	Issue_date         string `json:"issue_date"`
	Return_date        string `json:"return_date"`
	Actual_return_date string `json:"actual_return_date"`
	Fine               string `json:"fine"`
	Approved           string `json:"approved"`
}
type Order struct {
	ID                 int    `json:"id"`
	OrderRequestId     int    `json:"order_request_id" binding:"required"`
	Book_id            string `json:"book_id"`
	Student_id         int    `json:"student_id"`
	Issue_date         string `json:"issue_date"`
	Return_date        string `json:"return_date"`
	Actual_return_date string `json:"actual_return_date"`
	Fine               string `json:"fine"`
	Approved           string `json:"approved" binding:"required"`
}
type Orders struct {
	ID               int    `json:"id"`
	Book_id          string `json:"book_id"`
	Book_name        string `json:"book_name"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
	Student_id       int    `json:"student_id"`
	Order_request_id int    `json:"order_request_id"`
	Issue_date       string `json:"issue_date"`
	Return_date      string `json:"return_date"`
	Fine             string `json:"fine"`
	Return_status    string `json:"return_status"`
}

func main() {
	createDBConnection()
	// records := readCsvFile("./books.csv")
	// fmt.Println("Length of records is: ", len(records))
	// importCSV(records)
	//fmt.Println(records)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	// defer DB.Close()
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
func importCSV(records [][]string) {
	for i := 0; i < 1000; i++ {
		sqlStatement := `
		INSERT INTO books(book_id, book_name, book_author, book_cover_image)
		VALUES($1, $2, $3, $4)`
		_, err := DB.Exec(sqlStatement, records[i][0], records[i][1], records[i][2], records[i][5])
		if err != nil {
			log.Println("error in insert: ", err)
		}
	}
}
