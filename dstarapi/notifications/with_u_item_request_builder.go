package notifications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if76085177abdb9a89ef55b68403ab392fd7e735aee7094dcef57eecc3aedbb19 "github.com/0x090909/dstarapi/models/apimodel"
)

// WithUItemRequestBuilder builds and executes requests for operations under \notifications\{uid}
type WithUItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithUItemRequestBuilderGetQueryParameters get user notifications with pagination
type WithUItemRequestBuilderGetQueryParameters struct {
    // Count
    Count *int32 `uriparametername:"count"`
    // Offset
    Offset *int32 `uriparametername:"offset"`
}
// WithUItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithUItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithUItemRequestBuilderGetQueryParameters
}
// NewWithUItemRequestBuilderInternal instantiates a new WithUItemRequestBuilder and sets the default values.
func NewWithUItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUItemRequestBuilder) {
    m := &WithUItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/notifications/{uid}{?count*,offset*}", pathParameters),
    }
    return m
}
// NewWithUItemRequestBuilder instantiates a new WithUItemRequestBuilder and sets the default values.
func NewWithUItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithUItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get user notifications with pagination
// returns a PaginatedEventsable when successful
func (m *WithUItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithUItemRequestBuilderGetRequestConfiguration)(if76085177abdb9a89ef55b68403ab392fd7e735aee7094dcef57eecc3aedbb19.PaginatedEventsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if76085177abdb9a89ef55b68403ab392fd7e735aee7094dcef57eecc3aedbb19.CreatePaginatedEventsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if76085177abdb9a89ef55b68403ab392fd7e735aee7094dcef57eecc3aedbb19.PaginatedEventsable), nil
}
// ToGetRequestInformation get user notifications with pagination
// returns a *RequestInformation when successful
func (m *WithUItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithUItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WithUItemRequestBuilder when successful
func (m *WithUItemRequestBuilder) WithUrl(rawUrl string)(*WithUItemRequestBuilder) {
    return NewWithUItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
