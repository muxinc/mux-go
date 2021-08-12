# \AssetsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsset**](AssetsApi.md#CreateAsset) | **Post** /video/v1/assets | Create an asset
[**CreateAssetPlaybackId**](AssetsApi.md#CreateAssetPlaybackId) | **Post** /video/v1/assets/{ASSET_ID}/playback-ids | Create a playback ID
[**CreateAssetTrack**](AssetsApi.md#CreateAssetTrack) | **Post** /video/v1/assets/{ASSET_ID}/tracks | Create an asset track
[**DeleteAsset**](AssetsApi.md#DeleteAsset) | **Delete** /video/v1/assets/{ASSET_ID} | Delete an asset
[**DeleteAssetPlaybackId**](AssetsApi.md#DeleteAssetPlaybackId) | **Delete** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Delete a playback ID
[**DeleteAssetTrack**](AssetsApi.md#DeleteAssetTrack) | **Delete** /video/v1/assets/{ASSET_ID}/tracks/{TRACK_ID} | Delete an asset track
[**GetAsset**](AssetsApi.md#GetAsset) | **Get** /video/v1/assets/{ASSET_ID} | Retrieve an asset
[**GetAssetInputInfo**](AssetsApi.md#GetAssetInputInfo) | **Get** /video/v1/assets/{ASSET_ID}/input-info | Retrieve asset input info
[**GetAssetPlaybackId**](AssetsApi.md#GetAssetPlaybackId) | **Get** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Retrieve a playback ID
[**ListAssets**](AssetsApi.md#ListAssets) | **Get** /video/v1/assets | List assets
[**UpdateAssetMasterAccess**](AssetsApi.md#UpdateAssetMasterAccess) | **Put** /video/v1/assets/{ASSET_ID}/master-access | Update master access
[**UpdateAssetMp4Support**](AssetsApi.md#UpdateAssetMp4Support) | **Put** /video/v1/assets/{ASSET_ID}/mp4-support | Update MP4 support



## CreateAsset

> AssetResponse CreateAsset(ctx).CreateAssetRequest(createAssetRequest).Execute()

Create an asset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createAssetRequest := *openapiclient.NewCreateAssetRequest() // CreateAssetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.CreateAsset(context.Background()).CreateAssetRequest(createAssetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.CreateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAsset`: AssetResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.CreateAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetRequest** | [**CreateAssetRequest**](CreateAssetRequest.md) |  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetPlaybackId

> CreatePlaybackIDResponse CreateAssetPlaybackId(ctx, aSSETID).CreatePlaybackIDRequest(createPlaybackIDRequest).Execute()

Create a playback ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    createPlaybackIDRequest := *openapiclient.NewCreatePlaybackIDRequest() // CreatePlaybackIDRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.CreateAssetPlaybackId(context.Background(), aSSETID).CreatePlaybackIDRequest(createPlaybackIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.CreateAssetPlaybackId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssetPlaybackId`: CreatePlaybackIDResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.CreateAssetPlaybackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetPlaybackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPlaybackIDRequest** | [**CreatePlaybackIDRequest**](CreatePlaybackIDRequest.md) |  | 

### Return type

[**CreatePlaybackIDResponse**](CreatePlaybackIDResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetTrack

> CreateTrackResponse CreateAssetTrack(ctx, aSSETID).CreateTrackRequest(createTrackRequest).Execute()

Create an asset track

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    createTrackRequest := *openapiclient.NewCreateTrackRequest("Url_example", "Type_example", "TextType_example", "LanguageCode_example") // CreateTrackRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.CreateAssetTrack(context.Background(), aSSETID).CreateTrackRequest(createTrackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.CreateAssetTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssetTrack`: CreateTrackResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.CreateAssetTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTrackRequest** | [**CreateTrackRequest**](CreateTrackRequest.md) |  | 

### Return type

[**CreateTrackResponse**](CreateTrackResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAsset

> DeleteAsset(ctx, aSSETID).Execute()

Delete an asset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.DeleteAsset(context.Background(), aSSETID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssetPlaybackId

> DeleteAssetPlaybackId(ctx, aSSETID, pLAYBACKID).Execute()

Delete a playback ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    pLAYBACKID := "pLAYBACKID_example" // string | The live stream's playback ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.DeleteAssetPlaybackId(context.Background(), aSSETID, pLAYBACKID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.DeleteAssetPlaybackId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 
**pLAYBACKID** | **string** | The live stream&#39;s playback ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetPlaybackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssetTrack

> DeleteAssetTrack(ctx, aSSETID, tRACKID).Execute()

Delete an asset track

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    tRACKID := "tRACKID_example" // string | The track ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.DeleteAssetTrack(context.Background(), aSSETID, tRACKID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.DeleteAssetTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 
**tRACKID** | **string** | The track ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> AssetResponse GetAsset(ctx, aSSETID).Execute()

Retrieve an asset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.GetAsset(context.Background(), aSSETID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsset`: AssetResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetInputInfo

> GetAssetInputInfoResponse GetAssetInputInfo(ctx, aSSETID).Execute()

Retrieve asset input info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.GetAssetInputInfo(context.Background(), aSSETID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAssetInputInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetInputInfo`: GetAssetInputInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAssetInputInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetInputInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAssetInputInfoResponse**](GetAssetInputInfoResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetPlaybackId

> GetAssetPlaybackIDResponse GetAssetPlaybackId(ctx, aSSETID, pLAYBACKID).Execute()

Retrieve a playback ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    pLAYBACKID := "pLAYBACKID_example" // string | The live stream's playback ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.GetAssetPlaybackId(context.Background(), aSSETID, pLAYBACKID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAssetPlaybackId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetPlaybackId`: GetAssetPlaybackIDResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAssetPlaybackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 
**pLAYBACKID** | **string** | The live stream&#39;s playback ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetPlaybackIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetPlaybackIDResponse**](GetAssetPlaybackIDResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssets

> ListAssetsResponse ListAssets(ctx).Limit(limit).Page(page).LiveStreamId(liveStreamId).UploadId(uploadId).Execute()

List assets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | Number of items to include in the response (optional) (default to 25)
    page := int32(56) // int32 | Offset by this many pages, of the size of `limit` (optional) (default to 1)
    liveStreamId := "liveStreamId_example" // string | Filter response to return all the assets for this live stream only (optional)
    uploadId := "uploadId_example" // string | Filter response to return an asset created from this direct upload only (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.ListAssets(context.Background()).Limit(limit).Page(page).LiveStreamId(liveStreamId).UploadId(uploadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.ListAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssets`: ListAssetsResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.ListAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **liveStreamId** | **string** | Filter response to return all the assets for this live stream only | 
 **uploadId** | **string** | Filter response to return an asset created from this direct upload only | 

### Return type

[**ListAssetsResponse**](ListAssetsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetMasterAccess

> AssetResponse UpdateAssetMasterAccess(ctx, aSSETID).UpdateAssetMasterAccessRequest(updateAssetMasterAccessRequest).Execute()

Update master access



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    updateAssetMasterAccessRequest := *openapiclient.NewUpdateAssetMasterAccessRequest() // UpdateAssetMasterAccessRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.UpdateAssetMasterAccess(context.Background(), aSSETID).UpdateAssetMasterAccessRequest(updateAssetMasterAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.UpdateAssetMasterAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetMasterAccess`: AssetResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.UpdateAssetMasterAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetMasterAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAssetMasterAccessRequest** | [**UpdateAssetMasterAccessRequest**](UpdateAssetMasterAccessRequest.md) |  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetMp4Support

> AssetResponse UpdateAssetMp4Support(ctx, aSSETID).UpdateAssetMP4SupportRequest(updateAssetMP4SupportRequest).Execute()

Update MP4 support



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    aSSETID := "aSSETID_example" // string | The asset ID.
    updateAssetMP4SupportRequest := *openapiclient.NewUpdateAssetMP4SupportRequest() // UpdateAssetMP4SupportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssetsApi.UpdateAssetMp4Support(context.Background(), aSSETID).UpdateAssetMP4SupportRequest(updateAssetMP4SupportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.UpdateAssetMp4Support``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetMp4Support`: AssetResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.UpdateAssetMp4Support`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aSSETID** | **string** | The asset ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetMp4SupportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAssetMP4SupportRequest** | [**UpdateAssetMP4SupportRequest**](UpdateAssetMP4SupportRequest.md) |  | 

### Return type

[**AssetResponse**](AssetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

