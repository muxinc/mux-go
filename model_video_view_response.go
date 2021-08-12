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

// VideoViewResponse struct for VideoViewResponse
type VideoViewResponse struct {
	Data *VideoView `json:"data,omitempty"`
	Timeframe *[]int64 `json:"timeframe,omitempty"`
}

// NewVideoViewResponse instantiates a new VideoViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoViewResponse() *VideoViewResponse {
	this := VideoViewResponse{}
	return &this
}

// NewVideoViewResponseWithDefaults instantiates a new VideoViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoViewResponseWithDefaults() *VideoViewResponse {
	this := VideoViewResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VideoViewResponse) GetData() VideoView {
	if o == nil || o.Data == nil {
		var ret VideoView
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoViewResponse) GetDataOk() (*VideoView, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VideoViewResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given VideoView and assigns it to the Data field.
func (o *VideoViewResponse) SetData(v VideoView) {
	o.Data = &v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *VideoViewResponse) GetTimeframe() []int64 {
	if o == nil || o.Timeframe == nil {
		var ret []int64
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoViewResponse) GetTimeframeOk() (*[]int64, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *VideoViewResponse) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given []int64 and assigns it to the Timeframe field.
func (o *VideoViewResponse) SetTimeframe(v []int64) {
	o.Timeframe = &v
}

func (o VideoViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	return json.Marshal(toSerialize)
}

type NullableVideoViewResponse struct {
	value *VideoViewResponse
	isSet bool
}

func (v NullableVideoViewResponse) Get() *VideoViewResponse {
	return v.value
}

func (v *NullableVideoViewResponse) Set(val *VideoViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoViewResponse(val *VideoViewResponse) *NullableVideoViewResponse {
	return &NullableVideoViewResponse{value: val, isSet: true}
}

func (v NullableVideoViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


