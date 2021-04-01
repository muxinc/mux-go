# SimulcastTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the Simulcast Target | [optional] 
**Passthrough** | Pointer to **string** | Arbitrary Metadata set when creating a simulcast target. | [optional] 
**Status** | Pointer to **string** | The current status of the simulcast target. See Statuses below for detailed description.   * &#x60;idle&#x60;: Default status. When the parent live stream is in disconnected status, simulcast targets will be idle state.   * &#x60;starting&#x60;: The simulcast target transitions into this state when the parent live stream transition into connected state.   * &#x60;broadcasting&#x60;: The simulcast target has successfully connected to the third party live streaming service and is pushing video to that service.   * &#x60;errored&#x60;: The simulcast target encountered an error either while attempting to connect to the third party live streaming service, or mid-broadcasting. Compared to other errored statuses in the Mux Video API, a simulcast may transition back into the broadcasting state if a connection with the service can be re-established.  | [optional] 
**StreamKey** | Pointer to **string** | Stream Key represents an stream identifier for the third party live streaming service to simulcast the parent live stream too. | [optional] 
**Url** | Pointer to **string** | RTMP hostname including the application name for the third party live streaming service. | [optional] 

## Methods

### NewSimulcastTarget

`func NewSimulcastTarget() *SimulcastTarget`

NewSimulcastTarget instantiates a new SimulcastTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulcastTargetWithDefaults

`func NewSimulcastTargetWithDefaults() *SimulcastTarget`

NewSimulcastTargetWithDefaults instantiates a new SimulcastTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimulcastTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulcastTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulcastTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimulcastTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassthrough

`func (o *SimulcastTarget) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *SimulcastTarget) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *SimulcastTarget) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *SimulcastTarget) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetStatus

`func (o *SimulcastTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulcastTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulcastTarget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimulcastTarget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreamKey

`func (o *SimulcastTarget) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *SimulcastTarget) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *SimulcastTarget) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *SimulcastTarget) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.

### GetUrl

`func (o *SimulcastTarget) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimulcastTarget) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimulcastTarget) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SimulcastTarget) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


