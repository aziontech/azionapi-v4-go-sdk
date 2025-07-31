# PaginatedEdgeApplicationFunctionInstanceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]EdgeApplicationFunctionInstance**](EdgeApplicationFunctionInstance.md) |  | [optional] 

## Methods

### NewPaginatedEdgeApplicationFunctionInstanceList

`func NewPaginatedEdgeApplicationFunctionInstanceList() *PaginatedEdgeApplicationFunctionInstanceList`

NewPaginatedEdgeApplicationFunctionInstanceList instantiates a new PaginatedEdgeApplicationFunctionInstanceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEdgeApplicationFunctionInstanceListWithDefaults

`func NewPaginatedEdgeApplicationFunctionInstanceListWithDefaults() *PaginatedEdgeApplicationFunctionInstanceList`

NewPaginatedEdgeApplicationFunctionInstanceListWithDefaults instantiates a new PaginatedEdgeApplicationFunctionInstanceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedEdgeApplicationFunctionInstanceList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedEdgeApplicationFunctionInstanceList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedEdgeApplicationFunctionInstanceList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedEdgeApplicationFunctionInstanceList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedEdgeApplicationFunctionInstanceList) GetResults() []EdgeApplicationFunctionInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedEdgeApplicationFunctionInstanceList) GetResultsOk() (*[]EdgeApplicationFunctionInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedEdgeApplicationFunctionInstanceList) SetResults(v []EdgeApplicationFunctionInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedEdgeApplicationFunctionInstanceList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


