package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"ad-aggregator/internal/aggregator"
	"ad-aggregator/internal/output"
)

func main() {
	// Parse command line arguments
	inputPath := flag.String("input", "", "Path to input CSV file (required)")
	outputDir := flag.String("output", "results/", "Output directory for result files")
	flag.Parse()

	// Validate input
	if *inputPath == "" {
		fmt.Println("Error: --input flag is required")
		fmt.Println("Usage: aggregator --input <csv_file> [--output <output_dir>]")
		os.Exit(1)
	}

	// Check if input file exists
	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		fmt.Printf("Error: input file '%s' does not exist\n", *inputPath)
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Error: failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=================================================")
	fmt.Println("  Ad Performance Aggregator")
	fmt.Println("=================================================")
	fmt.Printf("Input file:  %s\n", *inputPath)
	fmt.Printf("Output dir:  %s\n", *outputDir)
	fmt.Println("-------------------------------------------------")

	// Record start time
	startTime := time.Now()

	// Get initial memory stats
	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// Process CSV file
	fmt.Println("Processing CSV file...")
	stats, err := aggregator.ProcessCSV(*inputPath)
	if err != nil {
		fmt.Printf("Error processing CSV: %v\n", err)
		os.Exit(1)
	}

	processTime := time.Since(startTime)
	fmt.Printf("Processed %d unique campaigns in %v\n", len(stats), processTime)

	// Write top 10 CTR
	ctrPath := filepath.Join(*outputDir, "top10_ctr.csv")
	fmt.Printf("Writing top 10 CTR to %s...\n", ctrPath)
	if err := output.WriteTop10CTR(stats, ctrPath); err != nil {
		fmt.Printf("Error writing CTR file: %v\n", err)
		os.Exit(1)
	}

	// Write top 10 CPA
	cpaPath := filepath.Join(*outputDir, "top10_cpa.csv")
	fmt.Printf("Writing top 10 CPA to %s...\n", cpaPath)
	if err := output.WriteTop10CPA(stats, cpaPath); err != nil {
		fmt.Printf("Error writing CPA file: %v\n", err)
		os.Exit(1)
	}

	// Get final memory stats
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	totalTime := time.Since(startTime)

	fmt.Println("-------------------------------------------------")
	fmt.Println("✅ Processing complete!")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Total processing time: %v\n", totalTime)
	fmt.Printf("Peak memory allocated: %.2f MB\n", float64(memStatsAfter.Alloc)/1024/1024)
	fmt.Printf("Total memory allocated: %.2f MB\n", float64(memStatsAfter.TotalAlloc)/1024/1024)
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Output files:\n")
	fmt.Printf("  - %s\n", ctrPath)
	fmt.Printf("  - %s\n", cpaPath)
	fmt.Println("=================================================")
}
