package dbcommunication

import (
	"database/sql"
	"fmt"
)

// DbConnect connects to database
func DbConnect() {
	_, err := sql.Open("mysql", "root:R00tPassword@/")
	if err != nil {
		fmt.Println("Connected to db")
	}
}
