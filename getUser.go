package main

import (
	"fmt"
)

func getUserByEmail(email string) User {
	user_data := User{}
	userSQL := "SELECT id, first_name, last_name,email,type FROM users WHERE email=$1"

	rows := DB.QueryRow(userSQL, email)

	fmt.Println("abc", user_data)
	fmt.Println("row", rows)

	rows.Scan(&user_data.Id, &user_data.First_name, &user_data.Last_name, &user_data.Email, user_data.Type)

	return user_data

}

func getUserByID(id int) User {
	user_data := User{}
	userSQL := "SELECT id, first_name, last_name,email,type FROM users WHERE id=$1"

	rows := DB.QueryRow(userSQL, id)

	fmt.Println("abc", user_data)
	//fmt.Println("row", rows)

	rows.Scan(&user_data.Id, &user_data.First_name, &user_data.Last_name, &user_data.Email, &user_data.Type)
	fmt.Println("abcd", user_data)
	return user_data

}
