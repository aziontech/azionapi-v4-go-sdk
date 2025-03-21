# PaginatedResponseListEdgeFunctionsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]ResponseListEdgeFunctions**](ResponseListEdgeFunctions.md) |  | [optional] 

## Methods

### NewPaginatedResponseListEdgeFunctionsList

`func NewPaginatedResponseListEdgeFunctionsList() *PaginatedResponseListEdgeFunctionsList`

NewPaginatedResponseListEdgeFunctionsList instantiates a new PaginatedResponseListEdgeFunctionsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseListEdgeFunctionsListWithDefaults

`func NewPaginatedResponseListEdgeFunctionsListWithDefaults() *PaginatedResponseListEdgeFunctionsList`

NewPaginatedResponseListEdgeFunctionsListWithDefaults instantiates a new PaginatedResponseListEdgeFunctionsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedResponseListEdgeFunctionsList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResponseListEdgeFunctionsList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResponseListEdgeFunctionsList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedResponseListEdgeFunctionsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseListEdgeFunctionsList) GetResults() []ResponseListEdgeFunctions`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseListEdgeFunctionsList) GetResultsOk() (*[]ResponseListEdgeFunctions, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseListEdgeFunctionsList) SetResults(v []ResponseListEdgeFunctions)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseListEdgeFunctionsList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


