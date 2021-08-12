# CreateTrackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Type** | **string** |  | 
**TextType** | **string** |  | 
**LanguageCode** | **string** | The language code value must be a valid BCP 47 specification compliant value. For example, en for English or en-US for the US version of English. | 
**Name** | Pointer to **string** | The name of the track containing a human-readable description. This value must be unique across all the text type and subtitles text type tracks. HLS manifest will associate subtitle text track with this value. For example, set the value to \&quot;English\&quot; for subtitles text track with language_code as en-US. If this parameter is not included, Mux will auto-populate based on the language_code value. | [optional] 
**ClosedCaptions** | Pointer to **bool** | Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). | [optional] 
**Passthrough** | Pointer to **string** | Arbitrary metadata set for the track either when creating the asset or track. | [optional] 

## Methods

### NewCreateTrackRequest

`func NewCreateTrackRequest(url string, type_ string, textType string, languageCode string, ) *CreateTrackRequest`

NewCreateTrackRequest instantiates a new CreateTrackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTrackRequestWithDefaults

`func NewCreateTrackRequestWithDefaults() *CreateTrackRequest`

NewCreateTrackRequestWithDefaults instantiates a new CreateTrackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateTrackRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateTrackRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateTrackRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *CreateTrackRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTrackRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTrackRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTextType

`func (o *CreateTrackRequest) GetTextType() string`

GetTextType returns the TextType field if non-nil, zero value otherwise.

### GetTextTypeOk

`func (o *CreateTrackRequest) GetTextTypeOk() (*string, bool)`

GetTextTypeOk returns a tuple with the TextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextType

`func (o *CreateTrackRequest) SetTextType(v string)`

SetTextType sets TextType field to given value.


### GetLanguageCode

`func (o *CreateTrackRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *CreateTrackRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *CreateTrackRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.


### GetName

`func (o *CreateTrackRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTrackRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTrackRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTrackRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClosedCaptions

`func (o *CreateTrackRequest) GetClosedCaptions() bool`

GetClosedCaptions returns the ClosedCaptions field if non-nil, zero value otherwise.

### GetClosedCaptionsOk

`func (o *CreateTrackRequest) GetClosedCaptionsOk() (*bool, bool)`

GetClosedCaptionsOk returns a tuple with the ClosedCaptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedCaptions

`func (o *CreateTrackRequest) SetClosedCaptions(v bool)`

SetClosedCaptions sets ClosedCaptions field to given value.

### HasClosedCaptions

`func (o *CreateTrackRequest) HasClosedCaptions() bool`

HasClosedCaptions returns a boolean if a field has been set.

### GetPassthrough

`func (o *CreateTrackRequest) GetPassthrough() string`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *CreateTrackRequest) GetPassthroughOk() (*string, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *CreateTrackRequest) SetPassthrough(v string)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *CreateTrackRequest) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


