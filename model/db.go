package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

const (
	driver      string = "mysql"
	credentials string = "root:R00tpassword@/EccelC1"
	database    string = "EccelC1"
)

// Db represent a global variable for storing a database connection.
var Db DB

// DB ...
type DB struct {
	database *sql.DB
}

// ActiveWorkdays - placeholder for storing active workdays retrieved from database
type ActiveWorkdays struct {
	Workdays map[int][5]string
}

// Connect connects to database and closes the connections after 10 seconds
func (db *DB) Connect() {
	var err error = nil
	db.database, err = sql.Open(driver, credentials)
	if err != nil {
		fmt.Println(err)
	}
}

// IsConnected verify database connection.
func (db *DB) IsConnected() bool {
	err := db.database.Ping()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// InsertIntoWorkday returns the sql command to start/stop time on workday.
func (db *DB) InsertIntoWorkday(deviceName, cardUID string) {
	command := "CALL INSERT_INTO_WORKDAY(\"" + deviceName + "\", \"" + cardUID + "\");"
	db.execute(command)
}

// Execute executes a command against database with no returning result set.
func (db *DB) execute(command string) {
	// TODO: make sure the command is proper executed, no error is triggered
	db.database.Exec(command)
}

// ExecuteQuery TODO: write about the behavior of this function
func (db *DB) executeQuery(command string) *sql.Rows {
	rows, err := db.database.Query(command)

	if err != nil {
		fmt.Println(err)
	}
	return rows
}

// RetrieveActiveWorkdays - TODO write about behavior
func (db *DB) RetrieveActiveWorkdays() map[int][]string {

	command := "CALL SELECT_ACTIVE_WORKDAY;"
	rows := db.executeQuery(command)

	table := make(map[int][]string)
	var id int
	var worker string
	var roNumber string
	var geNumber string
	var description string

	for rows.Next() {
		rows.Scan(&id, &worker, &roNumber, &geNumber, &description)
		table[id] = []string{worker, roNumber, geNumber, description}
	}
	return table
}

// RetrieveCurrentMonthTimeRaport - TODO write about behavior
func (db *DB) RetrieveCurrentMonthTimeRaport(workerID, currentMonth int) map[int][]string {

	command := "CALL SELECT_MONTH_TIME_RAPORT(" + strconv.Itoa(workerID) + ", " + strconv.Itoa(currentMonth) + ");"
	rows := db.executeQuery(command)

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
func (db *DB) RetrieveWorkerStatus(id int) (string, string) {

	var command string = "CALL SELECT_WORKER_STATUS(" + strconv.Itoa(id) + ");"
	rows := db.executeQuery(command)

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
func (db *DB) RetrieveActiveProjects() []Project {

	var command string = "CALL GET_ACTIVE_PROJECTS();"
	rows := db.executeQuery(command)

	projects := make([]Project, 0)

	for rows.Next() {
		var proj Project
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
func (db *DB) RetrieveAllWorkers() []Worker {

	command := "CALL GET_ALL_WORKERS();"
	rows := db.executeQuery(command)

	workers := make([]Worker, 0)

	for rows.Next() {
		var worker Worker
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

// GetUserByNameAndPassword TODO: write about function
func (db *DB) GetUserByNameAndPassword(name, password string) Worker {
	var worker Worker

	command := "SELECT * FROM WORKER WHERE NICKNAME = \"" + name + "\" AND PASSWORD = \"" + password + "\";"
	rows := db.executeQuery(command)

	for rows.Next() {
		rows.Scan(&worker.ID,
			&worker.FirstName,
			&worker.LastName,
			&worker.CardNumber,
			&worker.Position,
			&worker.IsActive,
			&worker.Nickname,
			&worker.Password,
			&worker.AccessLevel)
	}
	return worker
}

// RetrieveWorkerName returns worker's name based on id.
func (db *DB) RetrieveWorkerName(id int) string {

	firstName := ""
	lastName := ""

	command := "SELECT FIRSTNAME, LASTNAME FROM WORKER WHERE ID = " + strconv.Itoa(id) + ";"
	rows := db.executeQuery(command)

	for rows.Next() {
		rows.Scan(&firstName, &lastName)
	}
	return firstName + " " + lastName
}
