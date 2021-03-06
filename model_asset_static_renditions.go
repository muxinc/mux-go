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

// AssetStaticRenditions An object containing the current status of any static renditions (mp4s). The object does not exist if no static renditions have been requested. See [Download your videos](https://docs.mux.com/guides/video/download-your-videos) for more information.
type AssetStaticRenditions struct {
	// Indicates the status of downloadable MP4 versions of this asset.
	Status *string `json:"status,omitempty"`
	// Array of file objects.
	Files *[]AssetStaticRenditionsFiles `json:"files,omitempty"`
}

// NewAssetStaticRenditions instantiates a new AssetStaticRenditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetStaticRenditions() *AssetStaticRenditions {
	this := AssetStaticRenditions{}
	var status string = "disabled"
	this.Status = &status
	return &this
}

// NewAssetStaticRenditionsWithDefaults instantiates a new AssetStaticRenditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetStaticRenditionsWithDefaults() *AssetStaticRenditions {
	this := AssetStaticRenditions{}
	var status string = "disabled"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetStaticRenditions) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetStaticRenditions) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetStaticRenditions) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetStaticRenditions) SetStatus(v string) {
	o.Status = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *AssetStaticRenditions) GetFiles() []AssetStaticRenditionsFiles {
	if o == nil || o.Files == nil {
		var ret []AssetStaticRenditionsFiles
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetStaticRenditions) GetFilesOk() (*[]AssetStaticRenditionsFiles, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AssetStaticRenditions) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []AssetStaticRenditionsFiles and assigns it to the Files field.
func (o *AssetStaticRenditions) SetFiles(v []AssetStaticRenditionsFiles) {
	o.Files = &v
}

func (o AssetStaticRenditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return json.Marshal(toSerialize)
}

type NullableAssetStaticRenditions struct {
	value *AssetStaticRenditions
	isSet bool
}

func (v NullableAssetStaticRenditions) Get() *AssetStaticRenditions {
	return v.value
}

func (v *NullableAssetStaticRenditions) Set(val *AssetStaticRenditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetStaticRenditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetStaticRenditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetStaticRenditions(val *AssetStaticRenditions) *NullableAssetStaticRenditions {
	return &NullableAssetStaticRenditions{value: val, isSet: true}
}

func (v NullableAssetStaticRenditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetStaticRenditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


