package users

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemUploadRequestBuilder builds and executes requests for operations under \users\{uid}\upload
type ItemUploadRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewItemUploadRequestBuilderInternal instantiates a new ItemUploadRequestBuilder and sets the default values.
func NewItemUploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemUploadRequestBuilder {
	m := &ItemUploadRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/upload", pathParameters),
	}
	return m
}

// NewItemUploadRequestBuilder instantiates a new ItemUploadRequestBuilder and sets the default values.
func NewItemUploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemUploadRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemUploadRequestBuilderInternal(urlParams, requestAdapter)
}

// Image the image property
// returns a *ItemUploadImageRequestBuilder when successful
func (m *ItemUploadRequestBuilder) Image() *ItemUploadImageRequestBuilder {
	return NewItemUploadImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
