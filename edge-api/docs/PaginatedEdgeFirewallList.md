# PaginatedEdgeFirewallList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]EdgeFirewall**](EdgeFirewall.md) |  | [optional] 

## Methods

### NewPaginatedEdgeFirewallList

`func NewPaginatedEdgeFirewallList() *PaginatedEdgeFirewallList`

NewPaginatedEdgeFirewallList instantiates a new PaginatedEdgeFirewallList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedEdgeFirewallListWithDefaults

`func NewPaginatedEdgeFirewallListWithDefaults() *PaginatedEdgeFirewallList`

NewPaginatedEdgeFirewallListWithDefaults instantiates a new PaginatedEdgeFirewallList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedEdgeFirewallList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedEdgeFirewallList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedEdgeFirewallList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedEdgeFirewallList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedEdgeFirewallList) GetResults() []EdgeFirewall`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedEdgeFirewallList) GetResultsOk() (*[]EdgeFirewall, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedEdgeFirewallList) SetResults(v []EdgeFirewall)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedEdgeFirewallList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


