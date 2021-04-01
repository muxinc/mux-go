# DeliveryReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveStreamId** | Pointer to **string** | Unique identifier for the live stream that created the asset. | [optional] 
**AssetId** | Pointer to **string** | Unique identifier for the asset. | [optional] 
**Passthrough** | Pointer to **string** | The &#x60;passthrough&#x60; value for the asset. | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the asset was created. Measured in seconds since the Unix epoch. | [optional] 
**AssetState** | Pointer to **string** | The state of the asset. | [optional] 
**AssetDuration** | Pointer to **float64** | The duration of the asset in seconds. | [optional] 
**DeliveredSeconds** | Pointer to **float64** | Total number of delivered seconds during this time window. | [optional] 

## Methods

### NewDeliveryReport

`func NewDeliveryReport() *DeliveryReport`

NewDeliveryReport instantiates a new DeliveryReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryReportWithDefaults

`func NewDeliveryReportWithDefaults() *DeliveryReport`

NewDeliveryReportWithDefaults instantiates a new DeliveryReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveStreamId

`func (o *DeliveryReport) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *DeliveryReport) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *DeliveryReport) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *DeliveryReport) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetAssetId

`func (o *DeliveryReport) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *DeliveryReport) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *DeliveryReport) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *DeliveryReport) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetPassthrough

`func (o *DeliveryReport) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *DeliveryReport) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *DeliveryReport) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *DeliveryReport) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeliveryReport) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeliveryReport) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeliveryReport) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeliveryReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAssetState

`func (o *DeliveryReport) GetAssetState() string`

GetAssetState returns the AssetState field if non-nil, zero value otherwise.

### GetAssetStateOk

`func (o *DeliveryReport) GetAssetStateOk() (*string, bool)`

GetAssetStateOk returns a tuple with the AssetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetState

`func (o *DeliveryReport) SetAssetState(v string)`

SetAssetState sets AssetState field to given value.

### HasAssetState

`func (o *DeliveryReport) HasAssetState() bool`

HasAssetState returns a boolean if a field has been set.

### GetAssetDuration

`func (o *DeliveryReport) GetAssetDuration() float64`

GetAssetDuration returns the AssetDuration field if non-nil, zero value otherwise.

### GetAssetDurationOk

`func (o *DeliveryReport) GetAssetDurationOk() (*float64, bool)`

GetAssetDurationOk returns a tuple with the AssetDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetDuration

`func (o *DeliveryReport) SetAssetDuration(v float64)`

SetAssetDuration sets AssetDuration field to given value.

### HasAssetDuration

`func (o *DeliveryReport) HasAssetDuration() bool`

HasAssetDuration returns a boolean if a field has been set.

### GetDeliveredSeconds

`func (o *DeliveryReport) GetDeliveredSeconds() float64`

GetDeliveredSeconds returns the DeliveredSeconds field if non-nil, zero value otherwise.

### GetDeliveredSecondsOk

`func (o *DeliveryReport) GetDeliveredSecondsOk() (*float64, bool)`

GetDeliveredSecondsOk returns a tuple with the DeliveredSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredSeconds

`func (o *DeliveryReport) SetDeliveredSeconds(v float64)`

SetDeliveredSeconds sets DeliveredSeconds field to given value.

### HasDeliveredSeconds

`func (o *DeliveryReport) HasDeliveredSeconds() bool`

HasDeliveredSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


