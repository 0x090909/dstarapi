package backend_api_apimodel

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaginatedPosts struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The posts property
	posts []Postable
	// The remaining property
	remaining *int32
	// The total property
	total *int32
}

// NewPaginatedPosts instantiates a new PaginatedPosts and sets the default values.
func NewPaginatedPosts() *PaginatedPosts {
	m := &PaginatedPosts{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePaginatedPostsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaginatedPostsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPaginatedPosts(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaginatedPosts) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaginatedPosts) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["posts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetPosts(res)
		}
		return nil
	}
	res["remaining"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRemaining(val)
		}
		return nil
	}
	res["total"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTotal(val)
		}
		return nil
	}
	return res
}

// GetPosts gets the posts property value. The posts property
// returns a []Postable when successful
func (m *PaginatedPosts) GetPosts() []Postable {
	return m.posts
}

// GetRemaining gets the remaining property value. The remaining property
// returns a *int32 when successful
func (m *PaginatedPosts) GetRemaining() *int32 {
	return m.remaining
}

// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *PaginatedPosts) GetTotal() *int32 {
	return m.total
}

// Serialize serializes information the current object
func (m *PaginatedPosts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetPosts() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPosts()))
		for i, v := range m.GetPosts() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("posts", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("remaining", m.GetRemaining())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *PaginatedPosts) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetPosts sets the posts property value. The posts property
func (m *PaginatedPosts) SetPosts(value []Postable) {
	m.posts = value
}

// SetRemaining sets the remaining property value. The remaining property
func (m *PaginatedPosts) SetRemaining(value *int32) {
	m.remaining = value
}

// SetTotal sets the total property value. The total property
func (m *PaginatedPosts) SetTotal(value *int32) {
	m.total = value
}

type PaginatedPostsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetPosts() []Postable
	GetRemaining() *int32
	GetTotal() *int32
	SetPosts(value []Postable)
	SetRemaining(value *int32)
	SetTotal(value *int32)
}
