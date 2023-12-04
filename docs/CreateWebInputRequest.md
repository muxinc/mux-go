# CreateWebInputRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the Web Input. | [optional] 
**CreatedAt** | **string** | Time the Web Input was created, defined as a Unix timestamp (seconds since epoch). | [optional] 
**Url** | **string** | The URL for the Web Input to load. | [optional] 
**AutoLaunch** | **bool** | When set to &#x60;true&#x60; the Web Input will automatically launch and start streaming immediately after creation | [optional] 
**LiveStreamId** | **string** | The Live Stream ID to broadcast this Web Input to | [optional] 
**Status** | **string** |  | [optional] 
**Passthrough** | **string** | Arbitrary metadata that will be included in the Web Input details and related webhooks. Can be used to store your own ID for the Web Input. **Max: 255 characters**. | [optional] 
**Resolution** | **string** | The resolution of the viewport of the Web Input&#39;s browser instance. Defaults to 1920x1080 if not set. | [optional] [default to RESOLUTION__1920X1080]
**Timeout** | **int32** | The number of seconds that the Web Input should stream for before automatically shutting down. | [optional] [default to 3600]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


