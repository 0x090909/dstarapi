package backend_datamodel_model

import (
	i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1 "github.com/0x090909/dstarapi/models/gorm_io_gorm"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type User struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The bio property
	bio *string
	// The bronze_crowns property
	bronze_crowns *int32
	// The created_at property
	created_at *string
	// The date_of_birth property
	date_of_birth *string
	// The deleted_at property
	deleted_at i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
	// The device_token property
	device_token *string
	// The email property
	email *string
	// The favorite_posts property
	favorite_posts []Postable
	// The followers property
	followers []Userable
	// The gold_crowns property
	gold_crowns *int32
	// The image_profile_url property
	image_profile_url *string
	// The is_verified property
	is_verified *bool
	// The last_login property
	last_login *string
	// The last_name property
	last_name *string
	// The name property
	name *string
	// The phone_number property
	phone_number *string
	// The post_views property
	post_views []Postable
	// The preference property
	preference Preferenceable
	// The silver_crowns property
	silver_crowns *int32
	// The status property
	status *int32
	// The type property
	typeEscaped *int32
	// The uid property
	uid *string
	// The updated_at property
	updated_at *string
	// The username property
	username *string
	// The wallet_key property
	wallet_key *string
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

// GetBio gets the bio property value. The bio property
// returns a *string when successful
func (m *User) GetBio() *string {
	return m.bio
}

// GetBronzeCrowns gets the bronze_crowns property value. The bronze_crowns property
// returns a *int32 when successful
func (m *User) GetBronzeCrowns() *int32 {
	return m.bronze_crowns
}

// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *User) GetCreatedAt() *string {
	return m.created_at
}

// GetDateOfBirth gets the date_of_birth property value. The date_of_birth property
// returns a *string when successful
func (m *User) GetDateOfBirth() *string {
	return m.date_of_birth
}

// GetDeletedAt gets the deleted_at property value. The deleted_at property
// returns a DeletedAtable when successful
func (m *User) GetDeletedAt() i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable {
	return m.deleted_at
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
	res["bronze_crowns"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBronzeCrowns(val)
		}
		return nil
	}
	res["created_at"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedAt(val)
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
	res["deleted_at"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.CreateDeletedAtFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDeletedAt(val.(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable))
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
	res["gold_crowns"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGoldCrowns(val)
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
	res["last_login"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastLogin(val)
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
	res["post_views"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetPostViews(res)
		}
		return nil
	}
	res["preference"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreatePreferenceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPreference(val.(Preferenceable))
		}
		return nil
	}
	res["silver_crowns"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSilverCrowns(val)
		}
		return nil
	}
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
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
	res["updated_at"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUpdatedAt(val)
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
	res["wallet_key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWalletKey(val)
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

// GetGoldCrowns gets the gold_crowns property value. The gold_crowns property
// returns a *int32 when successful
func (m *User) GetGoldCrowns() *int32 {
	return m.gold_crowns
}

// GetImageProfileUrl gets the image_profile_url property value. The image_profile_url property
// returns a *string when successful
func (m *User) GetImageProfileUrl() *string {
	return m.image_profile_url
}

// GetIsVerified gets the is_verified property value. The is_verified property
// returns a *bool when successful
func (m *User) GetIsVerified() *bool {
	return m.is_verified
}

// GetLastLogin gets the last_login property value. The last_login property
// returns a *string when successful
func (m *User) GetLastLogin() *string {
	return m.last_login
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

// GetPostViews gets the post_views property value. The post_views property
// returns a []Postable when successful
func (m *User) GetPostViews() []Postable {
	return m.post_views
}

// GetPreference gets the preference property value. The preference property
// returns a Preferenceable when successful
func (m *User) GetPreference() Preferenceable {
	return m.preference
}

// GetSilverCrowns gets the silver_crowns property value. The silver_crowns property
// returns a *int32 when successful
func (m *User) GetSilverCrowns() *int32 {
	return m.silver_crowns
}

// GetStatus gets the status property value. The status property
// returns a *int32 when successful
func (m *User) GetStatus() *int32 {
	return m.status
}

// GetTypeEscaped gets the type property value. The type property
// returns a *int32 when successful
func (m *User) GetTypeEscaped() *int32 {
	return m.typeEscaped
}

// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *User) GetUid() *string {
	return m.uid
}

// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *User) GetUpdatedAt() *string {
	return m.updated_at
}

// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *User) GetUsername() *string {
	return m.username
}

// GetWalletKey gets the wallet_key property value. The wallet_key property
// returns a *string when successful
func (m *User) GetWalletKey() *string {
	return m.wallet_key
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *User) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("bio", m.GetBio())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("bronze_crowns", m.GetBronzeCrowns())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
		err := writer.WriteObjectValue("deleted_at", m.GetDeletedAt())
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
		err := writer.WriteInt32Value("gold_crowns", m.GetGoldCrowns())
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
		err := writer.WriteBoolValue("is_verified", m.GetIsVerified())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("last_login", m.GetLastLogin())
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
	if m.GetPostViews() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPostViews()))
		for i, v := range m.GetPostViews() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("post_views", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("preference", m.GetPreference())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("silver_crowns", m.GetSilverCrowns())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("status", m.GetStatus())
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
		err := writer.WriteStringValue("uid", m.GetUid())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
		err := writer.WriteStringValue("wallet_key", m.GetWalletKey())
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

// SetBio sets the bio property value. The bio property
func (m *User) SetBio(value *string) {
	m.bio = value
}

// SetBronzeCrowns sets the bronze_crowns property value. The bronze_crowns property
func (m *User) SetBronzeCrowns(value *int32) {
	m.bronze_crowns = value
}

// SetCreatedAt sets the created_at property value. The created_at property
func (m *User) SetCreatedAt(value *string) {
	m.created_at = value
}

// SetDateOfBirth sets the date_of_birth property value. The date_of_birth property
func (m *User) SetDateOfBirth(value *string) {
	m.date_of_birth = value
}

// SetDeletedAt sets the deleted_at property value. The deleted_at property
func (m *User) SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable) {
	m.deleted_at = value
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

// SetGoldCrowns sets the gold_crowns property value. The gold_crowns property
func (m *User) SetGoldCrowns(value *int32) {
	m.gold_crowns = value
}

// SetImageProfileUrl sets the image_profile_url property value. The image_profile_url property
func (m *User) SetImageProfileUrl(value *string) {
	m.image_profile_url = value
}

// SetIsVerified sets the is_verified property value. The is_verified property
func (m *User) SetIsVerified(value *bool) {
	m.is_verified = value
}

// SetLastLogin sets the last_login property value. The last_login property
func (m *User) SetLastLogin(value *string) {
	m.last_login = value
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

// SetPostViews sets the post_views property value. The post_views property
func (m *User) SetPostViews(value []Postable) {
	m.post_views = value
}

// SetPreference sets the preference property value. The preference property
func (m *User) SetPreference(value Preferenceable) {
	m.preference = value
}

// SetSilverCrowns sets the silver_crowns property value. The silver_crowns property
func (m *User) SetSilverCrowns(value *int32) {
	m.silver_crowns = value
}

// SetStatus sets the status property value. The status property
func (m *User) SetStatus(value *int32) {
	m.status = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *User) SetTypeEscaped(value *int32) {
	m.typeEscaped = value
}

// SetUid sets the uid property value. The uid property
func (m *User) SetUid(value *string) {
	m.uid = value
}

// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *User) SetUpdatedAt(value *string) {
	m.updated_at = value
}

// SetUsername sets the username property value. The username property
func (m *User) SetUsername(value *string) {
	m.username = value
}

// SetWalletKey sets the wallet_key property value. The wallet_key property
func (m *User) SetWalletKey(value *string) {
	m.wallet_key = value
}

// SetWebsite sets the website property value. The website property
func (m *User) SetWebsite(value *string) {
	m.website = value
}

type Userable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBio() *string
	GetBronzeCrowns() *int32
	GetCreatedAt() *string
	GetDateOfBirth() *string
	GetDeletedAt() i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
	GetDeviceToken() *string
	GetEmail() *string
	GetFavoritePosts() []Postable
	GetFollowers() []Userable
	GetGoldCrowns() *int32
	GetImageProfileUrl() *string
	GetIsVerified() *bool
	GetLastLogin() *string
	GetLastName() *string
	GetName() *string
	GetPhoneNumber() *string
	GetPostViews() []Postable
	GetPreference() Preferenceable
	GetSilverCrowns() *int32
	GetStatus() *int32
	GetTypeEscaped() *int32
	GetUid() *string
	GetUpdatedAt() *string
	GetUsername() *string
	GetWalletKey() *string
	GetWebsite() *string
	SetBio(value *string)
	SetBronzeCrowns(value *int32)
	SetCreatedAt(value *string)
	SetDateOfBirth(value *string)
	SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)
	SetDeviceToken(value *string)
	SetEmail(value *string)
	SetFavoritePosts(value []Postable)
	SetFollowers(value []Userable)
	SetGoldCrowns(value *int32)
	SetImageProfileUrl(value *string)
	SetIsVerified(value *bool)
	SetLastLogin(value *string)
	SetLastName(value *string)
	SetName(value *string)
	SetPhoneNumber(value *string)
	SetPostViews(value []Postable)
	SetPreference(value Preferenceable)
	SetSilverCrowns(value *int32)
	SetStatus(value *int32)
	SetTypeEscaped(value *int32)
	SetUid(value *string)
	SetUpdatedAt(value *string)
	SetUsername(value *string)
	SetWalletKey(value *string)
	SetWebsite(value *string)
}
