package backend_datamodel_model

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionEvent struct {
	// The additional_dataProperty property
	additional_dataProperty ActionDataable
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The post_id property
	post_id *int32
	// The type property
	typeEscaped *ActionType
	// The user_uid property
	user_uid *string
}

// NewActionEvent instantiates a new ActionEvent and sets the default values.
func NewActionEvent() *ActionEvent {
	m := &ActionEvent{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateActionEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewActionEvent(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionEvent) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAdditionalDataProperty gets the additional_data property value. The additional_dataProperty property
// returns a ActionDataable when successful
func (m *ActionEvent) GetAdditionalDataProperty() ActionDataable {
	return m.additional_dataProperty
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionEvent) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["additional_data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateActionDataFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAdditionalDataProperty(val.(ActionDataable))
		}
		return nil
	}
	res["post_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPostId(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseActionType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val.(*ActionType))
		}
		return nil
	}
	res["user_uid"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserUid(val)
		}
		return nil
	}
	return res
}

// GetPostId gets the post_id property value. The post_id property
// returns a *int32 when successful
func (m *ActionEvent) GetPostId() *int32 {
	return m.post_id
}

// GetTypeEscaped gets the type property value. The type property
// returns a *ActionType when successful
func (m *ActionEvent) GetTypeEscaped() *ActionType {
	return m.typeEscaped
}

// GetUserUid gets the user_uid property value. The user_uid property
// returns a *string when successful
func (m *ActionEvent) GetUserUid() *string {
	return m.user_uid
}

// Serialize serializes information the current object
func (m *ActionEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("additional_data", m.GetAdditionalDataProperty())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("post_id", m.GetPostId())
		if err != nil {
			return err
		}
	}
	if m.GetTypeEscaped() != nil {
		cast := (*m.GetTypeEscaped()).String()
		err := writer.WriteStringValue("type", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("user_uid", m.GetUserUid())
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
func (m *ActionEvent) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAdditionalDataProperty sets the additional_data property value. The additional_dataProperty property
func (m *ActionEvent) SetAdditionalDataProperty(value ActionDataable) {
	m.additional_dataProperty = value
}

// SetPostId sets the post_id property value. The post_id property
func (m *ActionEvent) SetPostId(value *int32) {
	m.post_id = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *ActionEvent) SetTypeEscaped(value *ActionType) {
	m.typeEscaped = value
}

// SetUserUid sets the user_uid property value. The user_uid property
func (m *ActionEvent) SetUserUid(value *string) {
	m.user_uid = value
}

type ActionEventable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAdditionalDataProperty() ActionDataable
	GetPostId() *int32
	GetTypeEscaped() *ActionType
	GetUserUid() *string
	SetAdditionalDataProperty(value ActionDataable)
	SetPostId(value *int32)
	SetTypeEscaped(value *ActionType)
	SetUserUid(value *string)
}
