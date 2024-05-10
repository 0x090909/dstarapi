package backend_datamodel_model

import (
	"errors"
)

type ScraperType int

const (
	ONETO1_SCRAPERTYPE ScraperType = iota
	GENERIC_SCRAPERTYPE
)

func (i ScraperType) String() string {
	return []string{"1to1", "generic"}[i]
}
func ParseScraperType(v string) (any, error) {
	result := ONETO1_SCRAPERTYPE
	switch v {
	case "1to1":
		result = ONETO1_SCRAPERTYPE
	case "generic":
		result = GENERIC_SCRAPERTYPE
	default:
		return 0, errors.New("Unknown ScraperType value: " + v)
	}
	return &result, nil
}
func SerializeScraperType(values []ScraperType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i ScraperType) isMultiValue() bool {
	return false
}
