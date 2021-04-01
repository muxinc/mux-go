# InputSettingsOverlaySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerticalAlign** | Pointer to **string** | Where the vertical positioning of the overlay/watermark should begin from. Defaults to &#x60;\&quot;top\&quot;&#x60; | [optional] 
**VerticalMargin** | Pointer to **string** | The distance from the vertical_align starting point and the image&#39;s closest edge. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). Negative values will move the overlay offscreen. In the case of &#39;middle&#39;, a positive value will shift the overlay towards the bottom and and a negative value will shift it towards the top. | [optional] 
**HorizontalAlign** | Pointer to **string** | Where the horizontal positioning of the overlay/watermark should begin from. | [optional] 
**HorizontalMargin** | Pointer to **string** | The distance from the horizontal_align starting point and the image&#39;s closest edge. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). Negative values will move the overlay offscreen. In the case of &#39;center&#39;, a positive value will shift the image towards the right and and a negative value will shift it towards the left. | [optional] 
**Width** | Pointer to **string** | How wide the overlay should appear. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). If both width and height are left blank the width will be the true pixels of the image, applied as if the video has been scaled to fit a 1920x1080 frame. If height is supplied with no width, the width will scale proportionally to the height. | [optional] 
**Height** | Pointer to **string** | How tall the overlay should appear. Can be expressed as a percent (\&quot;10%\&quot;) or as a pixel value (\&quot;100px\&quot;). If both width and height are left blank the height will be the true pixels of the image, applied as if the video has been scaled to fit a 1920x1080 frame. If width is supplied with no height, the height will scale proportionally to the width. | [optional] 
**Opacity** | Pointer to **string** | How opaque the overlay should appear, expressed as a percent. (Default 100%) | [optional] 

## Methods

### NewInputSettingsOverlaySettings

`func NewInputSettingsOverlaySettings() *InputSettingsOverlaySettings`

NewInputSettingsOverlaySettings instantiates a new InputSettingsOverlaySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSettingsOverlaySettingsWithDefaults

`func NewInputSettingsOverlaySettingsWithDefaults() *InputSettingsOverlaySettings`

NewInputSettingsOverlaySettingsWithDefaults instantiates a new InputSettingsOverlaySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerticalAlign

`func (o *InputSettingsOverlaySettings) GetVerticalAlign() string`

GetVerticalAlign returns the VerticalAlign field if non-nil, zero value otherwise.

### GetVerticalAlignOk

`func (o *InputSettingsOverlaySettings) GetVerticalAlignOk() (*string, bool)`

GetVerticalAlignOk returns a tuple with the VerticalAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalAlign

`func (o *InputSettingsOverlaySettings) SetVerticalAlign(v string)`

SetVerticalAlign sets VerticalAlign field to given value.

### HasVerticalAlign

`func (o *InputSettingsOverlaySettings) HasVerticalAlign() bool`

HasVerticalAlign returns a boolean if a field has been set.

### GetVerticalMargin

`func (o *InputSettingsOverlaySettings) GetVerticalMargin() string`

GetVerticalMargin returns the VerticalMargin field if non-nil, zero value otherwise.

### GetVerticalMarginOk

`func (o *InputSettingsOverlaySettings) GetVerticalMarginOk() (*string, bool)`

GetVerticalMarginOk returns a tuple with the VerticalMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalMargin

`func (o *InputSettingsOverlaySettings) SetVerticalMargin(v string)`

SetVerticalMargin sets VerticalMargin field to given value.

### HasVerticalMargin

`func (o *InputSettingsOverlaySettings) HasVerticalMargin() bool`

HasVerticalMargin returns a boolean if a field has been set.

### GetHorizontalAlign

`func (o *InputSettingsOverlaySettings) GetHorizontalAlign() string`

GetHorizontalAlign returns the HorizontalAlign field if non-nil, zero value otherwise.

### GetHorizontalAlignOk

`func (o *InputSettingsOverlaySettings) GetHorizontalAlignOk() (*string, bool)`

GetHorizontalAlignOk returns a tuple with the HorizontalAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAlign

`func (o *InputSettingsOverlaySettings) SetHorizontalAlign(v string)`

SetHorizontalAlign sets HorizontalAlign field to given value.

### HasHorizontalAlign

`func (o *InputSettingsOverlaySettings) HasHorizontalAlign() bool`

HasHorizontalAlign returns a boolean if a field has been set.

### GetHorizontalMargin

`func (o *InputSettingsOverlaySettings) GetHorizontalMargin() string`

GetHorizontalMargin returns the HorizontalMargin field if non-nil, zero value otherwise.

### GetHorizontalMarginOk

`func (o *InputSettingsOverlaySettings) GetHorizontalMarginOk() (*string, bool)`

GetHorizontalMarginOk returns a tuple with the HorizontalMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalMargin

`func (o *InputSettingsOverlaySettings) SetHorizontalMargin(v string)`

SetHorizontalMargin sets HorizontalMargin field to given value.

### HasHorizontalMargin

`func (o *InputSettingsOverlaySettings) HasHorizontalMargin() bool`

HasHorizontalMargin returns a boolean if a field has been set.

### GetWidth

`func (o *InputSettingsOverlaySettings) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InputSettingsOverlaySettings) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InputSettingsOverlaySettings) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InputSettingsOverlaySettings) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *InputSettingsOverlaySettings) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InputSettingsOverlaySettings) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InputSettingsOverlaySettings) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InputSettingsOverlaySettings) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetOpacity

`func (o *InputSettingsOverlaySettings) GetOpacity() string`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *InputSettingsOverlaySettings) GetOpacityOk() (*string, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *InputSettingsOverlaySettings) SetOpacity(v string)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *InputSettingsOverlaySettings) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


