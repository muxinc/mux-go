// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type GetAssetOrLiveStreamIdResponseData struct {
	// The Playback ID used to retrieve the corresponding asset or the live stream ID
	Id     string                                   `json:"id,omitempty"`
	Policy PlaybackPolicy                           `json:"policy,omitempty"`
	Object GetAssetOrLiveStreamIdResponseDataObject `json:"object,omitempty"`
}
