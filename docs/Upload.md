# Upload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the Direct Upload. | [optional] 
**Timeout** | **int32** | Max time in seconds for the signed upload URL to be valid. If a successful upload has not occurred before the timeout limit, the direct upload is marked &#x60;timed_out&#x60; | [optional] [default to 3600]
**Status** | **string** |  | [optional] 
**NewAssetSettings** | [**Asset**](Asset.md) |  | [optional] 
**AssetId** | **string** | Only set once the upload is in the &#x60;asset_created&#x60; state. | [optional] 
**Error** | [**UploadError**](Upload_error.md) |  | [optional] 
**CorsOrigin** | **string** | If the upload URL will be used in a browser, you must specify the origin in order for the signed URL to have the correct CORS headers. | [optional] 
**Url** | **string** | The URL to upload the associated source media to. | [optional] 
**Test** | **bool** | Indicates if this is a test Direct Upload, in which case the Asset that gets created will be a &#x60;test&#x60; Asset. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


