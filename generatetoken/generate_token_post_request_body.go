package generatetoken

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GenerateTokenPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Password
	pass *string
	// Username
	user *string
}

// NewGenerateTokenPostRequestBody instantiates a new GenerateTokenPostRequestBody and sets the default values.
func NewGenerateTokenPostRequestBody() *GenerateTokenPostRequestBody {
	m := &GenerateTokenPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateGenerateTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGenerateTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewGenerateTokenPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GenerateTokenPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GenerateTokenPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["pass"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPass(val)
		}
		return nil
	}
	res["user"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUser(val)
		}
		return nil
	}
	return res
}

// GetPass gets the pass property value. Password
// returns a *string when successful
func (m *GenerateTokenPostRequestBody) GetPass() *string {
	return m.pass
}

// GetUser gets the user property value. Username
// returns a *string when successful
func (m *GenerateTokenPostRequestBody) GetUser() *string {
	return m.user
}

// Serialize serializes information the current object
func (m *GenerateTokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("pass", m.GetPass())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("user", m.GetUser())
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
func (m *GenerateTokenPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetPass sets the pass property value. Password
func (m *GenerateTokenPostRequestBody) SetPass(value *string) {
	m.pass = value
}

// SetUser sets the user property value. Username
func (m *GenerateTokenPostRequestBody) SetUser(value *string) {
	m.user = value
}

type GenerateTokenPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetPass() *string
	GetUser() *string
	SetPass(value *string)
	SetUser(value *string)
}
