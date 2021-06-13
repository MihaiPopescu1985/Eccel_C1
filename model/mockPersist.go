package model

import (
	"errors"
	"log"
)

type MockPersist struct {
	isConnected bool
	workers     []Worker
	workdays    []Workday
	projects    []Project
	freeDays    []string
}

func (m *MockPersist) Init(file interface{}) error {
	log.Println("initialising mock persist")
	m.workers = make([]Worker, 0, 100)
	m.workdays = make([]Workday, 0, 100)
	m.projects = make([]Project, 0, 100)
	m.freeDays = make([]string, 0, 10)
	return nil
}

func (m *MockPersist) Connect() error {
	log.Println("connecting to mock persist")
	m.isConnected = true
	return nil
}

func (m *MockPersist) IsConnected() error {
	if !m.isConnected {
		return errors.New("mock persist is not connected")
	}
	return nil
}

func (m *MockPersist) InsertIntoWorkday(deviceName, cardUID string) error {
	log.Printf("inserting device: %v & card: %v into database\n", deviceName, cardUID)
	return nil
}

func (m *MockPersist) RetrieveActiveWorkdays() (map[int][]string, error) {
	return nil, nil
}

func (m *MockPersist) RetrieveCurrentMonthTimeRaport(workerID, currentMonth, currentYear string) ([][]string, error) {
	log.Println("this function must be implementer in order to be used")
	return nil, nil
}

func (m *MockPersist) RetrieveWorkerStatus(id string) (string, string, error) {
	log.Println("this function must be implemented in order to be used")
	return "", "", nil
}

func (m *MockPersist) RetrieveActiveProjects() ([]Project, error) {
	return m.projects, nil
}

func (m *MockPersist) RetrieveAllWorkers() ([]Worker, error) {
	return m.workers, nil
}

func (m *MockPersist) GetUserByNameAndPassword(name, password string) (*Worker, error) {
	for _, w := range m.workers {
		if w.Nickname == name && w.Password == password {
			return &w, nil
		}
	}
	return nil, errors.New("worker not found")
}

func (m *MockPersist) RetrieveWorkerName(id string) (string, error) {
	log.Println("implement this function first")
	return "", nil
}

func (m *MockPersist) RetrieveFreeDays() ([]string, error) {
	return m.freeDays, nil
}

func (m *MockPersist) RetrieveMinutesOvertime(workerID string) (string, error) {
	log.Println("implement this function first")
	return "", nil
}

func (m *MockPersist) AddWorkday(workerID, projectID string, startHour, stopHour string) error {
	log.Println("implement function before using it")
	return errors.New("implement function")
}

func (m *MockPersist) AddProject(project Project) error {
	log.Println("implement function before using it")
	return errors.New("implement function")
}

func (m *MockPersist) RetrieveAllPositions() (map[int]string, error) {
	log.Println("implement function before using it")
	return nil, errors.New("implement function")
}

func (m *MockPersist) AddWorker(worker Worker) error {
	m.workers = append(m.workers, worker)
	return nil
}

func (m *MockPersist) GetProject(projectID string) (*Project, error) {
	log.Println("implement function before using it")
	return nil, errors.New("implement function")
}

func (m *MockPersist) UpdateProject(project Project) error {
	log.Println("implement function before using it")
	return errors.New("implement function")
}

func (m *MockPersist) GetWorker(workerID string) (*Worker, error) {
	log.Println("implement function before using it")
	return nil, errors.New("implement function")
}

func (m *MockPersist) UpdateWorker(worker Worker) error {
	log.Println("implement function before using it")
	return errors.New("implement function")
}

func (m *MockPersist) RetrieveSentProjects() (map[Project]string, error) {
	log.Println("implement function before using it")
	return nil, errors.New("implement function")
}

func (m *MockPersist) DeleteFreeDay(freeDay string) error {
	log.Println("deleting from freedays")
	return nil
}
func (m *MockPersist) AddFreeDay(freeDay string) error {
	log.Println("implement function before using it")
	return errors.New("implement function")
}
