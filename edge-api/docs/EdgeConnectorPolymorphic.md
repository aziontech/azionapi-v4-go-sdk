# EdgeConnectorPolymorphic

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
**Attributes** | [**EdgeConnectorLiveIngestAttributes**](EdgeConnectorLiveIngestAttributes.md) |  | 

## Methods

### NewEdgeConnectorPolymorphic

`func NewEdgeConnectorPolymorphic(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, type_ string, attributes EdgeConnectorLiveIngestAttributes, ) *EdgeConnectorPolymorphic`

NewEdgeConnectorPolymorphic instantiates a new EdgeConnectorPolymorphic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorPolymorphicWithDefaults

`func NewEdgeConnectorPolymorphicWithDefaults() *EdgeConnectorPolymorphic`

NewEdgeConnectorPolymorphicWithDefaults instantiates a new EdgeConnectorPolymorphic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeConnectorPolymorphic) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeConnectorPolymorphic) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeConnectorPolymorphic) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeConnectorPolymorphic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorPolymorphic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorPolymorphic) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *EdgeConnectorPolymorphic) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeConnectorPolymorphic) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeConnectorPolymorphic) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeConnectorPolymorphic) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeConnectorPolymorphic) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeConnectorPolymorphic) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *EdgeConnectorPolymorphic) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorPolymorphic) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorPolymorphic) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorPolymorphic) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *EdgeConnectorPolymorphic) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeConnectorPolymorphic) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeConnectorPolymorphic) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetType

`func (o *EdgeConnectorPolymorphic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorPolymorphic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorPolymorphic) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeConnectorPolymorphic) GetAttributes() EdgeConnectorLiveIngestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeConnectorPolymorphic) GetAttributesOk() (*EdgeConnectorLiveIngestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeConnectorPolymorphic) SetAttributes(v EdgeConnectorLiveIngestAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


