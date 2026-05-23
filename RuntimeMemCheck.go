package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var m runtime.MemStats
	
	// 1. Check memory before allocation
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial Alloc: %v MB\n", m.Alloc / 1024 / 1024)

	// 2. Allocate a huge slice (Heap allocation)
	// This simulates a large MQTT message buffer
	buffer := make([]byte, 500 * 1024 * 1024) // 500 MB
	_ = buffer[0] // Touch it so the OS actually allocates it

	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc after 500MB buffer: %v MB\n", m.Alloc / 1024 / 1024)
	fmt.Printf("System Memory (OS view): %v MB\n", m.Sys / 1024 / 1024)

	fmt.Println("Waiting 10 seconds... check 'top' or 'htop' now!")
	time.Sleep(10 * time.Second)
}