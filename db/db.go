package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

/*
create table WORKER
(ID int not null primary key auto_increment,
FIRSTNAME varchar(15) not null,
LASTNAME varchar(20) not null,
POSITION varchar(15) not null,
CARDNUMBER varchar(14) not null);

create table PROJECT
(ID int not null primary key auto_increment,
GENUMBER varchar(15) not null,
RONUMBER varchar(15) not null,
DESCRIPTION varchar(100),
IPADDRESS varchar(15) not null,
DEVICENAME varchar(20) not null);

create table WORKDAY
(ID int not null primary key auto_increment,
WORKER int not null,
PROJECT int not null,
STARTTIME timestamp not null,
STOPTIME timestamp);

alter table WORKDAY add foreign key (WORKER) references WORKER(ID);
alter table WORKDAY add foreign key (PROJECT) references PROJECT(ID);
*/

var db *sql.DB = nil

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
	id        int
	worker    worker
	project   project
	startTime time.Time
	stopTime  time.Time
}

// Connect connects to database
func Connect() {
	var err error
	db, err = sql.Open("mysql", "root:R00tpassword@/")
	if err != nil {
		fmt.Println(err)
	}
}

// IsConnected verify database connection
func IsConnected() bool {
	err := db.Ping()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// InsertWorkday inserts a new workday into database.
// It first checks
func InsertWorkday(timeStamp time.Time, proj string, tag string) {

}
