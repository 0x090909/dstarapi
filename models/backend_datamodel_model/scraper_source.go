package backend_datamodel_model

import (
	"errors"
)

type ScraperSource int

const (
	YOUTUBE_SCRAPERSOURCE ScraperSource = iota
	INSTAGRAM_SCRAPERSOURCE
	TIKTOK_SCRAPERSOURCE
)

func (i ScraperSource) String() string {
	return []string{"youtube", "instagram", "tiktok"}[i]
}
func ParseScraperSource(v string) (any, error) {
	result := YOUTUBE_SCRAPERSOURCE
	switch v {
	case "youtube":
		result = YOUTUBE_SCRAPERSOURCE
	case "instagram":
		result = INSTAGRAM_SCRAPERSOURCE
	case "tiktok":
		result = TIKTOK_SCRAPERSOURCE
	default:
		return 0, errors.New("Unknown ScraperSource value: " + v)
	}
	return &result, nil
}
func SerializeScraperSource(values []ScraperSource) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ScraperSource) isMultiValue() bool {
	return false
}
