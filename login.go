package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(c *gin.Context) {
	reqBody := Login{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": parseError(err),
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := check_auth(reqBody)

	if result == true {
		user_data := getUserByEmail(reqBody.Email)
		c.SetCookie("id", strconv.Itoa(user_data.Id), 60*60, "", "", false, false)
		//c.Header("result", reqBody.Email)
		res := gin.H{
			"success": true,
			"message": "Login Successfull",
		}
		c.JSON(http.StatusOK, res)

		return
	} else {
		res := gin.H{
			"success": false,
			"message": "Invalid  Credential",
		}
		c.JSON(http.StatusOK, res)
		return
	}

}

func check_auth(reqBody Login) bool {

	var count int

	userSQL := "SELECT COUNT(*) FROM users WHERE email=$1 AND password=$2"

	row := DB.QueryRow(userSQL, reqBody.Email, reqBody.Password)
	err := row.Scan(&count)
	if err != nil {
		//log.Fatal(err)
	}
	fmt.Println("c", count)
	if count == 1 {
		return true
	} else {
		return false
	}

}
