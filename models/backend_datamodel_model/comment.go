package backend_datamodel_model

import (
	i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1 "github.com/0x090909/dstarapi/models/gorm_io_gorm"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Comment struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The author property
	author Userable
	// The author_uid property
	author_uid *string
	// The comment_likes property
	comment_likes []CommentLikeable
	// The createdAt property
	createdAt *string
	// The deletedAt property
	deletedAt i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
	// The description property
	description *string
	// The id property
	id *int32
	// The like property
	like *bool
	// The likes_count property
	likes_count *int32
	// The parent_id property
	parent_id *int32
	// The post_id property
	post_id *int32
	// The replies property
	replies []Commentable
	// The reply_count property
	reply_count *int32
	// The updatedAt property
	updatedAt *string
}

// NewComment instantiates a new Comment and sets the default values.
func NewComment() *Comment {
	m := &Comment{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewComment(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Comment) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAuthor gets the author property value. The author property
// returns a Userable when successful
func (m *Comment) GetAuthor() Userable {
	return m.author
}

// GetAuthorUid gets the author_uid property value. The author_uid property
// returns a *string when successful
func (m *Comment) GetAuthorUid() *string {
	return m.author_uid
}

// GetCommentLikes gets the comment_likes property value. The comment_likes property
// returns a []CommentLikeable when successful
func (m *Comment) GetCommentLikes() []CommentLikeable {
	return m.comment_likes
}

// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *string when successful
func (m *Comment) GetCreatedAt() *string {
	return m.createdAt
}

// GetDeletedAt gets the deletedAt property value. The deletedAt property
// returns a DeletedAtable when successful
func (m *Comment) GetDeletedAt() i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable {
	return m.deletedAt
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Comment) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Comment) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["author"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAuthor(val.(Userable))
		}
		return nil
	}
	res["author_uid"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAuthorUid(val)
		}
		return nil
	}
	res["comment_likes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCommentLikeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CommentLikeable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CommentLikeable)
				}
			}
			m.SetCommentLikes(res)
		}
		return nil
	}
	res["createdAt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedAt(val)
		}
		return nil
	}
	res["deletedAt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.CreateDeletedAtFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDeletedAt(val.(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable))
		}
		return nil
	}
	res["description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDescription(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["like"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLike(val)
		}
		return nil
	}
	res["likes_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLikesCount(val)
		}
		return nil
	}
	res["parent_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetParentId(val)
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
	res["replies"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCommentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]Commentable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(Commentable)
				}
			}
			m.SetReplies(res)
		}
		return nil
	}
	res["reply_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReplyCount(val)
		}
		return nil
	}
	res["updatedAt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *Comment) GetId() *int32 {
	return m.id
}

// GetLike gets the like property value. The like property
// returns a *bool when successful
func (m *Comment) GetLike() *bool {
	return m.like
}

// GetLikesCount gets the likes_count property value. The likes_count property
// returns a *int32 when successful
func (m *Comment) GetLikesCount() *int32 {
	return m.likes_count
}

// GetParentId gets the parent_id property value. The parent_id property
// returns a *int32 when successful
func (m *Comment) GetParentId() *int32 {
	return m.parent_id
}

// GetPostId gets the post_id property value. The post_id property
// returns a *int32 when successful
func (m *Comment) GetPostId() *int32 {
	return m.post_id
}

// GetReplies gets the replies property value. The replies property
// returns a []Commentable when successful
func (m *Comment) GetReplies() []Commentable {
	return m.replies
}

// GetReplyCount gets the reply_count property value. The reply_count property
// returns a *int32 when successful
func (m *Comment) GetReplyCount() *int32 {
	return m.reply_count
}

// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *string when successful
func (m *Comment) GetUpdatedAt() *string {
	return m.updatedAt
}

// Serialize serializes information the current object
func (m *Comment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("author", m.GetAuthor())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("author_uid", m.GetAuthorUid())
		if err != nil {
			return err
		}
	}
	if m.GetCommentLikes() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommentLikes()))
		for i, v := range m.GetCommentLikes() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("comment_likes", cast)
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
		err := writer.WriteStringValue("description", m.GetDescription())
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
		err := writer.WriteBoolValue("like", m.GetLike())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("likes_count", m.GetLikesCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("parent_id", m.GetParentId())
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
	if m.GetReplies() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReplies()))
		for i, v := range m.GetReplies() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("replies", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("reply_count", m.GetReplyCount())
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
func (m *Comment) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAuthor sets the author property value. The author property
func (m *Comment) SetAuthor(value Userable) {
	m.author = value
}

// SetAuthorUid sets the author_uid property value. The author_uid property
func (m *Comment) SetAuthorUid(value *string) {
	m.author_uid = value
}

// SetCommentLikes sets the comment_likes property value. The comment_likes property
func (m *Comment) SetCommentLikes(value []CommentLikeable) {
	m.comment_likes = value
}

// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *Comment) SetCreatedAt(value *string) {
	m.createdAt = value
}

// SetDeletedAt sets the deletedAt property value. The deletedAt property
func (m *Comment) SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable) {
	m.deletedAt = value
}

// SetDescription sets the description property value. The description property
func (m *Comment) SetDescription(value *string) {
	m.description = value
}

// SetId sets the id property value. The id property
func (m *Comment) SetId(value *int32) {
	m.id = value
}

// SetLike sets the like property value. The like property
func (m *Comment) SetLike(value *bool) {
	m.like = value
}

// SetLikesCount sets the likes_count property value. The likes_count property
func (m *Comment) SetLikesCount(value *int32) {
	m.likes_count = value
}

// SetParentId sets the parent_id property value. The parent_id property
func (m *Comment) SetParentId(value *int32) {
	m.parent_id = value
}

// SetPostId sets the post_id property value. The post_id property
func (m *Comment) SetPostId(value *int32) {
	m.post_id = value
}

// SetReplies sets the replies property value. The replies property
func (m *Comment) SetReplies(value []Commentable) {
	m.replies = value
}

// SetReplyCount sets the reply_count property value. The reply_count property
func (m *Comment) SetReplyCount(value *int32) {
	m.reply_count = value
}

// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *Comment) SetUpdatedAt(value *string) {
	m.updatedAt = value
}

type Commentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAuthor() Userable
	GetAuthorUid() *string
	GetCommentLikes() []CommentLikeable
	GetCreatedAt() *string
	GetDeletedAt() i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
	GetDescription() *string
	GetId() *int32
	GetLike() *bool
	GetLikesCount() *int32
	GetParentId() *int32
	GetPostId() *int32
	GetReplies() []Commentable
	GetReplyCount() *int32
	GetUpdatedAt() *string
	SetAuthor(value Userable)
	SetAuthorUid(value *string)
	SetCommentLikes(value []CommentLikeable)
	SetCreatedAt(value *string)
	SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)
	SetDescription(value *string)
	SetId(value *int32)
	SetLike(value *bool)
	SetLikesCount(value *int32)
	SetParentId(value *int32)
	SetPostId(value *int32)
	SetReplies(value []Commentable)
	SetReplyCount(value *int32)
	SetUpdatedAt(value *string)
}
