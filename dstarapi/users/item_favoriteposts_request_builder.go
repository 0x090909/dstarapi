package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
)

// ItemFavoritepostsRequestBuilder builds and executes requests for operations under \users\{uid}\favoriteposts
type ItemFavoritepostsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFavoritepostsRequestBuilderGetQueryParameters description
type ItemFavoritepostsRequestBuilderGetQueryParameters struct {
    // Limit
    Limit *int32 `uriparametername:"limit"`
    // Offset
    Offset *int32 `uriparametername:"offset"`
}
// ItemFavoritepostsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFavoritepostsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFavoritepostsRequestBuilderGetQueryParameters
}
// NewItemFavoritepostsRequestBuilderInternal instantiates a new ItemFavoritepostsRequestBuilder and sets the default values.
func NewItemFavoritepostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFavoritepostsRequestBuilder) {
    m := &ItemFavoritepostsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/favoriteposts?limit={limit}&offset={offset}", pathParameters),
    }
    return m
}
// NewItemFavoritepostsRequestBuilder instantiates a new ItemFavoritepostsRequestBuilder and sets the default values.
func NewItemFavoritepostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFavoritepostsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFavoritepostsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get description
// returns a []Postable when successful
func (m *ItemFavoritepostsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFavoritepostsRequestBuilderGetRequestConfiguration)([]ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreatePostFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Postable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Postable)
        }
    }
    return val, nil
}
// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemFavoritepostsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFavoritepostsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFavoritepostsRequestBuilder when successful
func (m *ItemFavoritepostsRequestBuilder) WithUrl(rawUrl string)(*ItemFavoritepostsRequestBuilder) {
    return NewItemFavoritepostsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
