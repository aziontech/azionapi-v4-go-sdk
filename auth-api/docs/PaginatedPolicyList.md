# PaginatedPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]Policy**](Policy.md) |  | [optional] 

## Methods

### NewPaginatedPolicyList

`func NewPaginatedPolicyList() *PaginatedPolicyList`

NewPaginatedPolicyList instantiates a new PaginatedPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPolicyListWithDefaults

`func NewPaginatedPolicyListWithDefaults() *PaginatedPolicyList`

NewPaginatedPolicyListWithDefaults instantiates a new PaginatedPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedPolicyList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedPolicyList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedPolicyList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedPolicyList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedPolicyList) GetResults() []Policy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedPolicyList) GetResultsOk() (*[]Policy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedPolicyList) SetResults(v []Policy)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedPolicyList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


