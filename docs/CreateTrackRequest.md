# CreateTrackRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the file that Mux should download and use. * For &#x60;audio&#x60; tracks, the URL is the location of the audio file for Mux to download, for example an M4A, WAV, or MP3 file. Mux supports most audio file formats and codecs, but for fastest processing, you should [use standard inputs wherever possible](https://docs.mux.com/guides/minimize-processing-time). * For &#x60;text&#x60; tracks, the URL is the location of subtitle/captions file. Mux supports [SubRip Text (SRT)](https://en.wikipedia.org/wiki/SubRip) and [Web Video Text Tracks](https://www.w3.org/TR/webvtt1/) formats for ingesting Subtitles and Closed Captions.  | [optional] 
**Type** | **string** |  | [optional] 
**TextType** | **string** |  | [optional] 
**LanguageCode** | **string** | The language code value must be a valid BCP 47 specification compliant value. For example, en for English or en-US for the US version of English. | [optional] 
**Name** | **string** | The name of the track containing a human-readable description. This value must be unique within each group of &#x60;text&#x60; or &#x60;audio&#x60; track types. The HLS manifest will associate the &#x60;text&#x60; or &#x60;audio&#x60; track with this value. For example, set the value to \&quot;English\&quot; for subtitles text track with &#x60;language_code&#x60; as en-US. If this parameter is not included, Mux will auto-populate a value based on the &#x60;language_code&#x60; value. | [optional] 
**ClosedCaptions** | **bool** | Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). | [optional] 
**Passthrough** | **string** | Arbitrary user-supplied metadata set for the track either when creating the asset or track. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


