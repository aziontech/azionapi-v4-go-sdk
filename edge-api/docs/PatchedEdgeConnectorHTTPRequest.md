# PatchedEdgeConnectorHTTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | Type of the edge connector  * &#x60;http&#x60; - HTTP * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | [optional] 
**Attributes** | Pointer to [**EdgeConnectorHTTPAttributesRequest**](EdgeConnectorHTTPAttributesRequest.md) |  | [optional] 

## Methods

### NewPatchedEdgeConnectorHTTPRequest

`func NewPatchedEdgeConnectorHTTPRequest() *PatchedEdgeConnectorHTTPRequest`

NewPatchedEdgeConnectorHTTPRequest instantiates a new PatchedEdgeConnectorHTTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeConnectorHTTPRequestWithDefaults

`func NewPatchedEdgeConnectorHTTPRequestWithDefaults() *PatchedEdgeConnectorHTTPRequest`

NewPatchedEdgeConnectorHTTPRequestWithDefaults instantiates a new PatchedEdgeConnectorHTTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeConnectorHTTPRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeConnectorHTTPRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeConnectorHTTPRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeConnectorHTTPRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeConnectorHTTPRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeConnectorHTTPRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeConnectorHTTPRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeConnectorHTTPRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchedEdgeConnectorHTTPRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedEdgeConnectorHTTPRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedEdgeConnectorHTTPRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedEdgeConnectorHTTPRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchedEdgeConnectorHTTPRequest) GetAttributes() EdgeConnectorHTTPAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedEdgeConnectorHTTPRequest) GetAttributesOk() (*EdgeConnectorHTTPAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedEdgeConnectorHTTPRequest) SetAttributes(v EdgeConnectorHTTPAttributesRequest)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedEdgeConnectorHTTPRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


