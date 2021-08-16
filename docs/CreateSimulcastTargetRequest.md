# CreateSimulcastTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passthrough** | Pointer to **string** | Arbitrary metadata set by you when creating a simulcast target. | [optional] 
**StreamKey** | Pointer to **string** | Stream Key represents a stream identifier on the third party live streaming service to send the parent live stream to. | [optional] 
**Url** | **string** | RTMP hostname including application name for the third party live streaming service. Example: &#39;rtmp://live.example.com/app&#39;. | 

## Methods

### NewCreateSimulcastTargetRequest

`func NewCreateSimulcastTargetRequest(url string, ) *CreateSimulcastTargetRequest`

NewCreateSimulcastTargetRequest instantiates a new CreateSimulcastTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSimulcastTargetRequestWithDefaults

`func NewCreateSimulcastTargetRequestWithDefaults() *CreateSimulcastTargetRequest`

NewCreateSimulcastTargetRequestWithDefaults instantiates a new CreateSimulcastTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassthrough

`func (o *CreateSimulcastTargetRequest) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *CreateSimulcastTargetRequest) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *CreateSimulcastTargetRequest) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *CreateSimulcastTargetRequest) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetStreamKey

`func (o *CreateSimulcastTargetRequest) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *CreateSimulcastTargetRequest) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *CreateSimulcastTargetRequest) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *CreateSimulcastTargetRequest) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.

### GetUrl

`func (o *CreateSimulcastTargetRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateSimulcastTargetRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateSimulcastTargetRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


