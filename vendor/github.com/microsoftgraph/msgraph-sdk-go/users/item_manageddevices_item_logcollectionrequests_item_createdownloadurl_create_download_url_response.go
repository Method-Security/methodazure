package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable instead.
type ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse struct {
    ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponse
}
// NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse instantiates a new ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse()(*ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse) {
    m := &ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse{
        ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponse: *NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponse(),
    }
    return m
}
// CreateItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable instead.
type ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlResponseable interface {
    ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}