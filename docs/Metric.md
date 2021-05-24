# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Measurement** | Pointer to **string** |  | [optional] 

## Methods

### NewMetric

`func NewMetric() *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Metric) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Metric) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Metric) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *Metric) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *Metric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Metric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Metric) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Metric) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Metric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetric

`func (o *Metric) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Metric) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Metric) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Metric) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetMeasurement

`func (o *Metric) GetMeasurement() string`

GetMeasurement returns the Measurement field if non-nil, zero value otherwise.

### GetMeasurementOk

`func (o *Metric) GetMeasurementOk() (*string, bool)`

GetMeasurementOk returns a tuple with the Measurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurement

`func (o *Metric) SetMeasurement(v string)`

SetMeasurement sets Measurement field to given value.

### HasMeasurement

`func (o *Metric) HasMeasurement() bool`

HasMeasurement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


