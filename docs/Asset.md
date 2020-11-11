# Asset

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**CreatedAt** | **string** |  | [optional] 
**DeletedAt** | **string** |  | [optional] 
**Status** | **string** |  | [optional] 
**Duration** | **float64** |  | [optional] 
**MaxStoredResolution** | **string** |  | [optional] 
**MaxStoredFrameRate** | **float64** |  | [optional] 
**AspectRatio** | **string** |  | [optional] 
**PlaybackIds** | [**[]PlaybackId**](PlaybackID.md) |  | [optional] 
**Tracks** | [**[]Track**](Track.md) |  | [optional] 
**Errors** | [**AssetErrors**](Asset_errors.md) |  | [optional] 
**PerTitleEncode** | **bool** |  | [optional] 
**IsLive** | **bool** |  | [optional] 
**Passthrough** | **string** |  | [optional] 
**LiveStreamId** | **string** |  | [optional] 
**Master** | [**AssetMaster**](Asset_master.md) |  | [optional] 
**MasterAccess** | **string** |  | [optional] [default to MASTER_ACCESS_NONE]
**Mp4Support** | **string** |  | [optional] [default to MP4_SUPPORT_NONE]
**NormalizeAudio** | **bool** |  | [optional] [default to false]
**StaticRenditions** | [**AssetStaticRenditions**](Asset_static_renditions.md) |  | [optional] 
**RecordingTimes** | [**[]AssetRecordingTimes**](Asset_recording_times.md) | An array of individual live stream recording sessions. A recording session is created on each encoder connection during the live stream | [optional] 
**NonStandardInputReasons** | [**AssetNonStandardInputReasons**](Asset_non_standard_input_reasons.md) |  | [optional] 
**Test** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


