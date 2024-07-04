package system

import (
	"fmt"
	"paging/memory"
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

		// var pid int
		// var process_size int

		switch option {
		case 1:
			memory.ShowMemory()
		case 2:
			fmt.Println("Option 2")
		case 3:
			fmt.Println("Option 3")
		case 0:
			return
		}
	}
}
