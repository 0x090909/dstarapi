package model

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EventContext struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The category property
    category *string
    // The diff property
    diff []string
    // The error_code property
    error_code *string
    // The likes property
    likes *int32
    // The position property
    position *int32
    // The previous_position property
    previous_position *int32
    // The score property
    score *int32
    // The views property
    views *int32
}
// NewEventContext instantiates a new EventContext and sets the default values.
func NewEventContext()(*EventContext) {
    m := &EventContext{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEventContextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEventContextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventContext(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EventContext) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCategory gets the category property value. The category property
// returns a *string when successful
func (m *EventContext) GetCategory()(*string) {
    return m.category
}
// GetDiff gets the diff property value. The diff property
// returns a []string when successful
func (m *EventContext) GetDiff()([]string) {
    return m.diff
}
// GetErrorCode gets the error_code property value. The error_code property
// returns a *string when successful
func (m *EventContext) GetErrorCode()(*string) {
    return m.error_code
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EventContext) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["diff"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDiff(res)
        }
        return nil
    }
    res["error_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["likes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLikes(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["previous_position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousPosition(val)
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["views"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViews(val)
        }
        return nil
    }
    return res
}
// GetLikes gets the likes property value. The likes property
// returns a *int32 when successful
func (m *EventContext) GetLikes()(*int32) {
    return m.likes
}
// GetPosition gets the position property value. The position property
// returns a *int32 when successful
func (m *EventContext) GetPosition()(*int32) {
    return m.position
}
// GetPreviousPosition gets the previous_position property value. The previous_position property
// returns a *int32 when successful
func (m *EventContext) GetPreviousPosition()(*int32) {
    return m.previous_position
}
// GetScore gets the score property value. The score property
// returns a *int32 when successful
func (m *EventContext) GetScore()(*int32) {
    return m.score
}
// GetViews gets the views property value. The views property
// returns a *int32 when successful
func (m *EventContext) GetViews()(*int32) {
    return m.views
}
// Serialize serializes information the current object
func (m *EventContext) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    if m.GetDiff() != nil {
        err := writer.WriteCollectionOfStringValues("diff", m.GetDiff())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error_code", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("likes", m.GetLikes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("previous_position", m.GetPreviousPosition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("views", m.GetViews())
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
func (m *EventContext) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCategory sets the category property value. The category property
func (m *EventContext) SetCategory(value *string)() {
    m.category = value
}
// SetDiff sets the diff property value. The diff property
func (m *EventContext) SetDiff(value []string)() {
    m.diff = value
}
// SetErrorCode sets the error_code property value. The error_code property
func (m *EventContext) SetErrorCode(value *string)() {
    m.error_code = value
}
// SetLikes sets the likes property value. The likes property
func (m *EventContext) SetLikes(value *int32)() {
    m.likes = value
}
// SetPosition sets the position property value. The position property
func (m *EventContext) SetPosition(value *int32)() {
    m.position = value
}
// SetPreviousPosition sets the previous_position property value. The previous_position property
func (m *EventContext) SetPreviousPosition(value *int32)() {
    m.previous_position = value
}
// SetScore sets the score property value. The score property
func (m *EventContext) SetScore(value *int32)() {
    m.score = value
}
// SetViews sets the views property value. The views property
func (m *EventContext) SetViews(value *int32)() {
    m.views = value
}
type EventContextable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*string)
    GetDiff()([]string)
    GetErrorCode()(*string)
    GetLikes()(*int32)
    GetPosition()(*int32)
    GetPreviousPosition()(*int32)
    GetScore()(*int32)
    GetViews()(*int32)
    SetCategory(value *string)()
    SetDiff(value []string)()
    SetErrorCode(value *string)()
    SetLikes(value *int32)()
    SetPosition(value *int32)()
    SetPreviousPosition(value *int32)()
    SetScore(value *int32)()
    SetViews(value *int32)()
}
