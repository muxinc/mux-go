# CreateAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to [**[]InputSettings**](InputSettings.md) | An array of objects that each describe an input file to be used to create the asset. As a shortcut, input can also be a string URL for a file when only one input file is used. See &#x60;input[].url&#x60; for requirements. | [optional] 
**PlaybackPolicy** | Pointer to [**[]PlaybackPolicy**](PlaybackPolicy.md) | An array of playback policy names that you want applied to this asset and available through &#x60;playback_ids&#x60;. Options include: &#x60;\&quot;public\&quot;&#x60; (anyone with the playback URL can stream the asset). And &#x60;\&quot;signed\&quot;&#x60; (an additional access token is required to play the asset). If no playback_policy is set, the asset will have no playback IDs and will therefore not be playable. For simplicity, a single string name can be used in place of the array in the case of only one playback policy. | [optional] 
**PerTitleEncode** | Pointer to **bool** |  | [optional] 
**Passthrough** | Pointer to **string** | Arbitrary metadata that will be included in the asset details and related webhooks. Can be used to store your own ID for a video along with the asset. **Max: 255 characters**. | [optional] 
**Mp4Support** | Pointer to **string** | Specify what level (if any) of support for mp4 playback. In most cases you should use our default HLS-based streaming playback ({playback_id}.m3u8) which can automatically adjust to viewers&#39; connection speeds, but an mp4 can be useful for some legacy devices or downloading for offline playback. See the [Download your videos guide](/guides/video/download-your-videos) for more information. | [optional] 
**NormalizeAudio** | Pointer to **bool** | Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets. | [optional] [default to false]
**MasterAccess** | Pointer to **string** | Specify what level (if any) of support for master access. Master access can be enabled temporarily for your asset to be downloaded. See the [Download your videos guide](/guides/video/download-your-videos) for more information. | [optional] 
**Test** | Pointer to **bool** | Marks the asset as a test asset when the value is set to true. A Test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test asset are watermarked with the Mux logo, limited to 10 seconds, deleted after 24 hrs. | [optional] 

## Methods

### NewCreateAssetRequest

`func NewCreateAssetRequest() *CreateAssetRequest`

NewCreateAssetRequest instantiates a new CreateAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetRequestWithDefaults

`func NewCreateAssetRequestWithDefaults() *CreateAssetRequest`

NewCreateAssetRequestWithDefaults instantiates a new CreateAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *CreateAssetRequest) GetInput() []InputSettings`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CreateAssetRequest) GetInputOk() (*[]InputSettings, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CreateAssetRequest) SetInput(v []InputSettings)`

SetInput sets Input field to given value.

### HasInput

`func (o *CreateAssetRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetPlaybackPolicy

`func (o *CreateAssetRequest) GetPlaybackPolicy() []PlaybackPolicy`

GetPlaybackPolicy returns the PlaybackPolicy field if non-nil, zero value otherwise.

### GetPlaybackPolicyOk

`func (o *CreateAssetRequest) GetPlaybackPolicyOk() (*[]PlaybackPolicy, bool)`

GetPlaybackPolicyOk returns a tuple with the PlaybackPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPolicy

`func (o *CreateAssetRequest) SetPlaybackPolicy(v []PlaybackPolicy)`

SetPlaybackPolicy sets PlaybackPolicy field to given value.

### HasPlaybackPolicy

`func (o *CreateAssetRequest) HasPlaybackPolicy() bool`

HasPlaybackPolicy returns a boolean if a field has been set.

### GetPerTitleEncode

`func (o *CreateAssetRequest) GetPerTitleEncode() bool`

GetPerTitleEncode returns the PerTitleEncode field if non-nil, zero value otherwise.

### GetPerTitleEncodeOk

`func (o *CreateAssetRequest) GetPerTitleEncodeOk() (*bool, bool)`

GetPerTitleEncodeOk returns a tuple with the PerTitleEncode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerTitleEncode

`func (o *CreateAssetRequest) SetPerTitleEncode(v bool)`

SetPerTitleEncode sets PerTitleEncode field to given value.

### HasPerTitleEncode

`func (o *CreateAssetRequest) HasPerTitleEncode() bool`

HasPerTitleEncode returns a boolean if a field has been set.

### GetPassthrough

`func (o *CreateAssetRequest) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *CreateAssetRequest) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *CreateAssetRequest) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *CreateAssetRequest) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetMp4Support

`func (o *CreateAssetRequest) GetMp4Support() string`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *CreateAssetRequest) GetMp4SupportOk() (*string, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *CreateAssetRequest) SetMp4Support(v string)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *CreateAssetRequest) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.

### GetNormalizeAudio

`func (o *CreateAssetRequest) GetNormalizeAudio() bool`

GetNormalizeAudio returns the NormalizeAudio field if non-nil, zero value otherwise.

### GetNormalizeAudioOk

`func (o *CreateAssetRequest) GetNormalizeAudioOk() (*bool, bool)`

GetNormalizeAudioOk returns a tuple with the NormalizeAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeAudio

`func (o *CreateAssetRequest) SetNormalizeAudio(v bool)`

SetNormalizeAudio sets NormalizeAudio field to given value.

### HasNormalizeAudio

`func (o *CreateAssetRequest) HasNormalizeAudio() bool`

HasNormalizeAudio returns a boolean if a field has been set.

### GetMasterAccess

`func (o *CreateAssetRequest) GetMasterAccess() string`

GetMasterAccess returns the MasterAccess field if non-nil, zero value otherwise.

### GetMasterAccessOk

`func (o *CreateAssetRequest) GetMasterAccessOk() (*string, bool)`

GetMasterAccessOk returns a tuple with the MasterAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccess

`func (o *CreateAssetRequest) SetMasterAccess(v string)`

SetMasterAccess sets MasterAccess field to given value.

### HasMasterAccess

`func (o *CreateAssetRequest) HasMasterAccess() bool`

HasMasterAccess returns a boolean if a field has been set.

### GetTest

`func (o *CreateAssetRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreateAssetRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreateAssetRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreateAssetRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


