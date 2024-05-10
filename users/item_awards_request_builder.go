package users

import (
	"context"
	ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
	i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAwardsRequestBuilder builds and executes requests for operations under \users\{uid}\awards
type ItemAwardsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemAwardsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAwardsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemAwardsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAwardsRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemAwardsRequestBuilderInternal instantiates a new ItemAwardsRequestBuilder and sets the default values.
func NewItemAwardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemAwardsRequestBuilder {
	m := &ItemAwardsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/awards", pathParameters),
	}
	return m
}

// NewItemAwardsRequestBuilder instantiates a new ItemAwardsRequestBuilder and sets the default values.
func NewItemAwardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemAwardsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemAwardsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get description
// returns a []Awardsable when successful
func (m *ItemAwardsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAwardsRequestBuilderGetRequestConfiguration) ([]ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Awardsable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreateAwardsFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	val := make([]ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Awardsable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Awardsable)
		}
	}
	return val, nil
}

// Put description
// returns a *string when successful
func (m *ItemAwardsRequestBuilder) Put(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Userable, requestConfiguration *ItemAwardsRequestBuilderPutRequestConfiguration) (*string, error) {
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
func (m *ItemAwardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAwardsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemAwardsRequestBuilder) ToPutRequestInformation(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Userable, requestConfiguration *ItemAwardsRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAwardsRequestBuilder when successful
func (m *ItemAwardsRequestBuilder) WithUrl(rawUrl string) *ItemAwardsRequestBuilder {
	return NewItemAwardsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
