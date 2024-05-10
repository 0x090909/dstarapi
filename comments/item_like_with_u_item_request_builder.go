package comments

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemLikeWithUItemRequestBuilder builds and executes requests for operations under \comments\{id-id}\like\{uid}
type ItemLikeWithUItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemLikeWithUItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLikeWithUItemRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemLikeWithUItemRequestBuilderInternal instantiates a new ItemLikeWithUItemRequestBuilder and sets the default values.
func NewItemLikeWithUItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemLikeWithUItemRequestBuilder {
	m := &ItemLikeWithUItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/comments/{id%2Did}/like/{uid}", pathParameters),
	}
	return m
}

// NewItemLikeWithUItemRequestBuilder instantiates a new ItemLikeWithUItemRequestBuilder and sets the default values.
func NewItemLikeWithUItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemLikeWithUItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemLikeWithUItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Post description
// returns a *string when successful
func (m *ItemLikeWithUItemRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemLikeWithUItemRequestBuilderPostRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
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

// ToPostRequestInformation description
// returns a *RequestInformation when successful
func (m *ItemLikeWithUItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemLikeWithUItemRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemLikeWithUItemRequestBuilder when successful
func (m *ItemLikeWithUItemRequestBuilder) WithUrl(rawUrl string) *ItemLikeWithUItemRequestBuilder {
	return NewItemLikeWithUItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
