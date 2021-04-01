# IncidentNotificationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]NotificationRule**](NotificationRule.md) |  | [optional] 
**PropertyId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 

## Methods

### NewIncidentNotificationRule

`func NewIncidentNotificationRule() *IncidentNotificationRule`

NewIncidentNotificationRule instantiates a new IncidentNotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentNotificationRuleWithDefaults

`func NewIncidentNotificationRuleWithDefaults() *IncidentNotificationRule`

NewIncidentNotificationRuleWithDefaults instantiates a new IncidentNotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IncidentNotificationRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentNotificationRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentNotificationRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncidentNotificationRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRules

`func (o *IncidentNotificationRule) GetRules() []NotificationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IncidentNotificationRule) GetRulesOk() (*[]NotificationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IncidentNotificationRule) SetRules(v []NotificationRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IncidentNotificationRule) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetPropertyId

`func (o *IncidentNotificationRule) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *IncidentNotificationRule) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *IncidentNotificationRule) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *IncidentNotificationRule) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetId

`func (o *IncidentNotificationRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentNotificationRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentNotificationRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentNotificationRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *IncidentNotificationRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IncidentNotificationRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IncidentNotificationRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IncidentNotificationRule) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


