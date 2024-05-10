package backend_datamodel_model

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Leaderboard struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The category property
	category *string
	// The last_name property
	last_name *string
	// The name property
	name *string
	// The post_id property
	post_id *int32
	// The rank property
	rank *int32
	// The score property
	score *int32
	// The uid property
	uid *string
	// The user_image_profile_url property
	user_image_profile_url *string
	// The username property
	username *string
	// The video_url property
	video_url *string
}

// NewLeaderboard instantiates a new Leaderboard and sets the default values.
func NewLeaderboard() *Leaderboard {
	m := &Leaderboard{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLeaderboardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLeaderboardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLeaderboard(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Leaderboard) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategory gets the category property value. The category property
// returns a *string when successful
func (m *Leaderboard) GetCategory() *string {
	return m.category
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Leaderboard) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["category"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCategory(val)
		}
		return nil
	}
	res["last_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastName(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
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
	res["rank"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRank(val)
		}
		return nil
	}
	res["score"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetScore(val)
		}
		return nil
	}
	res["uid"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUid(val)
		}
		return nil
	}
	res["user_image_profile_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserImageProfileUrl(val)
		}
		return nil
	}
	res["username"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUsername(val)
		}
		return nil
	}
	res["video_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetVideoUrl(val)
		}
		return nil
	}
	return res
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *Leaderboard) GetLastName() *string {
	return m.last_name
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Leaderboard) GetName() *string {
	return m.name
}

// GetPostId gets the post_id property value. The post_id property
// returns a *int32 when successful
func (m *Leaderboard) GetPostId() *int32 {
	return m.post_id
}

// GetRank gets the rank property value. The rank property
// returns a *int32 when successful
func (m *Leaderboard) GetRank() *int32 {
	return m.rank
}

// GetScore gets the score property value. The score property
// returns a *int32 when successful
func (m *Leaderboard) GetScore() *int32 {
	return m.score
}

// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *Leaderboard) GetUid() *string {
	return m.uid
}

// GetUserImageProfileUrl gets the user_image_profile_url property value. The user_image_profile_url property
// returns a *string when successful
func (m *Leaderboard) GetUserImageProfileUrl() *string {
	return m.user_image_profile_url
}

// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *Leaderboard) GetUsername() *string {
	return m.username
}

// GetVideoUrl gets the video_url property value. The video_url property
// returns a *string when successful
func (m *Leaderboard) GetVideoUrl() *string {
	return m.video_url
}

// Serialize serializes information the current object
func (m *Leaderboard) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("category", m.GetCategory())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("last_name", m.GetLastName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
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
		err := writer.WriteStringValue("uid", m.GetUid())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("username", m.GetUsername())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("user_image_profile_url", m.GetUserImageProfileUrl())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Leaderboard) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategory sets the category property value. The category property
func (m *Leaderboard) SetCategory(value *string) {
	m.category = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *Leaderboard) SetLastName(value *string) {
	m.last_name = value
}

// SetName sets the name property value. The name property
func (m *Leaderboard) SetName(value *string) {
	m.name = value
}

// SetPostId sets the post_id property value. The post_id property
func (m *Leaderboard) SetPostId(value *int32) {
	m.post_id = value
}

// SetRank sets the rank property value. The rank property
func (m *Leaderboard) SetRank(value *int32) {
	m.rank = value
}

// SetScore sets the score property value. The score property
func (m *Leaderboard) SetScore(value *int32) {
	m.score = value
}

// SetUid sets the uid property value. The uid property
func (m *Leaderboard) SetUid(value *string) {
	m.uid = value
}

// SetUserImageProfileUrl sets the user_image_profile_url property value. The user_image_profile_url property
func (m *Leaderboard) SetUserImageProfileUrl(value *string) {
	m.user_image_profile_url = value
}

// SetUsername sets the username property value. The username property
func (m *Leaderboard) SetUsername(value *string) {
	m.username = value
}

// SetVideoUrl sets the video_url property value. The video_url property
func (m *Leaderboard) SetVideoUrl(value *string) {
	m.video_url = value
}

type Leaderboardable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCategory() *string
	GetLastName() *string
	GetName() *string
	GetPostId() *int32
	GetRank() *int32
	GetScore() *int32
	GetUid() *string
	GetUserImageProfileUrl() *string
	GetUsername() *string
	GetVideoUrl() *string
	SetCategory(value *string)
	SetLastName(value *string)
	SetName(value *string)
	SetPostId(value *int32)
	SetRank(value *int32)
	SetScore(value *int32)
	SetUid(value *string)
	SetUserImageProfileUrl(value *string)
	SetUsername(value *string)
	SetVideoUrl(value *string)
}
