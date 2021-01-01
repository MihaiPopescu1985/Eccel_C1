package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

const (
	driver      string = "mysql"
	credentials string = "root:R00tpassword@/"
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

// InsertWorkday insert into database the time witch work is started or is stopped
func (dao *DAO) InsertWorkday(deviceName, cardUID string) string {
	var command string
	return command
}
