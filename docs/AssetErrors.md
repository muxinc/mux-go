# AssetErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of error that occurred for this asset. | [optional] 
**Messages** | Pointer to **[]string** | Error messages with more details. | [optional] 

## Methods

### NewAssetErrors

`func NewAssetErrors() *AssetErrors`

NewAssetErrors instantiates a new AssetErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetErrorsWithDefaults

`func NewAssetErrorsWithDefaults() *AssetErrors`

NewAssetErrorsWithDefaults instantiates a new AssetErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssetErrors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetErrors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetErrors) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssetErrors) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessages

`func (o *AssetErrors) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AssetErrors) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AssetErrors) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AssetErrors) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


