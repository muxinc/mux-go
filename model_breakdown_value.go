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

// BreakdownValue struct for BreakdownValue
type BreakdownValue struct {
	Views *int64 `json:"views,omitempty"`
	Value *float64 `json:"value,omitempty"`
	TotalWatchTime *int64 `json:"total_watch_time,omitempty"`
	NegativeImpact *int32 `json:"negative_impact,omitempty"`
	Field *string `json:"field,omitempty"`
}

// NewBreakdownValue instantiates a new BreakdownValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreakdownValue() *BreakdownValue {
	this := BreakdownValue{}
	return &this
}

// NewBreakdownValueWithDefaults instantiates a new BreakdownValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreakdownValueWithDefaults() *BreakdownValue {
	this := BreakdownValue{}
	return &this
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *BreakdownValue) GetViews() int64 {
	if o == nil || o.Views == nil {
		var ret int64
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BreakdownValue) GetViewsOk() (*int64, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *BreakdownValue) HasViews() bool {
	if o != nil && o.Views != nil {
		return true
	}

	return false
}

// SetViews gets a reference to the given int64 and assigns it to the Views field.
func (o *BreakdownValue) SetViews(v int64) {
	o.Views = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BreakdownValue) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BreakdownValue) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BreakdownValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BreakdownValue) SetValue(v float64) {
	o.Value = &v
}

// GetTotalWatchTime returns the TotalWatchTime field value if set, zero value otherwise.
func (o *BreakdownValue) GetTotalWatchTime() int64 {
	if o == nil || o.TotalWatchTime == nil {
		var ret int64
		return ret
	}
	return *o.TotalWatchTime
}

// GetTotalWatchTimeOk returns a tuple with the TotalWatchTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BreakdownValue) GetTotalWatchTimeOk() (*int64, bool) {
	if o == nil || o.TotalWatchTime == nil {
		return nil, false
	}
	return o.TotalWatchTime, true
}

// HasTotalWatchTime returns a boolean if a field has been set.
func (o *BreakdownValue) HasTotalWatchTime() bool {
	if o != nil && o.TotalWatchTime != nil {
		return true
	}

	return false
}

// SetTotalWatchTime gets a reference to the given int64 and assigns it to the TotalWatchTime field.
func (o *BreakdownValue) SetTotalWatchTime(v int64) {
	o.TotalWatchTime = &v
}

// GetNegativeImpact returns the NegativeImpact field value if set, zero value otherwise.
func (o *BreakdownValue) GetNegativeImpact() int32 {
	if o == nil || o.NegativeImpact == nil {
		var ret int32
		return ret
	}
	return *o.NegativeImpact
}

// GetNegativeImpactOk returns a tuple with the NegativeImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BreakdownValue) GetNegativeImpactOk() (*int32, bool) {
	if o == nil || o.NegativeImpact == nil {
		return nil, false
	}
	return o.NegativeImpact, true
}

// HasNegativeImpact returns a boolean if a field has been set.
func (o *BreakdownValue) HasNegativeImpact() bool {
	if o != nil && o.NegativeImpact != nil {
		return true
	}

	return false
}

// SetNegativeImpact gets a reference to the given int32 and assigns it to the NegativeImpact field.
func (o *BreakdownValue) SetNegativeImpact(v int32) {
	o.NegativeImpact = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *BreakdownValue) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BreakdownValue) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *BreakdownValue) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *BreakdownValue) SetField(v string) {
	o.Field = &v
}

func (o BreakdownValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Views != nil {
		toSerialize["views"] = o.Views
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.TotalWatchTime != nil {
		toSerialize["total_watch_time"] = o.TotalWatchTime
	}
	if o.NegativeImpact != nil {
		toSerialize["negative_impact"] = o.NegativeImpact
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	return json.Marshal(toSerialize)
}

type NullableBreakdownValue struct {
	value *BreakdownValue
	isSet bool
}

func (v NullableBreakdownValue) Get() *BreakdownValue {
	return v.value
}

func (v *NullableBreakdownValue) Set(val *BreakdownValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBreakdownValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBreakdownValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreakdownValue(val *BreakdownValue) *NullableBreakdownValue {
	return &NullableBreakdownValue{value: val, isSet: true}
}

func (v NullableBreakdownValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreakdownValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


