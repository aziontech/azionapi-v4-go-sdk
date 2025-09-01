# PaginatedFirewallList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]Firewall**](Firewall.md) |  | [optional] 

## Methods

### NewPaginatedFirewallList

`func NewPaginatedFirewallList() *PaginatedFirewallList`

NewPaginatedFirewallList instantiates a new PaginatedFirewallList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedFirewallListWithDefaults

`func NewPaginatedFirewallListWithDefaults() *PaginatedFirewallList`

NewPaginatedFirewallListWithDefaults instantiates a new PaginatedFirewallList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedFirewallList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedFirewallList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedFirewallList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedFirewallList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedFirewallList) GetResults() []Firewall`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedFirewallList) GetResultsOk() (*[]Firewall, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedFirewallList) SetResults(v []Firewall)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedFirewallList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


