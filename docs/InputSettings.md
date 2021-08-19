# InputSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The web address of the file that Mux should download and use. * For subtitles text tracks, the url is the location of subtitle/captions file. Mux supports [SubRip Text (SRT)](https://en.wikipedia.org/wiki/SubRip) and [Web Video Text Tracks](https://www.w3.org/TR/webvtt1/) format for ingesting Subtitles and Closed Captions. * For Watermarking or Overlay, the url is the location of the watermark image. * When creating clips from existing Mux assets, the url is defined with &#x60;mux://assets/{asset_id}&#x60; template where &#x60;asset_id&#x60; is the Asset Identifier for creating the clip from.  | [optional] 
**OverlaySettings** | [**InputSettingsOverlaySettings**](InputSettings_overlay_settings.md) |  | [optional] 
**StartTime** | **float64** | The time offset in seconds from the beginning of the video indicating the clip&#39;s starting marker. The default value is 0 when not included. This parameter is only applicable for creating clips when &#x60;input.url&#x60; has &#x60;mux://assets/{asset_id}&#x60; format. | [optional] 
**EndTime** | **float64** | The time offset in seconds from the beginning of the video, indicating the clip&#39;s ending marker. The default value is the duration of the video when not included. This parameter is only applicable for creating clips when &#x60;input.url&#x60; has &#x60;mux://assets/{asset_id}&#x60; format. | [optional] 
**Type** | **string** | This parameter is required for the &#x60;text&#x60; track type. | [optional] 
**TextType** | **string** | Type of text track. This parameter only supports subtitles value. For more information on Subtitles / Closed Captions, [see this blog post](https://mux.com/blog/subtitles-captions-webvtt-hls-and-those-magic-flags/). This parameter is required for &#x60;text&#x60; track type. | [optional] 
**LanguageCode** | **string** | The language code value must be a valid [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, en for English or en-US for the US version of English. This parameter is required for text type and subtitles text type track. | [optional] 
**Name** | **string** | The name of the track containing a human-readable description. This value must be unique across all text type and subtitles &#x60;text&#x60; type tracks. The hls manifest will associate a subtitle text track with this value. For example, the value should be \&quot;English\&quot; for subtitles text track with language_code as en. This optional parameter should be used only for &#x60;text&#x60; type and subtitles &#x60;text&#x60; type track. If this parameter is not included, Mux will auto-populate based on the &#x60;input[].language_code&#x60; value. | [optional] 
**ClosedCaptions** | **bool** | Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This optional parameter should be used for &#x60;text&#x60; type and subtitles &#x60;text&#x60; type tracks. | [optional] 
**Passthrough** | **string** | This optional parameter should be used for &#x60;text&#x60; type and subtitles &#x60;text&#x60; type tracks. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


