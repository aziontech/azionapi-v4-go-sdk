# PaginatedResponseListWAFRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]ResponseListWAFRule**](ResponseListWAFRule.md) |  | [optional] 

## Methods

### NewPaginatedResponseListWAFRuleList

`func NewPaginatedResponseListWAFRuleList() *PaginatedResponseListWAFRuleList`

NewPaginatedResponseListWAFRuleList instantiates a new PaginatedResponseListWAFRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseListWAFRuleListWithDefaults

`func NewPaginatedResponseListWAFRuleListWithDefaults() *PaginatedResponseListWAFRuleList`

NewPaginatedResponseListWAFRuleListWithDefaults instantiates a new PaginatedResponseListWAFRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedResponseListWAFRuleList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResponseListWAFRuleList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResponseListWAFRuleList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedResponseListWAFRuleList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseListWAFRuleList) GetResults() []ResponseListWAFRule`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseListWAFRuleList) GetResultsOk() (*[]ResponseListWAFRule, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseListWAFRuleList) SetResults(v []ResponseListWAFRule)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseListWAFRuleList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


