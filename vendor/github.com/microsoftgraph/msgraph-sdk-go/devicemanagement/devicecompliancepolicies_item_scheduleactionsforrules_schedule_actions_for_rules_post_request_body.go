package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody instantiates a new DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody and sets the default values.
func NewDevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody()(*DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) {
    m := &DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceComplianceScheduledActionForRules gets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
// returns a []DeviceComplianceScheduledActionForRuleable when successful
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) GetDeviceComplianceScheduledActionForRules()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable) {
    val, err := m.GetBackingStore().Get("deviceComplianceScheduledActionForRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceScheduledActionForRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
                }
            }
            m.SetDeviceComplianceScheduledActionForRules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceComplianceScheduledActionForRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceComplianceScheduledActionForRules()))
        for i, v := range m.GetDeviceComplianceScheduledActionForRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("deviceComplianceScheduledActionForRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceComplianceScheduledActionForRules sets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBody) SetDeviceComplianceScheduledActionForRules(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)() {
    err := m.GetBackingStore().Set("deviceComplianceScheduledActionForRules", value)
    if err != nil {
        panic(err)
    }
}
type DevicecompliancepoliciesItemScheduleactionsforrulesScheduleActionsForRulesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceComplianceScheduledActionForRules()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceComplianceScheduledActionForRules(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)()
}
