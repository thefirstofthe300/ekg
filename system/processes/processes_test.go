package processes

import (
	"os"
	"reflect"
	"testing"

	"github.com/shirou/gopsutil/process"
)

func SliceContains(slice []*process.Process, p *process.Process) bool {
	for i := range slice {
		if reflect.DeepEqual(slice[i], p) {
			return true
		}
	}
	return false
}

func TestProcess_Add(t *testing.T) {
	// gopsutil's NewProcess function will look use the HOST_PROC environmental
	// variable as its proc directory. Here we utilize to make testing easier
	hostProc := "testdata"
	os.Setenv("HOST_PROC", hostProc)
	defer os.Setenv("HOST_PROC", "")

	var testProcesses Processes
	var testList = []*process.Process{
		&process.Process{Pid: 1},
		&process.Process{Pid: 5},
		&process.Process{Pid: 14},
		&process.Process{Pid: 18},
		&process.Process{Pid: 41},
	}

	for _, proc := range []int32{1, 5, 14, 18, 41, 51} {
		err := testProcesses.Add(proc)
		if err != nil && err.Error() != "unable to create process object from pid: open testdata/51: no such file or directory" {
			t.Fatalf("could not add process to process list: %s", err)
		}
	}

	if reflect.DeepEqual(testProcesses.ProcessList, testList) == false {
		t.Fatal("process lists do no match:", testProcesses.ProcessList, testList)
	}
}

func TestNew(t *testing.T) {
	hostProc := "testdata"
	os.Setenv("HOST_PROC", hostProc)
	defer os.Setenv("HOST_PROC", "")

	var expectedProcesses = &Processes{
		ProcessList: []*process.Process{
			&process.Process{Pid: 1},
			&process.Process{Pid: 5},
			&process.Process{Pid: 14},
			&process.Process{Pid: 18},
			&process.Process{Pid: 41},
		},
	}

	testProcesses, err := New()

	if err != nil {
		t.Fatalf("could not create new processes struct: %s", err)
	}

	if len(expectedProcesses.ProcessList) == 0 && len(testProcesses.ProcessList) == 0 {
		return
	}

	if len(expectedProcesses.ProcessList) == 0 || len(testProcesses.ProcessList) == 0 {
		t.Fatal("The test processes", testProcesses, "did not meet the expected processes", expectedProcesses)
	}

	if len(expectedProcesses.ProcessList) != len(testProcesses.ProcessList) {
		t.Fatal("The test processes", testProcesses, "did not meet the expected processes", expectedProcesses)
	}

	for i := range expectedProcesses.ProcessList {
		if SliceContains(testProcesses.ProcessList, expectedProcesses.ProcessList[i]) == false {
			t.Fatal("The test processes", testProcesses, "did not meet the expected processes", expectedProcesses)
		}
	}
}
