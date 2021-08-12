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

// EnableLiveStreamResponse struct for EnableLiveStreamResponse
type EnableLiveStreamResponse struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewEnableLiveStreamResponse instantiates a new EnableLiveStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableLiveStreamResponse() *EnableLiveStreamResponse {
	this := EnableLiveStreamResponse{}
	return &this
}

// NewEnableLiveStreamResponseWithDefaults instantiates a new EnableLiveStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableLiveStreamResponseWithDefaults() *EnableLiveStreamResponse {
	this := EnableLiveStreamResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EnableLiveStreamResponse) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableLiveStreamResponse) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EnableLiveStreamResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *EnableLiveStreamResponse) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o EnableLiveStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEnableLiveStreamResponse struct {
	value *EnableLiveStreamResponse
	isSet bool
}

func (v NullableEnableLiveStreamResponse) Get() *EnableLiveStreamResponse {
	return v.value
}

func (v *NullableEnableLiveStreamResponse) Set(val *EnableLiveStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableLiveStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableLiveStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableLiveStreamResponse(val *EnableLiveStreamResponse) *NullableEnableLiveStreamResponse {
	return &NullableEnableLiveStreamResponse{value: val, isSet: true}
}

func (v NullableEnableLiveStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableLiveStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


