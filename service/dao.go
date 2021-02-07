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
	"strconv"

	"example.com/c1/model"
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

// CloseConnection closes the connection to database
func (dao *DAO) CloseConnection() {
	if dao.IsConnected() {
		dao.db.Close()
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

// Execute executes a command against database with no returning result set.
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
func (dao *DAO) RetrieveActiveWorkdays(rows *sql.Rows) map[int][]string {

	table := make(map[int][]string)
	var id int
	var worker string
	var roNumber string
	var geNumber string
	var description string
	var workedTime string

	for rows.Next() {
		rows.Scan(&id, &worker, &roNumber, &geNumber, &description, &workedTime)
		table[id] = []string{worker, roNumber, geNumber, description, workedTime}
	}
	return table
}

// RetrieveCurrentMonthTimeRaport - TODO write about behavior
func (dao *DAO) RetrieveCurrentMonthTimeRaport(rows *sql.Rows) map[int][]string {
	table := make(map[int][]string)
	var id int
	var geNo string
	var roNo string
	var description string
	var start string
	var stop string
	var minutes string

	for rows.Next() {
		rows.Scan(&id, &geNo, &roNo, &description, &start, &stop, &minutes)
		table[id] = []string{geNo, roNo, description, start, stop, toHoursAndMinutes(minutes)}
	}
	return table
}

// RetrieveWorkerStatus - TODO write about behavior
func (dao *DAO) RetrieveWorkerStatus(rows *sql.Rows) (string, string) {
	var status string
	var workedMinutes int

	rows.Next()
	rows.Scan(&status)
	rows.Next()
	rows.Scan(&workedMinutes)

	workedTime := toHoursAndMinutes(strconv.Itoa(workedMinutes))
	return status, workedTime
}

func toHoursAndMinutes(minutes string) string {

	workedMinutes, _ := strconv.Atoi(minutes)

	workedHours := workedMinutes / 60
	workedMinutes = workedMinutes - (workedHours * 60)

	workedTime := strconv.Itoa(workedHours) + "h" + strconv.Itoa(workedMinutes) + "m"
	return workedTime
}

// RetrieveActiveProjects ...
func (dao *DAO) RetrieveActiveProjects(rows *sql.Rows) []model.Project {
	projects := make([]model.Project, 0)

	for rows.Next() {
		var proj model.Project
		rows.Scan(&proj.ID,
			&proj.GeNumber,
			&proj.RoNumber,
			&proj.Description,
			&proj.DeviceID,
			&proj.IsActive,
			&proj.Begin,
			&proj.End)
		projects = append(projects, proj)
	}
	return projects
}

// RetrieveAllWorkers ...
func (dao *DAO) RetrieveAllWorkers(rows *sql.Rows) []model.Worker {
	workers := make([]model.Worker, 0)

	for rows.Next() {
		var worker model.Worker
		rows.Scan(&worker.ID,
			&worker.FirstName,
			&worker.LastName,
			&worker.CardNumber,
			&worker.Position,
			&worker.IsActive)
		workers = append(workers, worker)
	}
	return workers
}
