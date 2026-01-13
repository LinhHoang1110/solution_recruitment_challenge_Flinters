# Ad Performance Aggregator

A CLI application to process large CSV advertising data and generate aggregated analytics.

## Features

- Processes large CSV files (~1GB) with memory-efficient streaming
- Aggregates data by campaign_id
- Generates top 10 CTR and CPA reports
- Tracks processing time and memory usage

## Requirements

- Go 1.21 or later
- OR Docker

## Installation

### Option 1: Build from source

```bash
# Clone the repository
git clone <repository-url>
cd fv-sec-001-software-engineer-challenge

# Build the binary
go build -o aggregator .
```

### Option 2: Using Docker

```bash
# Build the Docker image
docker build -t ad-aggregator .
```

## Usage

### Running locally

```bash
# Basic usage
./aggregator --input ad_data.csv --output results/

# View help
./aggregator --help
```

### Running with Docker

```bash
# Run with mounted data directory
docker run -v $(pwd):/app/data ad-aggregator \
  --input /app/data/ad_data.csv \
  --output /app/data/results/
```

## Output Files

The program generates two CSV files in the output directory:

1. **top10_ctr.csv** - Top 10 campaigns with highest Click-Through Rate
2. **top10_cpa.csv** - Top 10 campaigns with lowest Cost Per Acquisition

### Output Format

| Column | Description |
|--------|-------------|
| campaign_id | Campaign identifier |
| total_impressions | Sum of impressions |
| total_clicks | Sum of clicks |
| total_spend | Sum of spend (USD) |
| total_conversions | Sum of conversions |
| CTR | Click-Through Rate (clicks/impressions) |
| CPA | Cost Per Acquisition (spend/conversions) |

## Performance

| Metric | Value |
|--------|-------|
| File Size | ~1GB (1,043,304,870 bytes) |
| Rows Processed | ~27 million |
| Unique Campaigns | 50 |
| Processing Time | **8.4 seconds** |
| Peak Memory Usage | **3.63 MB** |

## Running Tests

```bash
go test ./... -v
```

## Libraries Used

- Standard library only (no external dependencies)
  - `encoding/csv` - CSV parsing
  - `bufio` - Buffered I/O for performance
  - `flag` - CLI argument parsing
  - `runtime` - Memory statistics

## Project Structure

```
.
├── main.go                    # CLI entry point
├── go.mod                     # Go module file
├── Dockerfile                 # Docker configuration
├── README.md                  # This file
├── internal/
│   ├── aggregator/
│   │   ├── aggregator.go      # CSV processing logic
│   │   └── aggregator_test.go # Unit tests
│   ├── models/
│   │   ├── campaign.go        # Data structures
│   │   └── campaign_test.go   # Unit tests
│   └── output/
│       └── csv_writer.go      # CSV output generation
└── results/
    ├── top10_ctr.csv          # Output: Top 10 CTR
    └── top10_cpa.csv          # Output: Top 10 CPA
```

## Algorithm

1. **Stream Processing**: Read CSV line-by-line using buffered I/O
2. **Aggregation**: Maintain in-memory map of campaign statistics
3. **Sorting**: Sort campaigns by CTR (desc) and CPA (asc)
4. **Output**: Write top 10 results to CSV files

Memory efficiency is achieved by:
- Not loading entire file into memory
- Only storing aggregated stats (~50 campaigns)
- Using buffered I/O for optimal disk reads
