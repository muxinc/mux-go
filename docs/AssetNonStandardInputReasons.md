# AssetNonStandardInputReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoCodec** | Pointer to **string** | The video codec used on the input file. For example, the input file encoded with &#x60;hevc&#x60; video codec is non-standard and the value of this parameter is &#x60;hevc&#x60;. | [optional] 
**AudioCodec** | Pointer to **string** | The audio codec used on the input file. Non-AAC audio codecs are non-standard. | [optional] 
**VideoGopSize** | Pointer to **string** | The video key frame Interval (also called as Group of Picture or GOP) of the input file is &#x60;high&#x60;. This parameter is present when the gop is greater than 10 seconds. | [optional] 
**VideoFrameRate** | Pointer to **string** | The video frame rate of the input file. Video with average frames per second (fps) less than 10 or greater than 120 is non-standard. A &#x60;-1&#x60; frame rate value indicates Mux could not determine the frame rate of the video track. | [optional] 
**VideoResolution** | Pointer to **string** | The video resolution of the input file. Video resolution higher than 2048 pixels on any one dimension (height or width) is considered non-standard, The resolution value is presented as &#x60;width&#x60; x &#x60;height&#x60; in pixels. | [optional] 
**PixelAspectRatio** | Pointer to **string** | The video pixel aspect ratio of the input file. | [optional] 
**VideoEditList** | Pointer to **string** | Video Edit List reason indicates that the input file&#39;s video track contains a complex Edit Decision List. | [optional] 
**AudioEditList** | Pointer to **string** | Audio Edit List reason indicates that the input file&#39;s audio track contains a complex Edit Decision List. | [optional] 
**UnexpectedMediaFileParameters** | Pointer to **string** | A catch-all reason when the input file in created with non-standard encoding parameters. | [optional] 

## Methods

### NewAssetNonStandardInputReasons

`func NewAssetNonStandardInputReasons() *AssetNonStandardInputReasons`

NewAssetNonStandardInputReasons instantiates a new AssetNonStandardInputReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetNonStandardInputReasonsWithDefaults

`func NewAssetNonStandardInputReasonsWithDefaults() *AssetNonStandardInputReasons`

NewAssetNonStandardInputReasonsWithDefaults instantiates a new AssetNonStandardInputReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoCodec

`func (o *AssetNonStandardInputReasons) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *AssetNonStandardInputReasons) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *AssetNonStandardInputReasons) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *AssetNonStandardInputReasons) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *AssetNonStandardInputReasons) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *AssetNonStandardInputReasons) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *AssetNonStandardInputReasons) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *AssetNonStandardInputReasons) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoGopSize

`func (o *AssetNonStandardInputReasons) GetVideoGopSize() string`

GetVideoGopSize returns the VideoGopSize field if non-nil, zero value otherwise.

### GetVideoGopSizeOk

`func (o *AssetNonStandardInputReasons) GetVideoGopSizeOk() (*string, bool)`

GetVideoGopSizeOk returns a tuple with the VideoGopSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoGopSize

`func (o *AssetNonStandardInputReasons) SetVideoGopSize(v string)`

SetVideoGopSize sets VideoGopSize field to given value.

### HasVideoGopSize

`func (o *AssetNonStandardInputReasons) HasVideoGopSize() bool`

HasVideoGopSize returns a boolean if a field has been set.

### GetVideoFrameRate

`func (o *AssetNonStandardInputReasons) GetVideoFrameRate() string`

GetVideoFrameRate returns the VideoFrameRate field if non-nil, zero value otherwise.

### GetVideoFrameRateOk

`func (o *AssetNonStandardInputReasons) GetVideoFrameRateOk() (*string, bool)`

GetVideoFrameRateOk returns a tuple with the VideoFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFrameRate

`func (o *AssetNonStandardInputReasons) SetVideoFrameRate(v string)`

SetVideoFrameRate sets VideoFrameRate field to given value.

### HasVideoFrameRate

`func (o *AssetNonStandardInputReasons) HasVideoFrameRate() bool`

HasVideoFrameRate returns a boolean if a field has been set.

### GetVideoResolution

`func (o *AssetNonStandardInputReasons) GetVideoResolution() string`

GetVideoResolution returns the VideoResolution field if non-nil, zero value otherwise.

### GetVideoResolutionOk

`func (o *AssetNonStandardInputReasons) GetVideoResolutionOk() (*string, bool)`

GetVideoResolutionOk returns a tuple with the VideoResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoResolution

`func (o *AssetNonStandardInputReasons) SetVideoResolution(v string)`

SetVideoResolution sets VideoResolution field to given value.

### HasVideoResolution

`func (o *AssetNonStandardInputReasons) HasVideoResolution() bool`

HasVideoResolution returns a boolean if a field has been set.

### GetPixelAspectRatio

`func (o *AssetNonStandardInputReasons) GetPixelAspectRatio() string`

GetPixelAspectRatio returns the PixelAspectRatio field if non-nil, zero value otherwise.

### GetPixelAspectRatioOk

`func (o *AssetNonStandardInputReasons) GetPixelAspectRatioOk() (*string, bool)`

GetPixelAspectRatioOk returns a tuple with the PixelAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelAspectRatio

`func (o *AssetNonStandardInputReasons) SetPixelAspectRatio(v string)`

SetPixelAspectRatio sets PixelAspectRatio field to given value.

### HasPixelAspectRatio

`func (o *AssetNonStandardInputReasons) HasPixelAspectRatio() bool`

HasPixelAspectRatio returns a boolean if a field has been set.

### GetVideoEditList

`func (o *AssetNonStandardInputReasons) GetVideoEditList() string`

GetVideoEditList returns the VideoEditList field if non-nil, zero value otherwise.

### GetVideoEditListOk

`func (o *AssetNonStandardInputReasons) GetVideoEditListOk() (*string, bool)`

GetVideoEditListOk returns a tuple with the VideoEditList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEditList

`func (o *AssetNonStandardInputReasons) SetVideoEditList(v string)`

SetVideoEditList sets VideoEditList field to given value.

### HasVideoEditList

`func (o *AssetNonStandardInputReasons) HasVideoEditList() bool`

HasVideoEditList returns a boolean if a field has been set.

### GetAudioEditList

`func (o *AssetNonStandardInputReasons) GetAudioEditList() string`

GetAudioEditList returns the AudioEditList field if non-nil, zero value otherwise.

### GetAudioEditListOk

`func (o *AssetNonStandardInputReasons) GetAudioEditListOk() (*string, bool)`

GetAudioEditListOk returns a tuple with the AudioEditList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioEditList

`func (o *AssetNonStandardInputReasons) SetAudioEditList(v string)`

SetAudioEditList sets AudioEditList field to given value.

### HasAudioEditList

`func (o *AssetNonStandardInputReasons) HasAudioEditList() bool`

HasAudioEditList returns a boolean if a field has been set.

### GetUnexpectedMediaFileParameters

`func (o *AssetNonStandardInputReasons) GetUnexpectedMediaFileParameters() string`

GetUnexpectedMediaFileParameters returns the UnexpectedMediaFileParameters field if non-nil, zero value otherwise.

### GetUnexpectedMediaFileParametersOk

`func (o *AssetNonStandardInputReasons) GetUnexpectedMediaFileParametersOk() (*string, bool)`

GetUnexpectedMediaFileParametersOk returns a tuple with the UnexpectedMediaFileParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpectedMediaFileParameters

`func (o *AssetNonStandardInputReasons) SetUnexpectedMediaFileParameters(v string)`

SetUnexpectedMediaFileParameters sets UnexpectedMediaFileParameters field to given value.

### HasUnexpectedMediaFileParameters

`func (o *AssetNonStandardInputReasons) HasUnexpectedMediaFileParameters() bool`

HasUnexpectedMediaFileParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


