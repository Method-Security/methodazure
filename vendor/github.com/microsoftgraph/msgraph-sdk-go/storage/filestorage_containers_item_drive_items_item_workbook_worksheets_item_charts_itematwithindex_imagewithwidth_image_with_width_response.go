package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse struct {
    FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthGetResponse
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse()(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse{
        FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthGetResponse: *NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthGetResponseable instead.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthResponseable interface {
    FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthImageWithWidthGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}