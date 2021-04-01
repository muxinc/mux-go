# LiveStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the Live Stream. Max 255 characters. | [optional] 
**CreatedAt** | Pointer to **string** | Time the Live Stream was created, defined as a Unix timestamp (seconds since epoch). | [optional] 
**StreamKey** | Pointer to **string** | Unique key used for streaming to a Mux RTMP endpoint. This should be considered as sensitive as credentials, anyone with this stream key can begin streaming. | [optional] 
**ActiveAssetId** | Pointer to **string** | The Asset that is currently being created if there is an active broadcast. | [optional] 
**RecentAssetIds** | Pointer to **[]string** | An array of strings with the most recent Assets that were created from this live stream. | [optional] 
**Status** | Pointer to **string** | &#x60;idle&#x60; indicates that there is no active broadcast. &#x60;active&#x60; indicates that there is an active broadcast and &#x60;disabled&#x60; status indicates that no future RTMP streams can be published. | [optional] 
**PlaybackIds** | Pointer to [**[]PlaybackID**](PlaybackID.md) | An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/video/play-your-videos) for more details. | [optional] 
**NewAssetSettings** | Pointer to [**CreateAssetRequest**](CreateAssetRequest.md) |  | [optional] 
**Passthrough** | Pointer to **string** | Arbitrary metadata set for the asset. Max 255 characters. | [optional] 
**ReconnectWindow** | Pointer to **float32** | When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. **Min**: 0.1s. **Max**: 300s (5 minutes). | [optional] [default to 60]
**ReducedLatency** | Pointer to **bool** | Latency is the time from when the streamer does something in real life to when you see it happen in the player. Set this if you want lower latency for your live stream. **Note**: Reconnect windows are incompatible with Reduced Latency and will always be set to zero (0) seconds. See the [Reduce live stream latency guide](https://docs.mux.com/guides/video/reduce-live-stream-latency) to understand the tradeoffs. | [optional] 
**SimulcastTargets** | Pointer to [**[]SimulcastTarget**](SimulcastTarget.md) | Each Simulcast Target contains configuration details to broadcast (or \&quot;restream\&quot;) a live stream to a third-party streaming service. [See the Stream live to 3rd party platforms guide](https://docs.mux.com/guides/video/stream-live-to-3rd-party-platforms). | [optional] 
**Test** | Pointer to **bool** | True means this live stream is a test live stream. Test live streams can be used to help evaluate the Mux Video APIs for free. There is no limit on the number of test live streams, but they are watermarked with the Mux logo, limited to 5 minutes, and deleted after 24 hours. | [optional] 

## Methods

### NewLiveStream

`func NewLiveStream() *LiveStream`

NewLiveStream instantiates a new LiveStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamWithDefaults

`func NewLiveStreamWithDefaults() *LiveStream`

NewLiveStreamWithDefaults instantiates a new LiveStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveStream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveStream) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveStream) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveStream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LiveStream) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LiveStream) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LiveStream) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LiveStream) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStreamKey

`func (o *LiveStream) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *LiveStream) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *LiveStream) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *LiveStream) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.

### GetActiveAssetId

`func (o *LiveStream) GetActiveAssetId() string`

GetActiveAssetId returns the ActiveAssetId field if non-nil, zero value otherwise.

### GetActiveAssetIdOk

`func (o *LiveStream) GetActiveAssetIdOk() (*string, bool)`

GetActiveAssetIdOk returns a tuple with the ActiveAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAssetId

`func (o *LiveStream) SetActiveAssetId(v string)`

SetActiveAssetId sets ActiveAssetId field to given value.

### HasActiveAssetId

`func (o *LiveStream) HasActiveAssetId() bool`

HasActiveAssetId returns a boolean if a field has been set.

### GetRecentAssetIds

`func (o *LiveStream) GetRecentAssetIds() []string`

GetRecentAssetIds returns the RecentAssetIds field if non-nil, zero value otherwise.

### GetRecentAssetIdsOk

`func (o *LiveStream) GetRecentAssetIdsOk() (*[]string, bool)`

GetRecentAssetIdsOk returns a tuple with the RecentAssetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentAssetIds

`func (o *LiveStream) SetRecentAssetIds(v []string)`

SetRecentAssetIds sets RecentAssetIds field to given value.

### HasRecentAssetIds

`func (o *LiveStream) HasRecentAssetIds() bool`

HasRecentAssetIds returns a boolean if a field has been set.

### GetStatus

`func (o *LiveStream) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveStream) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveStream) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveStream) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPlaybackIds

`func (o *LiveStream) GetPlaybackIds() []PlaybackID`

GetPlaybackIds returns the PlaybackIds field if non-nil, zero value otherwise.

### GetPlaybackIdsOk

`func (o *LiveStream) GetPlaybackIdsOk() (*[]PlaybackID, bool)`

GetPlaybackIdsOk returns a tuple with the PlaybackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackIds

`func (o *LiveStream) SetPlaybackIds(v []PlaybackID)`

SetPlaybackIds sets PlaybackIds field to given value.

### HasPlaybackIds

`func (o *LiveStream) HasPlaybackIds() bool`

HasPlaybackIds returns a boolean if a field has been set.

### GetNewAssetSettings

`func (o *LiveStream) GetNewAssetSettings() CreateAssetRequest`

GetNewAssetSettings returns the NewAssetSettings field if non-nil, zero value otherwise.

### GetNewAssetSettingsOk

`func (o *LiveStream) GetNewAssetSettingsOk() (*CreateAssetRequest, bool)`

GetNewAssetSettingsOk returns a tuple with the NewAssetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssetSettings

`func (o *LiveStream) SetNewAssetSettings(v CreateAssetRequest)`

SetNewAssetSettings sets NewAssetSettings field to given value.

### HasNewAssetSettings

`func (o *LiveStream) HasNewAssetSettings() bool`

HasNewAssetSettings returns a boolean if a field has been set.

### GetPassthrough

`func (o *LiveStream) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *LiveStream) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *LiveStream) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *LiveStream) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetReconnectWindow

`func (o *LiveStream) GetReconnectWindow() float32`

GetReconnectWindow returns the ReconnectWindow field if non-nil, zero value otherwise.

### GetReconnectWindowOk

`func (o *LiveStream) GetReconnectWindowOk() (*float32, bool)`

GetReconnectWindowOk returns a tuple with the ReconnectWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconnectWindow

`func (o *LiveStream) SetReconnectWindow(v float32)`

SetReconnectWindow sets ReconnectWindow field to given value.

### HasReconnectWindow

`func (o *LiveStream) HasReconnectWindow() bool`

HasReconnectWindow returns a boolean if a field has been set.

### GetReducedLatency

`func (o *LiveStream) GetReducedLatency() bool`

GetReducedLatency returns the ReducedLatency field if non-nil, zero value otherwise.

### GetReducedLatencyOk

`func (o *LiveStream) GetReducedLatencyOk() (*bool, bool)`

GetReducedLatencyOk returns a tuple with the ReducedLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedLatency

`func (o *LiveStream) SetReducedLatency(v bool)`

SetReducedLatency sets ReducedLatency field to given value.

### HasReducedLatency

`func (o *LiveStream) HasReducedLatency() bool`

HasReducedLatency returns a boolean if a field has been set.

### GetSimulcastTargets

`func (o *LiveStream) GetSimulcastTargets() []SimulcastTarget`

GetSimulcastTargets returns the SimulcastTargets field if non-nil, zero value otherwise.

### GetSimulcastTargetsOk

`func (o *LiveStream) GetSimulcastTargetsOk() (*[]SimulcastTarget, bool)`

GetSimulcastTargetsOk returns a tuple with the SimulcastTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulcastTargets

`func (o *LiveStream) SetSimulcastTargets(v []SimulcastTarget)`

SetSimulcastTargets sets SimulcastTargets field to given value.

### HasSimulcastTargets

`func (o *LiveStream) HasSimulcastTargets() bool`

HasSimulcastTargets returns a boolean if a field has been set.

### GetTest

`func (o *LiveStream) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *LiveStream) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *LiveStream) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *LiveStream) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


