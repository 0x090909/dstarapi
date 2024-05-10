package users

import (
	"context"
	i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPreferenceRequestBuilder builds and executes requests for operations under \users\{uid}\preference
type ItemPreferenceRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemPreferenceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPreferenceRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemPreferenceRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPreferenceRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemPreferenceRequestBuilderInternal instantiates a new ItemPreferenceRequestBuilder and sets the default values.
func NewItemPreferenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPreferenceRequestBuilder {
	m := &ItemPreferenceRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/preference", pathParameters),
	}
	return m
}

// NewItemPreferenceRequestBuilder instantiates a new ItemPreferenceRequestBuilder and sets the default values.
func NewItemPreferenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemPreferenceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemPreferenceRequestBuilderInternal(urlParams, requestAdapter)
}

// Get description
// returns a Preferenceable when successful
func (m *ItemPreferenceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPreferenceRequestBuilderGetRequestConfiguration) (i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Preferenceable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.CreatePreferenceFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Preferenceable), nil
}

// Put description
// returns a *string when successful
func (m *ItemPreferenceRequestBuilder) Put(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Preferenceable, requestConfiguration *ItemPreferenceRequestBuilderPutRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
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

// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemPreferenceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPreferenceRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPutRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemPreferenceRequestBuilder) ToPutRequestInformation(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Preferenceable, requestConfiguration *ItemPreferenceRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPreferenceRequestBuilder when successful
func (m *ItemPreferenceRequestBuilder) WithUrl(rawUrl string) *ItemPreferenceRequestBuilder {
	return NewItemPreferenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
