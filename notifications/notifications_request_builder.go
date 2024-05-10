package notifications

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// NotificationsRequestBuilder builds and executes requests for operations under \notifications
type NotificationsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByUid gets an item from the github.com/0x090909/dstarapi.notifications.item collection
// returns a *WithUItemRequestBuilder when successful
func (m *NotificationsRequestBuilder) ByUid(uid string) *WithUItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if uid != "" {
		urlTplParams["uid"] = uid
	}
	return NewWithUItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewNotificationsRequestBuilderInternal instantiates a new NotificationsRequestBuilder and sets the default values.
func NewNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotificationsRequestBuilder {
	m := &NotificationsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/notifications", pathParameters),
	}
	return m
}

// NewNotificationsRequestBuilder instantiates a new NotificationsRequestBuilder and sets the default values.
func NewNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotificationsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
