# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the Asset. Max 255 characters. | [optional] 
**CreatedAt** | Pointer to **string** | Time the Asset was created, defined as a Unix timestamp (seconds since epoch). | [optional] 
**Status** | Pointer to **string** | The status of the asset. | [optional] 
**Duration** | Pointer to **float64** | The duration of the asset in seconds (max duration for a single asset is 24 hours). | [optional] 
**MaxStoredResolution** | Pointer to **string** | The maximum resolution that has been stored for the asset. The asset may be delivered at lower resolutions depending on the device and bandwidth, however it cannot be delivered at a higher value than is stored. | [optional] 
**MaxStoredFrameRate** | Pointer to **float64** | The maximum frame rate that has been stored for the asset. The asset may be delivered at lower frame rates depending on the device and bandwidth, however it cannot be delivered at a higher value than is stored. This field may return -1 if the frame rate of the input cannot be reliably determined.  | [optional] 
**AspectRatio** | Pointer to **string** | The aspect ratio of the asset in the form of &#x60;width:height&#x60;, for example &#x60;16:9&#x60;. | [optional] 
**PlaybackIds** | Pointer to [**[]PlaybackID**](PlaybackID.md) | An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/video/play-your-videos) for more details. | [optional] 
**Tracks** | Pointer to [**[]Track**](Track.md) | The individual media tracks that make up an asset. | [optional] 
**Errors** | Pointer to [**AssetErrors**](Asset_errors.md) |  | [optional] 
**PerTitleEncode** | Pointer to **bool** |  | [optional] 
**IsLive** | Pointer to **bool** | Whether the asset is created from a live stream and the live stream is currently &#x60;active&#x60; and not in &#x60;idle&#x60; state. | [optional] 
**Passthrough** | Pointer to **string** | Arbitrary metadata set for the asset. Max 255 characters. | [optional] 
**LiveStreamId** | Pointer to **string** | Unique identifier for the live stream. This is an optional parameter added when the asset is created from a live stream. | [optional] 
**Master** | Pointer to [**AssetMaster**](Asset_master.md) |  | [optional] 
**MasterAccess** | Pointer to **string** |  | [optional] [default to "none"]
**Mp4Support** | Pointer to **string** |  | [optional] [default to "none"]
**SourceAssetId** | Pointer to **string** | Asset Identifier of the video used as the source for creating the clip. | [optional] 
**NormalizeAudio** | Pointer to **bool** | Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets. | [optional] [default to false]
**StaticRenditions** | Pointer to [**AssetStaticRenditions**](Asset_static_renditions.md) |  | [optional] 
**RecordingTimes** | Pointer to [**[]AssetRecordingTimes**](AssetRecordingTimes.md) | An array of individual live stream recording sessions. A recording session is created on each encoder connection during the live stream | [optional] 
**NonStandardInputReasons** | Pointer to [**AssetNonStandardInputReasons**](Asset_non_standard_input_reasons.md) |  | [optional] 
**Test** | Pointer to **bool** | True means this live stream is a test asset. A test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test assets are watermarked with the Mux logo, limited to 10 seconds, and deleted after 24 hrs. | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Asset) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Asset) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Asset) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Asset) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Asset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Asset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuration

`func (o *Asset) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Asset) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Asset) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Asset) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMaxStoredResolution

`func (o *Asset) GetMaxStoredResolution() string`

GetMaxStoredResolution returns the MaxStoredResolution field if non-nil, zero value otherwise.

### GetMaxStoredResolutionOk

`func (o *Asset) GetMaxStoredResolutionOk() (*string, bool)`

GetMaxStoredResolutionOk returns a tuple with the MaxStoredResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredResolution

`func (o *Asset) SetMaxStoredResolution(v string)`

SetMaxStoredResolution sets MaxStoredResolution field to given value.

### HasMaxStoredResolution

`func (o *Asset) HasMaxStoredResolution() bool`

HasMaxStoredResolution returns a boolean if a field has been set.

### GetMaxStoredFrameRate

`func (o *Asset) GetMaxStoredFrameRate() float64`

GetMaxStoredFrameRate returns the MaxStoredFrameRate field if non-nil, zero value otherwise.

### GetMaxStoredFrameRateOk

`func (o *Asset) GetMaxStoredFrameRateOk() (*float64, bool)`

GetMaxStoredFrameRateOk returns a tuple with the MaxStoredFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredFrameRate

`func (o *Asset) SetMaxStoredFrameRate(v float64)`

SetMaxStoredFrameRate sets MaxStoredFrameRate field to given value.

### HasMaxStoredFrameRate

`func (o *Asset) HasMaxStoredFrameRate() bool`

HasMaxStoredFrameRate returns a boolean if a field has been set.

### GetAspectRatio

`func (o *Asset) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *Asset) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *Asset) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *Asset) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetPlaybackIds

`func (o *Asset) GetPlaybackIds() []PlaybackID`

GetPlaybackIds returns the PlaybackIds field if non-nil, zero value otherwise.

### GetPlaybackIdsOk

`func (o *Asset) GetPlaybackIdsOk() (*[]PlaybackID, bool)`

GetPlaybackIdsOk returns a tuple with the PlaybackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackIds

`func (o *Asset) SetPlaybackIds(v []PlaybackID)`

SetPlaybackIds sets PlaybackIds field to given value.

### HasPlaybackIds

`func (o *Asset) HasPlaybackIds() bool`

HasPlaybackIds returns a boolean if a field has been set.

### GetTracks

`func (o *Asset) GetTracks() []Track`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *Asset) GetTracksOk() (*[]Track, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *Asset) SetTracks(v []Track)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *Asset) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### GetErrors

`func (o *Asset) GetErrors() AssetErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Asset) GetErrorsOk() (*AssetErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Asset) SetErrors(v AssetErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Asset) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetPerTitleEncode

`func (o *Asset) GetPerTitleEncode() bool`

GetPerTitleEncode returns the PerTitleEncode field if non-nil, zero value otherwise.

### GetPerTitleEncodeOk

`func (o *Asset) GetPerTitleEncodeOk() (*bool, bool)`

GetPerTitleEncodeOk returns a tuple with the PerTitleEncode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTitleEncode

`func (o *Asset) SetPerTitleEncode(v bool)`

SetPerTitleEncode sets PerTitleEncode field to given value.

### HasPerTitleEncode

`func (o *Asset) HasPerTitleEncode() bool`

HasPerTitleEncode returns a boolean if a field has been set.

### GetIsLive

`func (o *Asset) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *Asset) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *Asset) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *Asset) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetPassthrough

`func (o *Asset) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *Asset) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *Asset) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *Asset) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *Asset) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *Asset) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *Asset) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *Asset) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetMaster

`func (o *Asset) GetMaster() AssetMaster`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *Asset) GetMasterOk() (*AssetMaster, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *Asset) SetMaster(v AssetMaster)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *Asset) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetMasterAccess

`func (o *Asset) GetMasterAccess() string`

GetMasterAccess returns the MasterAccess field if non-nil, zero value otherwise.

### GetMasterAccessOk

`func (o *Asset) GetMasterAccessOk() (*string, bool)`

GetMasterAccessOk returns a tuple with the MasterAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccess

`func (o *Asset) SetMasterAccess(v string)`

SetMasterAccess sets MasterAccess field to given value.

### HasMasterAccess

`func (o *Asset) HasMasterAccess() bool`

HasMasterAccess returns a boolean if a field has been set.

### GetMp4Support

`func (o *Asset) GetMp4Support() string`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *Asset) GetMp4SupportOk() (*string, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *Asset) SetMp4Support(v string)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *Asset) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.

### GetSourceAssetId

`func (o *Asset) GetSourceAssetId() string`

GetSourceAssetId returns the SourceAssetId field if non-nil, zero value otherwise.

### GetSourceAssetIdOk

`func (o *Asset) GetSourceAssetIdOk() (*string, bool)`

GetSourceAssetIdOk returns a tuple with the SourceAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAssetId

`func (o *Asset) SetSourceAssetId(v string)`

SetSourceAssetId sets SourceAssetId field to given value.

### HasSourceAssetId

`func (o *Asset) HasSourceAssetId() bool`

HasSourceAssetId returns a boolean if a field has been set.

### GetNormalizeAudio

`func (o *Asset) GetNormalizeAudio() bool`

GetNormalizeAudio returns the NormalizeAudio field if non-nil, zero value otherwise.

### GetNormalizeAudioOk

`func (o *Asset) GetNormalizeAudioOk() (*bool, bool)`

GetNormalizeAudioOk returns a tuple with the NormalizeAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeAudio

`func (o *Asset) SetNormalizeAudio(v bool)`

SetNormalizeAudio sets NormalizeAudio field to given value.

### HasNormalizeAudio

`func (o *Asset) HasNormalizeAudio() bool`

HasNormalizeAudio returns a boolean if a field has been set.

### GetStaticRenditions

`func (o *Asset) GetStaticRenditions() AssetStaticRenditions`

GetStaticRenditions returns the StaticRenditions field if non-nil, zero value otherwise.

### GetStaticRenditionsOk

`func (o *Asset) GetStaticRenditionsOk() (*AssetStaticRenditions, bool)`

GetStaticRenditionsOk returns a tuple with the StaticRenditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRenditions

`func (o *Asset) SetStaticRenditions(v AssetStaticRenditions)`

SetStaticRenditions sets StaticRenditions field to given value.

### HasStaticRenditions

`func (o *Asset) HasStaticRenditions() bool`

HasStaticRenditions returns a boolean if a field has been set.

### GetRecordingTimes

`func (o *Asset) GetRecordingTimes() []AssetRecordingTimes`

GetRecordingTimes returns the RecordingTimes field if non-nil, zero value otherwise.

### GetRecordingTimesOk

`func (o *Asset) GetRecordingTimesOk() (*[]AssetRecordingTimes, bool)`

GetRecordingTimesOk returns a tuple with the RecordingTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTimes

`func (o *Asset) SetRecordingTimes(v []AssetRecordingTimes)`

SetRecordingTimes sets RecordingTimes field to given value.

### HasRecordingTimes

`func (o *Asset) HasRecordingTimes() bool`

HasRecordingTimes returns a boolean if a field has been set.

### GetNonStandardInputReasons

`func (o *Asset) GetNonStandardInputReasons() AssetNonStandardInputReasons`

GetNonStandardInputReasons returns the NonStandardInputReasons field if non-nil, zero value otherwise.

### GetNonStandardInputReasonsOk

`func (o *Asset) GetNonStandardInputReasonsOk() (*AssetNonStandardInputReasons, bool)`

GetNonStandardInputReasonsOk returns a tuple with the NonStandardInputReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonStandardInputReasons

`func (o *Asset) SetNonStandardInputReasons(v AssetNonStandardInputReasons)`

SetNonStandardInputReasons sets NonStandardInputReasons field to given value.

### HasNonStandardInputReasons

`func (o *Asset) HasNonStandardInputReasons() bool`

HasNonStandardInputReasons returns a boolean if a field has been set.

### GetTest

`func (o *Asset) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Asset) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Asset) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *Asset) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


