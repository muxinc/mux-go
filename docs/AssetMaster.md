# AssetMaster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The temporary URL to the master version of the video, as an MP4 file. This URL will expire after 24 hours. | [optional] 

## Methods

### NewAssetMaster

`func NewAssetMaster() *AssetMaster`

NewAssetMaster instantiates a new AssetMaster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetMasterWithDefaults

`func NewAssetMasterWithDefaults() *AssetMaster`

NewAssetMasterWithDefaults instantiates a new AssetMaster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AssetMaster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetMaster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetMaster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetMaster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *AssetMaster) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AssetMaster) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AssetMaster) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AssetMaster) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


