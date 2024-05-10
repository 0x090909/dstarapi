package posts

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompetingRequestBuilder builds and executes requests for operations under \posts\competing
type CompetingRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByUid gets an item from the github.com/0x090909/dstarapi.posts.competing.item collection
// returns a *CompetingWithUItemRequestBuilder when successful
func (m *CompetingRequestBuilder) ByUid(uid string) *CompetingWithUItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if uid != "" {
		urlTplParams["uid"] = uid
	}
	return NewCompetingWithUItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCompetingRequestBuilderInternal instantiates a new CompetingRequestBuilder and sets the default values.
func NewCompetingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompetingRequestBuilder {
	m := &CompetingRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/posts/competing", pathParameters),
	}
	return m
}

// NewCompetingRequestBuilder instantiates a new CompetingRequestBuilder and sets the default values.
func NewCompetingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompetingRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCompetingRequestBuilderInternal(urlParams, requestAdapter)
}
