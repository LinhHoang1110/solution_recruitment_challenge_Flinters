package aggregator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProcessCSV(t *testing.T) {
	// Create a temporary test CSV file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_data.csv")

	content := `campaign_id,date,impressions,clicks,spend,conversions
CMP001,2025-01-01,12000,300,45.50,12
CMP002,2025-01-01,8000,120,28.00,4
CMP001,2025-01-02,14000,340,48.20,15
CMP003,2025-01-01,5000,60,15.00,3
CMP002,2025-01-02,8500,150,31.00,5
`
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Process the test file
	stats, err := ProcessCSV(testFile)
	if err != nil {
		t.Fatalf("ProcessCSV failed: %v", err)
	}

	// Verify results
	if len(stats) != 3 {
		t.Errorf("Expected 3 campaigns, got %d", len(stats))
	}

	// Check CMP001 aggregation
	cmp001 := stats["CMP001"]
	if cmp001 == nil {
		t.Fatal("CMP001 not found in results")
	}
	if cmp001.TotalImpressions != 26000 {
		t.Errorf("CMP001 impressions: expected 26000, got %d", cmp001.TotalImpressions)
	}
	if cmp001.TotalClicks != 640 {
		t.Errorf("CMP001 clicks: expected 640, got %d", cmp001.TotalClicks)
	}
	if cmp001.TotalConversions != 27 {
		t.Errorf("CMP001 conversions: expected 27, got %d", cmp001.TotalConversions)
	}

	// Check CMP002 aggregation
	cmp002 := stats["CMP002"]
	if cmp002 == nil {
		t.Fatal("CMP002 not found in results")
	}
	if cmp002.TotalImpressions != 16500 {
		t.Errorf("CMP002 impressions: expected 16500, got %d", cmp002.TotalImpressions)
	}
	if cmp002.TotalConversions != 9 {
		t.Errorf("CMP002 conversions: expected 9, got %d", cmp002.TotalConversions)
	}
}

func TestProcessCSVEmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "empty.csv")

	content := `campaign_id,date,impressions,clicks,spend,conversions
`
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	stats, err := ProcessCSV(testFile)
	if err != nil {
		t.Fatalf("ProcessCSV failed: %v", err)
	}

	if len(stats) != 0 {
		t.Errorf("Expected 0 campaigns, got %d", len(stats))
	}
}

func TestProcessCSVInvalidHeader(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "invalid.csv")

	content := `wrong_header,date,impressions,clicks,spend,conversions
CMP001,2025-01-01,12000,300,45.50,12
`
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err := ProcessCSV(testFile)
	if err == nil {
		t.Error("Expected error for invalid header, got nil")
	}
}
