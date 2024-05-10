package users

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemUnfollowRequestBuilder builds and executes requests for operations under \users\{uid}\unfollow
type ItemUnfollowRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByTargetuid gets an item from the github.com/0x090909/dstarapi.users.item.unfollow.item collection
// returns a *ItemUnfollowWithTargetuItemRequestBuilder when successful
func (m *ItemUnfollowRequestBuilder) ByTargetuid(targetuid string) *ItemUnfollowWithTargetuItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if targetuid != "" {
		urlTplParams["targetuid"] = targetuid
	}
	return NewItemUnfollowWithTargetuItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemUnfollowRequestBuilderInternal instantiates a new ItemUnfollowRequestBuilder and sets the default values.
func NewItemUnfollowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemUnfollowRequestBuilder {
	m := &ItemUnfollowRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/unfollow", pathParameters),
	}
	return m
}

// NewItemUnfollowRequestBuilder instantiates a new ItemUnfollowRequestBuilder and sets the default values.
func NewItemUnfollowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemUnfollowRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemUnfollowRequestBuilderInternal(urlParams, requestAdapter)
}
