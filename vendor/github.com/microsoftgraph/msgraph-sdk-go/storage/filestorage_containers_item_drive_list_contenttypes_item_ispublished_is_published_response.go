package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseable instead.
type FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse struct {
    FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponse
}
// NewFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse instantiates a new FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse()(*FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse) {
    m := &FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse{
        FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponse: *NewFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseable instead.
type FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedResponseable interface {
    FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
