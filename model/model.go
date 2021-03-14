package model

import (
	"time"
)

// Worker ...
type Worker struct {
	ID          int
	FirstName   string
	LastName    string
	CardNumber  string
	Position    string
	IsActive    bool
	Nickname    string
	Password    string
	AccessLevel int
}

// Project ...
type Project struct {
	ID          int
	GeNumber    string
	RoNumber    string
	Description string
	IPAddress   string
	DeviceID    int
	IsActive    bool
	Begin       string
	End         string
}

// Workday ...
type Workday struct {
	ID        int
	Worker    Worker
	Project   Project
	StartTime time.Time
	StopTime  time.Time
}

// Device ...
type Device struct {
	ID         int
	Name       string
	IP         string
	IsEndpoint bool
}

// ActiveWorkdays - placeholder for storing active workdays retrieved from database
type ActiveWorkdays struct {
	Workdays map[int][5]string
}
