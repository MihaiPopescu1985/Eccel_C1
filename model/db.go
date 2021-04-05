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

	_, err := db.database.Exec(command)

	if err != nil {
		util.Log.Panicln(err)
	}
}

// ExecuteQuery TODO: write about the behavior of this function
func (db *DB) executeQuery(command string) *sql.Rows {

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

	util.Log.Printf("Executing: %v \n", command)
	db.execute(command)
}

// RetrieveActiveWorkdays - TODO write about behavior
func (db *DB) RetrieveActiveWorkdays() map[int][]string {

	command := "CALL SELECT_ACTIVE_WORKDAY;"
	util.Log.Printf("Executing: %v \n", command)
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
func (db *DB) RetrieveCurrentMonthTimeRaport(workerID, currentMonth, currentYear string) map[string][]string {

	command := "CALL SELECT_MONTH_TIME_RAPORT(" + workerID + ", " + currentMonth + ", " + currentYear + ");"
	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	table := make(map[string][]string)
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
		table[strconv.Itoa(id)] = []string{geNo.String, roNo.String, description.String, start.String, stop.String, minutes.String}
		id++
	}
	return table
}

// RetrieveWorkerStatus - TODO write about behavior
func (db *DB) RetrieveWorkerStatus(id string) (string, string) {

	var command string = "CALL SELECT_WORKER_STATUS('" + id + "');"

	util.Log.Printf("Executing: %v \n", command)
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

	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	projects := make([]Project, 0)

	var (
		projID   sql.NullString
		geNo     sql.NullString
		roNo     sql.NullString
		descript sql.NullString
		devID    sql.NullString
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
			ID:          projID.String,
			GeNumber:    geNo.String,
			RoNumber:    roNo.String,
			Description: descript.String,
			DeviceID:    devID.String,
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

	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	var (
		workers = make([]Worker, 0)

		wID       sql.NullString
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
			ID:         wID.String,
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

	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	var (
		wID       sql.NullString
		wFirstN   sql.NullString
		wLastN    sql.NullString
		wCardNo   sql.NullString
		wPos      sql.NullString
		wIsActive sql.NullBool
		wNick     sql.NullString
		wPass     sql.NullString
		wAccess   sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(
			&wID, &wFirstN, &wLastN, &wCardNo, &wPos, &wIsActive, &wNick, &wPass, &wAccess); err != nil {
			util.Log.Panicln(err)
		}
	}
	return Worker{
		ID:          wID.String,
		FirstName:   wFirstN.String,
		LastName:    wLastN.String,
		CardNumber:  wCardNo.String,
		Position:    wPos.String,
		IsActive:    wIsActive.Bool,
		Nickname:    wNick.String,
		Password:    wPass.String,
		AccessLevel: wAccess.String,
	}
}

// RetrieveWorkerName returns worker's name based on id.
func (db *DB) RetrieveWorkerName(id string) string {
	var (
		firstName sql.NullString
		lastName  sql.NullString
	)
	command := "SELECT FIRSTNAME, LASTNAME FROM WORKER WHERE ID = '" + id + "';"

	util.Log.Printf("Executing: %v \n", command)
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

	command := "SELECT * FROM FREEDAYS;"

	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	var id sql.NullInt32
	var date sql.NullString

	table := make(map[int]string)

	for rows.Next() {
		if err := rows.Scan(&id, &date); err != nil {
			util.Log.Panicln(err)
		}
		table[int(id.Int32)] = date.String
	}
	return table
}

// RetrieveOvertime ...
func (db *DB) RetrieveOvertime(workerID string) string {

	command := "CALL GET_OVERTIME('" + workerID + "');"

	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	var overtime sql.NullString

	for rows.Next() {
		if err := rows.Scan(&overtime); err != nil {
			util.Log.Panicln(err)
		}
	}
	return overtime.String
}

func (db *DB) AddWorkday(workerID, projectID string, startHour, stopHour string) {

	var command strings.Builder
	command.WriteString("CALL ADD_NEW_WORKDAY('")
	command.WriteString(workerID)
	command.WriteString("', '")
	command.WriteString(projectID)
	command.WriteString("', '")
	command.WriteString(startHour)
	command.WriteString("', '")
	command.WriteString(stopHour)
	command.WriteString("');")

	util.Log.Printf("Executing: %v \n", command.String())
	db.execute(command.String())
}

func (db *DB) AddProject(project Project) {
	var command strings.Builder

	command.WriteString("CALL ADD_NEW_PROJECT('")
	command.WriteString(project.GeNumber)
	command.WriteString("', '")
	command.WriteString(project.RoNumber)
	command.WriteString("', '")
	command.WriteString(project.Description)
	command.WriteString("', '")
	command.WriteString(project.Begin)
	command.WriteString("');")

	util.Log.Printf("Executing: %v \n", command.String())
	db.execute(command.String())
}

func (db *DB) RetrieveAllPositions() map[int]string {

	command := ("CALL GET_ALL_POSITIONS();")
	positions := make(map[int]string)

	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

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

func (db *DB) AddWorker(worker Worker) {
	var command strings.Builder

	command.WriteString("CALL ADD_NEW_WORKER('")
	command.WriteString(worker.FirstName)
	command.WriteString("', '")
	command.WriteString(worker.LastName)
	command.WriteString("', '")
	command.WriteString(worker.CardNumber)
	command.WriteString("', '")
	command.WriteString(worker.Position)
	command.WriteString("', '")
	command.WriteString(worker.Nickname)
	command.WriteString("', '")
	command.WriteString(worker.Password)
	command.WriteString("', '")
	command.WriteString(worker.AccessLevel)
	command.WriteString("');")

	util.Log.Printf("Executing: %v \n", command.String())
	db.execute(command.String())
}

func (db *DB) GetProject(projectID string) Project {
	var (
		command strings.Builder
		rows    *sql.Rows

		id     sql.NullString
		geNo   sql.NullString
		roNo   sql.NullString
		desc   sql.NullString
		devID  sql.NullString
		active sql.NullBool
		begin  sql.NullString
		end    sql.NullString
	)
	command.WriteString("SELECT * FROM PROJECT WHERE ID='")
	command.WriteString(projectID)
	command.WriteString("';")

	util.Log.Printf("Executing: %v \n", command.String())
	rows = db.executeQuery(command.String())
	for rows.Next() {
		if err := rows.Scan(&id, &geNo, &roNo, &desc, &devID, &active, &begin, &end); err != nil {
			util.Log.Println(err)
		}
	}
	return Project{
		ID:          id.String,
		GeNumber:    geNo.String,
		RoNumber:    roNo.String,
		Description: desc.String,
		IPAddress:   "",
		DeviceID:    devID.String,
		IsActive:    active.Bool,
		Begin:       strings.Split(begin.String, " ")[0],
		End:         strings.Split(end.String, " ")[0],
	}
}

func (db *DB) UpdateProject(project Project) {

	var command strings.Builder

	command.WriteString("UPDATE PROJECT SET ")
	command.WriteString("GENUMBER='")
	command.WriteString(project.GeNumber)
	command.WriteString("', RONUMBER='")
	command.WriteString(project.RoNumber)
	command.WriteString("', DESCRIPTION='")
	command.WriteString(project.Description)
	command.WriteString("', DEVICEID='")
	command.WriteString(project.DeviceID)
	command.WriteString("', ACTIVE=")
	command.WriteString(strconv.FormatBool(project.IsActive))
	command.WriteString(", BEGIN=DATE('")
	command.WriteString(project.Begin)
	command.WriteString("'), END=DATE(")
	if project.End == "" {
		command.WriteString("NULL")
	} else {
		command.WriteString("'")
		command.WriteString(project.End)
		command.WriteString("'")
	}
	command.WriteString(") WHERE ID=")
	command.WriteString(project.ID)
	command.WriteString(";")

	util.Log.Printf("Executing: %v \n", command.String())
	db.execute(command.String())
}

func (db *DB) GetWorker(workerID string) Worker {
	var (
		command string = "SELECT * FROM WORKER WHERE ID = '" + workerID + "';"

		id     sql.NullString
		fName  sql.NullString
		lName  sql.NullString
		cardNo sql.NullString
		posID  sql.NullString
		active sql.NullString
		nick   sql.NullString
		pass   sql.NullString
		lvl    sql.NullString
	)
	util.Log.Printf("Executing: %v \n", command)
	rows := db.executeQuery(command)

	for rows.Next() {
		if err := rows.Scan(&id, &fName, &lName, &cardNo, &posID, &active, &nick, &pass, &lvl); err != nil {
			util.Log.Println(err)
		}
	}

	return Worker{
		ID:         id.String,
		FirstName:  fName.String,
		LastName:   lName.String,
		CardNumber: cardNo.String,
		Position:   posID.String,
		IsActive: func() bool {
			var IsActive bool
			var err error
			if IsActive, err = strconv.ParseBool(active.String); err != nil {
				util.Log.Println(err)
			}
			return IsActive
		}(),
		Nickname:    nick.String,
		Password:    pass.String,
		AccessLevel: lvl.String,
	}
}

func (db *DB) UpdateWorker(worker Worker) {
	var command strings.Builder

	command.WriteString("UPDATE WORKER SET FIRSTNAME='")
	command.WriteString(worker.FirstName)
	command.WriteString("', LASTNAME='")
	command.WriteString(worker.LastName)
	command.WriteString("', CARDNUMBER='")
	command.WriteString(worker.CardNumber)
	command.WriteString("', POSITIONID='")
	command.WriteString(worker.Position)
	command.WriteString("', ISACTIVE=")
	command.WriteString(strconv.FormatBool(worker.IsActive))
	command.WriteString(", NICKNAME='")
	command.WriteString(worker.Nickname)
	command.WriteString("', PASSWORD='")
	command.WriteString(worker.Password)
	command.WriteString("', ACCESSLEVEL='")
	command.WriteString(worker.AccessLevel)
	command.WriteString("' WHERE ID='")
	command.WriteString(worker.ID)
	command.WriteString("';")

	util.Log.Printf("Executing: %v \n", command.String())
	db.execute(command.String())
}
