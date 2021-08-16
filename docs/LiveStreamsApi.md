# \LiveStreamsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLiveStream**](LiveStreamsApi.md#CreateLiveStream) | **Post** /video/v1/live-streams | Create a live stream
[**CreateLiveStreamPlaybackId**](LiveStreamsApi.md#CreateLiveStreamPlaybackId) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids | Create a live stream playback ID
[**CreateLiveStreamSimulcastTarget**](LiveStreamsApi.md#CreateLiveStreamSimulcastTarget) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets | Create a live stream simulcast target
[**DeleteLiveStream**](LiveStreamsApi.md#DeleteLiveStream) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID} | Delete a live stream
[**DeleteLiveStreamPlaybackId**](LiveStreamsApi.md#DeleteLiveStreamPlaybackId) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID} | Delete a live stream playback ID
[**DeleteLiveStreamSimulcastTarget**](LiveStreamsApi.md#DeleteLiveStreamSimulcastTarget) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID} | Delete a Live Stream Simulcast Target
[**DisableLiveStream**](LiveStreamsApi.md#DisableLiveStream) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/disable | Disable a live stream
[**EnableLiveStream**](LiveStreamsApi.md#EnableLiveStream) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/enable | Enable a live stream
[**GetLiveStream**](LiveStreamsApi.md#GetLiveStream) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID} | Retrieve a live stream
[**GetLiveStreamSimulcastTarget**](LiveStreamsApi.md#GetLiveStreamSimulcastTarget) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID} | Retrieve a Live Stream Simulcast Target
[**ListLiveStreams**](LiveStreamsApi.md#ListLiveStreams) | **Get** /video/v1/live-streams | List live streams
[**ResetStreamKey**](LiveStreamsApi.md#ResetStreamKey) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/reset-stream-key | Reset a live stream’s stream key
[**SignalLiveStreamComplete**](LiveStreamsApi.md#SignalLiveStreamComplete) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/complete | Signal a live stream is finished



## CreateLiveStream

> LiveStreamResponse CreateLiveStream(ctx).CreateLiveStreamRequest(createLiveStreamRequest).Execute()

Create a live stream

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
    createLiveStreamRequest := *openapiclient.NewCreateLiveStreamRequest() // CreateLiveStreamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.CreateLiveStream(context.Background()).CreateLiveStreamRequest(createLiveStreamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.CreateLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStream`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.CreateLiveStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLiveStreamRequest** | [**CreateLiveStreamRequest**](CreateLiveStreamRequest.md) |  | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLiveStreamPlaybackId

> CreatePlaybackIDResponse CreateLiveStreamPlaybackId(ctx, lIVESTREAMID).CreatePlaybackIDRequest(createPlaybackIDRequest).Execute()

Create a live stream playback ID

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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID
    createPlaybackIDRequest := *openapiclient.NewCreatePlaybackIDRequest() // CreatePlaybackIDRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.CreateLiveStreamPlaybackId(context.Background(), lIVESTREAMID).CreatePlaybackIDRequest(createPlaybackIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.CreateLiveStreamPlaybackId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamPlaybackId`: CreatePlaybackIDResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.CreateLiveStreamPlaybackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamPlaybackIdRequest struct via the builder pattern


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


## CreateLiveStreamSimulcastTarget

> SimulcastTargetResponse CreateLiveStreamSimulcastTarget(ctx, lIVESTREAMID).CreateSimulcastTargetRequest(createSimulcastTargetRequest).Execute()

Create a live stream simulcast target



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID
    createSimulcastTargetRequest := *openapiclient.NewCreateSimulcastTargetRequest("Url_example") // CreateSimulcastTargetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.CreateLiveStreamSimulcastTarget(context.Background(), lIVESTREAMID).CreateSimulcastTargetRequest(createSimulcastTargetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.CreateLiveStreamSimulcastTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamSimulcastTarget`: SimulcastTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.CreateLiveStreamSimulcastTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamSimulcastTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSimulcastTargetRequest** | [**CreateSimulcastTargetRequest**](CreateSimulcastTargetRequest.md) |  | 

### Return type

[**SimulcastTargetResponse**](SimulcastTargetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLiveStream

> DeleteLiveStream(ctx, lIVESTREAMID).Execute()

Delete a live stream

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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.DeleteLiveStream(context.Background(), lIVESTREAMID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.DeleteLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLiveStreamRequest struct via the builder pattern


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


## DeleteLiveStreamPlaybackId

> DeleteLiveStreamPlaybackId(ctx, lIVESTREAMID, pLAYBACKID).Execute()

Delete a live stream playback ID

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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID
    pLAYBACKID := "pLAYBACKID_example" // string | The live stream's playback ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.DeleteLiveStreamPlaybackId(context.Background(), lIVESTREAMID, pLAYBACKID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.DeleteLiveStreamPlaybackId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 
**pLAYBACKID** | **string** | The live stream&#39;s playback ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLiveStreamPlaybackIdRequest struct via the builder pattern


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


## DeleteLiveStreamSimulcastTarget

> DeleteLiveStreamSimulcastTarget(ctx, lIVESTREAMID, sIMULCASTTARGETID).Execute()

Delete a Live Stream Simulcast Target



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID
    sIMULCASTTARGETID := "sIMULCASTTARGETID_example" // string | The ID of the simulcast target.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.DeleteLiveStreamSimulcastTarget(context.Background(), lIVESTREAMID, sIMULCASTTARGETID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.DeleteLiveStreamSimulcastTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 
**sIMULCASTTARGETID** | **string** | The ID of the simulcast target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLiveStreamSimulcastTargetRequest struct via the builder pattern


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


## DisableLiveStream

> DisableLiveStreamResponse DisableLiveStream(ctx, lIVESTREAMID).Execute()

Disable a live stream



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.DisableLiveStream(context.Background(), lIVESTREAMID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.DisableLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableLiveStream`: DisableLiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.DisableLiveStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisableLiveStreamResponse**](DisableLiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableLiveStream

> EnableLiveStreamResponse EnableLiveStream(ctx, lIVESTREAMID).Execute()

Enable a live stream



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.EnableLiveStream(context.Background(), lIVESTREAMID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.EnableLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableLiveStream`: EnableLiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.EnableLiveStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnableLiveStreamResponse**](EnableLiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStream

> LiveStreamResponse GetLiveStream(ctx, lIVESTREAMID).Execute()

Retrieve a live stream



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.GetLiveStream(context.Background(), lIVESTREAMID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.GetLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiveStream`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.GetLiveStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamSimulcastTarget

> SimulcastTargetResponse GetLiveStreamSimulcastTarget(ctx, lIVESTREAMID, sIMULCASTTARGETID).Execute()

Retrieve a Live Stream Simulcast Target



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID
    sIMULCASTTARGETID := "sIMULCASTTARGETID_example" // string | The ID of the simulcast target.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.GetLiveStreamSimulcastTarget(context.Background(), lIVESTREAMID, sIMULCASTTARGETID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.GetLiveStreamSimulcastTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiveStreamSimulcastTarget`: SimulcastTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.GetLiveStreamSimulcastTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 
**sIMULCASTTARGETID** | **string** | The ID of the simulcast target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveStreamSimulcastTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SimulcastTargetResponse**](SimulcastTargetResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLiveStreams

> ListLiveStreamsResponse ListLiveStreams(ctx).Limit(limit).Page(page).StreamKey(streamKey).Execute()

List live streams

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
    streamKey := "streamKey_example" // string | Filter response to return live stream for this stream key only (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.ListLiveStreams(context.Background()).Limit(limit).Page(page).StreamKey(streamKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.ListLiveStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLiveStreams`: ListLiveStreamsResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.ListLiveStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLiveStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items to include in the response | [default to 25]
 **page** | **int32** | Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **streamKey** | **string** | Filter response to return live stream for this stream key only | 

### Return type

[**ListLiveStreamsResponse**](ListLiveStreamsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetStreamKey

> LiveStreamResponse ResetStreamKey(ctx, lIVESTREAMID).Execute()

Reset a live stream’s stream key



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.ResetStreamKey(context.Background(), lIVESTREAMID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.ResetStreamKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetStreamKey`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.ResetStreamKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetStreamKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalLiveStreamComplete

> SignalLiveStreamCompleteResponse SignalLiveStreamComplete(ctx, lIVESTREAMID).Execute()

Signal a live stream is finished



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
    lIVESTREAMID := "lIVESTREAMID_example" // string | The live stream ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveStreamsApi.SignalLiveStreamComplete(context.Background(), lIVESTREAMID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreamsApi.SignalLiveStreamComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignalLiveStreamComplete`: SignalLiveStreamCompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreamsApi.SignalLiveStreamComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lIVESTREAMID** | **string** | The live stream ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignalLiveStreamCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignalLiveStreamCompleteResponse**](SignalLiveStreamCompleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

