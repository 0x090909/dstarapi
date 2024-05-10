package backend_api_apimodel

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaginatedComments struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The entries property
	entries []Commentable
	// The remaining property
	remaining *int32
	// The total property
	total *int32
}

// NewPaginatedComments instantiates a new PaginatedComments and sets the default values.
func NewPaginatedComments() *PaginatedComments {
	m := &PaginatedComments{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePaginatedCommentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaginatedCommentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPaginatedComments(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaginatedComments) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEntries gets the entries property value. The entries property
// returns a []Commentable when successful
func (m *PaginatedComments) GetEntries() []Commentable {
	return m.entries
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaginatedComments) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["entries"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetEntries(res)
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

// GetRemaining gets the remaining property value. The remaining property
// returns a *int32 when successful
func (m *PaginatedComments) GetRemaining() *int32 {
	return m.remaining
}

// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *PaginatedComments) GetTotal() *int32 {
	return m.total
}

// Serialize serializes information the current object
func (m *PaginatedComments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetEntries() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEntries()))
		for i, v := range m.GetEntries() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("entries", cast)
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
func (m *PaginatedComments) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEntries sets the entries property value. The entries property
func (m *PaginatedComments) SetEntries(value []Commentable) {
	m.entries = value
}

// SetRemaining sets the remaining property value. The remaining property
func (m *PaginatedComments) SetRemaining(value *int32) {
	m.remaining = value
}

// SetTotal sets the total property value. The total property
func (m *PaginatedComments) SetTotal(value *int32) {
	m.total = value
}

type PaginatedCommentsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEntries() []Commentable
	GetRemaining() *int32
	GetTotal() *int32
	SetEntries(value []Commentable)
	SetRemaining(value *int32)
	SetTotal(value *int32)
}
