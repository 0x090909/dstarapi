package backend_datamodel_model

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1 "github.com/0x090909/dstarapi/models/gorm_io_gorm"
)

type CommentLike struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The author_uid property
    author_uid *string
    // The comment_id property
    comment_id *int32
    // The createdAt property
    createdAt *string
    // The deletedAt property
    deletedAt i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
    // The id property
    id *int32
    // The updatedAt property
    updatedAt *string
}
// NewCommentLike instantiates a new CommentLike and sets the default values.
func NewCommentLike()(*CommentLike) {
    m := &CommentLike{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommentLikeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommentLikeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommentLike(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommentLike) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorUid gets the author_uid property value. The author_uid property
// returns a *string when successful
func (m *CommentLike) GetAuthorUid()(*string) {
    return m.author_uid
}
// GetCommentId gets the comment_id property value. The comment_id property
// returns a *int32 when successful
func (m *CommentLike) GetCommentId()(*int32) {
    return m.comment_id
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *string when successful
func (m *CommentLike) GetCreatedAt()(*string) {
    return m.createdAt
}
// GetDeletedAt gets the deletedAt property value. The deletedAt property
// returns a DeletedAtable when successful
func (m *CommentLike) GetDeletedAt()(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable) {
    return m.deletedAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommentLike) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author_uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorUid(val)
        }
        return nil
    }
    res["comment_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentId(val)
        }
        return nil
    }
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
    res["deletedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.CreateDeletedAtFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedAt(val.(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable))
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
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *CommentLike) GetId()(*int32) {
    return m.id
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *string when successful
func (m *CommentLike) GetUpdatedAt()(*string) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *CommentLike) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("author_uid", m.GetAuthorUid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("comment_id", m.GetCommentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deletedAt", m.GetDeletedAt())
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
        err := writer.WriteStringValue("updatedAt", m.GetUpdatedAt())
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
func (m *CommentLike) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorUid sets the author_uid property value. The author_uid property
func (m *CommentLike) SetAuthorUid(value *string)() {
    m.author_uid = value
}
// SetCommentId sets the comment_id property value. The comment_id property
func (m *CommentLike) SetCommentId(value *int32)() {
    m.comment_id = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *CommentLike) SetCreatedAt(value *string)() {
    m.createdAt = value
}
// SetDeletedAt sets the deletedAt property value. The deletedAt property
func (m *CommentLike) SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)() {
    m.deletedAt = value
}
// SetId sets the id property value. The id property
func (m *CommentLike) SetId(value *int32)() {
    m.id = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *CommentLike) SetUpdatedAt(value *string)() {
    m.updatedAt = value
}
type CommentLikeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorUid()(*string)
    GetCommentId()(*int32)
    GetCreatedAt()(*string)
    GetDeletedAt()(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)
    GetId()(*int32)
    GetUpdatedAt()(*string)
    SetAuthorUid(value *string)()
    SetCommentId(value *int32)()
    SetCreatedAt(value *string)()
    SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)()
    SetId(value *int32)()
    SetUpdatedAt(value *string)()
}
