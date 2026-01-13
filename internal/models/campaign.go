package models

// CampaignStats holds aggregated statistics for a campaign
type CampaignStats struct {
	CampaignID       string
	TotalImpressions int64
	TotalClicks      int64
	TotalSpend       float64
	TotalConversions int64
}

// CTR calculates Click-Through Rate: total_clicks / total_impressions
func (c *CampaignStats) CTR() float64 {
	if c.TotalImpressions == 0 {
		return 0
	}
	return float64(c.TotalClicks) / float64(c.TotalImpressions)
}

// CPA calculates Cost Per Acquisition: total_spend / total_conversions
// Returns nil if conversions is zero
func (c *CampaignStats) CPA() *float64 {
	if c.TotalConversions == 0 {
		return nil
	}
	cpa := c.TotalSpend / float64(c.TotalConversions)
	return &cpa
}
