package backend_datamodel_model

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationPreference struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The comments property
	comments *bool
	// The follows property
	follows *bool
	// The likes property
	likes *bool
	// The pause_all property
	pause_all *bool
	// The profile_views property
	profile_views *bool
	// The tags property
	tags *bool
}

// NewNotificationPreference instantiates a new NotificationPreference and sets the default values.
func NewNotificationPreference() *NotificationPreference {
	m := &NotificationPreference{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNotificationPreferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationPreferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotificationPreference(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NotificationPreference) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetComments gets the comments property value. The comments property
// returns a *bool when successful
func (m *NotificationPreference) GetComments() *bool {
	return m.comments
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationPreference) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["comments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetComments(val)
		}
		return nil
	}
	res["follows"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFollows(val)
		}
		return nil
	}
	res["likes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLikes(val)
		}
		return nil
	}
	res["pause_all"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPauseAll(val)
		}
		return nil
	}
	res["profile_views"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProfileViews(val)
		}
		return nil
	}
	res["tags"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTags(val)
		}
		return nil
	}
	return res
}

// GetFollows gets the follows property value. The follows property
// returns a *bool when successful
func (m *NotificationPreference) GetFollows() *bool {
	return m.follows
}

// GetLikes gets the likes property value. The likes property
// returns a *bool when successful
func (m *NotificationPreference) GetLikes() *bool {
	return m.likes
}

// GetPauseAll gets the pause_all property value. The pause_all property
// returns a *bool when successful
func (m *NotificationPreference) GetPauseAll() *bool {
	return m.pause_all
}

// GetProfileViews gets the profile_views property value. The profile_views property
// returns a *bool when successful
func (m *NotificationPreference) GetProfileViews() *bool {
	return m.profile_views
}

// GetTags gets the tags property value. The tags property
// returns a *bool when successful
func (m *NotificationPreference) GetTags() *bool {
	return m.tags
}

// Serialize serializes information the current object
func (m *NotificationPreference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("comments", m.GetComments())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("follows", m.GetFollows())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("likes", m.GetLikes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("pause_all", m.GetPauseAll())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("profile_views", m.GetProfileViews())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("tags", m.GetTags())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotificationPreference) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetComments sets the comments property value. The comments property
func (m *NotificationPreference) SetComments(value *bool) {
	m.comments = value
}

// SetFollows sets the follows property value. The follows property
func (m *NotificationPreference) SetFollows(value *bool) {
	m.follows = value
}

// SetLikes sets the likes property value. The likes property
func (m *NotificationPreference) SetLikes(value *bool) {
	m.likes = value
}

// SetPauseAll sets the pause_all property value. The pause_all property
func (m *NotificationPreference) SetPauseAll(value *bool) {
	m.pause_all = value
}

// SetProfileViews sets the profile_views property value. The profile_views property
func (m *NotificationPreference) SetProfileViews(value *bool) {
	m.profile_views = value
}

// SetTags sets the tags property value. The tags property
func (m *NotificationPreference) SetTags(value *bool) {
	m.tags = value
}

type NotificationPreferenceable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetComments() *bool
	GetFollows() *bool
	GetLikes() *bool
	GetPauseAll() *bool
	GetProfileViews() *bool
	GetTags() *bool
	SetComments(value *bool)
	SetFollows(value *bool)
	SetLikes(value *bool)
	SetPauseAll(value *bool)
	SetProfileViews(value *bool)
	SetTags(value *bool)
}
