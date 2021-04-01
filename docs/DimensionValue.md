# DimensionValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewDimensionValue

`func NewDimensionValue() *DimensionValue`

NewDimensionValue instantiates a new DimensionValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionValueWithDefaults

`func NewDimensionValueWithDefaults() *DimensionValue`

NewDimensionValueWithDefaults instantiates a new DimensionValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DimensionValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DimensionValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DimensionValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DimensionValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTotalCount

`func (o *DimensionValue) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DimensionValue) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DimensionValue) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DimensionValue) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


