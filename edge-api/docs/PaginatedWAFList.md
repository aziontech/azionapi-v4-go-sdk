# PaginatedWAFList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]WAF**](WAF.md) |  | [optional] 

## Methods

### NewPaginatedWAFList

`func NewPaginatedWAFList() *PaginatedWAFList`

NewPaginatedWAFList instantiates a new PaginatedWAFList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedWAFListWithDefaults

`func NewPaginatedWAFListWithDefaults() *PaginatedWAFList`

NewPaginatedWAFListWithDefaults instantiates a new PaginatedWAFList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedWAFList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedWAFList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedWAFList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedWAFList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedWAFList) GetResults() []WAF`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedWAFList) GetResultsOk() (*[]WAF, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedWAFList) SetResults(v []WAF)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedWAFList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


