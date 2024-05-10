package dstarapi

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i073398357348949c64d8ab9f7bd570a1a0060b70c88f477bfbb09bba5699cf40 "github.com/0x090909/dstarapi/comments"
    i2f8bb8723e3c954beea05ff56829755e6a7dc2127ebd38ae73305f4fc5146b39 "github.com/0x090909/dstarapi/notifications"
    i55c66ee0072b3268d6ac67b95b950bf6235906d43f8bf7af032fb63a04a6d0a8 "github.com/0x090909/dstarapi/actions"
    i5aad8d32236ed814e7512f117b5a38a945f51d9ea15af08638c480d9456883a1 "github.com/0x090909/dstarapi/contentjob"
    i778f0f01b09106348769da27c908aca950b5b959229ff2bf92ebf1524bfb39f8 "github.com/0x090909/dstarapi/upload"
    i7b43c66a71b7600594649aa4e7801834cd09c030c085467ef252a199b814ac89 "github.com/0x090909/dstarapi/posts"
    i86371f64b86a1fd1c8dee03f3ae652a70bf490d0327d07b4d1efa8373c953e75 "github.com/0x090909/dstarapi/post"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i89610620f84f4c6bf71916cd12d08498854b940782b567de5a4fec720f9f0a44 "github.com/0x090909/dstarapi/leaderboard"
    i8c127d7ff6243e8df3011c326fc569c638d28917e5307178babfa2c92b93d5d1 "github.com/0x090909/dstarapi/users"
    ib28131b64f983c0120477707deda924965947a95d39e825e06caa043ea6abb34 "github.com/0x090909/dstarapi/report"
    ifc17c062385653bb9f1866fd29a51b879b28fbf1af16e8dacc7700bea7d25546 "github.com/0x090909/dstarapi/generatetoken"
)

// DailyStarApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type DailyStarApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Actions the actions property
// returns a *ActionsRequestBuilder when successful
func (m *DailyStarApiClient) Actions()(*i55c66ee0072b3268d6ac67b95b950bf6235906d43f8bf7af032fb63a04a6d0a8.ActionsRequestBuilder) {
    return i55c66ee0072b3268d6ac67b95b950bf6235906d43f8bf7af032fb63a04a6d0a8.NewActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
// returns a *CommentsRequestBuilder when successful
func (m *DailyStarApiClient) Comments()(*i073398357348949c64d8ab9f7bd570a1a0060b70c88f477bfbb09bba5699cf40.CommentsRequestBuilder) {
    return i073398357348949c64d8ab9f7bd570a1a0060b70c88f477bfbb09bba5699cf40.NewCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDailyStarApiClient instantiates a new DailyStarApiClient and sets the default values.
func NewDailyStarApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyStarApiClient) {
    m := &DailyStarApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://ds-be-staging.up.railway.app")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// ContentJob the contentJob property
// returns a *ContentJobRequestBuilder when successful
func (m *DailyStarApiClient) ContentJob()(*i5aad8d32236ed814e7512f117b5a38a945f51d9ea15af08638c480d9456883a1.ContentJobRequestBuilder) {
    return i5aad8d32236ed814e7512f117b5a38a945f51d9ea15af08638c480d9456883a1.NewContentJobRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateToken the generateToken property
// returns a *GenerateTokenRequestBuilder when successful
func (m *DailyStarApiClient) GenerateToken()(*ifc17c062385653bb9f1866fd29a51b879b28fbf1af16e8dacc7700bea7d25546.GenerateTokenRequestBuilder) {
    return ifc17c062385653bb9f1866fd29a51b879b28fbf1af16e8dacc7700bea7d25546.NewGenerateTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Leaderboard the leaderboard property
// returns a *LeaderboardRequestBuilder when successful
func (m *DailyStarApiClient) Leaderboard()(*i89610620f84f4c6bf71916cd12d08498854b940782b567de5a4fec720f9f0a44.LeaderboardRequestBuilder) {
    return i89610620f84f4c6bf71916cd12d08498854b940782b567de5a4fec720f9f0a44.NewLeaderboardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
// returns a *NotificationsRequestBuilder when successful
func (m *DailyStarApiClient) Notifications()(*i2f8bb8723e3c954beea05ff56829755e6a7dc2127ebd38ae73305f4fc5146b39.NotificationsRequestBuilder) {
    return i2f8bb8723e3c954beea05ff56829755e6a7dc2127ebd38ae73305f4fc5146b39.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PostPath the postPath property
// returns a *PostRequestBuilder when successful
func (m *DailyStarApiClient) PostPath()(*i86371f64b86a1fd1c8dee03f3ae652a70bf490d0327d07b4d1efa8373c953e75.PostRequestBuilder) {
    return i86371f64b86a1fd1c8dee03f3ae652a70bf490d0327d07b4d1efa8373c953e75.NewPostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Posts the posts property
// returns a *PostsRequestBuilder when successful
func (m *DailyStarApiClient) Posts()(*i7b43c66a71b7600594649aa4e7801834cd09c030c085467ef252a199b814ac89.PostsRequestBuilder) {
    return i7b43c66a71b7600594649aa4e7801834cd09c030c085467ef252a199b814ac89.NewPostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Report the report property
// returns a *ReportRequestBuilder when successful
func (m *DailyStarApiClient) Report()(*ib28131b64f983c0120477707deda924965947a95d39e825e06caa043ea6abb34.ReportRequestBuilder) {
    return ib28131b64f983c0120477707deda924965947a95d39e825e06caa043ea6abb34.NewReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Upload the upload property
// returns a *UploadRequestBuilder when successful
func (m *DailyStarApiClient) Upload()(*i778f0f01b09106348769da27c908aca950b5b959229ff2bf92ebf1524bfb39f8.UploadRequestBuilder) {
    return i778f0f01b09106348769da27c908aca950b5b959229ff2bf92ebf1524bfb39f8.NewUploadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *DailyStarApiClient) Users()(*i8c127d7ff6243e8df3011c326fc569c638d28917e5307178babfa2c92b93d5d1.UsersRequestBuilder) {
    return i8c127d7ff6243e8df3011c326fc569c638d28917e5307178babfa2c92b93d5d1.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
