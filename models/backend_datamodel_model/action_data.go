package backend_datamodel_model

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionData struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The comment property
	comment *string
	// in seconds
	view_duration *float64
}

// NewActionData instantiates a new ActionData and sets the default values.
func NewActionData() *ActionData {
	m := &ActionData{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateActionDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewActionData(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionData) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetComment gets the comment property value. The comment property
// returns a *string when successful
func (m *ActionData) GetComment() *string {
	return m.comment
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionData) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["comment"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetComment(val)
		}
		return nil
	}
	res["view_duration"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetViewDuration(val)
		}
		return nil
	}
	return res
}

// GetViewDuration gets the view_duration property value. in seconds
// returns a *float64 when successful
func (m *ActionData) GetViewDuration() *float64 {
	return m.view_duration
}

// Serialize serializes information the current object
func (m *ActionData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("comment", m.GetComment())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("view_duration", m.GetViewDuration())
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
func (m *ActionData) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetComment sets the comment property value. The comment property
func (m *ActionData) SetComment(value *string) {
	m.comment = value
}

// SetViewDuration sets the view_duration property value. in seconds
func (m *ActionData) SetViewDuration(value *float64) {
	m.view_duration = value
}

type ActionDataable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetComment() *string
	GetViewDuration() *float64
	SetComment(value *string)
	SetViewDuration(value *float64)
}
