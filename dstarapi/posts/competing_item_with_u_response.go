package posts

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CompetingItemWithUGetResponseable instead.
type CompetingItemWithUResponse struct {
    CompetingItemWithUGetResponse
}
// NewCompetingItemWithUResponse instantiates a new CompetingItemWithUResponse and sets the default values.
func NewCompetingItemWithUResponse()(*CompetingItemWithUResponse) {
    m := &CompetingItemWithUResponse{
        CompetingItemWithUGetResponse: *NewCompetingItemWithUGetResponse(),
    }
    return m
}
// CreateCompetingItemWithUResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompetingItemWithUResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompetingItemWithUResponse(), nil
}
// Deprecated: This class is obsolete. Use CompetingItemWithUGetResponseable instead.
type CompetingItemWithUResponseable interface {
    CompetingItemWithUGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
