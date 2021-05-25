# ListDeliveryUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DeliveryReport**](DeliveryReport.md) |  | [optional] 
**TotalRowCount** | Pointer to **int64** |  | [optional] 
**Timeframe** | Pointer to **[]int64** |  | [optional] 
**Limit** | Pointer to **int64** | Number of assets returned in this response. Default value is 100. | [optional] 

## Methods

### NewListDeliveryUsageResponse

`func NewListDeliveryUsageResponse() *ListDeliveryUsageResponse`

NewListDeliveryUsageResponse instantiates a new ListDeliveryUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeliveryUsageResponseWithDefaults

`func NewListDeliveryUsageResponseWithDefaults() *ListDeliveryUsageResponse`

NewListDeliveryUsageResponseWithDefaults instantiates a new ListDeliveryUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDeliveryUsageResponse) GetData() []DeliveryReport`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDeliveryUsageResponse) GetDataOk() (*[]DeliveryReport, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDeliveryUsageResponse) SetData(v []DeliveryReport)`

SetData sets Data field to given value.

### HasData

`func (o *ListDeliveryUsageResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRowCount

`func (o *ListDeliveryUsageResponse) GetTotalRowCount() int64`

GetTotalRowCount returns the TotalRowCount field if non-nil, zero value otherwise.

### GetTotalRowCountOk

`func (o *ListDeliveryUsageResponse) GetTotalRowCountOk() (*int64, bool)`

GetTotalRowCountOk returns a tuple with the TotalRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRowCount

`func (o *ListDeliveryUsageResponse) SetTotalRowCount(v int64)`

SetTotalRowCount sets TotalRowCount field to given value.

### HasTotalRowCount

`func (o *ListDeliveryUsageResponse) HasTotalRowCount() bool`

HasTotalRowCount returns a boolean if a field has been set.

### GetTimeframe

`func (o *ListDeliveryUsageResponse) GetTimeframe() []int64`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ListDeliveryUsageResponse) GetTimeframeOk() (*[]int64, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ListDeliveryUsageResponse) SetTimeframe(v []int64)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ListDeliveryUsageResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetLimit

`func (o *ListDeliveryUsageResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListDeliveryUsageResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListDeliveryUsageResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListDeliveryUsageResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


