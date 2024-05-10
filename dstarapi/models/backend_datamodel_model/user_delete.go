package backend_datamodel_model

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserDelete struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The createdAt property
    createdAt *string
    // The delete_cause property
    delete_cause *int32
    // The feedback property
    feedback *string
    // The uid property
    uid *string
}
// NewUserDelete instantiates a new UserDelete and sets the default values.
func NewUserDelete()(*UserDelete) {
    m := &UserDelete{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserDeleteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserDeleteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserDelete(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserDelete) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *string when successful
func (m *UserDelete) GetCreatedAt()(*string) {
    return m.createdAt
}
// GetDeleteCause gets the delete_cause property value. The delete_cause property
// returns a *int32 when successful
func (m *UserDelete) GetDeleteCause()(*int32) {
    return m.delete_cause
}
// GetFeedback gets the feedback property value. The feedback property
// returns a *string when successful
func (m *UserDelete) GetFeedback()(*string) {
    return m.feedback
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserDelete) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["delete_cause"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteCause(val)
        }
        return nil
    }
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val)
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
        }
        return nil
    }
    return res
}
// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *UserDelete) GetUid()(*string) {
    return m.uid
}
// Serialize serializes information the current object
func (m *UserDelete) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("delete_cause", m.GetDeleteCause())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("feedback", m.GetFeedback())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uid", m.GetUid())
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
func (m *UserDelete) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *UserDelete) SetCreatedAt(value *string)() {
    m.createdAt = value
}
// SetDeleteCause sets the delete_cause property value. The delete_cause property
func (m *UserDelete) SetDeleteCause(value *int32)() {
    m.delete_cause = value
}
// SetFeedback sets the feedback property value. The feedback property
func (m *UserDelete) SetFeedback(value *string)() {
    m.feedback = value
}
// SetUid sets the uid property value. The uid property
func (m *UserDelete) SetUid(value *string)() {
    m.uid = value
}
type UserDeleteable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*string)
    GetDeleteCause()(*int32)
    GetFeedback()(*string)
    GetUid()(*string)
    SetCreatedAt(value *string)()
    SetDeleteCause(value *int32)()
    SetFeedback(value *string)()
    SetUid(value *string)()
}
