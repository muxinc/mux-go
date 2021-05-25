# Track

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the Track | [optional] 
**Type** | Pointer to **string** | The type of track | [optional] 
**Duration** | Pointer to **float64** | The duration in seconds of the track media. This parameter is not set for the &#x60;text&#x60; type track. This field is optional and may not be set. The top level &#x60;duration&#x60; field of an asset will always be set. | [optional] 
**MaxWidth** | Pointer to **int64** | The maximum width in pixels available for the track. Only set for the &#x60;video&#x60; type track. | [optional] 
**MaxHeight** | Pointer to **int64** | The maximum height in pixels available for the track. Only set for the &#x60;video&#x60; type track. | [optional] 
**MaxFrameRate** | Pointer to **float64** | The maximum frame rate available for the track. Only set for the &#x60;video&#x60; type track. This field may return &#x60;-1&#x60; if the frame rate of the input cannot be reliably determined. | [optional] 
**MaxChannels** | Pointer to **int64** | The maximum number of audio channels the track supports. Only set for the &#x60;audio&#x60; type track. | [optional] 
**MaxChannelLayout** | Pointer to **string** | Only set for the &#x60;audio&#x60; type track. | [optional] 
**TextType** | Pointer to **string** | This parameter is set only for the &#x60;text&#x60; type track. | [optional] 
**LanguageCode** | Pointer to **string** | The language code value represents [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, &#x60;en&#x60; for English or &#x60;en-US&#x60; for the US version of English. This parameter is set for &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. | [optional] 
**Name** | Pointer to **string** | The name of the track containing a human-readable description. The hls manifest will associate a subtitle text track with this value. For example, the value is \&quot;English\&quot; for subtitles text track for the &#x60;language_code&#x60; value of &#x60;en-US&#x60;. This parameter is set for the &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. | [optional] 
**ClosedCaptions** | Pointer to **bool** | Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This parameter is set for the &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. | [optional] 
**Passthrough** | Pointer to **string** | Arbitrary metadata set for the track either when creating the asset or track. This parameter is set for &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. Max 255 characters. | [optional] 

## Methods

### NewTrack

`func NewTrack() *Track`

NewTrack instantiates a new Track object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackWithDefaults

`func NewTrackWithDefaults() *Track`

NewTrackWithDefaults instantiates a new Track object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Track) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Track) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Track) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Track) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Track) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Track) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Track) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Track) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDuration

`func (o *Track) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Track) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Track) SetDuration(v float64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Track) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMaxWidth

`func (o *Track) GetMaxWidth() int64`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *Track) GetMaxWidthOk() (*int64, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *Track) SetMaxWidth(v int64)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *Track) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### GetMaxHeight

`func (o *Track) GetMaxHeight() int64`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *Track) GetMaxHeightOk() (*int64, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *Track) SetMaxHeight(v int64)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *Track) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### GetMaxFrameRate

`func (o *Track) GetMaxFrameRate() float64`

GetMaxFrameRate returns the MaxFrameRate field if non-nil, zero value otherwise.

### GetMaxFrameRateOk

`func (o *Track) GetMaxFrameRateOk() (*float64, bool)`

GetMaxFrameRateOk returns a tuple with the MaxFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFrameRate

`func (o *Track) SetMaxFrameRate(v float64)`

SetMaxFrameRate sets MaxFrameRate field to given value.

### HasMaxFrameRate

`func (o *Track) HasMaxFrameRate() bool`

HasMaxFrameRate returns a boolean if a field has been set.

### GetMaxChannels

`func (o *Track) GetMaxChannels() int64`

GetMaxChannels returns the MaxChannels field if non-nil, zero value otherwise.

### GetMaxChannelsOk

`func (o *Track) GetMaxChannelsOk() (*int64, bool)`

GetMaxChannelsOk returns a tuple with the MaxChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChannels

`func (o *Track) SetMaxChannels(v int64)`

SetMaxChannels sets MaxChannels field to given value.

### HasMaxChannels

`func (o *Track) HasMaxChannels() bool`

HasMaxChannels returns a boolean if a field has been set.

### GetMaxChannelLayout

`func (o *Track) GetMaxChannelLayout() string`

GetMaxChannelLayout returns the MaxChannelLayout field if non-nil, zero value otherwise.

### GetMaxChannelLayoutOk

`func (o *Track) GetMaxChannelLayoutOk() (*string, bool)`

GetMaxChannelLayoutOk returns a tuple with the MaxChannelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChannelLayout

`func (o *Track) SetMaxChannelLayout(v string)`

SetMaxChannelLayout sets MaxChannelLayout field to given value.

### HasMaxChannelLayout

`func (o *Track) HasMaxChannelLayout() bool`

HasMaxChannelLayout returns a boolean if a field has been set.

### GetTextType

`func (o *Track) GetTextType() string`

GetTextType returns the TextType field if non-nil, zero value otherwise.

### GetTextTypeOk

`func (o *Track) GetTextTypeOk() (*string, bool)`

GetTextTypeOk returns a tuple with the TextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextType

`func (o *Track) SetTextType(v string)`

SetTextType sets TextType field to given value.

### HasTextType

`func (o *Track) HasTextType() bool`

HasTextType returns a boolean if a field has been set.

### GetLanguageCode

`func (o *Track) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *Track) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *Track) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *Track) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetName

`func (o *Track) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Track) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Track) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Track) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClosedCaptions

`func (o *Track) GetClosedCaptions() bool`

GetClosedCaptions returns the ClosedCaptions field if non-nil, zero value otherwise.

### GetClosedCaptionsOk

`func (o *Track) GetClosedCaptionsOk() (*bool, bool)`

GetClosedCaptionsOk returns a tuple with the ClosedCaptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedCaptions

`func (o *Track) SetClosedCaptions(v bool)`

SetClosedCaptions sets ClosedCaptions field to given value.

### HasClosedCaptions

`func (o *Track) HasClosedCaptions() bool`

HasClosedCaptions returns a boolean if a field has been set.

### GetPassthrough

`func (o *Track) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *Track) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *Track) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *Track) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


