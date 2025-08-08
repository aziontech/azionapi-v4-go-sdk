# PaginatedCustomPageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]CustomPage**](CustomPage.md) |  | [optional] 

## Methods

### NewPaginatedCustomPageList

`func NewPaginatedCustomPageList() *PaginatedCustomPageList`

NewPaginatedCustomPageList instantiates a new PaginatedCustomPageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCustomPageListWithDefaults

`func NewPaginatedCustomPageListWithDefaults() *PaginatedCustomPageList`

NewPaginatedCustomPageListWithDefaults instantiates a new PaginatedCustomPageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedCustomPageList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedCustomPageList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedCustomPageList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedCustomPageList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedCustomPageList) GetResults() []CustomPage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCustomPageList) GetResultsOk() (*[]CustomPage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCustomPageList) SetResults(v []CustomPage)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedCustomPageList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


