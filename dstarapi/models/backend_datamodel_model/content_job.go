package backend_datamodel_model

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1 "github.com/0x090909/dstarapi/models/gorm_io_gorm"
)

type ContentJob struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The category property
    category *int32
    // The createdAt property
    createdAt *string
    // The deletedAt property
    deletedAt i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable
    // The id property
    id *int32
    // The scraper_source property
    scraper_source *ScraperSource
    // The scraper_type property
    scraper_type *ScraperType
    // The search_keywords property
    search_keywords *string
    // The status property
    status *JobStatus
    // The uid property
    uid *string
    // The updatedAt property
    updatedAt *string
}
// NewContentJob instantiates a new ContentJob and sets the default values.
func NewContentJob()(*ContentJob) {
    m := &ContentJob{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContentJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentJob(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContentJob) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCategory gets the category property value. The category property
// returns a *int32 when successful
func (m *ContentJob) GetCategory()(*int32) {
    return m.category
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *string when successful
func (m *ContentJob) GetCreatedAt()(*string) {
    return m.createdAt
}
// GetDeletedAt gets the deletedAt property value. The deletedAt property
// returns a DeletedAtable when successful
func (m *ContentJob) GetDeletedAt()(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable) {
    return m.deletedAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContentJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["scraper_source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScraperSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScraperSource(val.(*ScraperSource))
        }
        return nil
    }
    res["scraper_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScraperType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScraperType(val.(*ScraperType))
        }
        return nil
    }
    res["search_keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchKeywords(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseJobStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*JobStatus))
        }
        return nil
    }
    res["uid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUid(val)
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
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *ContentJob) GetId()(*int32) {
    return m.id
}
// GetScraperSource gets the scraper_source property value. The scraper_source property
// returns a *ScraperSource when successful
func (m *ContentJob) GetScraperSource()(*ScraperSource) {
    return m.scraper_source
}
// GetScraperType gets the scraper_type property value. The scraper_type property
// returns a *ScraperType when successful
func (m *ContentJob) GetScraperType()(*ScraperType) {
    return m.scraper_type
}
// GetSearchKeywords gets the search_keywords property value. The search_keywords property
// returns a *string when successful
func (m *ContentJob) GetSearchKeywords()(*string) {
    return m.search_keywords
}
// GetStatus gets the status property value. The status property
// returns a *JobStatus when successful
func (m *ContentJob) GetStatus()(*JobStatus) {
    return m.status
}
// GetUid gets the uid property value. The uid property
// returns a *string when successful
func (m *ContentJob) GetUid()(*string) {
    return m.uid
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *string when successful
func (m *ContentJob) GetUpdatedAt()(*string) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *ContentJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("category", m.GetCategory())
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
        err := writer.WriteObjectValue("deletedAt", m.GetDeletedAt())
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
    if m.GetScraperSource() != nil {
        cast := (*m.GetScraperSource()).String()
        err := writer.WriteStringValue("scraper_source", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScraperType() != nil {
        cast := (*m.GetScraperType()).String()
        err := writer.WriteStringValue("scraper_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("search_keywords", m.GetSearchKeywords())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
        err := writer.WriteStringValue("updatedAt", m.GetUpdatedAt())
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
func (m *ContentJob) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCategory sets the category property value. The category property
func (m *ContentJob) SetCategory(value *int32)() {
    m.category = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *ContentJob) SetCreatedAt(value *string)() {
    m.createdAt = value
}
// SetDeletedAt sets the deletedAt property value. The deletedAt property
func (m *ContentJob) SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)() {
    m.deletedAt = value
}
// SetId sets the id property value. The id property
func (m *ContentJob) SetId(value *int32)() {
    m.id = value
}
// SetScraperSource sets the scraper_source property value. The scraper_source property
func (m *ContentJob) SetScraperSource(value *ScraperSource)() {
    m.scraper_source = value
}
// SetScraperType sets the scraper_type property value. The scraper_type property
func (m *ContentJob) SetScraperType(value *ScraperType)() {
    m.scraper_type = value
}
// SetSearchKeywords sets the search_keywords property value. The search_keywords property
func (m *ContentJob) SetSearchKeywords(value *string)() {
    m.search_keywords = value
}
// SetStatus sets the status property value. The status property
func (m *ContentJob) SetStatus(value *JobStatus)() {
    m.status = value
}
// SetUid sets the uid property value. The uid property
func (m *ContentJob) SetUid(value *string)() {
    m.uid = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *ContentJob) SetUpdatedAt(value *string)() {
    m.updatedAt = value
}
type ContentJobable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*int32)
    GetCreatedAt()(*string)
    GetDeletedAt()(i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)
    GetId()(*int32)
    GetScraperSource()(*ScraperSource)
    GetScraperType()(*ScraperType)
    GetSearchKeywords()(*string)
    GetStatus()(*JobStatus)
    GetUid()(*string)
    GetUpdatedAt()(*string)
    SetCategory(value *int32)()
    SetCreatedAt(value *string)()
    SetDeletedAt(value i449f0abbc73d489f3793c3a55ad3eb2a7e83fc5ec786e2996b64d2d5d50119c1.DeletedAtable)()
    SetId(value *int32)()
    SetScraperSource(value *ScraperSource)()
    SetScraperType(value *ScraperType)()
    SetSearchKeywords(value *string)()
    SetStatus(value *JobStatus)()
    SetUid(value *string)()
    SetUpdatedAt(value *string)()
}
