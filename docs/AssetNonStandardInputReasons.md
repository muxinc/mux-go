# AssetNonStandardInputReasons

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoCodec** | **string** | The video codec used on the input file. For example, the input file encoded with &#x60;hevc&#x60; video codec is non-standard and the value of this parameter is &#x60;hevc&#x60;. | [optional] 
**AudioCodec** | **string** | The audio codec used on the input file. Non-AAC audio codecs are non-standard. | [optional] 
**VideoGopSize** | **string** | The video key frame Interval (also called as Group of Picture or GOP) of the input file is &#x60;high&#x60;. This parameter is present when the gop is greater than 10 seconds. | [optional] 
**VideoFrameRate** | **string** | The video frame rate of the input file. Video with average frames per second (fps) less than 10 or greater than 120 is non-standard. A &#x60;-1&#x60; frame rate value indicates Mux could not determine the frame rate of the video track. | [optional] 
**VideoResolution** | **string** | The video resolution of the input file. Video resolution higher than 2048 pixels on any one dimension (height or width) is considered non-standard, The resolution value is presented as &#x60;width&#x60; x &#x60;height&#x60; in pixels. | [optional] 
**VideoBitrate** | **string** | The video bitrate of the input file is &#x60;high&#x60;. This parameter is present when the average bitrate of any key frame interval (also known as Group of Pictures or GOP) is higher than what&#39;s considered standard which typically is 16 Mbps. | [optional] 
**PixelAspectRatio** | **string** | The video pixel aspect ratio of the input file. | [optional] 
**VideoEditList** | **string** | Video Edit List reason indicates that the input file&#39;s video track contains a complex Edit Decision List. | [optional] 
**AudioEditList** | **string** | Audio Edit List reason indicates that the input file&#39;s audio track contains a complex Edit Decision List. | [optional] 
**UnexpectedMediaFileParameters** | **string** | A catch-all reason when the input file in created with non-standard encoding parameters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


