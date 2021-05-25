# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the Direct Upload. | [optional] 
**Timeout** | Pointer to **int32** | Max time in seconds for the signed upload URL to be valid. If a successful upload has not occurred before the timeout limit, the direct upload is marked &#x60;timed_out&#x60; | [optional] [default to 3600]
**Status** | Pointer to **string** |  | [optional] 
**NewAssetSettings** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**AssetId** | Pointer to **string** | Only set once the upload is in the &#x60;asset_created&#x60; state. | [optional] 
**Error** | Pointer to [**UploadError**](Upload_error.md) |  | [optional] 
**CorsOrigin** | Pointer to **string** | If the upload URL will be used in a browser, you must specify the origin in order for the signed URL to have the correct CORS headers. | [optional] 
**Url** | Pointer to **string** | The URL to upload the associated source media to. | [optional] 
**Test** | Pointer to **bool** | Indicates if this is a test Direct Upload, in which case the Asset that gets created will be a &#x60;test&#x60; Asset. | [optional] 

## Methods

### NewUpload

`func NewUpload() *Upload`

NewUpload instantiates a new Upload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWithDefaults

`func NewUploadWithDefaults() *Upload`

NewUploadWithDefaults instantiates a new Upload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Upload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Upload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Upload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Upload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimeout

`func (o *Upload) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Upload) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Upload) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Upload) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetStatus

`func (o *Upload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Upload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Upload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Upload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNewAssetSettings

`func (o *Upload) GetNewAssetSettings() Asset`

GetNewAssetSettings returns the NewAssetSettings field if non-nil, zero value otherwise.

### GetNewAssetSettingsOk

`func (o *Upload) GetNewAssetSettingsOk() (*Asset, bool)`

GetNewAssetSettingsOk returns a tuple with the NewAssetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssetSettings

`func (o *Upload) SetNewAssetSettings(v Asset)`

SetNewAssetSettings sets NewAssetSettings field to given value.

### HasNewAssetSettings

`func (o *Upload) HasNewAssetSettings() bool`

HasNewAssetSettings returns a boolean if a field has been set.

### GetAssetId

`func (o *Upload) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Upload) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Upload) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Upload) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetError

`func (o *Upload) GetError() UploadError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Upload) GetErrorOk() (*UploadError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Upload) SetError(v UploadError)`

SetError sets Error field to given value.

### HasError

`func (o *Upload) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCorsOrigin

`func (o *Upload) GetCorsOrigin() string`

GetCorsOrigin returns the CorsOrigin field if non-nil, zero value otherwise.

### GetCorsOriginOk

`func (o *Upload) GetCorsOriginOk() (*string, bool)`

GetCorsOriginOk returns a tuple with the CorsOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigin

`func (o *Upload) SetCorsOrigin(v string)`

SetCorsOrigin sets CorsOrigin field to given value.

### HasCorsOrigin

`func (o *Upload) HasCorsOrigin() bool`

HasCorsOrigin returns a boolean if a field has been set.

### GetUrl

`func (o *Upload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Upload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Upload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Upload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTest

`func (o *Upload) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Upload) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Upload) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *Upload) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


