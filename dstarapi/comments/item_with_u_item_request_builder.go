package comments

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
)

// ItemWithUItemRequestBuilder builds and executes requests for operations under \comments\{id-id}\{uid}
type ItemWithUItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWithUItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithUItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemWithUItemRequestBuilderGetQueryParameters get all post comments or replies to a comment
type ItemWithUItemRequestBuilderGetQueryParameters struct {
    // Limit
    Limit *int32 `uriparametername:"limit"`
    // Offset
    Offset *int32 `uriparametername:"offset"`
    // Parent Comment ID
    Parent_id *int32 `uriparametername:"parent_id"`
}
// ItemWithUItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithUItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemWithUItemRequestBuilderGetQueryParameters
}
// NewItemWithUItemRequestBuilderInternal instantiates a new ItemWithUItemRequestBuilder and sets the default values.
func NewItemWithUItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithUItemRequestBuilder) {
    m := &ItemWithUItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/comments/{id%2Did}/{uid}?limit={limit}&offset={offset}{&parent_id*}", pathParameters),
    }
    return m
}
// NewItemWithUItemRequestBuilder instantiates a new ItemWithUItemRequestBuilder and sets the default values.
func NewItemWithUItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWithUItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWithUItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete description
// returns a *string when successful
func (m *ItemWithUItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemWithUItemRequestBuilderDeleteRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get all post comments or replies to a comment
// returns a PaginatedCommentsable when successful
func (m *ItemWithUItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWithUItemRequestBuilderGetRequestConfiguration)(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.PaginatedCommentsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreatePaginatedCommentsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.PaginatedCommentsable), nil
}
// ToDeleteRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemWithUItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemWithUItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/comments/{id%2Did}/{uid}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain")
    return requestInfo, nil
}
// ToGetRequestInformation get all post comments or replies to a comment
// returns a *RequestInformation when successful
func (m *ItemWithUItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWithUItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemWithUItemRequestBuilder when successful
func (m *ItemWithUItemRequestBuilder) WithUrl(rawUrl string)(*ItemWithUItemRequestBuilder) {
    return NewItemWithUItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
