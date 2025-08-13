# Asset

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the Asset. Max 255 characters. | [optional] 
**CreatedAt** | **string** | Time the Asset was created, defined as a Unix timestamp (seconds since epoch). | [optional] 
**Status** | **string** | The status of the asset. | [optional] 
**Duration** | **float64** | The duration of the asset in seconds (max duration for a single asset is 12 hours). | [optional] 
**MaxStoredResolution** | **string** | This field is deprecated. Please use &#x60;resolution_tier&#x60; instead. The maximum resolution that has been stored for the asset. The asset may be delivered at lower resolutions depending on the device and bandwidth, however it cannot be delivered at a higher value than is stored. | [optional] 
**ResolutionTier** | **string** | The resolution tier that the asset was ingested at, affecting billing for ingest &amp; storage. This field also represents the highest resolution tier that the content can be delivered at, however the actual resolution may be lower depending on the device, bandwidth, and exact resolution of the uploaded asset. | [optional] 
**MaxResolutionTier** | **string** | Max resolution tier can be used to control the maximum &#x60;resolution_tier&#x60; your asset is encoded, stored, and streamed at. If not set, this defaults to &#x60;1080p&#x60;. | [optional] 
**EncodingTier** | **string** | This field is deprecated. Please use &#x60;video_quality&#x60; instead. The encoding tier informs the cost, quality, and available platform features for the asset. The default encoding tier for an account can be set in the Mux Dashboard. [See the video quality guide for more details.](https://docs.mux.com/guides/use-video-quality-levels) | [optional] 
**VideoQuality** | **string** | The video quality controls the cost, quality, and available platform features for the asset. The default video quality for an account can be set in the Mux Dashboard. This field replaces the deprecated &#x60;encoding_tier&#x60; value. [See the video quality guide for more details.](https://docs.mux.com/guides/use-video-quality-levels) | [optional] 
**MaxStoredFrameRate** | **float64** | The maximum frame rate that has been stored for the asset. The asset may be delivered at lower frame rates depending on the device and bandwidth, however it cannot be delivered at a higher value than is stored. This field may return -1 if the frame rate of the input cannot be reliably determined. | [optional] 
**AspectRatio** | **string** | The aspect ratio of the asset in the form of &#x60;width:height&#x60;, for example &#x60;16:9&#x60;. | [optional] 
**PlaybackIds** | [**[]PlaybackId**](PlaybackID.md) | An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/play-your-videos) for more details. | [optional] 
**Tracks** | [**[]Track**](Track.md) | The individual media tracks that make up an asset. | [optional] 
**Errors** | [**AssetErrors**](Asset_errors.md) |  | [optional] 
**PerTitleEncode** | **bool** |  | [optional] 
**UploadId** | **string** | Unique identifier for the Direct Upload. This is an optional parameter added when the asset is created from a direct upload. | [optional] 
**IsLive** | **bool** | Indicates whether the live stream that created this asset is currently &#x60;active&#x60; and not in &#x60;idle&#x60; state. This is an optional parameter added when the asset is created from a live stream. | [optional] 
**Passthrough** | **string** | You can set this field to anything you want. It will be included in the asset details and related webhooks. If you&#39;re looking for more structured metadata, such as &#x60;title&#x60; or &#x60;external_id&#x60; , you can use the &#x60;meta&#x60; object instead. **Max: 255 characters**. | [optional] 
**LiveStreamId** | **string** | Unique identifier for the live stream. This is an optional parameter added when the asset is created from a live stream. | [optional] 
**Master** | [**AssetMaster**](Asset_master.md) |  | [optional] 
**MasterAccess** | **string** |  | [optional] [default to MASTER_ACCESS_NONE]
**Mp4Support** | **string** |  | [optional] [default to MP4_SUPPORT_NONE]
**SourceAssetId** | **string** | Asset Identifier of the video used as the source for creating the clip. | [optional] 
**NormalizeAudio** | **bool** | Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets. | [optional] [default to false]
**StaticRenditions** | [**AssetStaticRenditions**](Asset_static_renditions.md) |  | [optional] 
**RecordingTimes** | [**[]AssetRecordingTimes**](Asset_recording_times.md) | An array of individual live stream recording sessions. A recording session is created on each encoder connection during the live stream. Additionally any time slate media is inserted during brief interruptions in the live stream media or times when the live streaming software disconnects, a recording session representing the slate media will be added with a \&quot;slate\&quot; type. | [optional] 
**NonStandardInputReasons** | [**AssetNonStandardInputReasons**](Asset_non_standard_input_reasons.md) |  | [optional] 
**Test** | **bool** | True means this live stream is a test asset. A test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test assets are watermarked with the Mux logo, limited to 10 seconds, and deleted after 24 hrs. | [optional] 
**IngestType** | **string** | The type of ingest used to create the asset. | [optional] 
**Meta** | [**AssetMetadata**](AssetMetadata.md) |  | [optional] 
**Progress** | [**AssetProgress**](AssetProgress.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


