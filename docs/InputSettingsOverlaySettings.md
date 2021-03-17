# InputSettingsOverlaySettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerticalAlign** | **string** | Where the vertical positioning of the overlay/watermark should begin from. Defaults to &#x60;\&quot;top\&quot;&#x60; | [optional] 
**VerticalMargin** | **string** | The distance from the vertical_align starting point and the image&#39;s closest edge. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). Negative values will move the overlay offscreen. In the case of &#39;middle&#39;, a positive value will shift the overlay towards the bottom and and a negative value will shift it towards the top. | [optional] 
**HorizontalAlign** | **string** | Where the horizontal positioning of the overlay/watermark should begin from. | [optional] 
**HorizontalMargin** | **string** | The distance from the horizontal_align starting point and the image&#39;s closest edge. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). Negative values will move the overlay offscreen. In the case of &#39;center&#39;, a positive value will shift the image towards the right and and a negative value will shift it towards the left. | [optional] 
**Width** | **string** | How wide the overlay should appear. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). If both width and height are left blank the width will be the true pixels of the image, applied as if the video has been scaled to fit a 1920x1080 frame. If height is supplied with no width, the width will scale proportionally to the height. | [optional] 
**Height** | **string** | How tall the overlay should appear. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). If both width and height are left blank the height will be the true pixels of the image, applied as if the video has been scaled to fit a 1920x1080 frame. If width is supplied with no height, the height will scale proportionally to the width. | [optional] 
**Opacity** | **string** | How opaque the overlay should appear, expressed as a percent. (Default 100%) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


