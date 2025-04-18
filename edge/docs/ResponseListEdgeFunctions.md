# ResponseListEdgeFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Language** | Pointer to **string** | * &#x60;javascript&#x60; - JavaScript * &#x60;lua&#x60; - Lua | [optional] 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**InitiatorType** | Pointer to **string** | * &#x60;edge_application&#x60; - Edge Application * &#x60;edge_firewall&#x60; - Edge Firewall | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ReferenceCount** | **int64** |  | [readonly] 
**Version** | **string** | Installed version, which may not be the latest if the vendor has released updates since installation. | [readonly] 
**Vendor** | **string** |  | [readonly] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 

## Methods

### NewResponseListEdgeFunctions

`func NewResponseListEdgeFunctions(id int64, name string, referenceCount int64, version string, vendor string, lastEditor string, lastModified time.Time, productVersion string, ) *ResponseListEdgeFunctions`

NewResponseListEdgeFunctions instantiates a new ResponseListEdgeFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListEdgeFunctionsWithDefaults

`func NewResponseListEdgeFunctionsWithDefaults() *ResponseListEdgeFunctions`

NewResponseListEdgeFunctionsWithDefaults instantiates a new ResponseListEdgeFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListEdgeFunctions) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListEdgeFunctions) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListEdgeFunctions) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListEdgeFunctions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListEdgeFunctions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListEdgeFunctions) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *ResponseListEdgeFunctions) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResponseListEdgeFunctions) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResponseListEdgeFunctions) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ResponseListEdgeFunctions) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetJsonArgs

`func (o *ResponseListEdgeFunctions) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *ResponseListEdgeFunctions) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *ResponseListEdgeFunctions) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *ResponseListEdgeFunctions) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *ResponseListEdgeFunctions) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *ResponseListEdgeFunctions) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetInitiatorType

`func (o *ResponseListEdgeFunctions) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *ResponseListEdgeFunctions) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *ResponseListEdgeFunctions) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *ResponseListEdgeFunctions) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetActive

`func (o *ResponseListEdgeFunctions) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListEdgeFunctions) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListEdgeFunctions) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListEdgeFunctions) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetReferenceCount

`func (o *ResponseListEdgeFunctions) GetReferenceCount() int64`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *ResponseListEdgeFunctions) GetReferenceCountOk() (*int64, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *ResponseListEdgeFunctions) SetReferenceCount(v int64)`

SetReferenceCount sets ReferenceCount field to given value.


### GetVersion

`func (o *ResponseListEdgeFunctions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseListEdgeFunctions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseListEdgeFunctions) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVendor

`func (o *ResponseListEdgeFunctions) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ResponseListEdgeFunctions) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ResponseListEdgeFunctions) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetLastEditor

`func (o *ResponseListEdgeFunctions) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListEdgeFunctions) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListEdgeFunctions) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListEdgeFunctions) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListEdgeFunctions) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListEdgeFunctions) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *ResponseListEdgeFunctions) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseListEdgeFunctions) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseListEdgeFunctions) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


