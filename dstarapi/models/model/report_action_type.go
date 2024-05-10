package model
import (
    "errors"
)
type ReportActionType int

const (
    REPORT_USER_REPORTACTIONTYPE ReportActionType = iota
    REPORT_POST_REPORTACTIONTYPE
    REPORT_COMMENT_REPORTACTIONTYPE
    REPORT_PROBLEM_REPORTACTIONTYPE
    REPORT_FEEDBACK_REPORTACTIONTYPE
)

func (i ReportActionType) String() string {
    return []string{"report_user", "report_post", "report_comment", "report_problem", "report_feedback"}[i]
}
func ParseReportActionType(v string) (any, error) {
    result := REPORT_USER_REPORTACTIONTYPE
    switch v {
        case "report_user":
            result = REPORT_USER_REPORTACTIONTYPE
        case "report_post":
            result = REPORT_POST_REPORTACTIONTYPE
        case "report_comment":
            result = REPORT_COMMENT_REPORTACTIONTYPE
        case "report_problem":
            result = REPORT_PROBLEM_REPORTACTIONTYPE
        case "report_feedback":
            result = REPORT_FEEDBACK_REPORTACTIONTYPE
        default:
            return 0, errors.New("Unknown ReportActionType value: " + v)
    }
    return &result, nil
}
func SerializeReportActionType(values []ReportActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReportActionType) isMultiValue() bool {
    return false
}
