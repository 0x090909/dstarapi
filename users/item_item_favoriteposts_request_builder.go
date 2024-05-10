package users

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemFavoritepostsRequestBuilder builds and executes requests for operations under \users\{uid}\{postId}\favoriteposts
type ItemItemFavoritepostsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemFavoritepostsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemFavoritepostsRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemItemFavoritepostsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemFavoritepostsRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemFavoritepostsRequestBuilderInternal instantiates a new ItemItemFavoritepostsRequestBuilder and sets the default values.
func NewItemItemFavoritepostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemFavoritepostsRequestBuilder {
	m := &ItemItemFavoritepostsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{uid}/{postId}/favoriteposts", pathParameters),
	}
	return m
}

// NewItemItemFavoritepostsRequestBuilder instantiates a new ItemItemFavoritepostsRequestBuilder and sets the default values.
func NewItemItemFavoritepostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemFavoritepostsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemFavoritepostsRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete description
// returns a *string when successful
func (m *ItemItemFavoritepostsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemFavoritepostsRequestBuilderDeleteRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*string), nil
}

// Put description
// returns a *string when successful
func (m *ItemItemFavoritepostsRequestBuilder) Put(ctx context.Context, requestConfiguration *ItemItemFavoritepostsRequestBuilderPutRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*string), nil
}

// ToDeleteRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemItemFavoritepostsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemFavoritepostsRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	return requestInfo, nil
}

// ToPutRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemItemFavoritepostsRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *ItemItemFavoritepostsRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemFavoritepostsRequestBuilder when successful
func (m *ItemItemFavoritepostsRequestBuilder) WithUrl(rawUrl string) *ItemItemFavoritepostsRequestBuilder {
	return NewItemItemFavoritepostsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
