// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateAssetRequest struct {
	// An array of objects that each describe an input file to be used to create the asset. As a shortcut, input can also be a string URL for a file when only one input file is used. See `input[].url` for requirements.
	Input []InputSettings `json:"input,omitempty"`
	// An array of playback policy names that you want applied to this asset and available through `playback_ids`. Options include: `\"public\"` (anyone with the playback URL can stream the asset). And `\"signed\"` (an additional access token is required to play the asset). If no playback_policy is set, the asset will have no playback IDs and will therefore not be playable. For simplicity, a single string name can be used in place of the array in the case of only one playback policy.
	PlaybackPolicy []PlaybackPolicy `json:"playback_policy,omitempty"`
	PerTitleEncode bool             `json:"per_title_encode,omitempty"`
	// Arbitrary user-supplied metadata that will be included in the asset details and related webhooks. Can be used to store your own ID for a video along with the asset. **Max: 255 characters**.
	Passthrough string `json:"passthrough,omitempty"`
	// Specify what level (if any) of support for mp4 playback. In most cases you should use our default HLS-based streaming playback ({playback_id}.m3u8) which can automatically adjust to viewers' connection speeds, but an mp4 can be useful for some legacy devices or downloading for offline playback. See the [Download your videos guide](/guides/video/download-your-videos) for more information.
	Mp4Support string `json:"mp4_support,omitempty"`
	// Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets.
	NormalizeAudio bool `json:"normalize_audio,omitempty"`
	// Specify what level (if any) of support for master access. Master access can be enabled temporarily for your asset to be downloaded. See the [Download your videos guide](/guides/video/download-your-videos) for more information.
	MasterAccess string `json:"master_access,omitempty"`
	// Marks the asset as a test asset when the value is set to true. A Test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test asset are watermarked with the Mux logo, limited to 10 seconds, deleted after 24 hrs.
	Test bool `json:"test,omitempty"`
	// Max resolution tier can be used to control the maximum `resolution_tier` your asset is encoded, stored, and streamed at. If not set, this defaults to `1080p`.
	MaxResolutionTier string `json:"max_resolution_tier,omitempty"`
	// The encoding tier informs the cost, quality, and available platform features for the asset. By default the `smart` encoding tier is used. [See the guide for more details.](https://docs.mux.com/guides/video/use-encoding-tiers)
	EncodingTier string `json:"encoding_tier,omitempty"`
}
