package main

import (
	"fmt"
	"goenvs/server"
)

const (
	colorBlue = "\033[34m" // Blue color
)

func main() {
	server := server.Server()

	fmt.Println("🚀 Server running at http://localhost:8080")
	fmt.Printf("%s║ %-24s ║ %-24s ║ %-24s ║ %-24s ║  %-24s ║ %-10s ║ %-10s ║\n",
		colorBlue,
		"Allocated-Memory (MB)",
		"Os-Allocated-Memory (MB)",
		"TotalAllocated-Memory (MB)",
		"HeapAllocated-Memory (MB)",
		"CPU time used in GC (ms)",
		"NumGC",
		"Used Cpu")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
