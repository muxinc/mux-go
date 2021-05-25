# VideoViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**VideoView**](VideoView.md) |  | [optional] 
**Timeframe** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewVideoViewResponse

`func NewVideoViewResponse() *VideoViewResponse`

NewVideoViewResponse instantiates a new VideoViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoViewResponseWithDefaults

`func NewVideoViewResponseWithDefaults() *VideoViewResponse`

NewVideoViewResponseWithDefaults instantiates a new VideoViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VideoViewResponse) GetData() VideoView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VideoViewResponse) GetDataOk() (*VideoView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VideoViewResponse) SetData(v VideoView)`

SetData sets Data field to given value.

### HasData

`func (o *VideoViewResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimeframe

`func (o *VideoViewResponse) GetTimeframe() []int64`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *VideoViewResponse) GetTimeframeOk() (*[]int64, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *VideoViewResponse) SetTimeframe(v []int64)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *VideoViewResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


