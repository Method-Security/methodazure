package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable instead.
type FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse struct {
    FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponse
}
// NewFilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse instantiates a new FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse()(*FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse) {
    m := &FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse{
        FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponse: *NewFilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable instead.
type FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesResponseable interface {
    FilestorageContainersItemDriveListContenttypesGetcompatiblehubcontenttypesGetCompatibleHubContentTypesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}