package model

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"example.com/c1/util"
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

// Connect connects to database.
// The Open method is not enough to verify a
// database connection so a Ping method should be
// called. The Ping method is embedded inside
// IsConnected method.
func (db *DB) Connect() {

	var err error = nil
	db.database, err = sql.Open(driver, credentials)

	if err != nil {
		util.Log.Fatalln(err)
	}
	db.IsConnected()
}

// IsConnected verify database connection.
func (db *DB) IsConnected() bool {
	err := db.database.Ping()

	if err != nil {
		util.Log.Fatalln(err)
		return false
	}
	return true
}

// Execute executes a command against database with no returning result set.
func (db *DB) execute(command string) {

	util.Log.Printf("Executing: %v \n", command)
	_, err := db.database.Exec(command)

	if err != nil {
		util.Log.Panicln(err)
	}
}

// ExecuteQuery TODO: write about the behavior of this function
func (db *DB) executeQuery(command string) *sql.Rows {

	util.Log.Printf("Executing: %v \n", command)
	rows, err := db.database.Query(command)

	// Close rows after 5 seconds
	go closeRows(rows, time.Second*5)

	if err != nil {
		util.Log.Panicln(err)
	}
	return rows
}

// Helper function used for closing query rows after a certain time.
// Needed because database server will keep open a conection for every query.
func closeRows(rows *sql.Rows, timeSpan time.Duration) {
	time.Sleep(timeSpan)

	if err := rows.Close(); err != nil {
		util.Log.Println(err)
	}
}

// InsertIntoWorkday returns the sql command to start/stop time on workday.
func (db *DB) InsertIntoWorkday(deviceName, cardUID string) {
	command := "CALL INSERT_INTO_WORKDAY(\"" + deviceName + "\", \"" + cardUID + "\");"
	db.execute(command)
}

// RetrieveActiveWorkdays - TODO write about behavior
func (db *DB) RetrieveActiveWorkdays() map[int][]string {

	command := "CALL SELECT_ACTIVE_WORKDAY;"
	rows := db.executeQuery(command)

	table := make(map[int][]string)
	var (
		id          sql.NullInt32  //int
		worker      sql.NullString //string
		roNumber    sql.NullString //string
		geNumber    sql.NullString //string
		description sql.NullString //string
	)

	for rows.Next() {
		if err := rows.Scan(&id, &worker, &roNumber, &geNumber, &description); err != nil {
			util.Log.Panicln(err)
		}
		table[int(id.Int32)] = []string{worker.String, roNumber.String, geNumber.String, description.String}
	}
	return table
}

// RetrieveCurrentMonthTimeRaport - TODO write about behavior
func (db *DB) RetrieveCurrentMonthTimeRaport(workerID, currentMonth, currentYear int) map[int][]string {

	command := "CALL SELECT_MONTH_TIME_RAPORT(" + strconv.Itoa(workerID) + ", " + strconv.Itoa(currentMonth) + ", " + strconv.Itoa(currentYear) + ");"
	rows := db.executeQuery(command)

	table := make(map[int][]string)
	id := 0

	var (
		geNo        sql.NullString //string
		roNo        sql.NullString //string
		description sql.NullString //string
		start       sql.NullString //string
		stop        sql.NullString //string
		minutes     sql.NullString //string
	)

	for rows.Next() {
		if err := rows.Scan(&geNo, &roNo, &description, &start, &stop, &minutes); err != nil {
			util.Log.Panicln(err)
		}
		table[id] = []string{geNo.String, roNo.String, description.String, start.String, stop.String, minutes.String}
		id++
	}
	return table
}

// RetrieveWorkerStatus - TODO write about behavior
func (db *DB) RetrieveWorkerStatus(id int) (string, string) {

	var command string = "CALL SELECT_WORKER_STATUS(" + strconv.Itoa(id) + ");"
	rows := db.executeQuery(command)

	var status sql.NullString       //string
	var workedMinutes sql.NullInt32 //int

	rows.Next()
	if err := rows.Scan(&status); err != nil {
		util.Log.Panicln(err)
	}
	rows.Next()
	if err := rows.Scan(&workedMinutes); err != nil {
		util.Log.Panicln(err)
	}

	workedTime := strconv.Itoa(int(workedMinutes.Int32))
	return status.String, workedTime
}

// RetrieveActiveProjects ...
func (db *DB) RetrieveActiveProjects() []Project {

	var command string = "CALL GET_ACTIVE_PROJECTS();"
	rows := db.executeQuery(command)

	projects := make([]Project, 0)

	var (
		projID   sql.NullInt32
		geNo     sql.NullString
		roNo     sql.NullString
		descript sql.NullString
		devID    sql.NullInt32
		isActive sql.NullBool
		begin    sql.NullString
		end      sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(
			&projID, &geNo, &roNo, &descript, &devID, &isActive, &begin, &end); err != nil {
			util.Log.Panicln(err)
		}

		projects = append(projects, Project{
			ID:          int(projID.Int32),
			GeNumber:    geNo.String,
			RoNumber:    roNo.String,
			Description: descript.String,
			DeviceID:    int(devID.Int32),
			IsActive:    isActive.Bool,
			Begin:       begin.String,
			End:         end.String,
		})
	}
	return projects
}

// RetrieveAllWorkers ...
func (db *DB) RetrieveAllWorkers() []Worker {

	command := "CALL GET_ALL_WORKERS();"
	rows := db.executeQuery(command)

	var (
		workers = make([]Worker, 0)

		wID       sql.NullInt32
		wFirstN   sql.NullString
		wLastN    sql.NullString
		wCardNo   sql.NullString
		wPos      sql.NullString
		wIsActive sql.NullBool
	)

	for rows.Next() {
		if err := rows.Scan(
			&wID, &wFirstN, &wLastN, &wCardNo, &wPos, &wIsActive); err != nil {
			util.Log.Panicln(err)
		}

		workers = append(workers, Worker{
			ID:         int(wID.Int32),
			FirstName:  wFirstN.String,
			LastName:   wLastN.String,
			CardNumber: wCardNo.String,
			Position:   wPos.String,
			IsActive:   wIsActive.Bool,
		})
	}
	return workers
}

// GetUserByNameAndPassword TODO: write about function
func (db *DB) GetUserByNameAndPassword(name, password string) Worker {

	command := "SELECT * FROM WORKER WHERE NICKNAME = \"" + name + "\" AND PASSWORD = \"" + password + "\";"
	rows := db.executeQuery(command)

	var (
		wID       sql.NullInt32
		wFirstN   sql.NullString
		wLastN    sql.NullString
		wCardNo   sql.NullString
		wPos      sql.NullString
		wIsActive sql.NullBool
		wNick     sql.NullString
		wPass     sql.NullString
		wAccess   sql.NullInt32
	)

	for rows.Next() {
		if err := rows.Scan(
			&wID, &wFirstN, &wLastN, &wCardNo, &wPos, &wIsActive, &wNick, &wPass, &wAccess); err != nil {
			util.Log.Panicln(err)
		}
	}
	return Worker{
		ID:          int(wID.Int32),
		FirstName:   wFirstN.String,
		LastName:    wLastN.String,
		CardNumber:  wCardNo.String,
		Position:    wPos.String,
		IsActive:    wIsActive.Bool,
		Nickname:    wNick.String,
		Password:    wPass.String,
		AccessLevel: int(wAccess.Int32),
	}
}

// RetrieveWorkerName returns worker's name based on id.
func (db *DB) RetrieveWorkerName(id int) string {
	var (
		firstName sql.NullString
		lastName  sql.NullString
	)
	command := "SELECT FIRSTNAME, LASTNAME FROM WORKER WHERE ID = " + strconv.Itoa(id) + ";"
	rows := db.executeQuery(command)

	for rows.Next() {
		if err := rows.Scan(&firstName, &lastName); err != nil {
			util.Log.Panicln(err)
		}
	}
	return firstName.String + " " + lastName.String
}

// RetrieveFreeDays returns a map containing free days.
func (db *DB) RetrieveFreeDays() map[int]string {

	var (
		command = "SELECT * FROM FREEDAYS;"
		rows    = db.executeQuery(command)

		id   sql.NullInt32
		date sql.NullString

		table = make(map[int]string, 0)
	)

	for rows.Next() {
		if err := rows.Scan(&id, &date); err != nil {
			util.Log.Panicln(err)
		}
		table[int(id.Int32)] = date.String
	}
	return table
}

// RetrieveOvertime ...
func (db *DB) RetrieveOvertime(workerID int) string {
	var (
		command  = "CALL GET_OVERTIME(" + strconv.Itoa(workerID) + ");"
		overtime sql.NullString
		rows     = db.executeQuery(command)
	)
	for rows.Next() {
		if err := rows.Scan(&overtime); err != nil {
			util.Log.Panicln(err)
		}
	}
	return overtime.String
}

func (db *DB) AddWorkday(workerID, projectID int, startHour, stopHour string) {

	var command strings.Builder
	command.WriteString("CALL ADD_NEW_WORKDAY(")
	command.WriteString(strconv.Itoa(workerID))
	command.WriteString(", ")
	command.WriteString(strconv.Itoa(projectID))
	command.WriteString(", '")
	command.WriteString(startHour)
	command.WriteString("', '")
	command.WriteString(stopHour)
	command.WriteString("');")

	db.execute(command.String())
}

func (db *DB) AddProject(geNo, roNo, descr, startDate string) {
	var command strings.Builder

	command.WriteString("CALL ADD_NEW_PROJECT('")
	command.WriteString(geNo)
	command.WriteString("', '")
	command.WriteString(roNo)
	command.WriteString("', '")
	command.WriteString(descr)
	command.WriteString("', '")
	command.WriteString(startDate)
	command.WriteString("');")

	db.execute(command.String())
}

func (db *DB) RetrieveAllPositions() map[int]string {
	var (
		command   = ("CALL GET_ALL_POSITIONS()")
		positions = make(map[int]string, 0)
		rows      = db.executeQuery(command)
	)

	for rows.Next() {
		var pos sql.NullString
		var id sql.NullInt32

		if err := rows.Scan(&id, &pos); err != nil {
			util.Log.Println(err)
		}
		positions[int(id.Int32)] = pos.String
	}
	return positions
}

func (db *DB) AddWorker(firstName, lastName, cardNumber, position, nickName, password string) {
	var command strings.Builder

	command.WriteString("CALL ADD_NEW_WORKER('")
	command.WriteString(firstName)
	command.WriteString("', '")
	command.WriteString(lastName)
	command.WriteString("', '")
	command.WriteString(cardNumber)
	command.WriteString("', '")
	command.WriteString(position)
	command.WriteString("', '")
	command.WriteString(nickName)
	command.WriteString("', '")
	command.WriteString(password)
	command.WriteString("');")

	db.execute(command.String())
}
