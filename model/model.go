package model

import (
	"time"

	"example.com/c1/c1device"
)

// Worker ...
type Worker struct {
	ID         int
	FirstName  string
	LastName   string
	Position   string
	CardNumber string
	IsActive   bool
	Nickname   string
	Password   string
}

// Project ...
type Project struct {
	ID          int
	GeNumber    string
	RoNumber    string
	Description string
	IPAddress   string
	Device      c1device.C1Device
	IsActive    bool
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
