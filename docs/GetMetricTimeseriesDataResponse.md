# GetMetricTimeseriesDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[][]string** |  | [optional] 
**TotalRowCount** | Pointer to **int64** |  | [optional] 
**Timeframe** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewGetMetricTimeseriesDataResponse

`func NewGetMetricTimeseriesDataResponse() *GetMetricTimeseriesDataResponse`

NewGetMetricTimeseriesDataResponse instantiates a new GetMetricTimeseriesDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricTimeseriesDataResponseWithDefaults

`func NewGetMetricTimeseriesDataResponseWithDefaults() *GetMetricTimeseriesDataResponse`

NewGetMetricTimeseriesDataResponseWithDefaults instantiates a new GetMetricTimeseriesDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMetricTimeseriesDataResponse) GetData() [][]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMetricTimeseriesDataResponse) GetDataOk() (*[][]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMetricTimeseriesDataResponse) SetData(v [][]string)`

SetData sets Data field to given value.

### HasData

`func (o *GetMetricTimeseriesDataResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRowCount

`func (o *GetMetricTimeseriesDataResponse) GetTotalRowCount() int64`

GetTotalRowCount returns the TotalRowCount field if non-nil, zero value otherwise.

### GetTotalRowCountOk

`func (o *GetMetricTimeseriesDataResponse) GetTotalRowCountOk() (*int64, bool)`

GetTotalRowCountOk returns a tuple with the TotalRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRowCount

`func (o *GetMetricTimeseriesDataResponse) SetTotalRowCount(v int64)`

SetTotalRowCount sets TotalRowCount field to given value.

### HasTotalRowCount

`func (o *GetMetricTimeseriesDataResponse) HasTotalRowCount() bool`

HasTotalRowCount returns a boolean if a field has been set.

### GetTimeframe

`func (o *GetMetricTimeseriesDataResponse) GetTimeframe() []int64`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *GetMetricTimeseriesDataResponse) GetTimeframeOk() (*[]int64, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *GetMetricTimeseriesDataResponse) SetTimeframe(v []int64)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *GetMetricTimeseriesDataResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


