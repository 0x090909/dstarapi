package users

import (
	"context"
	ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemFollowersRequestBuilder builds and executes requests for operations under \users\{uid}\followers
type ItemFollowersRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemFollowersRequestBuilderGetQueryParameters description
type ItemFollowersRequestBuilderGetQueryParameters struct {
	// Limit
	Limit *int32 `uriparametername:"limit"`
	// Offset
	Offset *int32 `uriparametername:"offset"`
}

// ItemFollowersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowersRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemFollowersRequestBuilderGetQueryParameters
}

// NewItemFollowersRequestBuilderInternal instantiates a new ItemFollowersRequestBuilder and sets the default values.
func NewItemFollowersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemFollowersRequestBuilder {
	m := &ItemFollowersRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/followers{?limit*,offset*}", pathParameters),
	}
	return m
}

// NewItemFollowersRequestBuilder instantiates a new ItemFollowersRequestBuilder and sets the default values.
func NewItemFollowersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemFollowersRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemFollowersRequestBuilderInternal(urlParams, requestAdapter)
}

// Get description
// returns a []Userable when successful
func (m *ItemFollowersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFollowersRequestBuilderGetRequestConfiguration) ([]ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Userable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreateUserFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	val := make([]ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Userable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Userable)
		}
	}
	return val, nil
}

// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemFollowersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFollowersRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemFollowersRequestBuilder when successful
func (m *ItemFollowersRequestBuilder) WithUrl(rawUrl string) *ItemFollowersRequestBuilder {
	return NewItemFollowersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
