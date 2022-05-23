# Space

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the space. Max 255 characters. | 
**CreatedAt** | **string** | Time the space was created, defined as a Unix timestamp (seconds since epoch). | 
**Type** | [**SpaceType**](SpaceType.md) |  | 
**Status** | [**SpaceStatus**](SpaceStatus.md) |  | 
**Passthrough** | **string** | Arbitrary user-supplied metadata that will be included in the space details and related webhooks. Max: 255 characters. | [optional] 
**Broadcasts** | [**[]Broadcast**](Broadcast.md) | An array of broadcast destinations. | [optional] 
**ActiveSessionId** | **string** | Unique identifier for the current lifecycle of the space. Only set when the space is &#x60;active&#x60; and is set to a new value each time the space transitions from &#x60;idle&#x60; to &#x60;active&#x60;. This value is useful for logging and debugging issues. Max 255 characters. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


