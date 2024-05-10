package apimodel

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaginatedEvents struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The events property
	events []Eventable
	// The remaining property
	remaining *int32
	// The total property
	total *int32
}

// NewPaginatedEvents instantiates a new PaginatedEvents and sets the default values.
func NewPaginatedEvents() *PaginatedEvents {
	m := &PaginatedEvents{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreatePaginatedEventsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaginatedEventsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPaginatedEvents(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaginatedEvents) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEvents gets the events property value. The events property
// returns a []Eventable when successful
func (m *PaginatedEvents) GetEvents() []Eventable {
	return m.events
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaginatedEvents) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["events"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]Eventable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(Eventable)
				}
			}
			m.SetEvents(res)
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
func (m *PaginatedEvents) GetRemaining() *int32 {
	return m.remaining
}

// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *PaginatedEvents) GetTotal() *int32 {
	return m.total
}

// Serialize serializes information the current object
func (m *PaginatedEvents) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetEvents() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
		for i, v := range m.GetEvents() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("events", cast)
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
func (m *PaginatedEvents) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEvents sets the events property value. The events property
func (m *PaginatedEvents) SetEvents(value []Eventable) {
	m.events = value
}

// SetRemaining sets the remaining property value. The remaining property
func (m *PaginatedEvents) SetRemaining(value *int32) {
	m.remaining = value
}

// SetTotal sets the total property value. The total property
func (m *PaginatedEvents) SetTotal(value *int32) {
	m.total = value
}

type PaginatedEventsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEvents() []Eventable
	GetRemaining() *int32
	GetTotal() *int32
	SetEvents(value []Eventable)
	SetRemaining(value *int32)
	SetTotal(value *int32)
}
