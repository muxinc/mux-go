# InputFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerFormat** | Pointer to **string** |  | [optional] 
**Tracks** | Pointer to [**[]InputTrack**](InputTrack.md) |  | [optional] 

## Methods

### NewInputFile

`func NewInputFile() *InputFile`

NewInputFile instantiates a new InputFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFileWithDefaults

`func NewInputFileWithDefaults() *InputFile`

NewInputFileWithDefaults instantiates a new InputFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerFormat

`func (o *InputFile) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *InputFile) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *InputFile) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *InputFile) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### GetTracks

`func (o *InputFile) GetTracks() []InputTrack`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *InputFile) GetTracksOk() (*[]InputTrack, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *InputFile) SetTracks(v []InputTrack)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *InputFile) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


