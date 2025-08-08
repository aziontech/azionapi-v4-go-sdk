# EdgeConnectorStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** | Type of the edge connector  * &#x60;http&#x60; - HTTP * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Attributes** | [**EdgeConnectorStorageAttributesRequest**](EdgeConnectorStorageAttributesRequest.md) |  | 

## Methods

### NewEdgeConnectorStorageRequest

`func NewEdgeConnectorStorageRequest(name string, type_ string, attributes EdgeConnectorStorageAttributesRequest, ) *EdgeConnectorStorageRequest`

NewEdgeConnectorStorageRequest instantiates a new EdgeConnectorStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorStorageRequestWithDefaults

`func NewEdgeConnectorStorageRequestWithDefaults() *EdgeConnectorStorageRequest`

NewEdgeConnectorStorageRequestWithDefaults instantiates a new EdgeConnectorStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeConnectorStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorStorageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *EdgeConnectorStorageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorStorageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorStorageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorStorageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *EdgeConnectorStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeConnectorStorageRequest) GetAttributes() EdgeConnectorStorageAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeConnectorStorageRequest) GetAttributesOk() (*EdgeConnectorStorageAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeConnectorStorageRequest) SetAttributes(v EdgeConnectorStorageAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


