package system

import (
	"fmt"
	"paging/memory"
	"paging/process"
	"paging/utils"
)

func InitSystem() {
	for {
		fmt.Print("Define a physical memory size: ")
		fmt.Scan(&memory.PHYSICAL_MEMORY_SIZE)
		fmt.Print("Define a max logical memory size: ")
		fmt.Scan(&memory.MAX_LOGICAL_MEMORY_SIZE)
		fmt.Print("Define a frame/page size: ")
		fmt.Scan(&memory.FRAME_PAGE_SIZE)

		if !utils.IsPowerOfTwo(memory.PHYSICAL_MEMORY_SIZE) && !utils.IsPowerOfTwo(memory.MAX_LOGICAL_MEMORY_SIZE) && !utils.IsPowerOfTwo(memory.FRAME_PAGE_SIZE) {
			fmt.Println("The values must be power of two.")
			continue
		}

		if memory.FRAME_PAGE_SIZE > memory.PHYSICAL_MEMORY_SIZE {
			fmt.Println("Frame/Page size can't have a size larger than physical memory size.")
			continue
		}

		if memory.FRAME_PAGE_SIZE > memory.MAX_LOGICAL_MEMORY_SIZE {
			fmt.Println("Frame/Page size can't have a size larger than max logical memory size")
		}

		break
	}

	memory.InitPhysicalMemory()

	for {
		fmt.Println("[1] Show memory.")
		fmt.Println("[2] Create process.")
		fmt.Println("[3] Show page table.")
		fmt.Println("[0] Exit.")

		var option int

		fmt.Print("Option: ")
		fmt.Scan(&option)

		var pid int
		var process_size int

		switch option {
		case 1:
			memory.ShowMemory()
		case 2:
			fmt.Print("Enter a PID number: ")
			fmt.Scan(&pid)
			fmt.Print("Enter a size to the process: ")
			fmt.Scan(&process_size)

			if !utils.IsPowerOfTwo(process_size) {
				fmt.Println("The process must be a power of two.")
				break
			}

			if process_size > memory.MAX_LOGICAL_MEMORY_SIZE {
				fmt.Printf("The process sizes entered exceed memory limit! Enter a sizes lower than %d bytes.\n", process_size)
				break
			}

			if process_size < memory.FRAME_PAGE_SIZE {
				fmt.Println("The process size can't be lower than Frame/Page size.")
				break
			}

			if memory.FreeFrames < process_size/memory.FRAME_PAGE_SIZE {
				fmt.Println("There is not enough memory to allocate a process with that size.")
				break
			}

			err := process.CreateProcess(pid, process_size)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			fmt.Println("Process created!")
		case 3:
			fmt.Println("Option 3")
		case 0:
			return
		}
	}
}
