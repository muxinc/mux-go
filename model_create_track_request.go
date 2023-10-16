// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateTrackRequest struct {
	// The URL of the file that Mux should download and use. * For `audio` tracks, the URL is the location of the audio file for Mux to download, for example an M4A, WAV, or MP3 file. Mux supports most audio file formats and codecs, but for fastest processing, you should [use standard inputs wherever possible](https://docs.mux.com/guides/video/minimize-processing-time). * For `text` tracks, the URL is the location of subtitle/captions file. Mux supports [SubRip Text (SRT)](https://en.wikipedia.org/wiki/SubRip) and [Web Video Text Tracks](https://www.w3.org/TR/webvtt1/) formats for ingesting Subtitles and Closed Captions.
	Url      string `json:"url"`
	Type     string `json:"type"`
	TextType string `json:"text_type"`
	// The language code value must be a valid BCP 47 specification compliant value. For example, en for English or en-US for the US version of English.
	LanguageCode string `json:"language_code"`
	// The name of the track containing a human-readable description. This value must be unique within each group of `text` or `audio` track types. The HLS manifest will associate the `text` or `audio` track with this value. For example, set the value to \"English\" for subtitles text track with `language_code` as en-US. If this parameter is not included, Mux will auto-populate a value based on the `language_code` value.
	Name string `json:"name,omitempty"`
	// Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH).
	ClosedCaptions bool `json:"closed_captions,omitempty"`
	// Arbitrary user-supplied metadata set for the track either when creating the asset or track.
	Passthrough string `json:"passthrough,omitempty"`
}
