# PaginatedResponseListEdgeFirewallList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]ResponseListEdgeFirewall**](ResponseListEdgeFirewall.md) |  | [optional] 

## Methods

### NewPaginatedResponseListEdgeFirewallList

`func NewPaginatedResponseListEdgeFirewallList() *PaginatedResponseListEdgeFirewallList`

NewPaginatedResponseListEdgeFirewallList instantiates a new PaginatedResponseListEdgeFirewallList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseListEdgeFirewallListWithDefaults

`func NewPaginatedResponseListEdgeFirewallListWithDefaults() *PaginatedResponseListEdgeFirewallList`

NewPaginatedResponseListEdgeFirewallListWithDefaults instantiates a new PaginatedResponseListEdgeFirewallList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedResponseListEdgeFirewallList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResponseListEdgeFirewallList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResponseListEdgeFirewallList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedResponseListEdgeFirewallList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseListEdgeFirewallList) GetResults() []ResponseListEdgeFirewall`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseListEdgeFirewallList) GetResultsOk() (*[]ResponseListEdgeFirewall, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseListEdgeFirewallList) SetResults(v []ResponseListEdgeFirewall)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseListEdgeFirewallList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


