package backend_datamodel_model

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1 "github.com/0x090909/dstarapi/models/gorm_io_gorm"
)

type Post struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_comments property
    allow_comments *bool
    // The api_video_id property
    api_video_id *string
    // The author property
    author Userable
    // The author_uid property
    author_uid *string
    // The category property
    category *int32
    // The comments property
    comments []Commentable
    // The comments_count property
    comments_count *int32
    // The createdAt property
    createdAt *string
    // The current_rank property
    current_rank *int32
    // The deletedAt property
    deletedAt i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
    // The description property
    description *string
    // The dislikes property
    dislikes []Userable
    // The favourites_count property
    favourites_count *int32
    // The id property
    id *int32
    // Read only fields
    is_favourite *bool
    // The is_following_author property
    is_following_author *bool
    // The likes property
    likes []Userable
    // The likes_count property
    likes_count *int32
    // The score property
    score *int32
    // The shares property
    shares []Userable
    // The super_likes property
    super_likes []Userable
    // The thumbnail property
    thumbnail *string
    // The updatedAt property
    updatedAt *string
    // The video_length property
    video_length *float64
    // The video_url property
    video_url *string
    // The view_time property
    view_time *float64
    // The views property
    views []Userable
    // The views_count property
    views_count *int32
    // The vtl property
    vtl *float64
}
// NewPost instantiates a new Post and sets the default values.
func NewPost()(*Post) {
    m := &Post{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePostFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePostFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPost(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Post) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowComments gets the allow_comments property value. The allow_comments property
// returns a *bool when successful
func (m *Post) GetAllowComments()(*bool) {
    return m.allow_comments
}
// GetApiVideoId gets the api_video_id property value. The api_video_id property
// returns a *string when successful
func (m *Post) GetApiVideoId()(*string) {
    return m.api_video_id
}
// GetAuthor gets the author property value. The author property
// returns a Userable when successful
func (m *Post) GetAuthor()(Userable) {
    return m.author
}
// GetAuthorUid gets the author_uid property value. The author_uid property
// returns a *string when successful
func (m *Post) GetAuthorUid()(*string) {
    return m.author_uid
}
// GetCategory gets the category property value. The category property
// returns a *int32 when successful
func (m *Post) GetCategory()(*int32) {
    return m.category
}
// GetComments gets the comments property value. The comments property
// returns a []Commentable when successful
func (m *Post) GetComments()([]Commentable) {
    return m.comments
}
// GetCommentsCount gets the comments_count property value. The comments_count property
// returns a *int32 when successful
func (m *Post) GetCommentsCount()(*int32) {
    return m.comments_count
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *string when successful
func (m *Post) GetCreatedAt()(*string) {
    return m.createdAt
}
// GetCurrentRank gets the current_rank property value. The current_rank property
// returns a *int32 when successful
func (m *Post) GetCurrentRank()(*int32) {
    return m.current_rank
}
// GetDeletedAt gets the deletedAt property value. The deletedAt property
// returns a DeletedAtable when successful
func (m *Post) GetDeletedAt()(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable) {
    return m.deletedAt
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Post) GetDescription()(*string) {
    return m.description
}
// GetDislikes gets the dislikes property value. The dislikes property
// returns a []Userable when successful
func (m *Post) GetDislikes()([]Userable) {
    return m.dislikes
}
// GetFavouritesCount gets the favourites_count property value. The favourites_count property
// returns a *int32 when successful
func (m *Post) GetFavouritesCount()(*int32) {
    return m.favourites_count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Post) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allow_comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowComments(val)
        }
        return nil
    }
    res["api_video_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiVideoId(val)
        }
        return nil
    }
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(Userable))
        }
        return nil
    }
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
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetComments(res)
        }
        return nil
    }
    res["comments_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentsCount(val)
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
    res["current_rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentRank(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["dislikes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetDislikes(res)
        }
        return nil
    }
    res["favourites_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFavouritesCount(val)
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
    res["is_favourite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavourite(val)
        }
        return nil
    }
    res["is_following_author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFollowingAuthor(val)
        }
        return nil
    }
    res["likes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetLikes(res)
        }
        return nil
    }
    res["likes_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLikesCount(val)
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
    res["shares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetShares(res)
        }
        return nil
    }
    res["super_likes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetSuperLikes(res)
        }
        return nil
    }
    res["thumbnail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnail(val)
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
    res["video_length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoLength(val)
        }
        return nil
    }
    res["video_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoUrl(val)
        }
        return nil
    }
    res["view_time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewTime(val)
        }
        return nil
    }
    res["views"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetViews(res)
        }
        return nil
    }
    res["views_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewsCount(val)
        }
        return nil
    }
    res["vtl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVtl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Post) GetId()(*int32) {
    return m.id
}
// GetIsFavourite gets the is_favourite property value. Read only fields
// returns a *bool when successful
func (m *Post) GetIsFavourite()(*bool) {
    return m.is_favourite
}
// GetIsFollowingAuthor gets the is_following_author property value. The is_following_author property
// returns a *bool when successful
func (m *Post) GetIsFollowingAuthor()(*bool) {
    return m.is_following_author
}
// GetLikes gets the likes property value. The likes property
// returns a []Userable when successful
func (m *Post) GetLikes()([]Userable) {
    return m.likes
}
// GetLikesCount gets the likes_count property value. The likes_count property
// returns a *int32 when successful
func (m *Post) GetLikesCount()(*int32) {
    return m.likes_count
}
// GetScore gets the score property value. The score property
// returns a *int32 when successful
func (m *Post) GetScore()(*int32) {
    return m.score
}
// GetShares gets the shares property value. The shares property
// returns a []Userable when successful
func (m *Post) GetShares()([]Userable) {
    return m.shares
}
// GetSuperLikes gets the super_likes property value. The super_likes property
// returns a []Userable when successful
func (m *Post) GetSuperLikes()([]Userable) {
    return m.super_likes
}
// GetThumbnail gets the thumbnail property value. The thumbnail property
// returns a *string when successful
func (m *Post) GetThumbnail()(*string) {
    return m.thumbnail
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *string when successful
func (m *Post) GetUpdatedAt()(*string) {
    return m.updatedAt
}
// GetVideoLength gets the video_length property value. The video_length property
// returns a *float64 when successful
func (m *Post) GetVideoLength()(*float64) {
    return m.video_length
}
// GetVideoUrl gets the video_url property value. The video_url property
// returns a *string when successful
func (m *Post) GetVideoUrl()(*string) {
    return m.video_url
}
// GetViews gets the views property value. The views property
// returns a []Userable when successful
func (m *Post) GetViews()([]Userable) {
    return m.views
}
// GetViewsCount gets the views_count property value. The views_count property
// returns a *int32 when successful
func (m *Post) GetViewsCount()(*int32) {
    return m.views_count
}
// GetViewTime gets the view_time property value. The view_time property
// returns a *float64 when successful
func (m *Post) GetViewTime()(*float64) {
    return m.view_time
}
// GetVtl gets the vtl property value. The vtl property
// returns a *float64 when successful
func (m *Post) GetVtl()(*float64) {
    return m.vtl
}
// Serialize serializes information the current object
func (m *Post) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allow_comments", m.GetAllowComments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("api_video_id", m.GetApiVideoId())
        if err != nil {
            return err
        }
    }
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
    {
        err := writer.WriteInt32Value("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    if m.GetComments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("comments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("comments_count", m.GetCommentsCount())
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
        err := writer.WriteInt32Value("current_rank", m.GetCurrentRank())
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
    if m.GetDislikes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDislikes()))
        for i, v := range m.GetDislikes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dislikes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("favourites_count", m.GetFavouritesCount())
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
        err := writer.WriteBoolValue("is_favourite", m.GetIsFavourite())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_following_author", m.GetIsFollowingAuthor())
        if err != nil {
            return err
        }
    }
    if m.GetLikes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLikes()))
        for i, v := range m.GetLikes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("likes", cast)
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
        err := writer.WriteInt32Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    if m.GetShares() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShares()))
        for i, v := range m.GetShares() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("shares", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSuperLikes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSuperLikes()))
        for i, v := range m.GetSuperLikes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("super_likes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbnail", m.GetThumbnail())
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
        err := writer.WriteFloat64Value("video_length", m.GetVideoLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("video_url", m.GetVideoUrl())
        if err != nil {
            return err
        }
    }
    if m.GetViews() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetViews()))
        for i, v := range m.GetViews() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("views", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("views_count", m.GetViewsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("view_time", m.GetViewTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("vtl", m.GetVtl())
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
func (m *Post) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowComments sets the allow_comments property value. The allow_comments property
func (m *Post) SetAllowComments(value *bool)() {
    m.allow_comments = value
}
// SetApiVideoId sets the api_video_id property value. The api_video_id property
func (m *Post) SetApiVideoId(value *string)() {
    m.api_video_id = value
}
// SetAuthor sets the author property value. The author property
func (m *Post) SetAuthor(value Userable)() {
    m.author = value
}
// SetAuthorUid sets the author_uid property value. The author_uid property
func (m *Post) SetAuthorUid(value *string)() {
    m.author_uid = value
}
// SetCategory sets the category property value. The category property
func (m *Post) SetCategory(value *int32)() {
    m.category = value
}
// SetComments sets the comments property value. The comments property
func (m *Post) SetComments(value []Commentable)() {
    m.comments = value
}
// SetCommentsCount sets the comments_count property value. The comments_count property
func (m *Post) SetCommentsCount(value *int32)() {
    m.comments_count = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *Post) SetCreatedAt(value *string)() {
    m.createdAt = value
}
// SetCurrentRank sets the current_rank property value. The current_rank property
func (m *Post) SetCurrentRank(value *int32)() {
    m.current_rank = value
}
// SetDeletedAt sets the deletedAt property value. The deletedAt property
func (m *Post) SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)() {
    m.deletedAt = value
}
// SetDescription sets the description property value. The description property
func (m *Post) SetDescription(value *string)() {
    m.description = value
}
// SetDislikes sets the dislikes property value. The dislikes property
func (m *Post) SetDislikes(value []Userable)() {
    m.dislikes = value
}
// SetFavouritesCount sets the favourites_count property value. The favourites_count property
func (m *Post) SetFavouritesCount(value *int32)() {
    m.favourites_count = value
}
// SetId sets the id property value. The id property
func (m *Post) SetId(value *int32)() {
    m.id = value
}
// SetIsFavourite sets the is_favourite property value. Read only fields
func (m *Post) SetIsFavourite(value *bool)() {
    m.is_favourite = value
}
// SetIsFollowingAuthor sets the is_following_author property value. The is_following_author property
func (m *Post) SetIsFollowingAuthor(value *bool)() {
    m.is_following_author = value
}
// SetLikes sets the likes property value. The likes property
func (m *Post) SetLikes(value []Userable)() {
    m.likes = value
}
// SetLikesCount sets the likes_count property value. The likes_count property
func (m *Post) SetLikesCount(value *int32)() {
    m.likes_count = value
}
// SetScore sets the score property value. The score property
func (m *Post) SetScore(value *int32)() {
    m.score = value
}
// SetShares sets the shares property value. The shares property
func (m *Post) SetShares(value []Userable)() {
    m.shares = value
}
// SetSuperLikes sets the super_likes property value. The super_likes property
func (m *Post) SetSuperLikes(value []Userable)() {
    m.super_likes = value
}
// SetThumbnail sets the thumbnail property value. The thumbnail property
func (m *Post) SetThumbnail(value *string)() {
    m.thumbnail = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *Post) SetUpdatedAt(value *string)() {
    m.updatedAt = value
}
// SetVideoLength sets the video_length property value. The video_length property
func (m *Post) SetVideoLength(value *float64)() {
    m.video_length = value
}
// SetVideoUrl sets the video_url property value. The video_url property
func (m *Post) SetVideoUrl(value *string)() {
    m.video_url = value
}
// SetViews sets the views property value. The views property
func (m *Post) SetViews(value []Userable)() {
    m.views = value
}
// SetViewsCount sets the views_count property value. The views_count property
func (m *Post) SetViewsCount(value *int32)() {
    m.views_count = value
}
// SetViewTime sets the view_time property value. The view_time property
func (m *Post) SetViewTime(value *float64)() {
    m.view_time = value
}
// SetVtl sets the vtl property value. The vtl property
func (m *Post) SetVtl(value *float64)() {
    m.vtl = value
}
type Postable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowComments()(*bool)
    GetApiVideoId()(*string)
    GetAuthor()(Userable)
    GetAuthorUid()(*string)
    GetCategory()(*int32)
    GetComments()([]Commentable)
    GetCommentsCount()(*int32)
    GetCreatedAt()(*string)
    GetCurrentRank()(*int32)
    GetDeletedAt()(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)
    GetDescription()(*string)
    GetDislikes()([]Userable)
    GetFavouritesCount()(*int32)
    GetId()(*int32)
    GetIsFavourite()(*bool)
    GetIsFollowingAuthor()(*bool)
    GetLikes()([]Userable)
    GetLikesCount()(*int32)
    GetScore()(*int32)
    GetShares()([]Userable)
    GetSuperLikes()([]Userable)
    GetThumbnail()(*string)
    GetUpdatedAt()(*string)
    GetVideoLength()(*float64)
    GetVideoUrl()(*string)
    GetViews()([]Userable)
    GetViewsCount()(*int32)
    GetViewTime()(*float64)
    GetVtl()(*float64)
    SetAllowComments(value *bool)()
    SetApiVideoId(value *string)()
    SetAuthor(value Userable)()
    SetAuthorUid(value *string)()
    SetCategory(value *int32)()
    SetComments(value []Commentable)()
    SetCommentsCount(value *int32)()
    SetCreatedAt(value *string)()
    SetCurrentRank(value *int32)()
    SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)()
    SetDescription(value *string)()
    SetDislikes(value []Userable)()
    SetFavouritesCount(value *int32)()
    SetId(value *int32)()
    SetIsFavourite(value *bool)()
    SetIsFollowingAuthor(value *bool)()
    SetLikes(value []Userable)()
    SetLikesCount(value *int32)()
    SetScore(value *int32)()
    SetShares(value []Userable)()
    SetSuperLikes(value []Userable)()
    SetThumbnail(value *string)()
    SetUpdatedAt(value *string)()
    SetVideoLength(value *float64)()
    SetVideoUrl(value *string)()
    SetViews(value []Userable)()
    SetViewsCount(value *int32)()
    SetViewTime(value *float64)()
    SetVtl(value *float64)()
}
