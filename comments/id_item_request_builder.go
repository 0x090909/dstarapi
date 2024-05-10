package comments

import (
	"context"
	ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdItemRequestBuilder builds and executes requests for operations under \comments\{id-id}
type IdItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// IdItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByUid gets an item from the github.com/0x090909/dstarapi.comments.item.item collection
// returns a *ItemWithUItemRequestBuilder when successful
func (m *IdItemRequestBuilder) ByUid(uid string) *ItemWithUItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if uid != "" {
		urlTplParams["uid"] = uid
	}
	return NewItemWithUItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewIdItemRequestBuilderInternal instantiates a new IdItemRequestBuilder and sets the default values.
func NewIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *IdItemRequestBuilder {
	m := &IdItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/comments/{id%2Did}", pathParameters),
	}
	return m
}

// NewIdItemRequestBuilder instantiates a new IdItemRequestBuilder and sets the default values.
func NewIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *IdItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewIdItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Like the like property
// returns a *ItemLikeRequestBuilder when successful
func (m *IdItemRequestBuilder) Like() *ItemLikeRequestBuilder {
	return NewItemLikeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Put description
// returns a *string when successful
func (m *IdItemRequestBuilder) Put(ctx context.Context, body ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Commentable, requestConfiguration *IdItemRequestBuilderPutRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
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

// ToPutRequestInformation description
// returns a *RequestInformation when successful
func (m *IdItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Commentable, requestConfiguration *IdItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "text/plain")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdItemRequestBuilder when successful
func (m *IdItemRequestBuilder) WithUrl(rawUrl string) *IdItemRequestBuilder {
	return NewIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
