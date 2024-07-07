package process

import (
	"fmt"
	"math/rand"
	"paging/memory"
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
	if err := processAlreadyExists(pid); err != nil {
		return err
	}

	newProcess := &Process{pid, size, nil, nil, nil}
	initLogicalMemory(newProcess)
	initTablePage(newProcess)
	includeProcess(newProcess)

	return nil
}

func initLogicalMemory(process *Process) {
	process.LogicalMemory = make([]byte, process.Size)

	for i := 0; i < process.Size; i++ {
		randomByte := rand.Intn(25) + 65
		process.LogicalMemory[i] = byte(randomByte)
	}
}

func initTablePage(process *Process) {
	numPages := process.Size / memory.FRAME_PAGE_SIZE

	process.PageTableEntry = make([]int, numPages)
	pageAuxiliary := make([]byte, memory.FRAME_PAGE_SIZE)
	for i := 0; i < numPages; i++ {
		for offset := 0; offset < memory.FRAME_PAGE_SIZE; offset++ {
			pageAuxiliary[offset] = process.LogicalMemory[i*memory.FRAME_PAGE_SIZE+offset]
		}
		frameAllocated := memory.AllocateFrame(pageAuxiliary)
		process.PageTableEntry[i] = frameAllocated
	}
}

func ShowTablePage(pid int) (err error) {
	process, err := findProcess(pid)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Process size: %d\n", process.Size)
	fmt.Println("Page <=> Frame")
	numPages := process.Size / memory.FRAME_PAGE_SIZE
	for i := 0; i < numPages; i++ {
		frame := process.PageTableEntry[i]
		fmt.Printf("%d    <=> %d\n", i, frame)
	}

	return err
}

func findProcess(pid int) (process *Process, err error) {
	cursor := HeadProcess

	for cursor != nil {
		if cursor.Pid == pid {
			return cursor, nil
		}
		cursor = cursor.Next
	}

	err = fmt.Errorf("Process with PID %d not found.\n", pid)
	return nil, err
}

func includeProcess(process *Process) {
	if HeadProcess == nil {
		HeadProcess = process
		return
	}

	cursor := HeadProcess
	for cursor.Next != nil {
		cursor = cursor.Next
	}

	cursor.Next = process
}

func processAlreadyExists(pid int) (err error) {
	if HeadProcess == nil {
		return nil
	}

	cursor := HeadProcess
	for cursor != nil {
		if cursor.Pid == pid {
			return fmt.Errorf("Process with PID %d already created.\n", pid)
		}

		cursor = cursor.Next
	}

	return nil
}
