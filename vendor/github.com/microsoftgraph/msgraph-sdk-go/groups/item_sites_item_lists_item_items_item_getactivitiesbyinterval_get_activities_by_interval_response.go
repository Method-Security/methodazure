package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable instead.
type ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse struct {
    ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponse
}
// NewItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse instantiates a new ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse and sets the default values.
func NewItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse()(*ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse) {
    m := &ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse{
        ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponse: *NewItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable instead.
type ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable interface {
    ItemSitesItemListsItemItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
