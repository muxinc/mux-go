# ListDimensionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ListFiltersResponseData**](ListFiltersResponse_data.md) |  | [optional] 
**TotalRowCount** | Pointer to **int64** |  | [optional] 
**Timeframe** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewListDimensionsResponse

`func NewListDimensionsResponse() *ListDimensionsResponse`

NewListDimensionsResponse instantiates a new ListDimensionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDimensionsResponseWithDefaults

`func NewListDimensionsResponseWithDefaults() *ListDimensionsResponse`

NewListDimensionsResponseWithDefaults instantiates a new ListDimensionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDimensionsResponse) GetData() ListFiltersResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDimensionsResponse) GetDataOk() (*ListFiltersResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDimensionsResponse) SetData(v ListFiltersResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ListDimensionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRowCount

`func (o *ListDimensionsResponse) GetTotalRowCount() int64`

GetTotalRowCount returns the TotalRowCount field if non-nil, zero value otherwise.

### GetTotalRowCountOk

`func (o *ListDimensionsResponse) GetTotalRowCountOk() (*int64, bool)`

GetTotalRowCountOk returns a tuple with the TotalRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRowCount

`func (o *ListDimensionsResponse) SetTotalRowCount(v int64)`

SetTotalRowCount sets TotalRowCount field to given value.

### HasTotalRowCount

`func (o *ListDimensionsResponse) HasTotalRowCount() bool`

HasTotalRowCount returns a boolean if a field has been set.

### GetTimeframe

`func (o *ListDimensionsResponse) GetTimeframe() []int64`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ListDimensionsResponse) GetTimeframeOk() (*[]int64, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ListDimensionsResponse) SetTimeframe(v []int64)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ListDimensionsResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


