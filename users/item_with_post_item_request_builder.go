package users

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithPostItemRequestBuilder builds and executes requests for operations under \users\{uid}\{postId}
type ItemWithPostItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewItemWithPostItemRequestBuilderInternal instantiates a new ItemWithPostItemRequestBuilder and sets the default values.
func NewItemWithPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemWithPostItemRequestBuilder {
	m := &ItemWithPostItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/{postId}", pathParameters),
	}
	return m
}

// NewItemWithPostItemRequestBuilder instantiates a new ItemWithPostItemRequestBuilder and sets the default values.
func NewItemWithPostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemWithPostItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemWithPostItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Favoriteposts the favoriteposts property
// returns a *ItemItemFavoritepostsRequestBuilder when successful
func (m *ItemWithPostItemRequestBuilder) Favoriteposts() *ItemItemFavoritepostsRequestBuilder {
	return NewItemItemFavoritepostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
