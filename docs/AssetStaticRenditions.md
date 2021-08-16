# AssetStaticRenditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Indicates the status of downloadable MP4 versions of this asset. | [optional] [default to "disabled"]
**Files** | Pointer to [**[]AssetStaticRenditionsFiles**](AssetStaticRenditionsFiles.md) | Array of file objects. | [optional] 

## Methods

### NewAssetStaticRenditions

`func NewAssetStaticRenditions() *AssetStaticRenditions`

NewAssetStaticRenditions instantiates a new AssetStaticRenditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetStaticRenditionsWithDefaults

`func NewAssetStaticRenditionsWithDefaults() *AssetStaticRenditions`

NewAssetStaticRenditionsWithDefaults instantiates a new AssetStaticRenditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AssetStaticRenditions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetStaticRenditions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetStaticRenditions) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetStaticRenditions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFiles

`func (o *AssetStaticRenditions) GetFiles() []AssetStaticRenditionsFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AssetStaticRenditions) GetFilesOk() (*[]AssetStaticRenditionsFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AssetStaticRenditions) SetFiles(v []AssetStaticRenditionsFiles)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AssetStaticRenditions) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


