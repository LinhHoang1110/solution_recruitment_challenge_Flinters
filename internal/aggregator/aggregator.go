package aggregator

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"ad-aggregator/internal/models"
)

// ProcessCSV reads a CSV file and aggregates statistics by campaign_id
// Uses streaming approach for memory efficiency with large files
func ProcessCSV(inputPath string) (map[string]*models.CampaignStats, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Use buffered reader for better I/O performance
	reader := csv.NewReader(bufio.NewReaderSize(file, 64*1024)) // 64KB buffer

	// Read and validate header
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	// Validate expected columns
	if err := validateHeader(header); err != nil {
		return nil, err
	}

	// Map to store aggregated stats per campaign
	stats := make(map[string]*models.CampaignStats)

	// Process each row
	lineNum := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading line %d: %w", lineNum, err)
		}
		lineNum++

		if len(record) != 6 {
			continue // Skip malformed rows
		}

		// Parse fields
		campaignID := record[0]
		// date := record[1] // Not needed for aggregation

		impressions, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			continue // Skip rows with invalid impressions
		}

		clicks, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			continue // Skip rows with invalid clicks
		}

		spend, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			continue // Skip rows with invalid spend
		}

		conversions, err := strconv.ParseInt(record[5], 10, 64)
		if err != nil {
			continue // Skip rows with invalid conversions
		}

		// Aggregate
		if _, exists := stats[campaignID]; !exists {
			stats[campaignID] = &models.CampaignStats{
				CampaignID: campaignID,
			}
		}

		stats[campaignID].TotalImpressions += impressions
		stats[campaignID].TotalClicks += clicks
		stats[campaignID].TotalSpend += spend
		stats[campaignID].TotalConversions += conversions
	}

	return stats, nil
}

// validateHeader checks if the CSV has expected columns
func validateHeader(header []string) error {
	expected := []string{"campaign_id", "date", "impressions", "clicks", "spend", "conversions"}
	if len(header) != len(expected) {
		return fmt.Errorf("invalid header: expected %d columns, got %d", len(expected), len(header))
	}
	for i, col := range expected {
		if header[i] != col {
			return fmt.Errorf("invalid header: expected column %d to be '%s', got '%s'", i, col, header[i])
		}
	}
	return nil
}
