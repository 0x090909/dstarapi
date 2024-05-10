package contentjob

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContentJobRequestBuilder builds and executes requests for operations under \content-job
type ContentJobRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ById gets an item from the github.com/0x090909/dstarapi.contentJob.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ContentJobItemRequestBuilder when successful
func (m *ContentJobRequestBuilder) ById(id string)(*ContentJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewContentJobItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInteger gets an item from the github.com/0x090909/dstarapi.contentJob.item collection
// returns a *ContentJobItemRequestBuilder when successful
func (m *ContentJobRequestBuilder) ByIdInteger(id int32)(*ContentJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewContentJobItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewContentJobRequestBuilderInternal instantiates a new ContentJobRequestBuilder and sets the default values.
func NewContentJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentJobRequestBuilder) {
    m := &ContentJobRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/content-job", pathParameters),
    }
    return m
}
// NewContentJobRequestBuilder instantiates a new ContentJobRequestBuilder and sets the default values.
func NewContentJobRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentJobRequestBuilderInternal(urlParams, requestAdapter)
}
// Create the create property
// returns a *CreateRequestBuilder when successful
func (m *ContentJobRequestBuilder) Create()(*CreateRequestBuilder) {
    return NewCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeletePath the deletePath property
// returns a *DeleteRequestBuilder when successful
func (m *ContentJobRequestBuilder) DeletePath()(*DeleteRequestBuilder) {
    return NewDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Update the update property
// returns a *UpdateRequestBuilder when successful
func (m *ContentJobRequestBuilder) Update()(*UpdateRequestBuilder) {
    return NewUpdateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
