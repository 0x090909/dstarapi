package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemFollowRequestBuilder builds and executes requests for operations under \users\{uid}\follow
type ItemFollowRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByTargetuid gets an item from the github.com/0x090909/dstarapi.users.item.follow.item collection
// returns a *ItemFollowWithTargetuItemRequestBuilder when successful
func (m *ItemFollowRequestBuilder) ByTargetuid(targetuid string)(*ItemFollowWithTargetuItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if targetuid != "" {
        urlTplParams["targetuid"] = targetuid
    }
    return NewItemFollowWithTargetuItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemFollowRequestBuilderInternal instantiates a new ItemFollowRequestBuilder and sets the default values.
func NewItemFollowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowRequestBuilder) {
    m := &ItemFollowRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/follow", pathParameters),
    }
    return m
}
// NewItemFollowRequestBuilder instantiates a new ItemFollowRequestBuilder and sets the default values.
func NewItemFollowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowRequestBuilderInternal(urlParams, requestAdapter)
}
