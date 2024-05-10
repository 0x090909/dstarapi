package generatetoken

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i36cf3bd9ca5b242f4ea0ae9d17ea1ea49d66a818342df578db9868c6a12ed0f0 "github.com/0x090909/dstarapi/models/fiber"
)

// GenerateTokenRequestBuilder builds and executes requests for operations under \generate-token
type GenerateTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GenerateTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GenerateTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGenerateTokenRequestBuilderInternal instantiates a new GenerateTokenRequestBuilder and sets the default values.
func NewGenerateTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GenerateTokenRequestBuilder) {
    m := &GenerateTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/generate-token", pathParameters),
    }
    return m
}
// NewGenerateTokenRequestBuilder instantiates a new GenerateTokenRequestBuilder and sets the default values.
func NewGenerateTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GenerateTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generate Token to access secure endpoints
// returns a MapEscapedable when successful
func (m *GenerateTokenRequestBuilder) Post(ctx context.Context, body GenerateTokenPostRequestBodyable, requestConfiguration *GenerateTokenRequestBuilderPostRequestConfiguration)(i36cf3bd9ca5b242f4ea0ae9d17ea1ea49d66a818342df578db9868c6a12ed0f0.MapEscapedable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i36cf3bd9ca5b242f4ea0ae9d17ea1ea49d66a818342df578db9868c6a12ed0f0.CreateMapEscapedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i36cf3bd9ca5b242f4ea0ae9d17ea1ea49d66a818342df578db9868c6a12ed0f0.MapEscapedable), nil
}
// ToPostRequestInformation generate Token to access secure endpoints
// returns a *RequestInformation when successful
func (m *GenerateTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body GenerateTokenPostRequestBodyable, requestConfiguration *GenerateTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/x-www-form-urlencoded", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GenerateTokenRequestBuilder when successful
func (m *GenerateTokenRequestBuilder) WithUrl(rawUrl string)(*GenerateTokenRequestBuilder) {
    return NewGenerateTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
