# CreateLiveStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaybackPolicy** | Pointer to [**[]PlaybackPolicy**](PlaybackPolicy.md) |  | [optional] 
**NewAssetSettings** | Pointer to [**CreateAssetRequest**](CreateAssetRequest.md) |  | [optional] 
**ReconnectWindow** | Pointer to **float32** | When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. Defaults to 60 seconds on the API if not specified. | [optional] 
**Passthrough** | Pointer to **string** |  | [optional] 
**AudioOnly** | Pointer to **bool** | Force the live stream to only process the audio track when the value is set to true. Mux drops the video track if broadcasted. | [optional] 
**ReducedLatency** | Pointer to **bool** | Latency is the time from when the streamer does something in real life to when you see it happen in the player. Set this if you want lower latency for your live stream. Note: Reconnect windows are incompatible with Reduced Latency and will always be set to zero (0) seconds. Read more here: https://mux.com/blog/reduced-latency-for-mux-live-streaming-now-available/ | [optional] 
**Test** | Pointer to **bool** | Marks the live stream as a test live stream when the value is set to true. A test live stream can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test live streams created. Test live streams are watermarked with the Mux logo and limited to 5 minutes. The test live stream is disabled after the stream is active for 5 mins and the recorded asset also deleted after 24 hours. | [optional] 
**SimulcastTargets** | Pointer to [**[]CreateSimulcastTargetRequest**](CreateSimulcastTargetRequest.md) |  | [optional] 

## Methods

### NewCreateLiveStreamRequest

`func NewCreateLiveStreamRequest() *CreateLiveStreamRequest`

NewCreateLiveStreamRequest instantiates a new CreateLiveStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLiveStreamRequestWithDefaults

`func NewCreateLiveStreamRequestWithDefaults() *CreateLiveStreamRequest`

NewCreateLiveStreamRequestWithDefaults instantiates a new CreateLiveStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaybackPolicy

`func (o *CreateLiveStreamRequest) GetPlaybackPolicy() []PlaybackPolicy`

GetPlaybackPolicy returns the PlaybackPolicy field if non-nil, zero value otherwise.

### GetPlaybackPolicyOk

`func (o *CreateLiveStreamRequest) GetPlaybackPolicyOk() (*[]PlaybackPolicy, bool)`

GetPlaybackPolicyOk returns a tuple with the PlaybackPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPolicy

`func (o *CreateLiveStreamRequest) SetPlaybackPolicy(v []PlaybackPolicy)`

SetPlaybackPolicy sets PlaybackPolicy field to given value.

### HasPlaybackPolicy

`func (o *CreateLiveStreamRequest) HasPlaybackPolicy() bool`

HasPlaybackPolicy returns a boolean if a field has been set.

### GetNewAssetSettings

`func (o *CreateLiveStreamRequest) GetNewAssetSettings() CreateAssetRequest`

GetNewAssetSettings returns the NewAssetSettings field if non-nil, zero value otherwise.

### GetNewAssetSettingsOk

`func (o *CreateLiveStreamRequest) GetNewAssetSettingsOk() (*CreateAssetRequest, bool)`

GetNewAssetSettingsOk returns a tuple with the NewAssetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssetSettings

`func (o *CreateLiveStreamRequest) SetNewAssetSettings(v CreateAssetRequest)`

SetNewAssetSettings sets NewAssetSettings field to given value.

### HasNewAssetSettings

`func (o *CreateLiveStreamRequest) HasNewAssetSettings() bool`

HasNewAssetSettings returns a boolean if a field has been set.

### GetReconnectWindow

`func (o *CreateLiveStreamRequest) GetReconnectWindow() float32`

GetReconnectWindow returns the ReconnectWindow field if non-nil, zero value otherwise.

### GetReconnectWindowOk

`func (o *CreateLiveStreamRequest) GetReconnectWindowOk() (*float32, bool)`

GetReconnectWindowOk returns a tuple with the ReconnectWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconnectWindow

`func (o *CreateLiveStreamRequest) SetReconnectWindow(v float32)`

SetReconnectWindow sets ReconnectWindow field to given value.

### HasReconnectWindow

`func (o *CreateLiveStreamRequest) HasReconnectWindow() bool`

HasReconnectWindow returns a boolean if a field has been set.

### GetPassthrough

`func (o *CreateLiveStreamRequest) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *CreateLiveStreamRequest) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *CreateLiveStreamRequest) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *CreateLiveStreamRequest) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetAudioOnly

`func (o *CreateLiveStreamRequest) GetAudioOnly() bool`

GetAudioOnly returns the AudioOnly field if non-nil, zero value otherwise.

### GetAudioOnlyOk

`func (o *CreateLiveStreamRequest) GetAudioOnlyOk() (*bool, bool)`

GetAudioOnlyOk returns a tuple with the AudioOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioOnly

`func (o *CreateLiveStreamRequest) SetAudioOnly(v bool)`

SetAudioOnly sets AudioOnly field to given value.

### HasAudioOnly

`func (o *CreateLiveStreamRequest) HasAudioOnly() bool`

HasAudioOnly returns a boolean if a field has been set.

### GetReducedLatency

`func (o *CreateLiveStreamRequest) GetReducedLatency() bool`

GetReducedLatency returns the ReducedLatency field if non-nil, zero value otherwise.

### GetReducedLatencyOk

`func (o *CreateLiveStreamRequest) GetReducedLatencyOk() (*bool, bool)`

GetReducedLatencyOk returns a tuple with the ReducedLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedLatency

`func (o *CreateLiveStreamRequest) SetReducedLatency(v bool)`

SetReducedLatency sets ReducedLatency field to given value.

### HasReducedLatency

`func (o *CreateLiveStreamRequest) HasReducedLatency() bool`

HasReducedLatency returns a boolean if a field has been set.

### GetTest

`func (o *CreateLiveStreamRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreateLiveStreamRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreateLiveStreamRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreateLiveStreamRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetSimulcastTargets

`func (o *CreateLiveStreamRequest) GetSimulcastTargets() []CreateSimulcastTargetRequest`

GetSimulcastTargets returns the SimulcastTargets field if non-nil, zero value otherwise.

### GetSimulcastTargetsOk

`func (o *CreateLiveStreamRequest) GetSimulcastTargetsOk() (*[]CreateSimulcastTargetRequest, bool)`

GetSimulcastTargetsOk returns a tuple with the SimulcastTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulcastTargets

`func (o *CreateLiveStreamRequest) SetSimulcastTargets(v []CreateSimulcastTargetRequest)`

SetSimulcastTargets sets SimulcastTargets field to given value.

### HasSimulcastTargets

`func (o *CreateLiveStreamRequest) HasSimulcastTargets() bool`

HasSimulcastTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


