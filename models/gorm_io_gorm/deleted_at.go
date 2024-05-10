package gorm_io_gorm

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeletedAt struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The time property
	time *string
	// Valid is true if Time is not NULL
	valid *bool
}

// NewDeletedAt instantiates a new DeletedAt and sets the default values.
func NewDeletedAt() *DeletedAt {
	m := &DeletedAt{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateDeletedAtFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedAtFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewDeletedAt(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeletedAt) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeletedAt) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTime(val)
		}
		return nil
	}
	res["valid"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValid(val)
		}
		return nil
	}
	return res
}

// GetTime gets the time property value. The time property
// returns a *string when successful
func (m *DeletedAt) GetTime() *string {
	return m.time
}

// GetValid gets the valid property value. Valid is true if Time is not NULL
// returns a *bool when successful
func (m *DeletedAt) GetValid() *bool {
	return m.valid
}

// Serialize serializes information the current object
func (m *DeletedAt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("time", m.GetTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("valid", m.GetValid())
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
func (m *DeletedAt) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetTime sets the time property value. The time property
func (m *DeletedAt) SetTime(value *string) {
	m.time = value
}

// SetValid sets the valid property value. Valid is true if Time is not NULL
func (m *DeletedAt) SetValid(value *bool) {
	m.valid = value
}

type DeletedAtable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetTime() *string
	GetValid() *bool
	SetTime(value *string)
	SetValid(value *bool)
}
