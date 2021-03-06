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

// AssetResponse struct for AssetResponse
type AssetResponse struct {
	Data *Asset `json:"data,omitempty"`
}

// NewAssetResponse instantiates a new AssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetResponse() *AssetResponse {
	this := AssetResponse{}
	return &this
}

// NewAssetResponseWithDefaults instantiates a new AssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetResponseWithDefaults() *AssetResponse {
	this := AssetResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AssetResponse) GetData() Asset {
	if o == nil || o.Data == nil {
		var ret Asset
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetResponse) GetDataOk() (*Asset, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AssetResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Asset and assigns it to the Data field.
func (o *AssetResponse) SetData(v Asset) {
	o.Data = &v
}

func (o AssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAssetResponse struct {
	value *AssetResponse
	isSet bool
}

func (v NullableAssetResponse) Get() *AssetResponse {
	return v.value
}

func (v *NullableAssetResponse) Set(val *AssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetResponse(val *AssetResponse) *NullableAssetResponse {
	return &NullableAssetResponse{value: val, isSet: true}
}

func (v NullableAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


