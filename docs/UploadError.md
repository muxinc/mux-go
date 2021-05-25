# UploadError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Label for the specific error | [optional] 
**Message** | Pointer to **string** | Human readable error message | [optional] 

## Methods

### NewUploadError

`func NewUploadError() *UploadError`

NewUploadError instantiates a new UploadError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadErrorWithDefaults

`func NewUploadErrorWithDefaults() *UploadError`

NewUploadErrorWithDefaults instantiates a new UploadError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UploadError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UploadError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *UploadError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UploadError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UploadError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UploadError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


