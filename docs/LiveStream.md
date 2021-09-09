# LiveStream

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the Live Stream. Max 255 characters. | [optional] 
**CreatedAt** | **string** | Time the Live Stream was created, defined as a Unix timestamp (seconds since epoch). | [optional] 
**StreamKey** | **string** | Unique key used for streaming to a Mux RTMP endpoint. This should be considered as sensitive as credentials, anyone with this stream key can begin streaming. | [optional] 
**ActiveAssetId** | **string** | The Asset that is currently being created if there is an active broadcast. | [optional] 
**RecentAssetIds** | **[]string** | An array of strings with the most recent Assets that were created from this live stream. | [optional] 
**Status** | **string** | &#x60;idle&#x60; indicates that there is no active broadcast. &#x60;active&#x60; indicates that there is an active broadcast and &#x60;disabled&#x60; status indicates that no future RTMP streams can be published. | [optional] 
**PlaybackIds** | [**[]PlaybackId**](PlaybackID.md) | An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/video/play-your-videos) for more details. | [optional] 
**NewAssetSettings** | [**CreateAssetRequest**](CreateAssetRequest.md) |  | [optional] 
**Passthrough** | **string** | Arbitrary metadata set for the asset. Max 255 characters. | [optional] 
**AudioOnly** | **bool** | The live stream only processes the audio track if the value is set to true. Mux drops the video track if broadcasted. | [optional] 
**ReconnectWindow** | **float32** | When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. **Min**: 0.1s. **Max**: 300s (5 minutes). | [optional] [default to 60]
**ReducedLatency** | **bool** | Latency is the time from when the streamer does something in real life to when you see it happen in the player. Set this if you want lower latency for your live stream. **Note**: Reconnect windows are incompatible with Reduced Latency and will always be set to zero (0) seconds. See the [Reduce live stream latency guide](https://docs.mux.com/guides/video/reduce-live-stream-latency) to understand the tradeoffs. | [optional] 
**LowLatency** | **bool** | Latency is the time from when the streamer does something in real life to when you see it happen in the player. Setting this option will enable compatibility with the LL-HLS specification for low-latency streaming. This typically has lower latency than Reduced Latency streams, and cannot be combined with Reduced Latency. Note: Reconnect windows are incompatible with Low Latency and will always be set to zero (0) seconds. | [optional] 
**SimulcastTargets** | [**[]SimulcastTarget**](SimulcastTarget.md) | Each Simulcast Target contains configuration details to broadcast (or \&quot;restream\&quot;) a live stream to a third-party streaming service. [See the Stream live to 3rd party platforms guide](https://docs.mux.com/guides/video/stream-live-to-3rd-party-platforms). | [optional] 
**Test** | **bool** | True means this live stream is a test live stream. Test live streams can be used to help evaluate the Mux Video APIs for free. There is no limit on the number of test live streams, but they are watermarked with the Mux logo, and limited to 5 minutes. The test live stream is disabled after the stream is active for 5 mins and the recorded asset also deleted after 24 hours. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


