package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponseable instead.
type FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse struct {
    FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponse
}
// NewFilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse instantiates a new FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse and sets the default values.
func NewFilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse()(*FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse) {
    m := &FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse{
        FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponse: *NewFilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponseable instead.
type FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenResponseable interface {
    FilestorageContainersItemDriveListItemsDeltawithtokenDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}