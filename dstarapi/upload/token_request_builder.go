package upload

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
)

// TokenRequestBuilder builds and executes requests for operations under \upload\token
type TokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TokenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTokenRequestBuilderInternal instantiates a new TokenRequestBuilder and sets the default values.
func NewTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenRequestBuilder) {
    m := &TokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/upload/token", pathParameters),
    }
    return m
}
// NewTokenRequestBuilder instantiates a new TokenRequestBuilder and sets the default values.
func NewTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Get description
// returns a UploadTokenable when successful
func (m *TokenRequestBuilder) Get(ctx context.Context, requestConfiguration *TokenRequestBuilderGetRequestConfiguration)(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.UploadTokenable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreateUploadTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.UploadTokenable), nil
}
// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *TokenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TokenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TokenRequestBuilder when successful
func (m *TokenRequestBuilder) WithUrl(rawUrl string)(*TokenRequestBuilder) {
    return NewTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
