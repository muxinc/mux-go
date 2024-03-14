# SimulcastTarget

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the Simulcast Target | [optional] 
**Passthrough** | **string** | Arbitrary user-supplied metadata set when creating a simulcast target. | [optional] 
**Status** | **string** | The current status of the simulcast target. See Statuses below for detailed description.   * &#x60;idle&#x60;: Default status. When the parent live stream is in disconnected status, simulcast targets will be idle state.   * &#x60;starting&#x60;: The simulcast target transitions into this state when the parent live stream transition into connected state.   * &#x60;broadcasting&#x60;: The simulcast target has successfully connected to the third party live streaming service and is pushing video to that service.   * &#x60;errored&#x60;: The simulcast target encountered an error either while attempting to connect to the third party live streaming service, or mid-broadcasting. When a simulcast target has this status it will have an &#x60;error_severity&#x60; field with more details about the error.  | [optional] 
**StreamKey** | **string** | Stream Key represents a stream identifier on the third party live streaming service to send the parent live stream to. Only used for RTMP(s) simulcast destinations. | [optional] 
**Url** | **string** | The RTMP(s) or SRT endpoint for a simulcast destination. * For RTMP(s) destinations, this should include the application name for the third party live streaming service, for example: &#x60;rtmp://live.example.com/app&#x60;. * For SRT destinations, this should be a fully formed SRT connection string, for example: &#x60;srt://srt-live.example.com:1234?streamid&#x3D;{stream_key}&amp;passphrase&#x3D;{srt_passphrase}&#x60;.  Note: SRT simulcast targets can only be used when an source is connected over SRT.  | [optional] 
**ErrorSeverity** | **string** | The severity of the error encountered by the simulcast target. This field is only set when the simulcast target is in the &#x60;errored&#x60; status. See the values of severities below and their descriptions.   * &#x60;normal&#x60;: The simulcast target encountered an error either while attempting to connect to the third party live streaming service, or mid-broadcasting. A simulcast may transition back into the broadcasting state if a connection with the service can be re-established.   * &#x60;fatal&#x60;: The simulcast target is incompatible with the current input to the parent live stream. No further attempts to this simulcast target will be made for the current live stream asset.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


