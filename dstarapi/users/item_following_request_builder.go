package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
)

// ItemFollowingRequestBuilder builds and executes requests for operations under \users\{uid}\following
type ItemFollowingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFollowingRequestBuilderGetQueryParameters description
type ItemFollowingRequestBuilderGetQueryParameters struct {
    // Limit
    Limit *int32 `uriparametername:"limit"`
    // Offset
    Offset *int32 `uriparametername:"offset"`
}
// ItemFollowingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFollowingRequestBuilderGetQueryParameters
}
// NewItemFollowingRequestBuilderInternal instantiates a new ItemFollowingRequestBuilder and sets the default values.
func NewItemFollowingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowingRequestBuilder) {
    m := &ItemFollowingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/following{?limit*,offset*}", pathParameters),
    }
    return m
}
// NewItemFollowingRequestBuilder instantiates a new ItemFollowingRequestBuilder and sets the default values.
func NewItemFollowingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get description
// returns a []Userable when successful
func (m *ItemFollowingRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFollowingRequestBuilderGetRequestConfiguration)([]i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.CreateUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Userable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Userable)
        }
    }
    return val, nil
}
// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemFollowingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFollowingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFollowingRequestBuilder when successful
func (m *ItemFollowingRequestBuilder) WithUrl(rawUrl string)(*ItemFollowingRequestBuilder) {
    return NewItemFollowingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
