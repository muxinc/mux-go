# VideoViewEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewerTime** | Pointer to **int64** |  | [optional] 
**PlaybackTime** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EventTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewVideoViewEvent

`func NewVideoViewEvent() *VideoViewEvent`

NewVideoViewEvent instantiates a new VideoViewEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoViewEventWithDefaults

`func NewVideoViewEventWithDefaults() *VideoViewEvent`

NewVideoViewEventWithDefaults instantiates a new VideoViewEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewerTime

`func (o *VideoViewEvent) GetViewerTime() int64`

GetViewerTime returns the ViewerTime field if non-nil, zero value otherwise.

### GetViewerTimeOk

`func (o *VideoViewEvent) GetViewerTimeOk() (*int64, bool)`

GetViewerTimeOk returns a tuple with the ViewerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerTime

`func (o *VideoViewEvent) SetViewerTime(v int64)`

SetViewerTime sets ViewerTime field to given value.

### HasViewerTime

`func (o *VideoViewEvent) HasViewerTime() bool`

HasViewerTime returns a boolean if a field has been set.

### GetPlaybackTime

`func (o *VideoViewEvent) GetPlaybackTime() int64`

GetPlaybackTime returns the PlaybackTime field if non-nil, zero value otherwise.

### GetPlaybackTimeOk

`func (o *VideoViewEvent) GetPlaybackTimeOk() (*int64, bool)`

GetPlaybackTimeOk returns a tuple with the PlaybackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackTime

`func (o *VideoViewEvent) SetPlaybackTime(v int64)`

SetPlaybackTime sets PlaybackTime field to given value.

### HasPlaybackTime

`func (o *VideoViewEvent) HasPlaybackTime() bool`

HasPlaybackTime returns a boolean if a field has been set.

### GetName

`func (o *VideoViewEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoViewEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoViewEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoViewEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventTime

`func (o *VideoViewEvent) GetEventTime() int64`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *VideoViewEvent) GetEventTimeOk() (*int64, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *VideoViewEvent) SetEventTime(v int64)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *VideoViewEvent) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


