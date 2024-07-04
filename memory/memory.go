package memory

import "fmt"

type Node struct {
	IdSerial int
	Next     *Node
}

var PHYSICAL_MEMORY_SIZE int
var MAX_LOGICAL_MEMORY_SIZE int
var FRAME_PAGE_SIZE int

var PhysicalMemory []byte
var HeadFreeFrames *Node
var TotalFrames int
var FreeFrames int

func InitPhysicalMemory() {
	HeadFreeFrames = &Node{0, nil}
	TotalFrames = PHYSICAL_MEMORY_SIZE / FRAME_PAGE_SIZE
	FreeFrames = TotalFrames
	PhysicalMemory = make([]byte, PHYSICAL_MEMORY_SIZE)

	cursor := HeadFreeFrames
	for i := 0; i < TotalFrames; i++ {
		nextNode := &Node{i, nil}
		cursor.Next = nextNode
		cursor = cursor.Next
	}
}

func ShowMemory() {
	fmt.Printf("Free frames: %.2f%%\n", float64(FreeFrames)/float64(TotalFrames)*100)
	fmt.Println("Positon <=> Value")
	for i := 0; i < PHYSICAL_MEMORY_SIZE/FRAME_PAGE_SIZE; i++ {
		start := FRAME_PAGE_SIZE * i
		end := FRAME_PAGE_SIZE*i + FRAME_PAGE_SIZE - 1
		fmt.Printf("%d <=> %s\n", i, PhysicalMemory[start:end])
	}
	fmt.Println()
}
