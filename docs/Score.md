# Score

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WatchTime** | Pointer to **int64** |  | [optional] 
**ViewCount** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]Metric**](Metric.md) |  | [optional] 

## Methods

### NewScore

`func NewScore() *Score`

NewScore instantiates a new Score object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreWithDefaults

`func NewScoreWithDefaults() *Score`

NewScoreWithDefaults instantiates a new Score object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWatchTime

`func (o *Score) GetWatchTime() int64`

GetWatchTime returns the WatchTime field if non-nil, zero value otherwise.

### GetWatchTimeOk

`func (o *Score) GetWatchTimeOk() (*int64, bool)`

GetWatchTimeOk returns a tuple with the WatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchTime

`func (o *Score) SetWatchTime(v int64)`

SetWatchTime sets WatchTime field to given value.

### HasWatchTime

`func (o *Score) HasWatchTime() bool`

HasWatchTime returns a boolean if a field has been set.

### GetViewCount

`func (o *Score) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *Score) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *Score) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *Score) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetName

`func (o *Score) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Score) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Score) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Score) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *Score) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Score) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Score) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *Score) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMetric

`func (o *Score) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Score) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Score) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Score) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetItems

`func (o *Score) GetItems() []Metric`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Score) GetItemsOk() (*[]Metric, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Score) SetItems(v []Metric)`

SetItems sets Items field to given value.

### HasItems

`func (o *Score) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


