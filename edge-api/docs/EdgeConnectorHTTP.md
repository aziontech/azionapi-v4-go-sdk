# EdgeConnectorHTTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | [readonly] 
**Type** | **string** | Type of the edge connector  * &#x60;http&#x60; - HTTP * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | [**EdgeConnectorHTTPAttributes**](EdgeConnectorHTTPAttributes.md) |  | 

## Methods

### NewEdgeConnectorHTTP

`func NewEdgeConnectorHTTP(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, type_ string, attributes EdgeConnectorHTTPAttributes, ) *EdgeConnectorHTTP`

NewEdgeConnectorHTTP instantiates a new EdgeConnectorHTTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorHTTPWithDefaults

`func NewEdgeConnectorHTTPWithDefaults() *EdgeConnectorHTTP`

NewEdgeConnectorHTTPWithDefaults instantiates a new EdgeConnectorHTTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeConnectorHTTP) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeConnectorHTTP) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeConnectorHTTP) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeConnectorHTTP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorHTTP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorHTTP) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *EdgeConnectorHTTP) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeConnectorHTTP) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeConnectorHTTP) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeConnectorHTTP) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeConnectorHTTP) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeConnectorHTTP) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *EdgeConnectorHTTP) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorHTTP) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorHTTP) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorHTTP) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *EdgeConnectorHTTP) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeConnectorHTTP) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeConnectorHTTP) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetType

`func (o *EdgeConnectorHTTP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorHTTP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorHTTP) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeConnectorHTTP) GetAttributes() EdgeConnectorHTTPAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeConnectorHTTP) GetAttributesOk() (*EdgeConnectorHTTPAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeConnectorHTTP) SetAttributes(v EdgeConnectorHTTPAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


