package backend_datamodel_model

import (
	"errors"
)

type JobStatus int

const (
	STARTED_JOBSTATUS JobStatus = iota
	FINISHED_JOBSTATUS
)

func (i JobStatus) String() string {
	return []string{"started", "finished"}[i]
}
func ParseJobStatus(v string) (any, error) {
	result := STARTED_JOBSTATUS
	switch v {
	case "started":
		result = STARTED_JOBSTATUS
	case "finished":
		result = FINISHED_JOBSTATUS
	default:
		return 0, errors.New("Unknown JobStatus value: " + v)
	}
	return &result, nil
}
func SerializeJobStatus(values []JobStatus) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i JobStatus) isMultiValue() bool {
	return false
}
