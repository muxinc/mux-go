# DeliveryReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveStreamId** | **string** | Unique identifier for the live stream that created the asset. | [optional] 
**AssetId** | **string** | Unique identifier for the asset. | [optional] 
**Passthrough** | **string** | The &#x60;passthrough&#x60; value for the asset. | [optional] 
**CreatedAt** | **string** | Time at which the asset was created. Measured in seconds since the Unix epoch. | [optional] 
**DeletedAt** | **string** | If exists, time at which the asset was deleted. Measured in seconds since the Unix epoch. | [optional] 
**AssetState** | **string** | The state of the asset. | [optional] 
**AssetDuration** | **float64** | The duration of the asset in seconds. | [optional] 
**AssetResolutionTier** | **string** | The resolution tier that the asset was ingested at, affecting billing for ingest &amp; storage | [optional] 
**AssetEncodingTier** | **string** | This field is deprecated. Please use &#x60;asset_video_quality&#x60; instead. The encoding tier that the asset was ingested at. [See the video quality guide for more details.](https://docs.mux.com/guides/use-video-quality-levels) | [optional] 
**AssetVideoQuality** | **string** | The video quality that the asset was ingested at. This field replaces &#x60;asset_encoding_tier&#x60;. [See the video quality guide for more details.](https://docs.mux.com/guides/use-video-quality-levels) | [optional] 
**DeliveredSeconds** | **float64** | Total number of delivered seconds during this time window. | [optional] 
**DeliveredSecondsByResolution** | [**DeliveryReportDeliveredSecondsByResolution**](DeliveryReport_delivered_seconds_by_resolution.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


