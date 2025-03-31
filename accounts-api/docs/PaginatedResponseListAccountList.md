# PaginatedResponseListAccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ResponseListAccount**](ResponseListAccount.md) |  | [optional] 

## Methods

### NewPaginatedResponseListAccountList

`func NewPaginatedResponseListAccountList() *PaginatedResponseListAccountList`

NewPaginatedResponseListAccountList instantiates a new PaginatedResponseListAccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseListAccountListWithDefaults

`func NewPaginatedResponseListAccountListWithDefaults() *PaginatedResponseListAccountList`

NewPaginatedResponseListAccountListWithDefaults instantiates a new PaginatedResponseListAccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedResponseListAccountList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResponseListAccountList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResponseListAccountList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedResponseListAccountList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseListAccountList) GetResults() []ResponseListAccount`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseListAccountList) GetResultsOk() (*[]ResponseListAccount, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseListAccountList) SetResults(v []ResponseListAccount)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseListAccountList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


