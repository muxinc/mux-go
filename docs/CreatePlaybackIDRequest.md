# CreatePlaybackIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**PlaybackPolicy**](PlaybackPolicy.md) |  | [optional] 

## Methods

### NewCreatePlaybackIDRequest

`func NewCreatePlaybackIDRequest() *CreatePlaybackIDRequest`

NewCreatePlaybackIDRequest instantiates a new CreatePlaybackIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlaybackIDRequestWithDefaults

`func NewCreatePlaybackIDRequestWithDefaults() *CreatePlaybackIDRequest`

NewCreatePlaybackIDRequestWithDefaults instantiates a new CreatePlaybackIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *CreatePlaybackIDRequest) GetPolicy() PlaybackPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CreatePlaybackIDRequest) GetPolicyOk() (*PlaybackPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CreatePlaybackIDRequest) SetPolicy(v PlaybackPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CreatePlaybackIDRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


