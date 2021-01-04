package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

const (
	driver      string = "mysql"
	credentials string = "root:R00tpassword@/testEccelC1"
	database    string = "testEccelC1"
)

// DAO ...
type DAO struct {
	db *sql.DB
}

// Connect connects to database
func (dao *DAO) Connect() {
	var err error = nil
	dao.db, err = sql.Open(driver, credentials)
	if err != nil {
		fmt.Println(err)
	}
}

// IsConnected verify database connection.
func (dao *DAO) IsConnected() bool {
	err := dao.db.Ping()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// InsertIntoWorkday returns the sql command to start/stop time on workday
func (dao *DAO) InsertIntoWorkday(deviceName, cardUID string) string {
	var command string = "CALL INSERT_INTO_WORKDAY(\"" + deviceName + "\", \"" + cardUID + "\");"
	return command
}

// Execute executes a command against database
func (dao *DAO) Execute(command string) {
	dao.db.Exec(command)
}
