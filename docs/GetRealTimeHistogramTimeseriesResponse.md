# GetRealTimeHistogramTimeseriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**GetRealTimeHistogramTimeseriesResponseMeta**](GetRealTimeHistogramTimeseriesResponse_meta.md) |  | [optional] 
**Data** | Pointer to [**[]RealTimeHistogramTimeseriesDatapoint**](RealTimeHistogramTimeseriesDatapoint.md) |  | [optional] 
**TotalRowCount** | Pointer to **int64** |  | [optional] 
**Timeframe** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewGetRealTimeHistogramTimeseriesResponse

`func NewGetRealTimeHistogramTimeseriesResponse() *GetRealTimeHistogramTimeseriesResponse`

NewGetRealTimeHistogramTimeseriesResponse instantiates a new GetRealTimeHistogramTimeseriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRealTimeHistogramTimeseriesResponseWithDefaults

`func NewGetRealTimeHistogramTimeseriesResponseWithDefaults() *GetRealTimeHistogramTimeseriesResponse`

NewGetRealTimeHistogramTimeseriesResponseWithDefaults instantiates a new GetRealTimeHistogramTimeseriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetRealTimeHistogramTimeseriesResponse) GetMeta() GetRealTimeHistogramTimeseriesResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRealTimeHistogramTimeseriesResponse) GetMetaOk() (*GetRealTimeHistogramTimeseriesResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRealTimeHistogramTimeseriesResponse) SetMeta(v GetRealTimeHistogramTimeseriesResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetRealTimeHistogramTimeseriesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *GetRealTimeHistogramTimeseriesResponse) GetData() []RealTimeHistogramTimeseriesDatapoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRealTimeHistogramTimeseriesResponse) GetDataOk() (*[]RealTimeHistogramTimeseriesDatapoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRealTimeHistogramTimeseriesResponse) SetData(v []RealTimeHistogramTimeseriesDatapoint)`

SetData sets Data field to given value.

### HasData

`func (o *GetRealTimeHistogramTimeseriesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRowCount

`func (o *GetRealTimeHistogramTimeseriesResponse) GetTotalRowCount() int64`

GetTotalRowCount returns the TotalRowCount field if non-nil, zero value otherwise.

### GetTotalRowCountOk

`func (o *GetRealTimeHistogramTimeseriesResponse) GetTotalRowCountOk() (*int64, bool)`

GetTotalRowCountOk returns a tuple with the TotalRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRowCount

`func (o *GetRealTimeHistogramTimeseriesResponse) SetTotalRowCount(v int64)`

SetTotalRowCount sets TotalRowCount field to given value.

### HasTotalRowCount

`func (o *GetRealTimeHistogramTimeseriesResponse) HasTotalRowCount() bool`

HasTotalRowCount returns a boolean if a field has been set.

### GetTimeframe

`func (o *GetRealTimeHistogramTimeseriesResponse) GetTimeframe() []int64`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *GetRealTimeHistogramTimeseriesResponse) GetTimeframeOk() (*[]int64, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *GetRealTimeHistogramTimeseriesResponse) SetTimeframe(v []int64)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *GetRealTimeHistogramTimeseriesResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


