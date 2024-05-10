package posts

import (
	"context"
	i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// PostsRequestBuilder builds and executes requests for operations under \posts
type PostsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// PostsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// PostsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByIdId gets an item from the github.com/0x090909/dstarapi.posts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *IdItemRequestBuilder when successful
func (m *PostsRequestBuilder) ByIdId(idId string) *IdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if idId != "" {
		urlTplParams["id%2Did"] = idId
	}
	return NewIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByIdIdInteger gets an item from the github.com/0x090909/dstarapi.posts.item collection
// returns a *IdItemRequestBuilder when successful
func (m *PostsRequestBuilder) ByIdIdInteger(idId int32) *IdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["id%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(idId), 10)
	return NewIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// Competing the competing property
// returns a *CompetingRequestBuilder when successful
func (m *PostsRequestBuilder) Competing() *CompetingRequestBuilder {
	return NewCompetingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewPostsRequestBuilderInternal instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *PostsRequestBuilder {
	m := &PostsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/posts", pathParameters),
	}
	return m
}

// NewPostsRequestBuilder instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *PostsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewPostsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post description
// returns a *string when successful
func (m *PostsRequestBuilder) Post(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Postable, requestConfiguration *PostsRequestBuilderPostRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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
func (m *PostsRequestBuilder) Put(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Postable, requestConfiguration *PostsRequestBuilderPutRequestConfiguration) (*string, error) {
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

// ToPostRequestInformation description
// returns a *RequestInformation when successful
func (m *PostsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Postable, requestConfiguration *PostsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// ToPutRequestInformation description
// returns a *RequestInformation when successful
func (m *PostsRequestBuilder) ToPutRequestInformation(ctx context.Context, body i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.Postable, requestConfiguration *PostsRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PostsRequestBuilder when successful
func (m *PostsRequestBuilder) WithUrl(rawUrl string) *PostsRequestBuilder {
	return NewPostsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
