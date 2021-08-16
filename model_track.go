/*
 * Mux API
 *
 * Mux is how developers build online video. This API encompasses both Mux Video and Mux Data functionality to help you build your video-related projects better and faster than ever before. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package muxgo

import (
	"encoding/json"
)

// Track struct for Track
type Track struct {
	// Unique identifier for the Track
	Id *string `json:"id,omitempty"`
	// The type of track
	Type *string `json:"type,omitempty"`
	// The duration in seconds of the track media. This parameter is not set for the `text` type track. This field is optional and may not be set. The top level `duration` field of an asset will always be set.
	Duration *float64 `json:"duration,omitempty"`
	// The maximum width in pixels available for the track. Only set for the `video` type track.
	MaxWidth *int64 `json:"max_width,omitempty"`
	// The maximum height in pixels available for the track. Only set for the `video` type track.
	MaxHeight *int64 `json:"max_height,omitempty"`
	// The maximum frame rate available for the track. Only set for the `video` type track. This field may return `-1` if the frame rate of the input cannot be reliably determined.
	MaxFrameRate *float64 `json:"max_frame_rate,omitempty"`
	// The maximum number of audio channels the track supports. Only set for the `audio` type track.
	MaxChannels *int64 `json:"max_channels,omitempty"`
	// Only set for the `audio` type track.
	MaxChannelLayout *string `json:"max_channel_layout,omitempty"`
	// This parameter is set only for the `text` type track.
	TextType *string `json:"text_type,omitempty"`
	// The language code value represents [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, `en` for English or `en-US` for the US version of English. This parameter is set for `text` type and `subtitles` text type track.
	LanguageCode *string `json:"language_code,omitempty"`
	// The name of the track containing a human-readable description. The hls manifest will associate a subtitle text track with this value. For example, the value is \"English\" for subtitles text track for the `language_code` value of `en-US`. This parameter is set for the `text` type and `subtitles` text type track.
	Name *string `json:"name,omitempty"`
	// Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This parameter is set for the `text` type and `subtitles` text type track.
	ClosedCaptions *bool `json:"closed_captions,omitempty"`
	// Arbitrary metadata set for the track either when creating the asset or track. This parameter is set for `text` type and `subtitles` text type track. Max 255 characters.
	Passthrough *string `json:"passthrough,omitempty"`
}

// NewTrack instantiates a new Track object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrack() *Track {
	this := Track{}
	return &this
}

// NewTrackWithDefaults instantiates a new Track object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackWithDefaults() *Track {
	this := Track{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Track) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Track) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Track) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Track) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Track) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Track) SetType(v string) {
	o.Type = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Track) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Track) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *Track) SetDuration(v float64) {
	o.Duration = &v
}

// GetMaxWidth returns the MaxWidth field value if set, zero value otherwise.
func (o *Track) GetMaxWidth() int64 {
	if o == nil || o.MaxWidth == nil {
		var ret int64
		return ret
	}
	return *o.MaxWidth
}

// GetMaxWidthOk returns a tuple with the MaxWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetMaxWidthOk() (*int64, bool) {
	if o == nil || o.MaxWidth == nil {
		return nil, false
	}
	return o.MaxWidth, true
}

// HasMaxWidth returns a boolean if a field has been set.
func (o *Track) HasMaxWidth() bool {
	if o != nil && o.MaxWidth != nil {
		return true
	}

	return false
}

// SetMaxWidth gets a reference to the given int64 and assigns it to the MaxWidth field.
func (o *Track) SetMaxWidth(v int64) {
	o.MaxWidth = &v
}

// GetMaxHeight returns the MaxHeight field value if set, zero value otherwise.
func (o *Track) GetMaxHeight() int64 {
	if o == nil || o.MaxHeight == nil {
		var ret int64
		return ret
	}
	return *o.MaxHeight
}

// GetMaxHeightOk returns a tuple with the MaxHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetMaxHeightOk() (*int64, bool) {
	if o == nil || o.MaxHeight == nil {
		return nil, false
	}
	return o.MaxHeight, true
}

// HasMaxHeight returns a boolean if a field has been set.
func (o *Track) HasMaxHeight() bool {
	if o != nil && o.MaxHeight != nil {
		return true
	}

	return false
}

// SetMaxHeight gets a reference to the given int64 and assigns it to the MaxHeight field.
func (o *Track) SetMaxHeight(v int64) {
	o.MaxHeight = &v
}

// GetMaxFrameRate returns the MaxFrameRate field value if set, zero value otherwise.
func (o *Track) GetMaxFrameRate() float64 {
	if o == nil || o.MaxFrameRate == nil {
		var ret float64
		return ret
	}
	return *o.MaxFrameRate
}

// GetMaxFrameRateOk returns a tuple with the MaxFrameRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetMaxFrameRateOk() (*float64, bool) {
	if o == nil || o.MaxFrameRate == nil {
		return nil, false
	}
	return o.MaxFrameRate, true
}

// HasMaxFrameRate returns a boolean if a field has been set.
func (o *Track) HasMaxFrameRate() bool {
	if o != nil && o.MaxFrameRate != nil {
		return true
	}

	return false
}

// SetMaxFrameRate gets a reference to the given float64 and assigns it to the MaxFrameRate field.
func (o *Track) SetMaxFrameRate(v float64) {
	o.MaxFrameRate = &v
}

// GetMaxChannels returns the MaxChannels field value if set, zero value otherwise.
func (o *Track) GetMaxChannels() int64 {
	if o == nil || o.MaxChannels == nil {
		var ret int64
		return ret
	}
	return *o.MaxChannels
}

// GetMaxChannelsOk returns a tuple with the MaxChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetMaxChannelsOk() (*int64, bool) {
	if o == nil || o.MaxChannels == nil {
		return nil, false
	}
	return o.MaxChannels, true
}

// HasMaxChannels returns a boolean if a field has been set.
func (o *Track) HasMaxChannels() bool {
	if o != nil && o.MaxChannels != nil {
		return true
	}

	return false
}

// SetMaxChannels gets a reference to the given int64 and assigns it to the MaxChannels field.
func (o *Track) SetMaxChannels(v int64) {
	o.MaxChannels = &v
}

// GetMaxChannelLayout returns the MaxChannelLayout field value if set, zero value otherwise.
func (o *Track) GetMaxChannelLayout() string {
	if o == nil || o.MaxChannelLayout == nil {
		var ret string
		return ret
	}
	return *o.MaxChannelLayout
}

// GetMaxChannelLayoutOk returns a tuple with the MaxChannelLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetMaxChannelLayoutOk() (*string, bool) {
	if o == nil || o.MaxChannelLayout == nil {
		return nil, false
	}
	return o.MaxChannelLayout, true
}

// HasMaxChannelLayout returns a boolean if a field has been set.
func (o *Track) HasMaxChannelLayout() bool {
	if o != nil && o.MaxChannelLayout != nil {
		return true
	}

	return false
}

// SetMaxChannelLayout gets a reference to the given string and assigns it to the MaxChannelLayout field.
func (o *Track) SetMaxChannelLayout(v string) {
	o.MaxChannelLayout = &v
}

// GetTextType returns the TextType field value if set, zero value otherwise.
func (o *Track) GetTextType() string {
	if o == nil || o.TextType == nil {
		var ret string
		return ret
	}
	return *o.TextType
}

// GetTextTypeOk returns a tuple with the TextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetTextTypeOk() (*string, bool) {
	if o == nil || o.TextType == nil {
		return nil, false
	}
	return o.TextType, true
}

// HasTextType returns a boolean if a field has been set.
func (o *Track) HasTextType() bool {
	if o != nil && o.TextType != nil {
		return true
	}

	return false
}

// SetTextType gets a reference to the given string and assigns it to the TextType field.
func (o *Track) SetTextType(v string) {
	o.TextType = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *Track) GetLanguageCode() string {
	if o == nil || o.LanguageCode == nil {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetLanguageCodeOk() (*string, bool) {
	if o == nil || o.LanguageCode == nil {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *Track) HasLanguageCode() bool {
	if o != nil && o.LanguageCode != nil {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *Track) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Track) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Track) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Track) SetName(v string) {
	o.Name = &v
}

// GetClosedCaptions returns the ClosedCaptions field value if set, zero value otherwise.
func (o *Track) GetClosedCaptions() bool {
	if o == nil || o.ClosedCaptions == nil {
		var ret bool
		return ret
	}
	return *o.ClosedCaptions
}

// GetClosedCaptionsOk returns a tuple with the ClosedCaptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetClosedCaptionsOk() (*bool, bool) {
	if o == nil || o.ClosedCaptions == nil {
		return nil, false
	}
	return o.ClosedCaptions, true
}

// HasClosedCaptions returns a boolean if a field has been set.
func (o *Track) HasClosedCaptions() bool {
	if o != nil && o.ClosedCaptions != nil {
		return true
	}

	return false
}

// SetClosedCaptions gets a reference to the given bool and assigns it to the ClosedCaptions field.
func (o *Track) SetClosedCaptions(v bool) {
	o.ClosedCaptions = &v
}

// GetPassthrough returns the Passthrough field value if set, zero value otherwise.
func (o *Track) GetPassthrough() string {
	if o == nil || o.Passthrough == nil {
		var ret string
		return ret
	}
	return *o.Passthrough
}

// GetPassthroughOk returns a tuple with the Passthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Track) GetPassthroughOk() (*string, bool) {
	if o == nil || o.Passthrough == nil {
		return nil, false
	}
	return o.Passthrough, true
}

// HasPassthrough returns a boolean if a field has been set.
func (o *Track) HasPassthrough() bool {
	if o != nil && o.Passthrough != nil {
		return true
	}

	return false
}

// SetPassthrough gets a reference to the given string and assigns it to the Passthrough field.
func (o *Track) SetPassthrough(v string) {
	o.Passthrough = &v
}

func (o Track) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.MaxWidth != nil {
		toSerialize["max_width"] = o.MaxWidth
	}
	if o.MaxHeight != nil {
		toSerialize["max_height"] = o.MaxHeight
	}
	if o.MaxFrameRate != nil {
		toSerialize["max_frame_rate"] = o.MaxFrameRate
	}
	if o.MaxChannels != nil {
		toSerialize["max_channels"] = o.MaxChannels
	}
	if o.MaxChannelLayout != nil {
		toSerialize["max_channel_layout"] = o.MaxChannelLayout
	}
	if o.TextType != nil {
		toSerialize["text_type"] = o.TextType
	}
	if o.LanguageCode != nil {
		toSerialize["language_code"] = o.LanguageCode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClosedCaptions != nil {
		toSerialize["closed_captions"] = o.ClosedCaptions
	}
	if o.Passthrough != nil {
		toSerialize["passthrough"] = o.Passthrough
	}
	return json.Marshal(toSerialize)
}

type NullableTrack struct {
	value *Track
	isSet bool
}

func (v NullableTrack) Get() *Track {
	return v.value
}

func (v *NullableTrack) Set(val *Track) {
	v.value = val
	v.isSet = true
}

func (v NullableTrack) IsSet() bool {
	return v.isSet
}

func (v *NullableTrack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrack(val *Track) *NullableTrack {
	return &NullableTrack{value: val, isSet: true}
}

func (v NullableTrack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


