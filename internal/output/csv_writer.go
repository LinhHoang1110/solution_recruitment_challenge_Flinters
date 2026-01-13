package output

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

	"ad-aggregator/internal/models"
)

// WriteTop10CTR writes the top 10 campaigns with highest CTR to a CSV file
func WriteTop10CTR(stats map[string]*models.CampaignStats, outputPath string) error {
	// Convert map to slice for sorting
	campaigns := mapToSlice(stats)

	// Sort by CTR descending
	sort.Slice(campaigns, func(i, j int) bool {
		return campaigns[i].CTR() > campaigns[j].CTR()
	})

	// Take top 10
	top10 := campaigns
	if len(top10) > 10 {
		top10 = top10[:10]
	}

	return writeCSV(top10, outputPath)
}

// WriteTop10CPA writes the top 10 campaigns with lowest CPA to a CSV file
// Excludes campaigns with zero conversions
func WriteTop10CPA(stats map[string]*models.CampaignStats, outputPath string) error {
	// Convert map to slice, excluding zero conversions
	var campaigns []*models.CampaignStats
	for _, c := range stats {
		if c.TotalConversions > 0 {
			campaigns = append(campaigns, c)
		}
	}

	// Sort by CPA ascending (lowest first)
	sort.Slice(campaigns, func(i, j int) bool {
		cpaI := campaigns[i].CPA()
		cpaJ := campaigns[j].CPA()
		if cpaI == nil {
			return false
		}
		if cpaJ == nil {
			return true
		}
		return *cpaI < *cpaJ
	})

	// Take top 10
	top10 := campaigns
	if len(top10) > 10 {
		top10 = top10[:10]
	}

	return writeCSV(top10, outputPath)
}

// mapToSlice converts campaign stats map to slice
func mapToSlice(stats map[string]*models.CampaignStats) []*models.CampaignStats {
	result := make([]*models.CampaignStats, 0, len(stats))
	for _, c := range stats {
		result = append(result, c)
	}
	return result
}

// writeCSV writes campaign stats to a CSV file
func writeCSV(campaigns []*models.CampaignStats, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"campaign_id", "total_impressions", "total_clicks", "total_spend", "total_conversions", "CTR", "CPA"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write data rows
	for _, c := range campaigns {
		ctr := fmt.Sprintf("%.4f", c.CTR())
		cpa := "null"
		if c.CPA() != nil {
			cpa = fmt.Sprintf("%.2f", *c.CPA())
		}

		row := []string{
			c.CampaignID,
			fmt.Sprintf("%d", c.TotalImpressions),
			fmt.Sprintf("%d", c.TotalClicks),
			fmt.Sprintf("%.2f", c.TotalSpend),
			fmt.Sprintf("%d", c.TotalConversions),
			ctr,
			cpa,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}

	return nil
}
