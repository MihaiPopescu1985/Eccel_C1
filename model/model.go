package model

import (
	"database/sql"
)

// Worker ...
type Worker struct {
	ID          sql.NullInt32  //int
	FirstName   sql.NullString //string
	LastName    sql.NullString //string
	CardNumber  sql.NullString //string
	Position    sql.NullString //string
	IsActive    sql.NullBool   //bool
	Nickname    sql.NullString //string
	Password    sql.NullString //string
	AccessLevel sql.NullInt32  //byte
}

// Project ...
type Project struct {
	ID          sql.NullInt32  //int
	GeNumber    sql.NullString //string
	RoNumber    sql.NullString //string
	Description sql.NullString //string
	IPAddress   sql.NullString //string
	DeviceID    sql.NullInt32  //int
	IsActive    sql.NullBool   //bool
	Begin       sql.NullTime   //time.Time
	End         sql.NullTime   //time.Time
}

// Workday ...
type Workday struct {
	ID        sql.NullInt32 //int
	Worker    Worker
	Project   Project
	StartTime sql.NullTime //time.Time
	StopTime  sql.NullTime //time.Time
}

// Device ...
type Device struct {
	ID         sql.NullInt32  //int
	Name       sql.NullString //string
	IP         sql.NullString //string
	IsEndpoint sql.NullBool   //bool
}

// ActiveWorkdays - placeholder for storing active workdays retrieved from database
type ActiveWorkdays struct {
	Workdays map[sql.NullInt32][5]sql.NullString //string
}
