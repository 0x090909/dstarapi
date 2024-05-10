package backend_api_apimodel

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
)

type Post struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The allow_comments property
    allow_comments *bool
    // The author property
    author Userable
    // The category property
    category *int32
    // The comments_count property
    comments_count *int32
    // The description property
    description *string
    // The favourites_count property
    favourites_count *int32
    // The id property
    id *int32
    // The is_favourite property
    is_favourite *bool
    // The is_following_author property
    is_following_author *bool
    // The is_viewed property
    is_viewed *bool
    // The likes_count property
    likes_count *int32
    // The rank property
    rank *int32
    // The score property
    score *int32
    // The shared_count property
    shared_count *int32
    // The thumbnail property
    thumbnail *string
    // The video_length property
    video_length *float64
    // The video_url property
    video_url *string
    // The view_time property
    view_time *float64
    // The views_count property
    views_count *int32
    // placeholders for frontend
    vote_value *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ActionType
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
// GetAuthor gets the author property value. The author property
// returns a Userable when successful
func (m *Post) GetAuthor()(Userable) {
    return m.author
}
// GetCategory gets the category property value. The category property
// returns a *int32 when successful
func (m *Post) GetCategory()(*int32) {
    return m.category
}
// GetCommentsCount gets the comments_count property value. The comments_count property
// returns a *int32 when successful
func (m *Post) GetCommentsCount()(*int32) {
    return m.comments_count
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Post) GetDescription()(*string) {
    return m.description
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
    res["is_viewed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsViewed(val)
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
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
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
    res["shared_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedCount(val)
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
    res["vote_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ParseActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVoteValue(val.(*i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ActionType))
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
// GetIsFavourite gets the is_favourite property value. The is_favourite property
// returns a *bool when successful
func (m *Post) GetIsFavourite()(*bool) {
    return m.is_favourite
}
// GetIsFollowingAuthor gets the is_following_author property value. The is_following_author property
// returns a *bool when successful
func (m *Post) GetIsFollowingAuthor()(*bool) {
    return m.is_following_author
}
// GetIsViewed gets the is_viewed property value. The is_viewed property
// returns a *bool when successful
func (m *Post) GetIsViewed()(*bool) {
    return m.is_viewed
}
// GetLikesCount gets the likes_count property value. The likes_count property
// returns a *int32 when successful
func (m *Post) GetLikesCount()(*int32) {
    return m.likes_count
}
// GetRank gets the rank property value. The rank property
// returns a *int32 when successful
func (m *Post) GetRank()(*int32) {
    return m.rank
}
// GetScore gets the score property value. The score property
// returns a *int32 when successful
func (m *Post) GetScore()(*int32) {
    return m.score
}
// GetSharedCount gets the shared_count property value. The shared_count property
// returns a *int32 when successful
func (m *Post) GetSharedCount()(*int32) {
    return m.shared_count
}
// GetThumbnail gets the thumbnail property value. The thumbnail property
// returns a *string when successful
func (m *Post) GetThumbnail()(*string) {
    return m.thumbnail
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
// GetVoteValue gets the vote_value property value. placeholders for frontend
// returns a *ActionType when successful
func (m *Post) GetVoteValue()(*i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ActionType) {
    return m.vote_value
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
        err := writer.WriteObjectValue("author", m.GetAuthor())
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
    {
        err := writer.WriteInt32Value("comments_count", m.GetCommentsCount())
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
    {
        err := writer.WriteBoolValue("is_viewed", m.GetIsViewed())
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
        err := writer.WriteInt32Value("rank", m.GetRank())
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
        err := writer.WriteInt32Value("shared_count", m.GetSharedCount())
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
    if m.GetVoteValue() != nil {
        cast := (*m.GetVoteValue()).String()
        err := writer.WriteStringValue("vote_value", &cast)
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
// SetAuthor sets the author property value. The author property
func (m *Post) SetAuthor(value Userable)() {
    m.author = value
}
// SetCategory sets the category property value. The category property
func (m *Post) SetCategory(value *int32)() {
    m.category = value
}
// SetCommentsCount sets the comments_count property value. The comments_count property
func (m *Post) SetCommentsCount(value *int32)() {
    m.comments_count = value
}
// SetDescription sets the description property value. The description property
func (m *Post) SetDescription(value *string)() {
    m.description = value
}
// SetFavouritesCount sets the favourites_count property value. The favourites_count property
func (m *Post) SetFavouritesCount(value *int32)() {
    m.favourites_count = value
}
// SetId sets the id property value. The id property
func (m *Post) SetId(value *int32)() {
    m.id = value
}
// SetIsFavourite sets the is_favourite property value. The is_favourite property
func (m *Post) SetIsFavourite(value *bool)() {
    m.is_favourite = value
}
// SetIsFollowingAuthor sets the is_following_author property value. The is_following_author property
func (m *Post) SetIsFollowingAuthor(value *bool)() {
    m.is_following_author = value
}
// SetIsViewed sets the is_viewed property value. The is_viewed property
func (m *Post) SetIsViewed(value *bool)() {
    m.is_viewed = value
}
// SetLikesCount sets the likes_count property value. The likes_count property
func (m *Post) SetLikesCount(value *int32)() {
    m.likes_count = value
}
// SetRank sets the rank property value. The rank property
func (m *Post) SetRank(value *int32)() {
    m.rank = value
}
// SetScore sets the score property value. The score property
func (m *Post) SetScore(value *int32)() {
    m.score = value
}
// SetSharedCount sets the shared_count property value. The shared_count property
func (m *Post) SetSharedCount(value *int32)() {
    m.shared_count = value
}
// SetThumbnail sets the thumbnail property value. The thumbnail property
func (m *Post) SetThumbnail(value *string)() {
    m.thumbnail = value
}
// SetVideoLength sets the video_length property value. The video_length property
func (m *Post) SetVideoLength(value *float64)() {
    m.video_length = value
}
// SetVideoUrl sets the video_url property value. The video_url property
func (m *Post) SetVideoUrl(value *string)() {
    m.video_url = value
}
// SetViewsCount sets the views_count property value. The views_count property
func (m *Post) SetViewsCount(value *int32)() {
    m.views_count = value
}
// SetViewTime sets the view_time property value. The view_time property
func (m *Post) SetViewTime(value *float64)() {
    m.view_time = value
}
// SetVoteValue sets the vote_value property value. placeholders for frontend
func (m *Post) SetVoteValue(value *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ActionType)() {
    m.vote_value = value
}
type Postable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowComments()(*bool)
    GetAuthor()(Userable)
    GetCategory()(*int32)
    GetCommentsCount()(*int32)
    GetDescription()(*string)
    GetFavouritesCount()(*int32)
    GetId()(*int32)
    GetIsFavourite()(*bool)
    GetIsFollowingAuthor()(*bool)
    GetIsViewed()(*bool)
    GetLikesCount()(*int32)
    GetRank()(*int32)
    GetScore()(*int32)
    GetSharedCount()(*int32)
    GetThumbnail()(*string)
    GetVideoLength()(*float64)
    GetVideoUrl()(*string)
    GetViewsCount()(*int32)
    GetViewTime()(*float64)
    GetVoteValue()(*i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ActionType)
    SetAllowComments(value *bool)()
    SetAuthor(value Userable)()
    SetCategory(value *int32)()
    SetCommentsCount(value *int32)()
    SetDescription(value *string)()
    SetFavouritesCount(value *int32)()
    SetId(value *int32)()
    SetIsFavourite(value *bool)()
    SetIsFollowingAuthor(value *bool)()
    SetIsViewed(value *bool)()
    SetLikesCount(value *int32)()
    SetRank(value *int32)()
    SetScore(value *int32)()
    SetSharedCount(value *int32)()
    SetThumbnail(value *string)()
    SetVideoLength(value *float64)()
    SetVideoUrl(value *string)()
    SetViewsCount(value *int32)()
    SetViewTime(value *float64)()
    SetVoteValue(value *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ActionType)()
}
