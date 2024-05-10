package backend_api_apimodel

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CheckUsername struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The available property
	available *bool
	// The suggested_username property
	suggested_username *string
}

// NewCheckUsername instantiates a new CheckUsername and sets the default values.
func NewCheckUsername() *CheckUsername {
	m := &CheckUsername{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCheckUsernameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckUsernameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCheckUsername(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckUsername) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAvailable gets the available property value. The available property
// returns a *bool when successful
func (m *CheckUsername) GetAvailable() *bool {
	return m.available
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckUsername) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["available"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAvailable(val)
		}
		return nil
	}
	res["suggested_username"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSuggestedUsername(val)
		}
		return nil
	}
	return res
}

// GetSuggestedUsername gets the suggested_username property value. The suggested_username property
// returns a *string when successful
func (m *CheckUsername) GetSuggestedUsername() *string {
	return m.suggested_username
}

// Serialize serializes information the current object
func (m *CheckUsername) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("available", m.GetAvailable())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("suggested_username", m.GetSuggestedUsername())
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
func (m *CheckUsername) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAvailable sets the available property value. The available property
func (m *CheckUsername) SetAvailable(value *bool) {
	m.available = value
}

// SetSuggestedUsername sets the suggested_username property value. The suggested_username property
func (m *CheckUsername) SetSuggestedUsername(value *string) {
	m.suggested_username = value
}

type CheckUsernameable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAvailable() *bool
	GetSuggestedUsername() *string
	SetAvailable(value *bool)
	SetSuggestedUsername(value *string)
}
