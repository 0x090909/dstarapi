package backend_datamodel_model

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Preference struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The categories property
	categories []int32
	// The notifications property
	notifications NotificationPreferenceable
	// The onboarding_showed property
	onboarding_showed *bool
	// The right_handed property
	right_handed *bool
}

// NewPreference instantiates a new Preference and sets the default values.
func NewPreference() *Preference {
	m := &Preference{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePreferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePreferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPreference(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Preference) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategories gets the categories property value. The categories property
// returns a []int32 when successful
func (m *Preference) GetCategories() []int32 {
	return m.categories
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Preference) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["categories"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("int32")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]int32, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*int32))
				}
			}
			m.SetCategories(res)
		}
		return nil
	}
	res["notifications"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNotificationPreferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotifications(val.(NotificationPreferenceable))
		}
		return nil
	}
	res["onboarding_showed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOnboardingShowed(val)
		}
		return nil
	}
	res["right_handed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRightHanded(val)
		}
		return nil
	}
	return res
}

// GetNotifications gets the notifications property value. The notifications property
// returns a NotificationPreferenceable when successful
func (m *Preference) GetNotifications() NotificationPreferenceable {
	return m.notifications
}

// GetOnboardingShowed gets the onboarding_showed property value. The onboarding_showed property
// returns a *bool when successful
func (m *Preference) GetOnboardingShowed() *bool {
	return m.onboarding_showed
}

// GetRightHanded gets the right_handed property value. The right_handed property
// returns a *bool when successful
func (m *Preference) GetRightHanded() *bool {
	return m.right_handed
}

// Serialize serializes information the current object
func (m *Preference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCategories() != nil {
		err := writer.WriteCollectionOfInt32Values("categories", m.GetCategories())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("notifications", m.GetNotifications())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("onboarding_showed", m.GetOnboardingShowed())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("right_handed", m.GetRightHanded())
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
func (m *Preference) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategories sets the categories property value. The categories property
func (m *Preference) SetCategories(value []int32) {
	m.categories = value
}

// SetNotifications sets the notifications property value. The notifications property
func (m *Preference) SetNotifications(value NotificationPreferenceable) {
	m.notifications = value
}

// SetOnboardingShowed sets the onboarding_showed property value. The onboarding_showed property
func (m *Preference) SetOnboardingShowed(value *bool) {
	m.onboarding_showed = value
}

// SetRightHanded sets the right_handed property value. The right_handed property
func (m *Preference) SetRightHanded(value *bool) {
	m.right_handed = value
}

type Preferenceable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCategories() []int32
	GetNotifications() NotificationPreferenceable
	GetOnboardingShowed() *bool
	GetRightHanded() *bool
	SetCategories(value []int32)
	SetNotifications(value NotificationPreferenceable)
	SetOnboardingShowed(value *bool)
	SetRightHanded(value *bool)
}
