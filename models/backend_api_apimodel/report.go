package backend_api_apimodel

import (
	i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08 "github.com/0x090909/dstarapi/models/backend_datamodel_model"
	i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7 "github.com/0x090909/dstarapi/models/model"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Report struct {
	// The action property
	action *i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ReportActionType
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The comment_id property
	comment_id *int32
	// The description property
	description *string
	// The post_id property
	post_id *int32
	// The reason property
	reason *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ReportReason
	// The reported_uid property
	reported_uid *string
}

// NewReport instantiates a new Report and sets the default values.
func NewReport() *Report {
	m := &Report{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewReport(), nil
}

// GetAction gets the action property value. The action property
// returns a *ReportActionType when successful
func (m *Report) GetAction() *i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ReportActionType {
	return m.action
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Report) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCommentId gets the comment_id property value. The comment_id property
// returns a *int32 when successful
func (m *Report) GetCommentId() *int32 {
	return m.comment_id
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *Report) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Report) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["action"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ParseReportActionType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAction(val.(*i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ReportActionType))
		}
		return nil
	}
	res["comment_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCommentId(val)
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
	res["reason"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ParseReportReason)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReason(val.(*i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ReportReason))
		}
		return nil
	}
	res["reported_uid"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReportedUid(val)
		}
		return nil
	}
	return res
}

// GetPostId gets the post_id property value. The post_id property
// returns a *int32 when successful
func (m *Report) GetPostId() *int32 {
	return m.post_id
}

// GetReason gets the reason property value. The reason property
// returns a *ReportReason when successful
func (m *Report) GetReason() *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ReportReason {
	return m.reason
}

// GetReportedUid gets the reported_uid property value. The reported_uid property
// returns a *string when successful
func (m *Report) GetReportedUid() *string {
	return m.reported_uid
}

// Serialize serializes information the current object
func (m *Report) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAction() != nil {
		cast := (*m.GetAction()).String()
		err := writer.WriteStringValue("action", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("comment_id", m.GetCommentId())
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
		err := writer.WriteInt32Value("post_id", m.GetPostId())
		if err != nil {
			return err
		}
	}
	if m.GetReason() != nil {
		cast := (*m.GetReason()).String()
		err := writer.WriteStringValue("reason", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("reported_uid", m.GetReportedUid())
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

// SetAction sets the action property value. The action property
func (m *Report) SetAction(value *i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ReportActionType) {
	m.action = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Report) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCommentId sets the comment_id property value. The comment_id property
func (m *Report) SetCommentId(value *int32) {
	m.comment_id = value
}

// SetDescription sets the description property value. The description property
func (m *Report) SetDescription(value *string) {
	m.description = value
}

// SetPostId sets the post_id property value. The post_id property
func (m *Report) SetPostId(value *int32) {
	m.post_id = value
}

// SetReason sets the reason property value. The reason property
func (m *Report) SetReason(value *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ReportReason) {
	m.reason = value
}

// SetReportedUid sets the reported_uid property value. The reported_uid property
func (m *Report) SetReportedUid(value *string) {
	m.reported_uid = value
}

type Reportable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAction() *i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ReportActionType
	GetCommentId() *int32
	GetDescription() *string
	GetPostId() *int32
	GetReason() *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ReportReason
	GetReportedUid() *string
	SetAction(value *i2c508c0cd86c7fab487d349cc29968385358e8062d73fe756e3fe541b62ea5d7.ReportActionType)
	SetCommentId(value *int32)
	SetDescription(value *string)
	SetPostId(value *int32)
	SetReason(value *i3c7c38f07907ad248d25ebc6f36e01f59271c2fbe64e2e996c12e980ea57ea08.ReportReason)
	SetReportedUid(value *string)
}
