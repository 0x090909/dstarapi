package users

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCheckRequestBuilder builds and executes requests for operations under \users\{uid}\check
type ItemCheckRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByUsername gets an item from the github.com/0x090909/dstarapi.users.item.check.item collection
// returns a *ItemCheckWithUsernameItemRequestBuilder when successful
func (m *ItemCheckRequestBuilder) ByUsername(username string) *ItemCheckWithUsernameItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if username != "" {
		urlTplParams["username"] = username
	}
	return NewItemCheckWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemCheckRequestBuilderInternal instantiates a new ItemCheckRequestBuilder and sets the default values.
func NewItemCheckRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCheckRequestBuilder {
	m := &ItemCheckRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/check", pathParameters),
	}
	return m
}

// NewItemCheckRequestBuilder instantiates a new ItemCheckRequestBuilder and sets the default values.
func NewItemCheckRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCheckRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCheckRequestBuilderInternal(urlParams, requestAdapter)
}
