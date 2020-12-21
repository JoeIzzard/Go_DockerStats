package dockerstats

// ----- Internal Structs -----
type rawStatus struct {
	Name string
	ID string
	CPUPerc string
	MemUsage string
	NetIO string
	BlockIO string
	MemPerc string
	PIDs int
}

type rawState struct {
	Status string
	Running bool
	Paused bool
	Restarting bool
	OOMKilled bool
	Dead bool
	Pid int
	ExitCode int
	Error string
	StartedAt string
	FinishedAt string
}

type rawHealth struct {
	Status string
	FailingStreak int
	Log []LogEntry
}

// ----- Exported Structs -----
type Status struct {
	Name string
	ID string
	CPU float64
	Memory Memory
	NetIO IO
	BlockIO IO
	PIDs int
	State State
	ExitCode int
	PID int
	StartedAt string
	FinishedAt string
	Health Health
}

type State struct {
	Message string
	Running bool
	Paused bool
	Restarting bool
	OOMKilled bool
	Dead bool
}

type Health struct {
	Enabled bool
	FailStreak int
	Status string
	Log []LogEntry
}

type Memory struct {
	Using Data
	Limit Data
	Percent float64
}

type IO struct {
	In Data
	Out Data
}

type Data struct {
	Value float64
	Base string
}

type LogEntry struct {
	StartedAt string
	FinishedAt string
	ExitCode int
	Output string
}