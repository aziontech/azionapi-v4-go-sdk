# PaginatedEdgeApplicationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]EdgeApplication**](EdgeApplication.md) |  | [optional] 

## Methods

### NewPaginatedEdgeApplicationList

`func NewPaginatedEdgeApplicationList() *PaginatedEdgeApplicationList`

NewPaginatedEdgeApplicationList instantiates a new PaginatedEdgeApplicationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEdgeApplicationListWithDefaults

`func NewPaginatedEdgeApplicationListWithDefaults() *PaginatedEdgeApplicationList`

NewPaginatedEdgeApplicationListWithDefaults instantiates a new PaginatedEdgeApplicationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedEdgeApplicationList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedEdgeApplicationList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedEdgeApplicationList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedEdgeApplicationList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedEdgeApplicationList) GetResults() []EdgeApplication`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedEdgeApplicationList) GetResultsOk() (*[]EdgeApplication, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedEdgeApplicationList) SetResults(v []EdgeApplication)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedEdgeApplicationList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


