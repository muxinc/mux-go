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

// RealTimeHistogramTimeseriesBucket struct for RealTimeHistogramTimeseriesBucket
type RealTimeHistogramTimeseriesBucket struct {
	Start *int64 `json:"start,omitempty"`
	End *int64 `json:"end,omitempty"`
}

// NewRealTimeHistogramTimeseriesBucket instantiates a new RealTimeHistogramTimeseriesBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealTimeHistogramTimeseriesBucket() *RealTimeHistogramTimeseriesBucket {
	this := RealTimeHistogramTimeseriesBucket{}
	return &this
}

// NewRealTimeHistogramTimeseriesBucketWithDefaults instantiates a new RealTimeHistogramTimeseriesBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealTimeHistogramTimeseriesBucketWithDefaults() *RealTimeHistogramTimeseriesBucket {
	this := RealTimeHistogramTimeseriesBucket{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *RealTimeHistogramTimeseriesBucket) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealTimeHistogramTimeseriesBucket) GetStartOk() (*int64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *RealTimeHistogramTimeseriesBucket) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *RealTimeHistogramTimeseriesBucket) SetStart(v int64) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *RealTimeHistogramTimeseriesBucket) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealTimeHistogramTimeseriesBucket) GetEndOk() (*int64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *RealTimeHistogramTimeseriesBucket) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *RealTimeHistogramTimeseriesBucket) SetEnd(v int64) {
	o.End = &v
}

func (o RealTimeHistogramTimeseriesBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableRealTimeHistogramTimeseriesBucket struct {
	value *RealTimeHistogramTimeseriesBucket
	isSet bool
}

func (v NullableRealTimeHistogramTimeseriesBucket) Get() *RealTimeHistogramTimeseriesBucket {
	return v.value
}

func (v *NullableRealTimeHistogramTimeseriesBucket) Set(val *RealTimeHistogramTimeseriesBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableRealTimeHistogramTimeseriesBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableRealTimeHistogramTimeseriesBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealTimeHistogramTimeseriesBucket(val *RealTimeHistogramTimeseriesBucket) *NullableRealTimeHistogramTimeseriesBucket {
	return &NullableRealTimeHistogramTimeseriesBucket{value: val, isSet: true}
}

func (v NullableRealTimeHistogramTimeseriesBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealTimeHistogramTimeseriesBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


