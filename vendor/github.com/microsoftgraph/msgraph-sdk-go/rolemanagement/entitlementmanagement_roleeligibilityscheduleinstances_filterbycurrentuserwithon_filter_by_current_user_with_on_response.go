package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse struct {
    EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse()(*EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse) {
    m := &EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse{
        EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse: *NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable instead.
type EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnResponseable interface {
    EntitlementmanagementRoleeligibilityscheduleinstancesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}