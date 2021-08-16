# Incident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **float64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SampleSizeUnit** | Pointer to **string** |  | [optional] 
**SampleSize** | Pointer to **int64** |  | [optional] 
**ResolvedAt** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**[]IncidentNotification**](IncidentNotification.md) |  | [optional] 
**NotificationRules** | Pointer to [**[]IncidentNotificationRule**](IncidentNotificationRule.md) |  | [optional] 
**Measurement** | Pointer to **string** |  | [optional] 
**MeasuredValueOnClose** | Pointer to **float64** |  | [optional] 
**MeasuredValue** | Pointer to **float64** |  | [optional] 
**IncidentKey** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Breakdowns** | Pointer to [**[]IncidentBreakdown**](IncidentBreakdown.md) |  | [optional] 
**AffectedViewsPerHourOnOpen** | Pointer to **int64** |  | [optional] 
**AffectedViewsPerHour** | Pointer to **int64** |  | [optional] 
**AffectedViews** | Pointer to **int64** |  | [optional] 

## Methods

### NewIncident

`func NewIncident() *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *Incident) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Incident) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Incident) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Incident) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetStatus

`func (o *Incident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Incident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Incident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Incident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *Incident) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Incident) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Incident) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Incident) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSeverity

`func (o *Incident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Incident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Incident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Incident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSampleSizeUnit

`func (o *Incident) GetSampleSizeUnit() string`

GetSampleSizeUnit returns the SampleSizeUnit field if non-nil, zero value otherwise.

### GetSampleSizeUnitOk

`func (o *Incident) GetSampleSizeUnitOk() (*string, bool)`

GetSampleSizeUnitOk returns a tuple with the SampleSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSizeUnit

`func (o *Incident) SetSampleSizeUnit(v string)`

SetSampleSizeUnit sets SampleSizeUnit field to given value.

### HasSampleSizeUnit

`func (o *Incident) HasSampleSizeUnit() bool`

HasSampleSizeUnit returns a boolean if a field has been set.

### GetSampleSize

`func (o *Incident) GetSampleSize() int64`

GetSampleSize returns the SampleSize field if non-nil, zero value otherwise.

### GetSampleSizeOk

`func (o *Incident) GetSampleSizeOk() (*int64, bool)`

GetSampleSizeOk returns a tuple with the SampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSize

`func (o *Incident) SetSampleSize(v int64)`

SetSampleSize sets SampleSize field to given value.

### HasSampleSize

`func (o *Incident) HasSampleSize() bool`

HasSampleSize returns a boolean if a field has been set.

### GetResolvedAt

`func (o *Incident) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Incident) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Incident) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *Incident) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetNotifications

`func (o *Incident) GetNotifications() []IncidentNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Incident) GetNotificationsOk() (*[]IncidentNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Incident) SetNotifications(v []IncidentNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Incident) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetNotificationRules

`func (o *Incident) GetNotificationRules() []IncidentNotificationRule`

GetNotificationRules returns the NotificationRules field if non-nil, zero value otherwise.

### GetNotificationRulesOk

`func (o *Incident) GetNotificationRulesOk() (*[]IncidentNotificationRule, bool)`

GetNotificationRulesOk returns a tuple with the NotificationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRules

`func (o *Incident) SetNotificationRules(v []IncidentNotificationRule)`

SetNotificationRules sets NotificationRules field to given value.

### HasNotificationRules

`func (o *Incident) HasNotificationRules() bool`

HasNotificationRules returns a boolean if a field has been set.

### GetMeasurement

`func (o *Incident) GetMeasurement() string`

GetMeasurement returns the Measurement field if non-nil, zero value otherwise.

### GetMeasurementOk

`func (o *Incident) GetMeasurementOk() (*string, bool)`

GetMeasurementOk returns a tuple with the Measurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurement

`func (o *Incident) SetMeasurement(v string)`

SetMeasurement sets Measurement field to given value.

### HasMeasurement

`func (o *Incident) HasMeasurement() bool`

HasMeasurement returns a boolean if a field has been set.

### GetMeasuredValueOnClose

`func (o *Incident) GetMeasuredValueOnClose() float64`

GetMeasuredValueOnClose returns the MeasuredValueOnClose field if non-nil, zero value otherwise.

### GetMeasuredValueOnCloseOk

`func (o *Incident) GetMeasuredValueOnCloseOk() (*float64, bool)`

GetMeasuredValueOnCloseOk returns a tuple with the MeasuredValueOnClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredValueOnClose

`func (o *Incident) SetMeasuredValueOnClose(v float64)`

SetMeasuredValueOnClose sets MeasuredValueOnClose field to given value.

### HasMeasuredValueOnClose

`func (o *Incident) HasMeasuredValueOnClose() bool`

HasMeasuredValueOnClose returns a boolean if a field has been set.

### GetMeasuredValue

`func (o *Incident) GetMeasuredValue() float64`

GetMeasuredValue returns the MeasuredValue field if non-nil, zero value otherwise.

### GetMeasuredValueOk

`func (o *Incident) GetMeasuredValueOk() (*float64, bool)`

GetMeasuredValueOk returns a tuple with the MeasuredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredValue

`func (o *Incident) SetMeasuredValue(v float64)`

SetMeasuredValue sets MeasuredValue field to given value.

### HasMeasuredValue

`func (o *Incident) HasMeasuredValue() bool`

HasMeasuredValue returns a boolean if a field has been set.

### GetIncidentKey

`func (o *Incident) GetIncidentKey() string`

GetIncidentKey returns the IncidentKey field if non-nil, zero value otherwise.

### GetIncidentKeyOk

`func (o *Incident) GetIncidentKeyOk() (*string, bool)`

GetIncidentKeyOk returns a tuple with the IncidentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentKey

`func (o *Incident) SetIncidentKey(v string)`

SetIncidentKey sets IncidentKey field to given value.

### HasIncidentKey

`func (o *Incident) HasIncidentKey() bool`

HasIncidentKey returns a boolean if a field has been set.

### GetImpact

`func (o *Incident) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *Incident) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *Incident) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *Incident) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetId

`func (o *Incident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incident) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Incident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrorDescription

`func (o *Incident) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *Incident) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *Incident) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *Incident) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetDescription

`func (o *Incident) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Incident) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Incident) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Incident) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBreakdowns

`func (o *Incident) GetBreakdowns() []IncidentBreakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *Incident) GetBreakdownsOk() (*[]IncidentBreakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *Incident) SetBreakdowns(v []IncidentBreakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *Incident) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.

### GetAffectedViewsPerHourOnOpen

`func (o *Incident) GetAffectedViewsPerHourOnOpen() int64`

GetAffectedViewsPerHourOnOpen returns the AffectedViewsPerHourOnOpen field if non-nil, zero value otherwise.

### GetAffectedViewsPerHourOnOpenOk

`func (o *Incident) GetAffectedViewsPerHourOnOpenOk() (*int64, bool)`

GetAffectedViewsPerHourOnOpenOk returns a tuple with the AffectedViewsPerHourOnOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedViewsPerHourOnOpen

`func (o *Incident) SetAffectedViewsPerHourOnOpen(v int64)`

SetAffectedViewsPerHourOnOpen sets AffectedViewsPerHourOnOpen field to given value.

### HasAffectedViewsPerHourOnOpen

`func (o *Incident) HasAffectedViewsPerHourOnOpen() bool`

HasAffectedViewsPerHourOnOpen returns a boolean if a field has been set.

### GetAffectedViewsPerHour

`func (o *Incident) GetAffectedViewsPerHour() int64`

GetAffectedViewsPerHour returns the AffectedViewsPerHour field if non-nil, zero value otherwise.

### GetAffectedViewsPerHourOk

`func (o *Incident) GetAffectedViewsPerHourOk() (*int64, bool)`

GetAffectedViewsPerHourOk returns a tuple with the AffectedViewsPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedViewsPerHour

`func (o *Incident) SetAffectedViewsPerHour(v int64)`

SetAffectedViewsPerHour sets AffectedViewsPerHour field to given value.

### HasAffectedViewsPerHour

`func (o *Incident) HasAffectedViewsPerHour() bool`

HasAffectedViewsPerHour returns a boolean if a field has been set.

### GetAffectedViews

`func (o *Incident) GetAffectedViews() int64`

GetAffectedViews returns the AffectedViews field if non-nil, zero value otherwise.

### GetAffectedViewsOk

`func (o *Incident) GetAffectedViewsOk() (*int64, bool)`

GetAffectedViewsOk returns a tuple with the AffectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedViews

`func (o *Incident) SetAffectedViews(v int64)`

SetAffectedViews sets AffectedViews field to given value.

### HasAffectedViews

`func (o *Incident) HasAffectedViews() bool`

HasAffectedViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


