package service

/*

Two devices are assigned for entrance and exit.
The worker is comming to work, enabling one of this devices.
After that point, the worker is able to start working.

In order for the worker to start working,
he must enable another device.

*/

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

// ActiveWorkdays - placeholder for storing active workdays retrieved from database
type ActiveWorkdays struct {
	Workdays map[int][5]string
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

// InsertIntoWorkday returns the sql command to start/stop time on workday.
func (dao *DAO) InsertIntoWorkday(deviceName, cardUID string) string {
	return "CALL INSERT_INTO_WORKDAY(\"" + deviceName + "\", \"" + cardUID + "\");"
}

// SelectActiveWorkday returns the sql command to select active workers
func (dao *DAO) SelectActiveWorkday() string {
	return "CALL SELECT_ACTIVE_WORKDAY;"
}

// Execute executes a command against database
func (dao *DAO) Execute(command string) {
	// TODO: make sure the command is proper executed, no error is triggered
	dao.db.Exec(command)
}

// ExecuteQuery TODO: write about the behavior of this function
func (dao *DAO) ExecuteQuery(command string) *sql.Rows {
	rows, err := dao.db.Query(command)

	if err != nil {
		fmt.Println(err)
	}

	return rows
}

// RetrieveActiveWorkdays - TODO write about behavior
func (dao *DAO) RetrieveActiveWorkdays(rows *sql.Rows) ([]string, map[int]string) {
	return nil, nil
}
