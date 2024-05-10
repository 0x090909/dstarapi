package posts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompetingWithUItemRequestBuilder builds and executes requests for operations under \posts\competing\{uid}
type CompetingWithUItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompetingWithUItemRequestBuilderGetQueryParameters description
type CompetingWithUItemRequestBuilderGetQueryParameters struct {
    // yesterday
    Yesterday *bool `uriparametername:"yesterday"`
}
// CompetingWithUItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompetingWithUItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompetingWithUItemRequestBuilderGetQueryParameters
}
// NewCompetingWithUItemRequestBuilderInternal instantiates a new CompetingWithUItemRequestBuilder and sets the default values.
func NewCompetingWithUItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompetingWithUItemRequestBuilder) {
    m := &CompetingWithUItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/posts/competing/{uid}{?yesterday*}", pathParameters),
    }
    return m
}
// NewCompetingWithUItemRequestBuilder instantiates a new CompetingWithUItemRequestBuilder and sets the default values.
func NewCompetingWithUItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompetingWithUItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompetingWithUItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get description
// Deprecated: This method is obsolete. Use GetAsWithUGetResponse instead.
// returns a CompetingItemWithUResponseable when successful
func (m *CompetingWithUItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompetingWithUItemRequestBuilderGetRequestConfiguration)(CompetingItemWithUResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompetingItemWithUResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompetingItemWithUResponseable), nil
}
// GetAsWithUGetResponse description
// returns a CompetingItemWithUGetResponseable when successful
func (m *CompetingWithUItemRequestBuilder) GetAsWithUGetResponse(ctx context.Context, requestConfiguration *CompetingWithUItemRequestBuilderGetRequestConfiguration)(CompetingItemWithUGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompetingItemWithUGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompetingItemWithUGetResponseable), nil
}
// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *CompetingWithUItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompetingWithUItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompetingWithUItemRequestBuilder when successful
func (m *CompetingWithUItemRequestBuilder) WithUrl(rawUrl string)(*CompetingWithUItemRequestBuilder) {
    return NewCompetingWithUItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
