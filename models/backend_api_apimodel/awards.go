package backend_api_apimodel

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Awards struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The bronze_crowns property
	bronze_crowns *int32
	// The gold_crowns property
	gold_crowns *int32
	// The silver_crowns property
	silver_crowns *int32
}

// NewAwards instantiates a new Awards and sets the default values.
func NewAwards() *Awards {
	m := &Awards{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAwardsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwardsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAwards(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Awards) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBronzeCrowns gets the bronze_crowns property value. The bronze_crowns property
// returns a *int32 when successful
func (m *Awards) GetBronzeCrowns() *int32 {
	return m.bronze_crowns
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Awards) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	return res
}

// GetGoldCrowns gets the gold_crowns property value. The gold_crowns property
// returns a *int32 when successful
func (m *Awards) GetGoldCrowns() *int32 {
	return m.gold_crowns
}

// GetSilverCrowns gets the silver_crowns property value. The silver_crowns property
// returns a *int32 when successful
func (m *Awards) GetSilverCrowns() *int32 {
	return m.silver_crowns
}

// Serialize serializes information the current object
func (m *Awards) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("bronze_crowns", m.GetBronzeCrowns())
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
		err := writer.WriteInt32Value("silver_crowns", m.GetSilverCrowns())
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
func (m *Awards) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBronzeCrowns sets the bronze_crowns property value. The bronze_crowns property
func (m *Awards) SetBronzeCrowns(value *int32) {
	m.bronze_crowns = value
}

// SetGoldCrowns sets the gold_crowns property value. The gold_crowns property
func (m *Awards) SetGoldCrowns(value *int32) {
	m.gold_crowns = value
}

// SetSilverCrowns sets the silver_crowns property value. The silver_crowns property
func (m *Awards) SetSilverCrowns(value *int32) {
	m.silver_crowns = value
}

type Awardsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBronzeCrowns() *int32
	GetGoldCrowns() *int32
	GetSilverCrowns() *int32
	SetBronzeCrowns(value *int32)
	SetGoldCrowns(value *int32)
	SetSilverCrowns(value *int32)
}
