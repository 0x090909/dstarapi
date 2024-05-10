package apimodel

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7 "github.com/0x090909/dstarapi/models/model"
    i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5 "github.com/0x090909/dstarapi/models/types"
)

type Event struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The context property
    context i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.EventContextable
    // The id property
    id *int32
    // The origin property
    origin *int32
    // The seen_by_recipient property
    seen_by_recipient *bool
    // The seen_by_recipient_timestamp property
    seen_by_recipient_timestamp i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable
    // The timestamp property
    timestamp i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable
    // The type property
    typeEscaped *int32
}
// NewEvent instantiates a new Event and sets the default values.
func NewEvent()(*Event) {
    m := &Event{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Event) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContext gets the context property value. The context property
// returns a EventContextable when successful
func (m *Event) GetContext()(i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.EventContextable) {
    return m.context
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Event) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.CreateEventContextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val.(i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.EventContextable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["origin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrigin(val)
        }
        return nil
    }
    res["seen_by_recipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeenByRecipient(val)
        }
        return nil
    }
    res["seen_by_recipient_timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.CreateJSONTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeenByRecipientTimestamp(val.(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable))
        }
        return nil
    }
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.CreateJSONTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val.(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Event) GetId()(*int32) {
    return m.id
}
// GetOrigin gets the origin property value. The origin property
// returns a *int32 when successful
func (m *Event) GetOrigin()(*int32) {
    return m.origin
}
// GetSeenByRecipient gets the seen_by_recipient property value. The seen_by_recipient property
// returns a *bool when successful
func (m *Event) GetSeenByRecipient()(*bool) {
    return m.seen_by_recipient
}
// GetSeenByRecipientTimestamp gets the seen_by_recipient_timestamp property value. The seen_by_recipient_timestamp property
// returns a JSONTimeable when successful
func (m *Event) GetSeenByRecipientTimestamp()(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable) {
    return m.seen_by_recipient_timestamp
}
// GetTimestamp gets the timestamp property value. The timestamp property
// returns a JSONTimeable when successful
func (m *Event) GetTimestamp()(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable) {
    return m.timestamp
}
// GetTypeEscaped gets the type property value. The type property
// returns a *int32 when successful
func (m *Event) GetTypeEscaped()(*int32) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Event) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("context", m.GetContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("origin", m.GetOrigin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("seen_by_recipient", m.GetSeenByRecipient())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("seen_by_recipient_timestamp", m.GetSeenByRecipientTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timestamp", m.GetTimestamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("type", m.GetTypeEscaped())
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
func (m *Event) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContext sets the context property value. The context property
func (m *Event) SetContext(value i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.EventContextable)() {
    m.context = value
}
// SetId sets the id property value. The id property
func (m *Event) SetId(value *int32)() {
    m.id = value
}
// SetOrigin sets the origin property value. The origin property
func (m *Event) SetOrigin(value *int32)() {
    m.origin = value
}
// SetSeenByRecipient sets the seen_by_recipient property value. The seen_by_recipient property
func (m *Event) SetSeenByRecipient(value *bool)() {
    m.seen_by_recipient = value
}
// SetSeenByRecipientTimestamp sets the seen_by_recipient_timestamp property value. The seen_by_recipient_timestamp property
func (m *Event) SetSeenByRecipientTimestamp(value i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable)() {
    m.seen_by_recipient_timestamp = value
}
// SetTimestamp sets the timestamp property value. The timestamp property
func (m *Event) SetTimestamp(value i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable)() {
    m.timestamp = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Event) SetTypeEscaped(value *int32)() {
    m.typeEscaped = value
}
type Eventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.EventContextable)
    GetId()(*int32)
    GetOrigin()(*int32)
    GetSeenByRecipient()(*bool)
    GetSeenByRecipientTimestamp()(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable)
    GetTimestamp()(i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable)
    GetTypeEscaped()(*int32)
    SetContext(value i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.EventContextable)()
    SetId(value *int32)()
    SetOrigin(value *int32)()
    SetSeenByRecipient(value *bool)()
    SetSeenByRecipientTimestamp(value i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable)()
    SetTimestamp(value i47c923e7a77b37536a2867d55d0a56204f43c4f1e45f2065dc204c0f6f3c1bc5.JSONTimeable)()
    SetTypeEscaped(value *int32)()
}
