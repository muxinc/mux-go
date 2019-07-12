# CreateLiveStreamRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaybackPolicy** | [**[]PlaybackPolicy**](PlaybackPolicy.md) |  | [optional] 
**NewAssetSettings** | [**CreateAssetRequest**](CreateAssetRequest.md) |  | [optional] 
**ReconnectWindow** | **float32** | When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. Default: 60 seconds | [optional] [default to 60]
**Passthrough** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


