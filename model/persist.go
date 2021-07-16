package model

type Persist interface {
	Init(file interface{}) error
	Connect() error
	IsConnected() error
	InsertIntoWorkday(deviceName, cardUID string) error
	RetrieveActiveWorkdays() (map[int][]string, error)
	RetrieveCurrentMonthTimeRaport(workerID, currentMonth, currentYear string) ([][]string, error)
	RetrieveWorkerStatus(id string) (string, string, error)
	RetrieveActiveProjects() ([]Project, error)
	RetrieveAllWorkers() ([]Worker, error)
	GetUserByNameAndPassword(name, password string) (*Worker, error)
	GetUserByNickname(nickname string) (*Worker, error)
	RetrieveWorkerName(id string) (string, error)
	RetrieveFreeDays() ([]string, error)
	RetrieveMinutesOvertime(workerID string) (string, error)
	AddWorkday(workerID, projectID string, startHour, stopHour string) error
	AddProject(project Project) error
	RetrieveAllPositions() (map[int]string, error)
	AddWorker(worker Worker) error
	GetProject(projectID string) (*Project, error)
	UpdateProject(project Project) error
	GetWorker(workerID string) (*Worker, error)
	UpdateWorker(worker Worker) error
	RetrieveSentProjects() (map[Project]string, error)
	DeleteFreeDay(freeDay string) error
	AddFreeDay(freeDay string) error
	RetrieveActiveWorkers() (map[string][]string, error)
	GetTodayBreak(workerID string) (string, error)
}
