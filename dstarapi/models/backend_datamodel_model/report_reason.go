package backend_datamodel_model
import (
    "errors"
)
type ReportReason int

const (
    VIOLENCE_REPORTREASON ReportReason = iota
    SEXUAL_REPORTREASON
    COPYRIGHT_REPORTREASON
    FRAUD_REPORTREASON
    OTHER_REPORTREASON
)

func (i ReportReason) String() string {
    return []string{"violence", "sexual", "copyright", "fraud", "other"}[i]
}
func ParseReportReason(v string) (any, error) {
    result := VIOLENCE_REPORTREASON
    switch v {
        case "violence":
            result = VIOLENCE_REPORTREASON
        case "sexual":
            result = SEXUAL_REPORTREASON
        case "copyright":
            result = COPYRIGHT_REPORTREASON
        case "fraud":
            result = FRAUD_REPORTREASON
        case "other":
            result = OTHER_REPORTREASON
        default:
            return 0, errors.New("Unknown ReportReason value: " + v)
    }
    return &result, nil
}
func SerializeReportReason(values []ReportReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReportReason) isMultiValue() bool {
    return false
}
