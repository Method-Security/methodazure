package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable instead.
type ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse struct {
    ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponse
}
// NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse instantiates a new ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse and sets the default values.
func NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse()(*ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse) {
    m := &ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse{
        ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponse: *NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponse(),
    }
    return m
}
// CreateItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable instead.
type ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserResponseable interface {
    ItemCalendargroupsItemCalendarsItemAllowedcalendarsharingroleswithuserAllowedCalendarSharingRolesWithUserGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}