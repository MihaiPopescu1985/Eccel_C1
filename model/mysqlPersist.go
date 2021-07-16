package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
)

const settingsFile = "model/settings.json"

// MysqlDB
type MysqlDB struct {
	database     *sql.DB
	settings     dbSettings
	settingsFile string
}

type dbSettings struct {
	driver   string
	settings string
}

// Init is setting up the file where database connection settings are.
// If no file is provided, the mode/settings.json file is used.
// It is done so for testing.
func (db *MysqlDB) Init(file interface{}) error {
	if file == "" {
		db.settingsFile = settingsFile
	} else {
		db.settingsFile = fmt.Sprint(file)
	}
	return db.getDBSettings()
}

func (db *MysqlDB) getDBSettings() error {
	type sett struct {
		Driver   string
		User     string
		Password string
		URL      string
		Name     string
	}

	fileSettings, err := db.readSettingsFromFile()
	if err != nil {
		return err
	}
	var settings sett

	if err := json.Unmarshal([]byte(fileSettings), &settings); err != nil {
		return err
	}
	db.settings.driver = settings.Driver
	db.settings.settings = settings.User + ":" + settings.Password + "@" + settings.URL + settings.Name
	return err
}

func (db *MysqlDB) readSettingsFromFile() (string, error) {
	settings, err := ioutil.ReadFile(db.settingsFile)
	if err != nil {
		return "", err
	}
	return string(settings), nil
}

// Connect connects to database.
// The Open method is not enough to verify a
// database connection so a Ping method should be
// called. The Ping method is embedded inside
// IsConnected method.
func (db *MysqlDB) Connect() error {

	var err error = nil
	db.database, err = sql.Open(db.settings.driver, db.settings.settings)

	if err != nil {
		return err
	}
	return db.IsConnected()
}

// IsConnected verify database connection.
func (db *MysqlDB) IsConnected() error {
	return db.database.Ping()
}

// Execute executes a command against database with no returning result set.
func (db *MysqlDB) execute(command string) error {

	_, err := db.database.Exec(command)
	return err
}

// ExecuteQuery TODO: write about the behavior of this function
func (db *MysqlDB) executeQuery(command string) (*sql.Rows, error) {

	rows, err := db.database.Query(command)

	// Close rows after 5 seconds
	go closeRows(rows, time.Second*5)

	return rows, err
}

// Helper function used for closing query rows after a certain time.
// Needed because database server will keep open a conection for every query.
func closeRows(rows *sql.Rows, timeSpan time.Duration) {
	defer func() {
		if recover := recover(); recover != nil {
			log.Println(recover)
		}
	}()
	time.Sleep(timeSpan)

	if err := rows.Close(); err != nil {
		log.Println(err)
	}
}

// InsertIntoWorkday returns the sql command to start/stop time on workday.
func (db *MysqlDB) InsertIntoWorkday(deviceName, cardUID string) error {
	command := "CALL INSERT_INTO_WORKDAY(\"" + deviceName + "\", \"" + cardUID + "\");"

	log.Printf("Executing: %v \n", command)
	return db.execute(command)
}

// RetrieveActiveWorkdays - TODO write about behavior
func (db *MysqlDB) RetrieveActiveWorkdays() (map[int][]string, error) {

	command := "CALL SELECT_ACTIVE_WORKDAY;"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

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
			return nil, err
		}
		table[int(id.Int32)] = []string{worker.String, roNumber.String, geNumber.String, description.String}
	}
	return table, nil
}

// RetrieveCurrentMonthTimeRaport - TODO write about behavior
func (db *MysqlDB) RetrieveCurrentMonthTimeRaport(workerID, currentMonth, currentYear string) ([][]string, error) {

	command := "CALL SELECT_MONTH_TIME_RAPORT(" + workerID + ", " + currentMonth + ", " + currentYear + ");"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	table := make([][]string, 0)
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
			return nil, err
		}
		table = append(table, []string{geNo.String, roNo.String, description.String, start.String, stop.String, minutes.String})
	}
	return table, nil
}

// RetrieveWorkerStatus - TODO write about behavior
func (db *MysqlDB) RetrieveWorkerStatus(id string) (string, string, error) {

	var command string = "CALL SELECT_WORKER_STATUS('" + id + "');"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return "", "", err
	}

	var status sql.NullString       //string
	var workedMinutes sql.NullInt32 //int

	rows.Next()
	if err := rows.Scan(&status); err != nil {
		return "", "", err
	}
	rows.Next()
	if err := rows.Scan(&workedMinutes); err != nil {
		return "", "", err
	}

	workedTime := strconv.Itoa(int(workedMinutes.Int32))
	return status.String, workedTime, nil
}

// RetrieveActiveProjects ...
func (db *MysqlDB) RetrieveActiveProjects() ([]Project, error) {

	var command string = "CALL GET_ACTIVE_PROJECTS();"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

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
			return nil, err
		}

		projects = append(projects, Project{
			ID:          projID.String,
			GeNumber:    geNo.String,
			RoNumber:    roNo.String,
			Description: descript.String,
			DeviceID:    devID.String,
			IsActive:    isActive.Bool,
			Begin:       strings.Split(begin.String, " ")[0],
			End:         strings.Split(end.String, " ")[0],
		})
	}
	return projects, nil
}

// RetrieveAllWorkers ...
func (db *MysqlDB) RetrieveAllWorkers() ([]Worker, error) {

	command := "CALL GET_ALL_WORKERS();"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

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
			return nil, err
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
	return workers, nil
}

// GetUserByNameAndPassword TODO: write about function
func (db *MysqlDB) GetUserByNameAndPassword(name, password string) (*Worker, error) {

	command := "SELECT * FROM WORKER WHERE NICKNAME = \"" + name + "\" AND PASSWORD = \"" + password + "\";"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	var (
		wID            sql.NullString
		wFirstN        sql.NullString
		wLastN         sql.NullString
		wCardNo        sql.NullString
		wPos           sql.NullString
		wIsActive      sql.NullBool
		wNick          sql.NullString
		wPass          sql.NullString
		wAccess        sql.NullString
		wHire          sql.NullString
		wCloseContract sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(
			&wID, &wFirstN, &wLastN, &wCardNo, &wPos, &wIsActive, &wNick, &wPass, &wAccess, &wHire, &wCloseContract); err != nil {
			return nil, err
		}
	}

	if wID.String == "" {
		return nil, errors.New("user not found")
	}

	return &Worker{
		ID:                wID.String,
		FirstName:         wFirstN.String,
		LastName:          wLastN.String,
		CardNumber:        wCardNo.String,
		Position:          wPos.String,
		IsActive:          wIsActive.Bool,
		Nickname:          wNick.String,
		Password:          wPass.String,
		AccessLevel:       wAccess.String,
		HireDate:          wHire.String,
		CloseContractDate: wCloseContract.String,
	}, nil
}

func (db *MysqlDB) GetUserByNickname(nickname string) (*Worker, error) {
	command := "SELECT * FROM WORKER WHERE NICKNAME = \"" + nickname + "\";"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	var (
		wID            sql.NullString
		wFirstN        sql.NullString
		wLastN         sql.NullString
		wCardNo        sql.NullString
		wPos           sql.NullString
		wIsActive      sql.NullBool
		wNick          sql.NullString
		wPass          sql.NullString
		wAccess        sql.NullString
		wHire          sql.NullString
		wCloseContract sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(
			&wID, &wFirstN, &wLastN, &wCardNo, &wPos, &wIsActive, &wNick, &wPass, &wAccess, &wHire, &wCloseContract); err != nil {
			return nil, err
		}
	}

	if wID.String == "" {
		return nil, errors.New("user not found")
	}

	return &Worker{
		ID:                wID.String,
		FirstName:         wFirstN.String,
		LastName:          wLastN.String,
		CardNumber:        wCardNo.String,
		Position:          wPos.String,
		IsActive:          wIsActive.Bool,
		Nickname:          wNick.String,
		Password:          wPass.String,
		AccessLevel:       wAccess.String,
		HireDate:          wHire.String,
		CloseContractDate: wCloseContract.String,
	}, nil
}

// RetrieveWorkerName returns worker's name based on id.
func (db *MysqlDB) RetrieveWorkerName(id string) (string, error) {
	var (
		firstName sql.NullString
		lastName  sql.NullString
	)
	command := "SELECT FIRSTNAME, LASTNAME FROM WORKER WHERE ID = '" + id + "';"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		if err := rows.Scan(&firstName, &lastName); err != nil {
			return "", err
		}
	}
	return firstName.String + " " + lastName.String, nil
}

// RetrieveFreeDays returns a map containing free days.
func (db *MysqlDB) RetrieveFreeDays() ([]string, error) {

	command := "SELECT * FROM FREEDAYS ORDER BY DATE ASC;"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	var id sql.NullString
	var date sql.NullString

	table := make([]string, 0)

	for rows.Next() {
		if err := rows.Scan(&id, &date); err != nil {
			return nil, err
		}
		table = append(table, date.String)
	}
	return table, nil
}

// RetrieveOvertime ...
func (db *MysqlDB) RetrieveMinutesOvertime(workerID string) (string, error) {

	command := "CALL GET_OVERTIME('" + workerID + "');"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		log.Println(err)
		return "0", nil
	}

	var overtime sql.NullString

	for rows.Next() {
		if err := rows.Scan(&overtime); err != nil {
			log.Println(err)
			return "0", nil
		}
	}
	return overtime.String, nil
}

func (db *MysqlDB) AddWorkday(workerID, projectID string, startHour, stopHour string) error {

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

	log.Printf("Executing: %v \n", command.String())
	return db.execute(command.String())
}

func (db *MysqlDB) AddProject(project Project) error {
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

	log.Printf("Executing: %v \n", command.String())
	return db.execute(command.String())
}

func (db *MysqlDB) RetrieveAllPositions() (map[int]string, error) {

	command := ("CALL GET_ALL_POSITIONS();")
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}
	positions := make(map[int]string)

	for rows.Next() {
		var pos sql.NullString
		var id sql.NullInt32

		if err := rows.Scan(&id, &pos); err != nil {
			return nil, err
		}
		positions[int(id.Int32)] = pos.String
	}
	return positions, nil
}

func (db *MysqlDB) AddWorker(worker Worker) error {
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

	log.Printf("Executing: %v \n", command.String())
	return db.execute(command.String())
}

func (db *MysqlDB) GetProject(projectID string) (*Project, error) {
	var (
		command strings.Builder

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

	log.Printf("Executing: %v \n", command.String())
	rows, err := db.executeQuery(command.String())
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&id, &geNo, &roNo, &desc, &devID, &active, &begin, &end); err != nil {
			return nil, err
		}
	}
	return &Project{
		ID:          id.String,
		GeNumber:    geNo.String,
		RoNumber:    roNo.String,
		Description: desc.String,
		IPAddress:   "",
		DeviceID:    devID.String,
		IsActive:    active.Bool,
		Begin:       strings.Split(begin.String, " ")[0],
		End:         strings.Split(end.String, " ")[0],
	}, nil
}

func (db *MysqlDB) UpdateProject(project Project) error {

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

	log.Printf("Executing: %v \n", command.String())
	return db.execute(command.String())
}

func (db *MysqlDB) GetWorker(workerID string) (*Worker, error) {
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
		hire   sql.NullString
		close  sql.NullString
	)
	log.Printf("Executing: %v \n", command)
	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &fName, &lName, &cardNo, &posID, &active, &nick, &pass, &lvl, &hire, &close); err != nil {
			return nil, err
		}
	}

	if id.String == "" {
		return nil, errors.New("user not found")
	}

	return &Worker{
		ID:         id.String,
		FirstName:  fName.String,
		LastName:   lName.String,
		CardNumber: cardNo.String,
		Position:   posID.String,
		IsActive: func() bool {
			var IsActive bool
			var err error
			if IsActive, err = strconv.ParseBool(active.String); err != nil {
				log.Println(err)
			}
			return IsActive
		}(),
		Nickname:          nick.String,
		Password:          pass.String,
		AccessLevel:       lvl.String,
		HireDate:          hire.String,
		CloseContractDate: close.String,
	}, nil
}

func (db *MysqlDB) UpdateWorker(worker Worker) error {
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

	log.Printf("Executing: %v \n", command.String())
	return db.execute(command.String())
}

func (db *MysqlDB) RetrieveSentProjects() (map[Project]string, error) {

	command := "CALL GET_DELIVERED_PROJECTS;"
	log.Printf("Executing: %v \n", command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	var (
		projects = make(map[Project]string)
		id       sql.NullString
		geNo     sql.NullString
		roNo     sql.NullString
		desc     sql.NullString
		devID    sql.NullString
		active   sql.NullString
		begin    sql.NullString
		end      sql.NullString
		wMin     sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(&id, &geNo, &roNo, &desc, &devID, &active, &begin, &end, &wMin); err != nil {
			return nil, err
		}
		projects[Project{
			ID:          id.String,
			GeNumber:    geNo.String,
			RoNumber:    roNo.String,
			Description: desc.String,
			IPAddress:   "",
			DeviceID:    devID.String,
			IsActive: func() bool {
				isActive, err := strconv.ParseBool(active.String)
				if err != nil {
					log.Println(err)
				}
				return isActive
			}(),
			Begin: begin.String,
			End:   end.String,
		}] = wMin.String
	}
	return projects, nil
}

func (db *MysqlDB) DeleteFreeDay(freeDay string) error {
	command := "DELETE FROM FREEDAYS WHERE DATE = '" + freeDay + "';"
	log.Println(command)
	return db.execute(command)
}

func (db *MysqlDB) AddFreeDay(freeDay string) error {
	command := "INSERT INTO FREEDAYS (DATE) VALUES ('" + freeDay + "');"
	log.Println(command)
	return db.execute(command)
}

func (db *MysqlDB) RetrieveActiveWorkers() (map[string][]string, error) {
	command := "CALL WORKERS_DAY_RAPORT();"
	log.Println(command)

	rows, err := db.executeQuery(command)
	if err != nil {
		return nil, err
	}

	var (
		activeWorkerStatus = make(map[string][]string)
		name               sql.NullString
		project            sql.NullString
		startTime          sql.NullString
		minutes            sql.NullString
		status             sql.NullString
		overtime           sql.NullString
		breakTime          sql.NullString
	)

	for rows.Next() {
		if err := rows.Scan(&name, &project, &startTime, &minutes, &status, &overtime, &breakTime); err != nil {
			return nil, err
		}
		activeWorkerStatus[name.String] = []string{project.String, startTime.String, minutes.String, status.String, overtime.String, breakTime.String}
	}
	return activeWorkerStatus, nil
}

func (db *MysqlDB) GetTodayBreak(workerID string) (string, error) {
	command := "SELECT SUM(TIMESTAMPDIFF(MINUTE, STARTTIME, IFNULL(STOPTIME, NOW()))) AS BREAK FROM WORKDAY WHERE WORKERID=" + workerID + " AND DATE(STARTTIME)=DATE(NOW()) AND PROJECTID=1;"
	log.Println(command)
	db.executeQuery(command)

	row, err := db.executeQuery(command)
	if err != nil {
		return "", err
	}
	var breakTime sql.NullString
	row.Next()
	if err = row.Scan(&breakTime); err != nil {
		return "", err
	}
	return breakTime.String, nil
}
