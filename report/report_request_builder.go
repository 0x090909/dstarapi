package report

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ReportRequestBuilder builds and executes requests for operations under \report
type ReportRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByUid gets an item from the github.com/0x090909/dstarapi.report.item collection
// returns a *WithUItemRequestBuilder when successful
func (m *ReportRequestBuilder) ByUid(uid string) *WithUItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if uid != "" {
		urlTplParams["uid"] = uid
	}
	return NewWithUItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewReportRequestBuilderInternal instantiates a new ReportRequestBuilder and sets the default values.
func NewReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ReportRequestBuilder {
	m := &ReportRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/report", pathParameters),
	}
	return m
}

// NewReportRequestBuilder instantiates a new ReportRequestBuilder and sets the default values.
func NewReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ReportRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewReportRequestBuilderInternal(urlParams, requestAdapter)
}
