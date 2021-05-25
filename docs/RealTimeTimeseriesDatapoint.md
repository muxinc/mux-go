# RealTimeTimeseriesDatapoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float64** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**ConcurentViewers** | Pointer to **int64** |  | [optional] 

## Methods

### NewRealTimeTimeseriesDatapoint

`func NewRealTimeTimeseriesDatapoint() *RealTimeTimeseriesDatapoint`

NewRealTimeTimeseriesDatapoint instantiates a new RealTimeTimeseriesDatapoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeTimeseriesDatapointWithDefaults

`func NewRealTimeTimeseriesDatapointWithDefaults() *RealTimeTimeseriesDatapoint`

NewRealTimeTimeseriesDatapointWithDefaults instantiates a new RealTimeTimeseriesDatapoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RealTimeTimeseriesDatapoint) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RealTimeTimeseriesDatapoint) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RealTimeTimeseriesDatapoint) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *RealTimeTimeseriesDatapoint) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDate

`func (o *RealTimeTimeseriesDatapoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RealTimeTimeseriesDatapoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RealTimeTimeseriesDatapoint) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RealTimeTimeseriesDatapoint) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetConcurentViewers

`func (o *RealTimeTimeseriesDatapoint) GetConcurentViewers() int64`

GetConcurentViewers returns the ConcurentViewers field if non-nil, zero value otherwise.

### GetConcurentViewersOk

`func (o *RealTimeTimeseriesDatapoint) GetConcurentViewersOk() (*int64, bool)`

GetConcurentViewersOk returns a tuple with the ConcurentViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurentViewers

`func (o *RealTimeTimeseriesDatapoint) SetConcurentViewers(v int64)`

SetConcurentViewers sets ConcurentViewers field to given value.

### HasConcurentViewers

`func (o *RealTimeTimeseriesDatapoint) HasConcurentViewers() bool`

HasConcurentViewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


