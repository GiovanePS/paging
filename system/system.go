package system

import "fmt"

var PHYSICAL_MEMORY_SIZE int

func Init_system() {
	invalid_values := true

	for invalid_values {
		fmt.Println("Defina a physical memory size: ")
		fmt.Scan(&PHYSICAL_MEMORY_SIZE)
		fmt.Print(PHYSICAL_MEMORY_SIZE)
	}
}
