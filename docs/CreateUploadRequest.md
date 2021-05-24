# CreateUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | Max time in seconds for the signed upload URL to be valid. If a successful upload has not occurred before the timeout limit, the direct upload is marked &#x60;timed_out&#x60; | [optional] [default to 3600]
**CorsOrigin** | Pointer to **string** | If the upload URL will be used in a browser, you must specify the origin in order for the signed URL to have the correct CORS headers. | [optional] 
**NewAssetSettings** | [**CreateAssetRequest**](CreateAssetRequest.md) |  | 
**Test** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateUploadRequest

`func NewCreateUploadRequest(newAssetSettings CreateAssetRequest, ) *CreateUploadRequest`

NewCreateUploadRequest instantiates a new CreateUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUploadRequestWithDefaults

`func NewCreateUploadRequestWithDefaults() *CreateUploadRequest`

NewCreateUploadRequestWithDefaults instantiates a new CreateUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *CreateUploadRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateUploadRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateUploadRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CreateUploadRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetCorsOrigin

`func (o *CreateUploadRequest) GetCorsOrigin() string`

GetCorsOrigin returns the CorsOrigin field if non-nil, zero value otherwise.

### GetCorsOriginOk

`func (o *CreateUploadRequest) GetCorsOriginOk() (*string, bool)`

GetCorsOriginOk returns a tuple with the CorsOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigin

`func (o *CreateUploadRequest) SetCorsOrigin(v string)`

SetCorsOrigin sets CorsOrigin field to given value.

### HasCorsOrigin

`func (o *CreateUploadRequest) HasCorsOrigin() bool`

HasCorsOrigin returns a boolean if a field has been set.

### GetNewAssetSettings

`func (o *CreateUploadRequest) GetNewAssetSettings() CreateAssetRequest`

GetNewAssetSettings returns the NewAssetSettings field if non-nil, zero value otherwise.

### GetNewAssetSettingsOk

`func (o *CreateUploadRequest) GetNewAssetSettingsOk() (*CreateAssetRequest, bool)`

GetNewAssetSettingsOk returns a tuple with the NewAssetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAssetSettings

`func (o *CreateUploadRequest) SetNewAssetSettings(v CreateAssetRequest)`

SetNewAssetSettings sets NewAssetSettings field to given value.


### GetTest

`func (o *CreateUploadRequest) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CreateUploadRequest) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CreateUploadRequest) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CreateUploadRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


