package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable instead.
type ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse struct {
    ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponse
}
// NewItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse instantiates a new ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse and sets the default values.
func NewItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse()(*ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse) {
    m := &ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse{
        ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponse: *NewItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable instead.
type ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable interface {
    ItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
