# RealTimeHistogramTimeseriesDatapoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**Sum** | Pointer to **int64** |  | [optional] 
**P95** | Pointer to **float64** |  | [optional] 
**Median** | Pointer to **float64** |  | [optional] 
**MaxPercentage** | Pointer to **float64** |  | [optional] 
**BucketValues** | Pointer to [**[]RealTimeHistogramTimeseriesBucketValues**](RealTimeHistogramTimeseriesBucketValues.md) |  | [optional] 
**Average** | Pointer to **float64** |  | [optional] 

## Methods

### NewRealTimeHistogramTimeseriesDatapoint

`func NewRealTimeHistogramTimeseriesDatapoint() *RealTimeHistogramTimeseriesDatapoint`

NewRealTimeHistogramTimeseriesDatapoint instantiates a new RealTimeHistogramTimeseriesDatapoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealTimeHistogramTimeseriesDatapointWithDefaults

`func NewRealTimeHistogramTimeseriesDatapointWithDefaults() *RealTimeHistogramTimeseriesDatapoint`

NewRealTimeHistogramTimeseriesDatapointWithDefaults instantiates a new RealTimeHistogramTimeseriesDatapoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *RealTimeHistogramTimeseriesDatapoint) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RealTimeHistogramTimeseriesDatapoint) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RealTimeHistogramTimeseriesDatapoint) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RealTimeHistogramTimeseriesDatapoint) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSum

`func (o *RealTimeHistogramTimeseriesDatapoint) GetSum() int64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *RealTimeHistogramTimeseriesDatapoint) GetSumOk() (*int64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *RealTimeHistogramTimeseriesDatapoint) SetSum(v int64)`

SetSum sets Sum field to given value.

### HasSum

`func (o *RealTimeHistogramTimeseriesDatapoint) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetP95

`func (o *RealTimeHistogramTimeseriesDatapoint) GetP95() float64`

GetP95 returns the P95 field if non-nil, zero value otherwise.

### GetP95Ok

`func (o *RealTimeHistogramTimeseriesDatapoint) GetP95Ok() (*float64, bool)`

GetP95Ok returns a tuple with the P95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95

`func (o *RealTimeHistogramTimeseriesDatapoint) SetP95(v float64)`

SetP95 sets P95 field to given value.

### HasP95

`func (o *RealTimeHistogramTimeseriesDatapoint) HasP95() bool`

HasP95 returns a boolean if a field has been set.

### GetMedian

`func (o *RealTimeHistogramTimeseriesDatapoint) GetMedian() float64`

GetMedian returns the Median field if non-nil, zero value otherwise.

### GetMedianOk

`func (o *RealTimeHistogramTimeseriesDatapoint) GetMedianOk() (*float64, bool)`

GetMedianOk returns a tuple with the Median field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian

`func (o *RealTimeHistogramTimeseriesDatapoint) SetMedian(v float64)`

SetMedian sets Median field to given value.

### HasMedian

`func (o *RealTimeHistogramTimeseriesDatapoint) HasMedian() bool`

HasMedian returns a boolean if a field has been set.

### GetMaxPercentage

`func (o *RealTimeHistogramTimeseriesDatapoint) GetMaxPercentage() float64`

GetMaxPercentage returns the MaxPercentage field if non-nil, zero value otherwise.

### GetMaxPercentageOk

`func (o *RealTimeHistogramTimeseriesDatapoint) GetMaxPercentageOk() (*float64, bool)`

GetMaxPercentageOk returns a tuple with the MaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPercentage

`func (o *RealTimeHistogramTimeseriesDatapoint) SetMaxPercentage(v float64)`

SetMaxPercentage sets MaxPercentage field to given value.

### HasMaxPercentage

`func (o *RealTimeHistogramTimeseriesDatapoint) HasMaxPercentage() bool`

HasMaxPercentage returns a boolean if a field has been set.

### GetBucketValues

`func (o *RealTimeHistogramTimeseriesDatapoint) GetBucketValues() []RealTimeHistogramTimeseriesBucketValues`

GetBucketValues returns the BucketValues field if non-nil, zero value otherwise.

### GetBucketValuesOk

`func (o *RealTimeHistogramTimeseriesDatapoint) GetBucketValuesOk() (*[]RealTimeHistogramTimeseriesBucketValues, bool)`

GetBucketValuesOk returns a tuple with the BucketValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketValues

`func (o *RealTimeHistogramTimeseriesDatapoint) SetBucketValues(v []RealTimeHistogramTimeseriesBucketValues)`

SetBucketValues sets BucketValues field to given value.

### HasBucketValues

`func (o *RealTimeHistogramTimeseriesDatapoint) HasBucketValues() bool`

HasBucketValues returns a boolean if a field has been set.

### GetAverage

`func (o *RealTimeHistogramTimeseriesDatapoint) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *RealTimeHistogramTimeseriesDatapoint) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *RealTimeHistogramTimeseriesDatapoint) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *RealTimeHistogramTimeseriesDatapoint) HasAverage() bool`

HasAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


