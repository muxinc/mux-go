# CreateAssetRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**[]InputSettings**](InputSettings.md) | An array of objects that each describe an input file to be used to create the asset. As a shortcut, input can also be a string URL for a file when only one input file is used. See &#x60;input[].url&#x60; for requirements. | [optional] 
**PlaybackPolicy** | [**[]PlaybackPolicy**](PlaybackPolicy.md) | An array of playback policy names that you want applied to this asset and available through &#x60;playback_ids&#x60;. Options include: &#x60;\&quot;public\&quot;&#x60; (anyone with the playback URL can stream the asset). And &#x60;\&quot;signed\&quot;&#x60; (an additional access token is required to play the asset). If no playback_policy is set, the asset will have no playback IDs and will therefore not be playable. For simplicity, a single string name can be used in place of the array in the case of only one playback policy. | [optional] 
**PerTitleEncode** | **bool** |  | [optional] 
**Passthrough** | **string** | Arbitrary user-supplied metadata that will be included in the asset details and related webhooks. Can be used to store your own ID for a video along with the asset. **Max: 255 characters**. | [optional] 
**Mp4Support** | **string** | Specify what level (if any) of support for mp4 playback. In most cases you should use our default HLS-based streaming playback ({playback_id}.m3u8) which can automatically adjust to viewers&#39; connection speeds, but an mp4 can be useful for some legacy devices or downloading for offline playback. See the [Download your videos guide](/guides/video/download-your-videos) for more information. | [optional] 
**NormalizeAudio** | **bool** | Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets. | [optional] [default to false]
**MasterAccess** | **string** | Specify what level (if any) of support for master access. Master access can be enabled temporarily for your asset to be downloaded. See the [Download your videos guide](/guides/video/download-your-videos) for more information. | [optional] 
**Test** | **bool** | Marks the asset as a test asset when the value is set to true. A Test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test asset are watermarked with the Mux logo, limited to 10 seconds, deleted after 24 hrs. | [optional] 
**MaxResolutionTier** | **string** | Max resolution tier can be used to control the maximum &#x60;resolution_tier&#x60; your asset is encoded, stored, and streamed at. If not set, this defaults to &#x60;1080p&#x60;. | [optional] 
**EncodingTier** | **string** | The encoding tier informs the cost, quality, and available platform features for the asset. By default the &#x60;smart&#x60; encoding tier is used. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


