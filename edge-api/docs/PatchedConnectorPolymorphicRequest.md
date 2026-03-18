# PatchedConnectorPolymorphicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | Type of the connector  * &#x60;http&#x60; - HTTP * &#x60;storage&#x60; - Storage * &#x60;live_ingest&#x60; - Live Ingest | [optional] 
**Attributes** | Pointer to [**ConnectorStorageAttributesRequest**](ConnectorStorageAttributesRequest.md) |  | [optional] 

## Methods

### NewPatchedConnectorPolymorphicRequest

`func NewPatchedConnectorPolymorphicRequest() *PatchedConnectorPolymorphicRequest`

NewPatchedConnectorPolymorphicRequest instantiates a new PatchedConnectorPolymorphicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConnectorPolymorphicRequestWithDefaults

`func NewPatchedConnectorPolymorphicRequestWithDefaults() *PatchedConnectorPolymorphicRequest`

NewPatchedConnectorPolymorphicRequestWithDefaults instantiates a new PatchedConnectorPolymorphicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedConnectorPolymorphicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConnectorPolymorphicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConnectorPolymorphicRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConnectorPolymorphicRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedConnectorPolymorphicRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedConnectorPolymorphicRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedConnectorPolymorphicRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedConnectorPolymorphicRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchedConnectorPolymorphicRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedConnectorPolymorphicRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedConnectorPolymorphicRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedConnectorPolymorphicRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchedConnectorPolymorphicRequest) GetAttributes() ConnectorStorageAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedConnectorPolymorphicRequest) GetAttributesOk() (*ConnectorStorageAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedConnectorPolymorphicRequest) SetAttributes(v ConnectorStorageAttributesRequest)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedConnectorPolymorphicRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


