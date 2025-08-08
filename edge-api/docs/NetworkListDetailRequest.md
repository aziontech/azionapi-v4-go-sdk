# NetworkListDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** | * &#x60;asn&#x60; - ASN * &#x60;countries&#x60; - Countries * &#x60;ip_cidr&#x60; - IP/CIDR | 
**Items** | **[]string** |  | 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkListDetailRequest

`func NewNetworkListDetailRequest(name string, type_ string, items []string, ) *NetworkListDetailRequest`

NewNetworkListDetailRequest instantiates a new NetworkListDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkListDetailRequestWithDefaults

`func NewNetworkListDetailRequestWithDefaults() *NetworkListDetailRequest`

NewNetworkListDetailRequestWithDefaults instantiates a new NetworkListDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkListDetailRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkListDetailRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkListDetailRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NetworkListDetailRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkListDetailRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkListDetailRequest) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *NetworkListDetailRequest) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkListDetailRequest) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkListDetailRequest) SetItems(v []string)`

SetItems sets Items field to given value.


### GetActive

`func (o *NetworkListDetailRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkListDetailRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkListDetailRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkListDetailRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


