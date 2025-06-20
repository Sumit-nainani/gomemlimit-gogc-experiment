package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

type dataUsed struct {
	Data int `json:"data"`
}

func PrintMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	color := colorGreen
	if m.Alloc > 50*1024*1024 { // > 50MB
		color = colorYellow
	}
	if m.Alloc > 90*1024*1024 { // > 90MB
		color = colorRed
	}

	fmt.Printf("%s║ %-24.2f ║ %-24.2f ║ %-26.2f ║ %-25.2f ║ %-25.2f ║ %-10d ║ %-10d ║\n", color,
		float64(m.Alloc)/(1024*1024),
		float64(m.Sys)/(1024*1024),
		float64(m.TotalAlloc)/(1024*1024),
		float64(m.HeapAlloc)/(1024*1024),
		float64(m.PauseTotalNs)/1e6,
		m.NumGC,
		runtime.GOMAXPROCS(runtime.NumCPU()),
	)
}

func HeavyHandler(w http.ResponseWriter, r *http.Request) {
	var dataused dataUsed
	json.NewDecoder(r.Body).Decode(&dataused)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		data := make([]int, dataused.Data)
		for i := range data {
			data[i] = i
		}
	}()

	wg.Wait()
	// 🔥 Final Memory Snapshot
	PrintMemStats()
	fmt.Fprintf(w, "✅ Successfully handled heavy request\n")
}
