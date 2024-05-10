package leaderboard

import (
	"context"
	i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LeaderboardRequestBuilder builds and executes requests for operations under \leaderboard
type LeaderboardRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// LeaderboardRequestBuilderGetQueryParameters description
type LeaderboardRequestBuilderGetQueryParameters struct {
	// Category
	Category *int32 `uriparametername:"category"`
	// Limit
	Limit *int32 `uriparametername:"limit"`
	// Offset
	Offset *int32 `uriparametername:"offset"`
	// User
	Uid *string `uriparametername:"uid"`
	// Yesterday Leaderboard
	Yesterday *bool `uriparametername:"yesterday"`
}

// LeaderboardRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LeaderboardRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *LeaderboardRequestBuilderGetQueryParameters
}

// NewLeaderboardRequestBuilderInternal instantiates a new LeaderboardRequestBuilder and sets the default values.
func NewLeaderboardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LeaderboardRequestBuilder {
	m := &LeaderboardRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/leaderboard?category={category}&limit={limit}&offset={offset}&uid={uid}{&yesterday*}", pathParameters),
	}
	return m
}

// NewLeaderboardRequestBuilder instantiates a new LeaderboardRequestBuilder and sets the default values.
func NewLeaderboardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LeaderboardRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLeaderboardRequestBuilderInternal(urlParams, requestAdapter)
}

// Get description
// returns a PaginatedLeaderboardable when successful
func (m *LeaderboardRequestBuilder) Get(ctx context.Context, requestConfiguration *LeaderboardRequestBuilderGetRequestConfiguration) (i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.PaginatedLeaderboardable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.CreatePaginatedLeaderboardFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.PaginatedLeaderboardable), nil
}

// ToGetRequestInformation description
// returns a *RequestInformation when successful
func (m *LeaderboardRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LeaderboardRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LeaderboardRequestBuilder when successful
func (m *LeaderboardRequestBuilder) WithUrl(rawUrl string) *LeaderboardRequestBuilder {
	return NewLeaderboardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
