# IncidentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Incident**](Incident.md) |  | [optional] 
**Timeframe** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewIncidentResponse

`func NewIncidentResponse() *IncidentResponse`

NewIncidentResponse instantiates a new IncidentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentResponseWithDefaults

`func NewIncidentResponseWithDefaults() *IncidentResponse`

NewIncidentResponseWithDefaults instantiates a new IncidentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentResponse) GetData() Incident`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentResponse) GetDataOk() (*Incident, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentResponse) SetData(v Incident)`

SetData sets Data field to given value.

### HasData

`func (o *IncidentResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTimeframe

`func (o *IncidentResponse) GetTimeframe() []int64`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *IncidentResponse) GetTimeframeOk() (*[]int64, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *IncidentResponse) SetTimeframe(v []int64)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *IncidentResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


