package comments

import (
	"context"
	ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1 "github.com/0x090909/dstarapi/models/backend_api_apimodel"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CommentsRequestBuilder builds and executes requests for operations under \comments
type CommentsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CommentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommentsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByIdId gets an item from the github.com/0x090909/dstarapi.comments.item collection
// returns a *IdItemRequestBuilder when successful
func (m *CommentsRequestBuilder) ByIdId(idId string) *IdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if idId != "" {
		urlTplParams["id%2Did"] = idId
	}
	return NewIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCommentsRequestBuilderInternal instantiates a new CommentsRequestBuilder and sets the default values.
func NewCommentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CommentsRequestBuilder {
	m := &CommentsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/comments", pathParameters),
	}
	return m
}

// NewCommentsRequestBuilder instantiates a new CommentsRequestBuilder and sets the default values.
func NewCommentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CommentsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCommentsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post description
// returns a *string when successful
func (m *CommentsRequestBuilder) Post(ctx context.Context, body ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Commentable, requestConfiguration *CommentsRequestBuilderPostRequestConfiguration) (*string, error) {
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

// ToPostRequestInformation description
// returns a *RequestInformation when successful
func (m *CommentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia5e5c5a00aaeeae36464d84d89824912189cd78e368489dce036c5d454722fc1.Commentable, requestConfiguration *CommentsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CommentsRequestBuilder when successful
func (m *CommentsRequestBuilder) WithUrl(rawUrl string) *CommentsRequestBuilder {
	return NewCommentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
