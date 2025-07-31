# NetworkListDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Type** | **string** | * &#x60;asn&#x60; - ASN * &#x60;countries&#x60; - Countries * &#x60;ip_cidr&#x60; - IP/CIDR | 
**Items** | **[]string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkListDetail

`func NewNetworkListDetail(id int64, name string, type_ string, items []string, lastEditor string, lastModified time.Time, ) *NetworkListDetail`

NewNetworkListDetail instantiates a new NetworkListDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkListDetailWithDefaults

`func NewNetworkListDetailWithDefaults() *NetworkListDetail`

NewNetworkListDetailWithDefaults instantiates a new NetworkListDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkListDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkListDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkListDetail) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *NetworkListDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkListDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkListDetail) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NetworkListDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkListDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkListDetail) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *NetworkListDetail) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkListDetail) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkListDetail) SetItems(v []string)`

SetItems sets Items field to given value.


### GetLastEditor

`func (o *NetworkListDetail) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *NetworkListDetail) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *NetworkListDetail) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *NetworkListDetail) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *NetworkListDetail) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *NetworkListDetail) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *NetworkListDetail) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkListDetail) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkListDetail) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkListDetail) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


