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
	AccessLevel byte
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
	Begin       time.Time
	End         time.Time
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
