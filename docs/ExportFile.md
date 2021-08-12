# ExportFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewExportFile

`func NewExportFile() *ExportFile`

NewExportFile instantiates a new ExportFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportFileWithDefaults

`func NewExportFileWithDefaults() *ExportFile`

NewExportFileWithDefaults instantiates a new ExportFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ExportFile) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExportFile) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExportFile) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExportFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *ExportFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExportFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPath

`func (o *ExportFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ExportFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ExportFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ExportFile) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


