package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}