package processes

import (
	"github.com/shirou/gopsutil/process"
	"log"
)

type Processes struct {
	ProcessList []*process.Process
}

func AddProcessFromPid(p *Processes, pid int32) {
	process, err := process.NewProcess(pid)

	if err != nil {
		log.Fatal("Unable to create process object from pid: %s", err)
	}

	p.ProcessList = append(p.ProcessList, process)
}

// Returns a list of currently running processes
func New() *Processes {
	var processes Processes

	pids, err := process.Pids()

	if err != nil {
		log.Fatalf("Unable to get PIDs: %s", err)
	}

	for _, pid := range pids {
		AddProcessFromPid(&processes, pid)
	}

	return &processes
}
