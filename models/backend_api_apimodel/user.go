package backend_api_apimodel

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type User struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The awards property
	awards Awardsable
	// The bio property
	bio *string
	// The date_of_birth property
	date_of_birth *string
	// The device_token property
	device_token *string
	// The email property
	email *string
	// The favorite_posts property
	favorite_posts []Postable
	// The followers property
	followers []Userable
	// The followers_count property
	followers_count *int32
	// The following property
	following []Userable
	// The following_count property
	following_count *int32
	// The image_profile_url property
	image_profile_url *string
	// Only used in the context of a user profile page to indicate if the visitor is following the user
	is_following *bool
	// The is_verified property
	is_verified *bool
	// The last_name property
	last_name *string
	// The name property
	name *string
	// The phone_number property
	phone_number *string
	// The score property
	score *int32
	// The uid property
	uid *string
	// The username property
	username *string
	// The website property
	website *string
}

// NewUser instantiates a new User and sets the default values.
func NewUser() *User {
	m := &User{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUser(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *User) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAwards gets the awards property value. The awards property
// returns a Awardsable when successful
func (m *User) GetAwards() Awardsable {
	return m.awards
}

// GetBio gets the bio property value. The bio property
// returns a *string when successful
func (m *User) GetBio() *string {
	return m.bio
}

// GetDateOfBirth gets the date_of_birth property value. The date_of_birth property
// returns a *string when successful
func (m *User) GetDateOfBirth() *string {
	return m.date_of_birth
}

// GetDeviceToken gets the device_token property value. The device_token property
// returns a *string when successful
func (m *User) GetDeviceToken() *string {
	return m.device_token
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *User) GetEmail() *string {
	return m.email
}

// GetFavoritePosts gets the favorite_posts property value. The favorite_posts property
// returns a []Postable when successful
func (m *User) GetFavoritePosts() []Postable {
	return m.favorite_posts
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *User) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["awards"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAwardsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAwards(val.(Awardsable))
		}
		return nil
	}
	res["bio"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBio(val)
		}
		return nil
	}
	res["date_of_birth"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateOfBirth(val)
		}
		return nil
	}
	res["device_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDeviceToken(val)
		}
		return nil
	}
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val)
		}
		return nil
	}
	res["favorite_posts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreatePostFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]Postable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(Postable)
				}
			}
			m.SetFavoritePosts(res)
		}
		return nil
	}
	res["followers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetFollowers(res)
		}
		return nil
	}
	res["followers_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFollowersCount(val)
		}
		return nil
	}
	res["following"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetFollowing(res)
		}
		return nil
	}
	res["following_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFollowingCount(val)
		}
		return nil
	}
	res["image_profile_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetImageProfileUrl(val)
		}
		return nil
	}
	res["is_following"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsFollowing(val)
		}
		return nil
	}
	res["is_verified"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsVerified(val)
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
	res["phone_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneNumber(val)
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
	res["website"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWebsite(val)
		}
		return nil
	}
	return res
}

// GetFollowers gets the followers property value. The followers property
// returns a []Userable when successful
func (m *User) GetFollowers() []Userable {
	return m.followers
}

// GetFollowersCount gets the followers_count property value. The followers_count property
// returns a *int32 when successful
func (m *User) GetFollowersCount() *int32 {
	return m.followers_count
}

// GetFollowing gets the following property value. The following property
// returns a []Userable when successful
func (m *User) GetFollowing() []Userable {
	return m.following
}

// GetFollowingCount gets the following_count property value. The following_count property
// returns a *int32 when successful
func (m *User) GetFollowingCount() *int32 {
	return m.following_count
}

// GetImageProfileUrl gets the image_profile_url property value. The image_profile_url property
// returns a *string when successful
func (m *User) GetImageProfileUrl() *string {
	return m.image_profile_url
}

// GetIsFollowing gets the is_following property value. Only used in the context of a user profile page to indicate if the visitor is following the user
// returns a *bool when successful
func (m *User) GetIsFollowing() *bool {
	return m.is_following
}

// GetIsVerified gets the is_verified property value. The is_verified property
// returns a *bool when successful
func (m *User) GetIsVerified() *bool {
	return m.is_verified
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *User) GetLastName() *string {
	return m.last_name
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *User) GetName() *string {
	return m.name
}

// GetPhoneNumber gets the phone_number property value. The phone_number property
// returns a *string when successful
func (m *User) GetPhoneNumber() *string {
	return m.phone_number
}

// GetScore gets the score property value. The score property
// returns a *int32 when successful
func (m *User) GetScore() *int32 {
	return m.score
}

// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *User) GetUid() *string {
	return m.uid
}

// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *User) GetUsername() *string {
	return m.username
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *User) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("awards", m.GetAwards())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("bio", m.GetBio())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("date_of_birth", m.GetDateOfBirth())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("device_token", m.GetDeviceToken())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	if m.GetFavoritePosts() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFavoritePosts()))
		for i, v := range m.GetFavoritePosts() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("favorite_posts", cast)
		if err != nil {
			return err
		}
	}
	if m.GetFollowers() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFollowers()))
		for i, v := range m.GetFollowers() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("followers", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("followers_count", m.GetFollowersCount())
		if err != nil {
			return err
		}
	}
	if m.GetFollowing() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFollowing()))
		for i, v := range m.GetFollowing() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("following", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("following_count", m.GetFollowingCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("image_profile_url", m.GetImageProfileUrl())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_following", m.GetIsFollowing())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_verified", m.GetIsVerified())
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
		err := writer.WriteStringValue("phone_number", m.GetPhoneNumber())
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
		err := writer.WriteStringValue("website", m.GetWebsite())
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
func (m *User) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAwards sets the awards property value. The awards property
func (m *User) SetAwards(value Awardsable) {
	m.awards = value
}

// SetBio sets the bio property value. The bio property
func (m *User) SetBio(value *string) {
	m.bio = value
}

// SetDateOfBirth sets the date_of_birth property value. The date_of_birth property
func (m *User) SetDateOfBirth(value *string) {
	m.date_of_birth = value
}

// SetDeviceToken sets the device_token property value. The device_token property
func (m *User) SetDeviceToken(value *string) {
	m.device_token = value
}

// SetEmail sets the email property value. The email property
func (m *User) SetEmail(value *string) {
	m.email = value
}

// SetFavoritePosts sets the favorite_posts property value. The favorite_posts property
func (m *User) SetFavoritePosts(value []Postable) {
	m.favorite_posts = value
}

// SetFollowers sets the followers property value. The followers property
func (m *User) SetFollowers(value []Userable) {
	m.followers = value
}

// SetFollowersCount sets the followers_count property value. The followers_count property
func (m *User) SetFollowersCount(value *int32) {
	m.followers_count = value
}

// SetFollowing sets the following property value. The following property
func (m *User) SetFollowing(value []Userable) {
	m.following = value
}

// SetFollowingCount sets the following_count property value. The following_count property
func (m *User) SetFollowingCount(value *int32) {
	m.following_count = value
}

// SetImageProfileUrl sets the image_profile_url property value. The image_profile_url property
func (m *User) SetImageProfileUrl(value *string) {
	m.image_profile_url = value
}

// SetIsFollowing sets the is_following property value. Only used in the context of a user profile page to indicate if the visitor is following the user
func (m *User) SetIsFollowing(value *bool) {
	m.is_following = value
}

// SetIsVerified sets the is_verified property value. The is_verified property
func (m *User) SetIsVerified(value *bool) {
	m.is_verified = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *User) SetLastName(value *string) {
	m.last_name = value
}

// SetName sets the name property value. The name property
func (m *User) SetName(value *string) {
	m.name = value
}

// SetPhoneNumber sets the phone_number property value. The phone_number property
func (m *User) SetPhoneNumber(value *string) {
	m.phone_number = value
}

// SetScore sets the score property value. The score property
func (m *User) SetScore(value *int32) {
	m.score = value
}

// SetUid sets the uid property value. The uid property
func (m *User) SetUid(value *string) {
	m.uid = value
}

// SetUsername sets the username property value. The username property
func (m *User) SetUsername(value *string) {
	m.username = value
}

// SetWebsite sets the website property value. The website property
func (m *User) SetWebsite(value *string) {
	m.website = value
}

type Userable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAwards() Awardsable
	GetBio() *string
	GetDateOfBirth() *string
	GetDeviceToken() *string
	GetEmail() *string
	GetFavoritePosts() []Postable
	GetFollowers() []Userable
	GetFollowersCount() *int32
	GetFollowing() []Userable
	GetFollowingCount() *int32
	GetImageProfileUrl() *string
	GetIsFollowing() *bool
	GetIsVerified() *bool
	GetLastName() *string
	GetName() *string
	GetPhoneNumber() *string
	GetScore() *int32
	GetUid() *string
	GetUsername() *string
	GetWebsite() *string
	SetAwards(value Awardsable)
	SetBio(value *string)
	SetDateOfBirth(value *string)
	SetDeviceToken(value *string)
	SetEmail(value *string)
	SetFavoritePosts(value []Postable)
	SetFollowers(value []Userable)
	SetFollowersCount(value *int32)
	SetFollowing(value []Userable)
	SetFollowingCount(value *int32)
	SetImageProfileUrl(value *string)
	SetIsFollowing(value *bool)
	SetIsVerified(value *bool)
	SetLastName(value *string)
	SetName(value *string)
	SetPhoneNumber(value *string)
	SetScore(value *int32)
	SetUid(value *string)
	SetUsername(value *string)
	SetWebsite(value *string)
}
