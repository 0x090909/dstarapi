package comments

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemLikeRequestBuilder builds and executes requests for operations under \comments\{id-id}\like
type ItemLikeRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByUid gets an item from the github.com/0x090909/dstarapi.comments.item.like.item collection
// returns a *ItemLikeWithUItemRequestBuilder when successful
func (m *ItemLikeRequestBuilder) ByUid(uid string) *ItemLikeWithUItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if uid != "" {
		urlTplParams["uid"] = uid
	}
	return NewItemLikeWithUItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemLikeRequestBuilderInternal instantiates a new ItemLikeRequestBuilder and sets the default values.
func NewItemLikeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemLikeRequestBuilder {
	m := &ItemLikeRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/comments/{id%2Did}/like", pathParameters),
	}
	return m
}

// NewItemLikeRequestBuilder instantiates a new ItemLikeRequestBuilder and sets the default values.
func NewItemLikeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemLikeRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemLikeRequestBuilderInternal(urlParams, requestAdapter)
}
