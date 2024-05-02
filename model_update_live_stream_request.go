// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type UpdateLiveStreamRequest struct {
	// Arbitrary user-supplied metadata set for the live stream. Max 255 characters. In order to clear this value, the field should be included with an empty-string value.
	Passthrough string `json:"passthrough,omitempty"`
	// Latency is the time from when the streamer transmits a frame of video to when you see it in the player. Set this as an alternative to setting low latency or reduced latency flags.
	LatencyMode string `json:"latency_mode,omitempty"`
	// When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset.  If not specified directly, Standard Latency streams have a Reconnect Window of 60 seconds; Reduced and Low Latency streams have a default of 0 seconds, or no Reconnect Window. For that reason, we suggest specifying a value other than zero for Reduced and Low Latency streams.  Reduced and Low Latency streams with a Reconnect Window greater than zero will insert slate media into the recorded asset while waiting for the streaming software to reconnect or when there are brief interruptions in the live stream media. When using a Reconnect Window setting higher than 60 seconds with a Standard Latency stream, we highly recommend enabling slate with the `use_slate_for_standard_latency` option.
	ReconnectWindow float32 `json:"reconnect_window,omitempty"`
	// By default, Standard Latency live streams do not have slate media inserted while waiting for live streaming software to reconnect to Mux. Setting this to true enables slate insertion on a Standard Latency stream.
	UseSlateForStandardLatency bool `json:"use_slate_for_standard_latency,omitempty"`
	// The URL of the image file that Mux should download and use as slate media during interruptions of the live stream media. This file will be downloaded each time a new recorded asset is created from the live stream. Set this to a blank string to clear the value so that the default slate media will be used.
	ReconnectSlateUrl string `json:"reconnect_slate_url,omitempty"`
	// The time in seconds a live stream may be continuously active before being disconnected. Defaults to 12 hours.
	MaxContinuousDuration int32                            `json:"max_continuous_duration,omitempty"`
	NewAssetSettings      UpdateLiveStreamNewAssetSettings `json:"new_asset_settings,omitempty"`
}
