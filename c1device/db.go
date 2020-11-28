package c1device

import (
	"database/sql"
	"fmt"
	"time"
)

var position = [5]string{"electric", "mecanic", "software", "proiectare", "operational"}

type worker struct {
	id         int
	firstName  string
	lastName   string
	position   string
	cardNumber string
}

type project struct {
	id          int
	geNumber    string
	roNumber    string
	description string
	ipAddress   string
	deviceName  string
}

type workday struct {
	id      int
	worker  worker
	project project
	start   time.Time
	stop    time.Time
}

// DbConnect connects to database
func DbConnect() {
	_, err := sql.Open("mysql", "root:R00tPassword@/")
	if err != nil {
		fmt.Println("Connected to db")
	}
}
