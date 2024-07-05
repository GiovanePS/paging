package process

import (
	"fmt"
	"math/rand"
)

type Process struct {
	Pid            int
	Size           int
	LogicalMemory  []byte
	PageTableEntry []int
	Next           *Process
}

var HeadProcess *Process

func CreateProcess(pid int, size int) error {
	if err := findProcess(pid); err != nil {
		return err
	}

	newProcess := &Process{pid, size, nil, nil, nil}
	initLocalMemory(newProcess)

	return nil
}

func initLocalMemory(process *Process) {
	process.LogicalMemory = make([]byte, process.Size)

	for i := 0; i < process.Size; i++ {
		randomByte := rand.Intn(25) + 65
		process.LogicalMemory[i] = byte(randomByte)
	}
}

func findProcess(pid int) (err error) {
	if HeadProcess == nil {
		return nil
	}

	processCreated := false

	cursor := HeadProcess
	for cursor.Next != nil {
		if cursor.Pid == pid {
			processCreated = true
			break
		}

		cursor = cursor.Next
	}

	if cursor.Next == nil && cursor.Pid != pid {
		return nil
	} else {
		processCreated = true
	}

	if processCreated {
		err = fmt.Errorf("Process with PID %d already created.\n", pid)
		return err
	}

	return nil
}
