# StaticRendition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the static rendition file | [optional] 
**Ext** | **string** | Extension of the static rendition file | [optional] 
**Height** | **int32** | The height of the static rendition&#39;s file in pixels | [optional] 
**Width** | **int32** | The width of the static rendition&#39;s file in pixels | [optional] 
**Bitrate** | **int64** | The bitrate in bits per second | [optional] 
**Filesize** | **string** | The file size in bytes | [optional] 
**Type** | **string** | Indicates the static rendition type of this specific MP4 version of this asset. This field is only valid for &#x60;static_renditions&#x60;, not for &#x60;mp4_support&#x60;. | [optional] 
**Status** | **string** | Indicates the status of this specific MP4 version of this asset. This field is only valid for &#x60;static_renditions&#x60;, not for &#x60;mp4_support&#x60;. * &#x60;ready&#x60; indicates the MP4 has been generated and is ready for download * &#x60;preparing&#x60; indicates the asset has not been ingested or the static rendition is still being generated after an asset is ready * &#x60;skipped&#x60; indicates the static rendition will not be generated because the requested resolution conflicts with the asset attributes after the asset has been ingested * &#x60;errored&#x60; indicates the static rendition cannot be generated. For example, an asset could not be ingested  | [optional] 
**ResolutionTier** | **string** | Indicates the resolution tier of this specific MP4 version of this asset. This field is only valid for &#x60;static_renditions&#x60;, not for &#x60;mp4_support&#x60;. | [optional] 
**Resolution** | **string** | Indicates the resolution of this specific MP4 version of this asset. This field is only valid for &#x60;static_renditions&#x60;, not for &#x60;mp4_support&#x60;. | [optional] 
**Id** | **string** | The ID of this static rendition, used in managing this static rendition. This field is only valid for &#x60;static_renditions&#x60;, not for &#x60;mp4_support&#x60;. | [optional] 
**Passthrough** | **string** | Arbitrary user-supplied metadata set for the static rendition. Max 255 characters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


