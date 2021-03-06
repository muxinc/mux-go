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

// ListLiveStreamsResponse struct for ListLiveStreamsResponse
type ListLiveStreamsResponse struct {
	Data *[]LiveStream `json:"data,omitempty"`
}

// NewListLiveStreamsResponse instantiates a new ListLiveStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLiveStreamsResponse() *ListLiveStreamsResponse {
	this := ListLiveStreamsResponse{}
	return &this
}

// NewListLiveStreamsResponseWithDefaults instantiates a new ListLiveStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLiveStreamsResponseWithDefaults() *ListLiveStreamsResponse {
	this := ListLiveStreamsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListLiveStreamsResponse) GetData() []LiveStream {
	if o == nil || o.Data == nil {
		var ret []LiveStream
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLiveStreamsResponse) GetDataOk() (*[]LiveStream, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListLiveStreamsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []LiveStream and assigns it to the Data field.
func (o *ListLiveStreamsResponse) SetData(v []LiveStream) {
	o.Data = &v
}

func (o ListLiveStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListLiveStreamsResponse struct {
	value *ListLiveStreamsResponse
	isSet bool
}

func (v NullableListLiveStreamsResponse) Get() *ListLiveStreamsResponse {
	return v.value
}

func (v *NullableListLiveStreamsResponse) Set(val *ListLiveStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListLiveStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListLiveStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLiveStreamsResponse(val *ListLiveStreamsResponse) *NullableListLiveStreamsResponse {
	return &NullableListLiveStreamsResponse{value: val, isSet: true}
}

func (v NullableListLiveStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLiveStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


