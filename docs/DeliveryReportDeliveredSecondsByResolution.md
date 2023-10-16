# DeliveryReportDeliveredSecondsByResolution

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier2160p** | **float64** | Total number of delivered seconds during this time window that had a resolution larger than the 1440p tier (over 4,194,304 pixels total). | [optional] 
**Tier1440p** | **float64** | Total number of delivered seconds during this time window that had a resolution larger than the 1080p tier but less than or equal to the 2160p tier (over 2,073,600 and &lt;&#x3D; 4,194,304 pixels total). | [optional] 
**Tier1080p** | **float64** | Total number of delivered seconds during this time window that had a resolution larger than the 720p tier but less than or equal to the 1440p tier (over 921,600 and &lt;&#x3D; 2,073,600 pixels total). | [optional] 
**Tier720p** | **float64** | Total number of delivered seconds during this time window that had a resolution within the 720p tier (up to 921,600 pixels total, based on typical 1280x720). | [optional] 
**TierAudioOnly** | **float64** | Total number of delivered seconds during this time window that had a resolution of audio only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


