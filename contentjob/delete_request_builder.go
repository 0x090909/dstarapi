package contentjob

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// DeleteRequestBuilder builds and executes requests for operations under \content-job\delete
type DeleteRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ById gets an item from the github.com/0x090909/dstarapi.contentJob.delete.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *DeleteDeleteItemRequestBuilder when successful
func (m *DeleteRequestBuilder) ById(id string) *DeleteDeleteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["id"] = id
	}
	return NewDeleteDeleteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByIdInteger gets an item from the github.com/0x090909/dstarapi.contentJob.delete.item collection
// returns a *DeleteDeleteItemRequestBuilder when successful
func (m *DeleteRequestBuilder) ByIdInteger(id int32) *DeleteDeleteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
	return NewDeleteDeleteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewDeleteRequestBuilderInternal instantiates a new DeleteRequestBuilder and sets the default values.
func NewDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DeleteRequestBuilder {
	m := &DeleteRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/content-job/delete", pathParameters),
	}
	return m
}

// NewDeleteRequestBuilder instantiates a new DeleteRequestBuilder and sets the default values.
func NewDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DeleteRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
