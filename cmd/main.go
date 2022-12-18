package main

import (
	"card-search/pkg/mock_data"
	"card-search/pkg/prefixtree"

	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {

	tree := prefixtree.New()

	for id, vc := range mock_data.Id2VC {
		// index some mock_data
		tree.Add(strings.ToLower(vc), id)
	}

	// DEBUG Section
	tree.Output() //may take a looooong time
	fmt.Println("=======================")
	fmt.Println(" ")
	PrintMemUsage()
	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()

	// Attempt to find various prefixes in the tree, and output the payload.
	fmt.Printf("%-18s %-8s\n", "Search String", "Payloads")
	fmt.Printf("%-18s %-8s\n", "------", "----")

	for _, s := range []string{
		"127",
		"882",
		"42",
	} {
		start := time.Now()
		data, _ := tree.Find(s)

		fmt.Println("Search time: ", time.Since(start))
		fmt.Printf("%-18s %-8v\n", s, data)
	}

}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
