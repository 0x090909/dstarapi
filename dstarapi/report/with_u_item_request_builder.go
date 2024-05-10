package report

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
)

// WithUItemRequestBuilder builds and executes requests for operations under \report\{uid}
type WithUItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithUItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithUItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWithUItemRequestBuilderInternal instantiates a new WithUItemRequestBuilder and sets the default values.
func NewWithUItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUItemRequestBuilder) {
    m := &WithUItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/report/{uid}", pathParameters),
    }
    return m
}
// NewWithUItemRequestBuilder instantiates a new WithUItemRequestBuilder and sets the default values.
func NewWithUItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithUItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post description
// returns a *string when successful
func (m *WithUItemRequestBuilder) Post(ctx context.Context, body ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Reportable, requestConfiguration *WithUItemRequestBuilderPostRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation description
// returns a *RequestInformation when successful
func (m *WithUItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Reportable, requestConfiguration *WithUItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WithUItemRequestBuilder when successful
func (m *WithUItemRequestBuilder) WithUrl(rawUrl string)(*WithUItemRequestBuilder) {
    return NewWithUItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
