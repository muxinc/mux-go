# AssetRecordingTimes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **time.Time** | The time at which the recording for the live stream started. The time value is Unix epoch time represented in ISO 8601 format. | [optional] 
**Duration** | Pointer to **float64** | The duration of the live stream recorded. The time value is in seconds. | [optional] 

## Methods

### NewAssetRecordingTimes

`func NewAssetRecordingTimes() *AssetRecordingTimes`

NewAssetRecordingTimes instantiates a new AssetRecordingTimes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetRecordingTimesWithDefaults

`func NewAssetRecordingTimesWithDefaults() *AssetRecordingTimes`

NewAssetRecordingTimesWithDefaults instantiates a new AssetRecordingTimes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *AssetRecordingTimes) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AssetRecordingTimes) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AssetRecordingTimes) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AssetRecordingTimes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetDuration

`func (o *AssetRecordingTimes) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AssetRecordingTimes) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AssetRecordingTimes) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AssetRecordingTimes) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


