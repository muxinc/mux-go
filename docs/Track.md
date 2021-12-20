# Track

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the Track | [optional] 
**Type** | **string** | The type of track | [optional] 
**Duration** | **float64** | The duration in seconds of the track media. This parameter is not set for the &#x60;text&#x60; type track. This field is optional and may not be set. The top level &#x60;duration&#x60; field of an asset will always be set. | [optional] 
**MaxWidth** | **int64** | The maximum width in pixels available for the track. Only set for the &#x60;video&#x60; type track. | [optional] 
**MaxHeight** | **int64** | The maximum height in pixels available for the track. Only set for the &#x60;video&#x60; type track. | [optional] 
**MaxFrameRate** | **float64** | The maximum frame rate available for the track. Only set for the &#x60;video&#x60; type track. This field may return &#x60;-1&#x60; if the frame rate of the input cannot be reliably determined. | [optional] 
**MaxChannels** | **int64** | The maximum number of audio channels the track supports. Only set for the &#x60;audio&#x60; type track. | [optional] 
**MaxChannelLayout** | **string** | Only set for the &#x60;audio&#x60; type track. | [optional] 
**TextType** | **string** | This parameter is set only for the &#x60;text&#x60; type track. | [optional] 
**LanguageCode** | **string** | The language code value represents [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, &#x60;en&#x60; for English or &#x60;en-US&#x60; for the US version of English. This parameter is set for &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. | [optional] 
**Name** | **string** | The name of the track containing a human-readable description. The hls manifest will associate a subtitle text track with this value. For example, the value is \&quot;English\&quot; for subtitles text track for the &#x60;language_code&#x60; value of &#x60;en-US&#x60;. This parameter is set for the &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. | [optional] 
**ClosedCaptions** | **bool** | Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This parameter is set for the &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. | [optional] 
**Passthrough** | **string** | Arbitrary user-supplied metadata set for the track either when creating the asset or track. This parameter is set for &#x60;text&#x60; type and &#x60;subtitles&#x60; text type track. Max 255 characters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


