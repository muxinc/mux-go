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

// ListSigningKeysResponse struct for ListSigningKeysResponse
type ListSigningKeysResponse struct {
	Data *[]SigningKey `json:"data,omitempty"`
}

// NewListSigningKeysResponse instantiates a new ListSigningKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSigningKeysResponse() *ListSigningKeysResponse {
	this := ListSigningKeysResponse{}
	return &this
}

// NewListSigningKeysResponseWithDefaults instantiates a new ListSigningKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSigningKeysResponseWithDefaults() *ListSigningKeysResponse {
	this := ListSigningKeysResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSigningKeysResponse) GetData() []SigningKey {
	if o == nil || o.Data == nil {
		var ret []SigningKey
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSigningKeysResponse) GetDataOk() (*[]SigningKey, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSigningKeysResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SigningKey and assigns it to the Data field.
func (o *ListSigningKeysResponse) SetData(v []SigningKey) {
	o.Data = &v
}

func (o ListSigningKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListSigningKeysResponse struct {
	value *ListSigningKeysResponse
	isSet bool
}

func (v NullableListSigningKeysResponse) Get() *ListSigningKeysResponse {
	return v.value
}

func (v *NullableListSigningKeysResponse) Set(val *ListSigningKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSigningKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSigningKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSigningKeysResponse(val *ListSigningKeysResponse) *NullableListSigningKeysResponse {
	return &NullableListSigningKeysResponse{value: val, isSet: true}
}

func (v NullableListSigningKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSigningKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


