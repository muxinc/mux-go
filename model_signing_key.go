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

// SigningKey struct for SigningKey
type SigningKey struct {
	// Unique identifier for the Signing Key.
	Id *string `json:"id,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	CreatedAt *string `json:"created_at,omitempty"`
	// A Base64 encoded private key that can be used with the RS256 algorithm when creating a [JWT](https://jwt.io/). **Note that this value is only returned once when creating a URL signing key.**
	PrivateKey *string `json:"private_key,omitempty"`
}

// NewSigningKey instantiates a new SigningKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKey() *SigningKey {
	this := SigningKey{}
	return &this
}

// NewSigningKeyWithDefaults instantiates a new SigningKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeyWithDefaults() *SigningKey {
	this := SigningKey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SigningKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SigningKey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SigningKey) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SigningKey) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKey) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SigningKey) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SigningKey) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *SigningKey) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKey) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *SigningKey) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *SigningKey) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o SigningKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableSigningKey struct {
	value *SigningKey
	isSet bool
}

func (v NullableSigningKey) Get() *SigningKey {
	return v.value
}

func (v *NullableSigningKey) Set(val *SigningKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningKey(val *SigningKey) *NullableSigningKey {
	return &NullableSigningKey{value: val, isSet: true}
}

func (v NullableSigningKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


