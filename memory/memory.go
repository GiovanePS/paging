package memory

import (
	"fmt"
	"math/rand"
)

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

func AllocateFrame(page []byte) int {
	previous := getSomeFrameToAllocate()
	var cursor *Node

	if previous.Next == nil {
		cursor = previous
	} else {
		cursor = previous.Next
	}

	frameStart := cursor.IdSerial * FRAME_PAGE_SIZE

	for offset := 0; offset < FRAME_PAGE_SIZE; offset++ {
		PhysicalMemory[frameStart+offset] = page[offset]
	}

	FreeFrames--
	frameAllocated := cursor.IdSerial
	if cursor.Next != nil {
		previous.Next = cursor.Next
	} else {
		previous.Next = nil
	}

	return frameAllocated
}

func getSomeFrameToAllocate() *Node {
	if float64(FreeFrames)/float64(TotalFrames) >= 0.15 {
		// While runs until draw a free frame
		for {
			frame := rand.Intn(FreeFrames)
			cursor := HeadFreeFrames
			if cursor.IdSerial == frame {
				HeadFreeFrames = HeadFreeFrames.Next
				cursor.Next = nil
				return cursor
			}

			for cursor.Next != nil {
				if cursor.Next.IdSerial >= frame {
					// Returning the previous frame, to delete the choosed frame
					return cursor
				}

				cursor = cursor.Next
			}
		}

	} else {
		cursor := HeadFreeFrames
		HeadFreeFrames = HeadFreeFrames.Next
		cursor.Next = nil
		return cursor
	}
}

func ShowMemory() {
	fmt.Printf("Free frames: %.2f%%\n", float64(FreeFrames)/float64(TotalFrames)*100)
	fmt.Println("Positon <=> Value")
	for i := 0; i < PHYSICAL_MEMORY_SIZE/FRAME_PAGE_SIZE; i++ {
		start := FRAME_PAGE_SIZE * i
		end := FRAME_PAGE_SIZE*i + FRAME_PAGE_SIZE
		fmt.Printf("%d <=> %s\n", i, PhysicalMemory[start:end])
	}
	fmt.Println()
}
