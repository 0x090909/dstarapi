package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
)

// WithUItemRequestBuilder builds and executes requests for operations under \users\{uid}
type WithUItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithUItemRequestBuilderGetQueryParameters description
type WithUItemRequestBuilderGetQueryParameters struct {
    // Visitor UID
    Visitor_uid *string `uriparametername:"visitor_uid"`
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
// Awards the awards property
// returns a *ItemAwardsRequestBuilder when successful
func (m *WithUItemRequestBuilder) Awards()(*ItemAwardsRequestBuilder) {
    return NewItemAwardsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByPostId gets an item from the github.com/0x090909/dstarapi.users.item.item collection
// returns a *ItemWithPostItemRequestBuilder when successful
func (m *WithUItemRequestBuilder) ByPostId(postId string)(*ItemWithPostItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if postId != "" {
        urlTplParams["postId"] = postId
    }
    return NewItemWithPostItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Check the check property
// returns a *ItemCheckRequestBuilder when successful
func (m *WithUItemRequestBuilder) Check()(*ItemCheckRequestBuilder) {
    return NewItemCheckRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithUItemRequestBuilderInternal instantiates a new WithUItemRequestBuilder and sets the default values.
func NewWithUItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUItemRequestBuilder) {
    m := &WithUItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}{?visitor_uid*}", pathParameters),
    }
    return m
}
// NewWithUItemRequestBuilder instantiates a new WithUItemRequestBuilder and sets the default values.
func NewWithUItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithUItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Favoriteposts the favoriteposts property
// returns a *ItemFavoritepostsRequestBuilder when successful
func (m *WithUItemRequestBuilder) Favoriteposts()(*ItemFavoritepostsRequestBuilder) {
    return NewItemFavoritepostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Follow the follow property
// returns a *ItemFollowRequestBuilder when successful
func (m *WithUItemRequestBuilder) Follow()(*ItemFollowRequestBuilder) {
    return NewItemFollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Followers the followers property
// returns a *ItemFollowersRequestBuilder when successful
func (m *WithUItemRequestBuilder) Followers()(*ItemFollowersRequestBuilder) {
    return NewItemFollowersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Following the following property
// returns a *ItemFollowingRequestBuilder when successful
func (m *WithUItemRequestBuilder) Following()(*ItemFollowingRequestBuilder) {
    return NewItemFollowingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get description
// returns a Userable when successful
func (m *WithUItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithUItemRequestBuilderGetRequestConfiguration)(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreateUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Userable), nil
}
// Preference the preference property
// returns a *ItemPreferenceRequestBuilder when successful
func (m *WithUItemRequestBuilder) Preference()(*ItemPreferenceRequestBuilder) {
    return NewItemPreferenceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation description
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
// Unfollow the unfollow property
// returns a *ItemUnfollowRequestBuilder when successful
func (m *WithUItemRequestBuilder) Unfollow()(*ItemUnfollowRequestBuilder) {
    return NewItemUnfollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Upload the upload property
// returns a *ItemUploadRequestBuilder when successful
func (m *WithUItemRequestBuilder) Upload()(*ItemUploadRequestBuilder) {
    return NewItemUploadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithUItemRequestBuilder when successful
func (m *WithUItemRequestBuilder) WithUrl(rawUrl string)(*WithUItemRequestBuilder) {
    return NewWithUItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
