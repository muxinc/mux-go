# ListVideoViewExportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ExportDate**](ExportDate.md) |  | [optional] 
**TotalRowCount** | Pointer to **int32** |  | [optional] 
**Timeframe** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewListVideoViewExportsResponse

`func NewListVideoViewExportsResponse() *ListVideoViewExportsResponse`

NewListVideoViewExportsResponse instantiates a new ListVideoViewExportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVideoViewExportsResponseWithDefaults

`func NewListVideoViewExportsResponseWithDefaults() *ListVideoViewExportsResponse`

NewListVideoViewExportsResponseWithDefaults instantiates a new ListVideoViewExportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListVideoViewExportsResponse) GetData() []ExportDate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListVideoViewExportsResponse) GetDataOk() (*[]ExportDate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListVideoViewExportsResponse) SetData(v []ExportDate)`

SetData sets Data field to given value.

### HasData

`func (o *ListVideoViewExportsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRowCount

`func (o *ListVideoViewExportsResponse) GetTotalRowCount() int32`

GetTotalRowCount returns the TotalRowCount field if non-nil, zero value otherwise.

### GetTotalRowCountOk

`func (o *ListVideoViewExportsResponse) GetTotalRowCountOk() (*int32, bool)`

GetTotalRowCountOk returns a tuple with the TotalRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRowCount

`func (o *ListVideoViewExportsResponse) SetTotalRowCount(v int32)`

SetTotalRowCount sets TotalRowCount field to given value.

### HasTotalRowCount

`func (o *ListVideoViewExportsResponse) HasTotalRowCount() bool`

HasTotalRowCount returns a boolean if a field has been set.

### GetTimeframe

`func (o *ListVideoViewExportsResponse) GetTimeframe() []int32`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ListVideoViewExportsResponse) GetTimeframeOk() (*[]int32, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ListVideoViewExportsResponse) SetTimeframe(v []int32)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ListVideoViewExportsResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


