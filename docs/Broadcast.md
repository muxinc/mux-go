# Broadcast

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the broadcast. Max 255 characters. | 
**Passthrough** | **string** | Arbitrary user-supplied metadata that will be included in the broadcast details and related webhooks. Max: 255 characters. | [optional] 
**LiveStreamId** | **string** | The ID of the live stream that the broadcast will be sent to. | 
**Status** | [**BroadcastStatus**](BroadcastStatus.md) |  | 
**Layout** | [**BroadcastLayout**](BroadcastLayout.md) |  | 
**Background** | **string** | URL of an image to display as the background of the broadcast. Its dimensions should match the provided resolution. | [optional] 
**Resolution** | [**BroadcastResolution**](BroadcastResolution.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


