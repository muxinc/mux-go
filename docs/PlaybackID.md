# PlaybackID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the PlaybackID | [optional] 
**Policy** | Pointer to [**PlaybackPolicy**](PlaybackPolicy.md) |  | [optional] 

## Methods

### NewPlaybackID

`func NewPlaybackID() *PlaybackID`

NewPlaybackID instantiates a new PlaybackID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybackIDWithDefaults

`func NewPlaybackIDWithDefaults() *PlaybackID`

NewPlaybackIDWithDefaults instantiates a new PlaybackID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaybackID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaybackID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaybackID) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaybackID) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicy

`func (o *PlaybackID) GetPolicy() PlaybackPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PlaybackID) GetPolicyOk() (*PlaybackPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PlaybackID) SetPolicy(v PlaybackPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PlaybackID) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


