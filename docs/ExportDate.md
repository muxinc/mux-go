# ExportDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportDate** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]ExportFile**](ExportFile.md) |  | [optional] 

## Methods

### NewExportDate

`func NewExportDate() *ExportDate`

NewExportDate instantiates a new ExportDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportDateWithDefaults

`func NewExportDateWithDefaults() *ExportDate`

NewExportDateWithDefaults instantiates a new ExportDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportDate

`func (o *ExportDate) GetExportDate() string`

GetExportDate returns the ExportDate field if non-nil, zero value otherwise.

### GetExportDateOk

`func (o *ExportDate) GetExportDateOk() (*string, bool)`

GetExportDateOk returns a tuple with the ExportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportDate

`func (o *ExportDate) SetExportDate(v string)`

SetExportDate sets ExportDate field to given value.

### HasExportDate

`func (o *ExportDate) HasExportDate() bool`

HasExportDate returns a boolean if a field has been set.

### GetFiles

`func (o *ExportDate) GetFiles() []ExportFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExportDate) GetFilesOk() (*[]ExportFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExportDate) SetFiles(v []ExportFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExportDate) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


