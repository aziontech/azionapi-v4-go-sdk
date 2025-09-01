# ConnectorLiveIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | 
**Type** | **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | [**ConnectorLiveIngestAttributes**](ConnectorLiveIngestAttributes.md) |  | 

## Methods

### NewConnectorLiveIngest

`func NewConnectorLiveIngest(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, type_ string, attributes ConnectorLiveIngestAttributes, ) *ConnectorLiveIngest`

NewConnectorLiveIngest instantiates a new ConnectorLiveIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorLiveIngestWithDefaults

`func NewConnectorLiveIngestWithDefaults() *ConnectorLiveIngest`

NewConnectorLiveIngestWithDefaults instantiates a new ConnectorLiveIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorLiveIngest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorLiveIngest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorLiveIngest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectorLiveIngest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorLiveIngest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorLiveIngest) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ConnectorLiveIngest) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ConnectorLiveIngest) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ConnectorLiveIngest) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ConnectorLiveIngest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ConnectorLiveIngest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ConnectorLiveIngest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *ConnectorLiveIngest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorLiveIngest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorLiveIngest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorLiveIngest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *ConnectorLiveIngest) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ConnectorLiveIngest) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ConnectorLiveIngest) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetType

`func (o *ConnectorLiveIngest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectorLiveIngest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectorLiveIngest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ConnectorLiveIngest) GetAttributes() ConnectorLiveIngestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConnectorLiveIngest) GetAttributesOk() (*ConnectorLiveIngestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConnectorLiveIngest) SetAttributes(v ConnectorLiveIngestAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


