package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var CurrentUser = make(map[string]User)

func isLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware
		cookie_id, err := c.Cookie("id")
		fmt.Println("cookie value: ", cookie_id)

		if err != nil || cookie_id == "" {

			fmt.Println(err)
			res := gin.H{
				"message": "Access denied",
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
		id, err2 := strconv.Atoi(cookie_id)
		data := getUserByID(id)
		CurrentUser["currentUser"] = data
		fmt.Println("current data", CurrentUser)
		fmt.Println(CurrentUser["currentUser"].First_name)
		if err2 != nil {

			fmt.Println(err)
			res := gin.H{
				"message": "Access denied",
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}

	}
}

func isAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware
		cookie_id, err := c.Cookie("id")
		fmt.Println("cookie value: ", cookie_id)

		if err != nil || cookie_id == "" {

			fmt.Println(err)
			res := gin.H{
				"message": "Access denied",
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
		id, err2 := strconv.Atoi(cookie_id)
		if err2 != nil {

			fmt.Println(err)
			res := gin.H{
				"message": "Access denied",
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
		data := getUserByID(id)
		CurrentUser["currentUser"] = data

		if CurrentUser["currentUser"].Type != 1 {

			fmt.Println(err)
			res := gin.H{
				"message": "Access denied",
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}

	}
}
