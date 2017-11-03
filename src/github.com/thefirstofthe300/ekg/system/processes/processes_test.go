package processes

import (
	"github.com/shirou/gopsutil/process"
	"os"
	"reflect"
	"testing"
)

func SliceContains(slice []*process.Process, p *process.Process) bool {
	for i, _ := range slice {
		if reflect.DeepEqual(slice[i], p) {
			return true
		}
	}
	return false
}

func TestAddProcessFromPid(t *testing.T) {
	// gopsutil's NewProcess function will look use the HOST_PROC environmental
	// variable as its proc directory. Here we utilize to make testing easier
	host_proc := "testdata"
	os.Setenv("HOST_PROC", host_proc)
	defer os.Setenv("HOST_PROC", "")

	var testProcesses Processes

	var testList = []*process.Process{
		&process.Process{Pid: 1},
		&process.Process{Pid: 5},
		&process.Process{Pid: 14},
		&process.Process{Pid: 18},
		&process.Process{Pid: 41},
	}

	AddProcessFromPid(&testProcesses, 1)
	AddProcessFromPid(&testProcesses, 5)
	AddProcessFromPid(&testProcesses, 14)
	AddProcessFromPid(&testProcesses, 18)
	AddProcessFromPid(&testProcesses, 41)

	if reflect.DeepEqual(testProcesses.ProcessList, testList) == false {
		t.Fatal("Process lists do no match:", testProcesses.ProcessList, testList)
	}
}

func TestNewProcesses(t *testing.T) {
	host_proc := "testdata"
	os.Setenv("HOST_PROC", host_proc)
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

	testProcesses := New()

	if len(expectedProcesses.ProcessList) == 0 && len(testProcesses.ProcessList) == 0 {
		return
	}

	if len(expectedProcesses.ProcessList) == 0 || len(testProcesses.ProcessList) == 0 {
		t.Fatal("The test processes", testProcesses, "did not meet the expected processes", expectedProcesses)
	}

	if len(expectedProcesses.ProcessList) != len(testProcesses.ProcessList) {
		t.Fatal("The test processes", testProcesses, "did not meet the expected processes", expectedProcesses)
	}

	for i, _ := range expectedProcesses.ProcessList {
		if SliceContains(testProcesses.ProcessList, expectedProcesses.ProcessList[i]) == false {
			t.Fatal("The test processes", testProcesses, "did not meet the expected processes", expectedProcesses)
		}
	}
}
