package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseable instead.
type ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse struct {
    ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponse
}
// NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse instantiates a new ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse and sets the default values.
func NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse()(*ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse) {
    m := &ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse{
        ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponse: *NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponse(),
    }
    return m
}
// CreateItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseable instead.
type ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseable interface {
    ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}