package upload

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UploadRequestBuilder builds and executes requests for operations under \upload
type UploadRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// UploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UploadRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewUploadRequestBuilderInternal instantiates a new UploadRequestBuilder and sets the default values.
func NewUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UploadRequestBuilder {
	m := &UploadRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/upload", pathParameters),
	}
	return m
}

// NewUploadRequestBuilder instantiates a new UploadRequestBuilder and sets the default values.
func NewUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UploadRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUploadRequestBuilderInternal(urlParams, requestAdapter)
}

// Post description
// returns a *string when successful
func (m *UploadRequestBuilder) Post(ctx context.Context, body *string, requestConfiguration *UploadRequestBuilderPostRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// Token the token property
// returns a *TokenRequestBuilder when successful
func (m *UploadRequestBuilder) Token() *TokenRequestBuilder {
	return NewTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToPostRequestInformation description
// returns a *RequestInformation when successful
func (m *UploadRequestBuilder) ToPostRequestInformation(ctx context.Context, body *string, requestConfiguration *UploadRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	requestInfo.SetContentFromScalar(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UploadRequestBuilder when successful
func (m *UploadRequestBuilder) WithUrl(rawUrl string) *UploadRequestBuilder {
	return NewUploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
