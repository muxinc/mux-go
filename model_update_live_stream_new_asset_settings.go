// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

// Updates the new asset settings to use to generate a new asset for this live stream. Only the `mp4_support` and `master_access` settings may be updated.
type UpdateLiveStreamNewAssetSettings struct {
	// Specify what level of support for mp4 playback should be added to new assets generated from this live stream. * The `none` option disables MP4 support for new assets. MP4 files will not be produced for an asset generated from this live stream. * The `capped-1080p` option produces a single MP4 file, called `capped-1080p.mp4`, with the video resolution capped at 1080p. This option produces an `audio.m4a` file for an audio-only asset. * The `audio-only` option produces a single M4A file, called `audio.m4a` for a video or an audio-only asset. MP4 generation will error when this option is specified for a video-only asset. * The `audio-only,capped-1080p` option produces both the `audio.m4a` and `capped-1080p.mp4` files. Only the `capped-1080p.mp4` file is produced for a video-only asset, while only the `audio.m4a` file is produced for an audio-only asset. * The `standard`(deprecated) option produces up to three MP4 files with different levels of resolution (`high.mp4`, `medium.mp4`, `low.mp4`, or `audio.m4a` for an audio-only asset).
	Mp4Support string `json:"mp4_support,omitempty"`
	// Add or remove access to the master version of the video.
	MasterAccess string `json:"master_access,omitempty"`
}
