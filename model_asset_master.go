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

// AssetMaster An object containing the current status of Master Access and the link to the Master MP4 file when ready. This object does not exist if `master_acess` is set to `none` and when the temporary URL expires.
type AssetMaster struct {
	Status *string `json:"status,omitempty"`
	// The temporary URL to the master version of the video, as an MP4 file. This URL will expire after 24 hours.
	Url *string `json:"url,omitempty"`
}

// NewAssetMaster instantiates a new AssetMaster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetMaster() *AssetMaster {
	this := AssetMaster{}
	return &this
}

// NewAssetMasterWithDefaults instantiates a new AssetMaster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetMasterWithDefaults() *AssetMaster {
	this := AssetMaster{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetMaster) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMaster) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetMaster) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetMaster) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AssetMaster) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMaster) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AssetMaster) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AssetMaster) SetUrl(v string) {
	o.Url = &v
}

func (o AssetMaster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableAssetMaster struct {
	value *AssetMaster
	isSet bool
}

func (v NullableAssetMaster) Get() *AssetMaster {
	return v.value
}

func (v *NullableAssetMaster) Set(val *AssetMaster) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetMaster) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetMaster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetMaster(val *AssetMaster) *NullableAssetMaster {
	return &NullableAssetMaster{value: val, isSet: true}
}

func (v NullableAssetMaster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetMaster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


