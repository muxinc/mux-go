# InputSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The web address of the file that Mux should download and use. * For subtitles text tracks, the url is the location of subtitle/captions file. Mux supports [SubRip Text (SRT)](https://en.wikipedia.org/wiki/SubRip) and [Web Video Text Tracks](https://www.w3.org/TR/webvtt1/) format for ingesting Subtitles and Closed Captions. * For Watermarking or Overlay, the url is the location of the watermark image. * When creating clips from existing Mux assets, the url is defined with &#x60;mux://assets/{asset_id}&#x60; template where &#x60;asset_id&#x60; is the Asset Identifier for creating the clip from.  | [optional] 
**OverlaySettings** | Pointer to [**InputSettingsOverlaySettings**](InputSettings_overlay_settings.md) |  | [optional] 
**StartTime** | Pointer to **float64** | The time offset in seconds from the beginning of the video indicating the clip&#39;s starting marker. The default value is 0 when not included. This parameter is only applicable for creating clips when &#x60;input.url&#x60; has &#x60;mux://assets/{asset_id}&#x60; format. | [optional] 
**EndTime** | Pointer to **float64** | The time offset in seconds from the beginning of the video, indicating the clip&#39;s ending marker. The default value is the duration of the video when not included. This parameter is only applicable for creating clips when &#x60;input.url&#x60; has &#x60;mux://assets/{asset_id}&#x60; format. | [optional] 
**Type** | Pointer to **string** | This parameter is required for the &#x60;text&#x60; track type. | [optional] 
**TextType** | Pointer to **string** | Type of text track. This parameter only supports subtitles value. For more information on Subtitles / Closed Captions, [see this blog post](https://mux.com/blog/subtitles-captions-webvtt-hls-and-those-magic-flags/). This parameter is required for &#x60;text&#x60; track type. | [optional] 
**LanguageCode** | Pointer to **string** | The language code value must be a valid [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, en for English or en-US for the US version of English. This parameter is required for text type and subtitles text type track. | [optional] 
**Name** | Pointer to **string** | The name of the track containing a human-readable description. This value must be unique across all text type and subtitles &#x60;text&#x60; type tracks. The hls manifest will associate a subtitle text track with this value. For example, the value should be \&quot;English\&quot; for subtitles text track with language_code as en. This optional parameter should be used only for &#x60;text&#x60; type and subtitles &#x60;text&#x60; type track. If this parameter is not included, Mux will auto-populate based on the &#x60;input[].language_code&#x60; value. | [optional] 
**ClosedCaptions** | Pointer to **bool** | Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This optional parameter should be used for &#x60;text&#x60; type and subtitles &#x60;text&#x60; type tracks. | [optional] 
**Passthrough** | Pointer to **string** | This optional parameter should be used for &#x60;text&#x60; type and subtitles &#x60;text&#x60; type tracks. | [optional] 

## Methods

### NewInputSettings

`func NewInputSettings() *InputSettings`

NewInputSettings instantiates a new InputSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSettingsWithDefaults

`func NewInputSettingsWithDefaults() *InputSettings`

NewInputSettingsWithDefaults instantiates a new InputSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InputSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InputSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InputSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InputSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOverlaySettings

`func (o *InputSettings) GetOverlaySettings() InputSettingsOverlaySettings`

GetOverlaySettings returns the OverlaySettings field if non-nil, zero value otherwise.

### GetOverlaySettingsOk

`func (o *InputSettings) GetOverlaySettingsOk() (*InputSettingsOverlaySettings, bool)`

GetOverlaySettingsOk returns a tuple with the OverlaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlaySettings

`func (o *InputSettings) SetOverlaySettings(v InputSettingsOverlaySettings)`

SetOverlaySettings sets OverlaySettings field to given value.

### HasOverlaySettings

`func (o *InputSettings) HasOverlaySettings() bool`

HasOverlaySettings returns a boolean if a field has been set.

### GetStartTime

`func (o *InputSettings) GetStartTime() float64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InputSettings) GetStartTimeOk() (*float64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InputSettings) SetStartTime(v float64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InputSettings) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *InputSettings) GetEndTime() float64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InputSettings) GetEndTimeOk() (*float64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InputSettings) SetEndTime(v float64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InputSettings) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetType

`func (o *InputSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InputSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTextType

`func (o *InputSettings) GetTextType() string`

GetTextType returns the TextType field if non-nil, zero value otherwise.

### GetTextTypeOk

`func (o *InputSettings) GetTextTypeOk() (*string, bool)`

GetTextTypeOk returns a tuple with the TextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextType

`func (o *InputSettings) SetTextType(v string)`

SetTextType sets TextType field to given value.

### HasTextType

`func (o *InputSettings) HasTextType() bool`

HasTextType returns a boolean if a field has been set.

### GetLanguageCode

`func (o *InputSettings) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *InputSettings) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *InputSettings) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *InputSettings) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetName

`func (o *InputSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InputSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InputSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InputSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClosedCaptions

`func (o *InputSettings) GetClosedCaptions() bool`

GetClosedCaptions returns the ClosedCaptions field if non-nil, zero value otherwise.

### GetClosedCaptionsOk

`func (o *InputSettings) GetClosedCaptionsOk() (*bool, bool)`

GetClosedCaptionsOk returns a tuple with the ClosedCaptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedCaptions

`func (o *InputSettings) SetClosedCaptions(v bool)`

SetClosedCaptions sets ClosedCaptions field to given value.

### HasClosedCaptions

`func (o *InputSettings) HasClosedCaptions() bool`

HasClosedCaptions returns a boolean if a field has been set.

### GetPassthrough

`func (o *InputSettings) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *InputSettings) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *InputSettings) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *InputSettings) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


