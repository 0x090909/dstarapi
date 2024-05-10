package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemUnfollowWithTargetuItemRequestBuilder builds and executes requests for operations under \users\{uid}\unfollow\{targetuid}
type ItemUnfollowWithTargetuItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUnfollowWithTargetuItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUnfollowWithTargetuItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUnfollowWithTargetuItemRequestBuilderInternal instantiates a new ItemUnfollowWithTargetuItemRequestBuilder and sets the default values.
func NewItemUnfollowWithTargetuItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnfollowWithTargetuItemRequestBuilder) {
    m := &ItemUnfollowWithTargetuItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/unfollow/{targetuid}", pathParameters),
    }
    return m
}
// NewItemUnfollowWithTargetuItemRequestBuilder instantiates a new ItemUnfollowWithTargetuItemRequestBuilder and sets the default values.
func NewItemUnfollowWithTargetuItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnfollowWithTargetuItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnfollowWithTargetuItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post description
// returns a *string when successful
func (m *ItemUnfollowWithTargetuItemRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemUnfollowWithTargetuItemRequestBuilderPostRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
func (m *ItemUnfollowWithTargetuItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemUnfollowWithTargetuItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemUnfollowWithTargetuItemRequestBuilder when successful
func (m *ItemUnfollowWithTargetuItemRequestBuilder) WithUrl(rawUrl string)(*ItemUnfollowWithTargetuItemRequestBuilder) {
    return NewItemUnfollowWithTargetuItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
