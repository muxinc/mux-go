/*
 * Mux API
 *
 * Mux is how developers build online video. This API encompasses both Mux Video and Mux Data functionality to help you build your video-related projects better and faster than ever before.
 *
 * API version: v1
 * Contact: devex@mux.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package muxgo

import (
	"encoding/json"
)

// InputFile struct for InputFile
type InputFile struct {
	ContainerFormat *string `json:"container_format,omitempty"`
	Tracks *[]InputTrack `json:"tracks,omitempty"`
}

// NewInputFile instantiates a new InputFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputFile() *InputFile {
	this := InputFile{}
	return &this
}

// NewInputFileWithDefaults instantiates a new InputFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputFileWithDefaults() *InputFile {
	this := InputFile{}
	return &this
}

// GetContainerFormat returns the ContainerFormat field value if set, zero value otherwise.
func (o *InputFile) GetContainerFormat() string {
	if o == nil || o.ContainerFormat == nil {
		var ret string
		return ret
	}
	return *o.ContainerFormat
}

// GetContainerFormatOk returns a tuple with the ContainerFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputFile) GetContainerFormatOk() (*string, bool) {
	if o == nil || o.ContainerFormat == nil {
		return nil, false
	}
	return o.ContainerFormat, true
}

// HasContainerFormat returns a boolean if a field has been set.
func (o *InputFile) HasContainerFormat() bool {
	if o != nil && o.ContainerFormat != nil {
		return true
	}

	return false
}

// SetContainerFormat gets a reference to the given string and assigns it to the ContainerFormat field.
func (o *InputFile) SetContainerFormat(v string) {
	o.ContainerFormat = &v
}

// GetTracks returns the Tracks field value if set, zero value otherwise.
func (o *InputFile) GetTracks() []InputTrack {
	if o == nil || o.Tracks == nil {
		var ret []InputTrack
		return ret
	}
	return *o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputFile) GetTracksOk() (*[]InputTrack, bool) {
	if o == nil || o.Tracks == nil {
		return nil, false
	}
	return o.Tracks, true
}

// HasTracks returns a boolean if a field has been set.
func (o *InputFile) HasTracks() bool {
	if o != nil && o.Tracks != nil {
		return true
	}

	return false
}

// SetTracks gets a reference to the given []InputTrack and assigns it to the Tracks field.
func (o *InputFile) SetTracks(v []InputTrack) {
	o.Tracks = &v
}

func (o InputFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerFormat != nil {
		toSerialize["container_format"] = o.ContainerFormat
	}
	if o.Tracks != nil {
		toSerialize["tracks"] = o.Tracks
	}
	return json.Marshal(toSerialize)
}

type NullableInputFile struct {
	value *InputFile
	isSet bool
}

func (v NullableInputFile) Get() *InputFile {
	return v.value
}

func (v *NullableInputFile) Set(val *InputFile) {
	v.value = val
	v.isSet = true
}

func (v NullableInputFile) IsSet() bool {
	return v.isSet
}

func (v *NullableInputFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputFile(val *InputFile) *NullableInputFile {
	return &NullableInputFile{value: val, isSet: true}
}

func (v NullableInputFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


