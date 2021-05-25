# InputInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**InputSettings**](InputSettings.md) |  | [optional] 
**File** | Pointer to [**InputFile**](InputFile.md) |  | [optional] 

## Methods

### NewInputInfo

`func NewInputInfo() *InputInfo`

NewInputInfo instantiates a new InputInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputInfoWithDefaults

`func NewInputInfoWithDefaults() *InputInfo`

NewInputInfoWithDefaults instantiates a new InputInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *InputInfo) GetSettings() InputSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *InputInfo) GetSettingsOk() (*InputSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *InputInfo) SetSettings(v InputSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *InputInfo) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetFile

`func (o *InputInfo) GetFile() InputFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *InputInfo) GetFileOk() (*InputFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *InputInfo) SetFile(v InputFile)`

SetFile sets File field to given value.

### HasFile

`func (o *InputInfo) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


