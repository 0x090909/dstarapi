package users

import (
	"context"
	ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCheckWithUsernameItemRequestBuilder builds and executes requests for operations under \users\{uid}\check\{username}
type ItemCheckWithUsernameItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCheckWithUsernameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCheckWithUsernameItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemCheckWithUsernameItemRequestBuilderInternal instantiates a new ItemCheckWithUsernameItemRequestBuilder and sets the default values.
func NewItemCheckWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCheckWithUsernameItemRequestBuilder {
	m := &ItemCheckWithUsernameItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/check/{username}", pathParameters),
	}
	return m
}

// NewItemCheckWithUsernameItemRequestBuilder instantiates a new ItemCheckWithUsernameItemRequestBuilder and sets the default values.
func NewItemCheckWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCheckWithUsernameItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCheckWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get checkUsername checks if a username is available and if not finds an available username similar to the one provided
// returns a CheckUsernameable when successful
func (m *ItemCheckWithUsernameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCheckWithUsernameItemRequestBuilderGetRequestConfiguration) (ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CheckUsernameable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreateCheckUsernameFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CheckUsernameable), nil
}

// ToGetRequestInformation checkUsername checks if a username is available and if not finds an available username similar to the one provided
// returns a *RequestInformation when successful
func (m *ItemCheckWithUsernameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCheckWithUsernameItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCheckWithUsernameItemRequestBuilder when successful
func (m *ItemCheckWithUsernameItemRequestBuilder) WithUrl(rawUrl string) *ItemCheckWithUsernameItemRequestBuilder {
	return NewItemCheckWithUsernameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
