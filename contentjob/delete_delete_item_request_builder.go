package contentjob

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeleteDeleteItemRequestBuilder builds and executes requests for operations under \content-job\delete\{id}
type DeleteDeleteItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DeleteDeleteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteDeleteItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewDeleteDeleteItemRequestBuilderInternal instantiates a new DeleteDeleteItemRequestBuilder and sets the default values.
func NewDeleteDeleteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DeleteDeleteItemRequestBuilder {
	m := &DeleteDeleteItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/content-job/delete/{id}", pathParameters),
	}
	return m
}

// NewDeleteDeleteItemRequestBuilder instantiates a new DeleteDeleteItemRequestBuilder and sets the default values.
func NewDeleteDeleteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DeleteDeleteItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDeleteDeleteItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete content job
// returns a *string when successful
func (m *DeleteDeleteItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeleteDeleteItemRequestBuilderDeleteRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*string), nil
}

// ToDeleteRequestInformation delete content job
// returns a *RequestInformation when successful
func (m *DeleteDeleteItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeleteDeleteItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeleteDeleteItemRequestBuilder when successful
func (m *DeleteDeleteItemRequestBuilder) WithUrl(rawUrl string) *DeleteDeleteItemRequestBuilder {
	return NewDeleteDeleteItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
