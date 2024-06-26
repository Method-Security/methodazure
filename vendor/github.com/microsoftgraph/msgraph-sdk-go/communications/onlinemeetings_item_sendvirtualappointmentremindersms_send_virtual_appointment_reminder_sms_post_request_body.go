package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody instantiates a new OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody and sets the default values.
func NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody()(*OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) {
    m := &OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAttendees gets the attendees property value. The attendees property
// returns a []AttendeeNotificationInfoable when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetAttendees()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable) {
    val, err := m.GetBackingStore().Get("attendees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttendeeNotificationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable)
                }
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["remindBeforeTimeInMinutesType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseRemindBeforeTimeInMinutesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindBeforeTimeInMinutesType(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemindBeforeTimeInMinutesType))
        }
        return nil
    }
    return res
}
// GetRemindBeforeTimeInMinutesType gets the remindBeforeTimeInMinutesType property value. The remindBeforeTimeInMinutesType property
// returns a *RemindBeforeTimeInMinutesType when successful
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) GetRemindBeforeTimeInMinutesType()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemindBeforeTimeInMinutesType) {
    val, err := m.GetBackingStore().Get("remindBeforeTimeInMinutesType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemindBeforeTimeInMinutesType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemindBeforeTimeInMinutesType() != nil {
        cast := (*m.GetRemindBeforeTimeInMinutesType()).String()
        err := writer.WriteStringValue("remindBeforeTimeInMinutesType", &cast)
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
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendees sets the attendees property value. The attendees property
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetAttendees(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable)() {
    err := m.GetBackingStore().Set("attendees", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetRemindBeforeTimeInMinutesType sets the remindBeforeTimeInMinutesType property value. The remindBeforeTimeInMinutesType property
func (m *OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBody) SetRemindBeforeTimeInMinutesType(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemindBeforeTimeInMinutesType)() {
    err := m.GetBackingStore().Set("remindBeforeTimeInMinutesType", value)
    if err != nil {
        panic(err)
    }
}
type OnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetRemindBeforeTimeInMinutesType()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemindBeforeTimeInMinutesType)
    SetAttendees(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeNotificationInfoable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetRemindBeforeTimeInMinutesType(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemindBeforeTimeInMinutesType)()
}
