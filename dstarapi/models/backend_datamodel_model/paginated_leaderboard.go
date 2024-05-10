package backend_datamodel_model

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaginatedLeaderboard struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The entries property
    entries []Leaderboardable
    // The remaining property
    remaining *int32
    // The total property
    total *int32
    // The user_competing_posts_count property
    user_competing_posts_count *int32
    // The user_entry property
    user_entry Leaderboardable
}
// NewPaginatedLeaderboard instantiates a new PaginatedLeaderboard and sets the default values.
func NewPaginatedLeaderboard()(*PaginatedLeaderboard) {
    m := &PaginatedLeaderboard{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePaginatedLeaderboardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaginatedLeaderboardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPaginatedLeaderboard(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaginatedLeaderboard) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEntries gets the entries property value. The entries property
// returns a []Leaderboardable when successful
func (m *PaginatedLeaderboard) GetEntries()([]Leaderboardable) {
    return m.entries
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaginatedLeaderboard) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLeaderboardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Leaderboardable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Leaderboardable)
                }
            }
            m.SetEntries(res)
        }
        return nil
    }
    res["remaining"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemaining(val)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["user_competing_posts_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCompetingPostsCount(val)
        }
        return nil
    }
    res["user_entry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLeaderboardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEntry(val.(Leaderboardable))
        }
        return nil
    }
    return res
}
// GetRemaining gets the remaining property value. The remaining property
// returns a *int32 when successful
func (m *PaginatedLeaderboard) GetRemaining()(*int32) {
    return m.remaining
}
// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *PaginatedLeaderboard) GetTotal()(*int32) {
    return m.total
}
// GetUserCompetingPostsCount gets the user_competing_posts_count property value. The user_competing_posts_count property
// returns a *int32 when successful
func (m *PaginatedLeaderboard) GetUserCompetingPostsCount()(*int32) {
    return m.user_competing_posts_count
}
// GetUserEntry gets the user_entry property value. The user_entry property
// returns a Leaderboardable when successful
func (m *PaginatedLeaderboard) GetUserEntry()(Leaderboardable) {
    return m.user_entry
}
// Serialize serializes information the current object
func (m *PaginatedLeaderboard) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("user_competing_posts_count", m.GetUserCompetingPostsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user_entry", m.GetUserEntry())
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
func (m *PaginatedLeaderboard) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEntries sets the entries property value. The entries property
func (m *PaginatedLeaderboard) SetEntries(value []Leaderboardable)() {
    m.entries = value
}
// SetRemaining sets the remaining property value. The remaining property
func (m *PaginatedLeaderboard) SetRemaining(value *int32)() {
    m.remaining = value
}
// SetTotal sets the total property value. The total property
func (m *PaginatedLeaderboard) SetTotal(value *int32)() {
    m.total = value
}
// SetUserCompetingPostsCount sets the user_competing_posts_count property value. The user_competing_posts_count property
func (m *PaginatedLeaderboard) SetUserCompetingPostsCount(value *int32)() {
    m.user_competing_posts_count = value
}
// SetUserEntry sets the user_entry property value. The user_entry property
func (m *PaginatedLeaderboard) SetUserEntry(value Leaderboardable)() {
    m.user_entry = value
}
type PaginatedLeaderboardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEntries()([]Leaderboardable)
    GetRemaining()(*int32)
    GetTotal()(*int32)
    GetUserCompetingPostsCount()(*int32)
    GetUserEntry()(Leaderboardable)
    SetEntries(value []Leaderboardable)()
    SetRemaining(value *int32)()
    SetTotal(value *int32)()
    SetUserCompetingPostsCount(value *int32)()
    SetUserEntry(value Leaderboardable)()
}
