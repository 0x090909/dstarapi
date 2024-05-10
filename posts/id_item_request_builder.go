package posts

import (
	"context"
	ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// IdItemRequestBuilder builds and executes requests for operations under \posts\{id-id}
type IdItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// IdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// IdItemRequestBuilderGetQueryParameters description
type IdItemRequestBuilderGetQueryParameters struct {
	// Limit
	Limit *int32 `uriparametername:"limit"`
	// Offset
	Offset *int32 `uriparametername:"offset"`
}

// IdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *IdItemRequestBuilderGetQueryParameters
}

// ByCategoryId gets an item from the github.com/0x090909/dstarapi.posts.item.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemWithCategoryItemRequestBuilder when successful
func (m *IdItemRequestBuilder) ByCategoryId(categoryId string) *ItemWithCategoryItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if categoryId != "" {
		urlTplParams["categoryId"] = categoryId
	}
	return NewItemWithCategoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByCategoryIdInteger gets an item from the github.com/0x090909/dstarapi.posts.item.item collection
// returns a *ItemWithCategoryItemRequestBuilder when successful
func (m *IdItemRequestBuilder) ByCategoryIdInteger(categoryId int32) *ItemWithCategoryItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["categoryId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(categoryId), 10)
	return NewItemWithCategoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewIdItemRequestBuilderInternal instantiates a new IdItemRequestBuilder and sets the default values.
func NewIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *IdItemRequestBuilder {
	m := &IdItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/posts/{id%2Did}{?limit*,offset*}", pathParameters),
	}
	return m
}

// NewIdItemRequestBuilder instantiates a new IdItemRequestBuilder and sets the default values.
func NewIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *IdItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewIdItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete description
// returns a *string when successful
func (m *IdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdItemRequestBuilderDeleteRequestConfiguration) (*string, error) {
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

// Get description
// returns a PaginatedPostsable when successful
func (m *IdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IdItemRequestBuilderGetRequestConfiguration) (ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.PaginatedPostsable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.CreatePaginatedPostsFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.PaginatedPostsable), nil
}

// Thumbnail the thumbnail property
// returns a *ItemThumbnailRequestBuilder when successful
func (m *IdItemRequestBuilder) Thumbnail() *ItemThumbnailRequestBuilder {
	return NewItemThumbnailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation description
// returns a *RequestInformation when successful
func (m *IdItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IdItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/posts/{id%2Did}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	return requestInfo, nil
}

// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *IdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IdItemRequestBuilder when successful
func (m *IdItemRequestBuilder) WithUrl(rawUrl string) *IdItemRequestBuilder {
	return NewIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
