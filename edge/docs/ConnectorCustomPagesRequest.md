# ConnectorCustomPagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeConnector** | Pointer to **NullableInt64** |  | [optional] 
**Pages** | [**[]PageRequest**](PageRequest.md) |  | 

## Methods

### NewConnectorCustomPagesRequest

`func NewConnectorCustomPagesRequest(pages []PageRequest, ) *ConnectorCustomPagesRequest`

NewConnectorCustomPagesRequest instantiates a new ConnectorCustomPagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCustomPagesRequestWithDefaults

`func NewConnectorCustomPagesRequestWithDefaults() *ConnectorCustomPagesRequest`

NewConnectorCustomPagesRequestWithDefaults instantiates a new ConnectorCustomPagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeConnector

`func (o *ConnectorCustomPagesRequest) GetEdgeConnector() int64`

GetEdgeConnector returns the EdgeConnector field if non-nil, zero value otherwise.

### GetEdgeConnectorOk

`func (o *ConnectorCustomPagesRequest) GetEdgeConnectorOk() (*int64, bool)`

GetEdgeConnectorOk returns a tuple with the EdgeConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeConnector

`func (o *ConnectorCustomPagesRequest) SetEdgeConnector(v int64)`

SetEdgeConnector sets EdgeConnector field to given value.

### HasEdgeConnector

`func (o *ConnectorCustomPagesRequest) HasEdgeConnector() bool`

HasEdgeConnector returns a boolean if a field has been set.

### SetEdgeConnectorNil

`func (o *ConnectorCustomPagesRequest) SetEdgeConnectorNil(b bool)`

 SetEdgeConnectorNil sets the value for EdgeConnector to be an explicit nil

### UnsetEdgeConnector
`func (o *ConnectorCustomPagesRequest) UnsetEdgeConnector()`

UnsetEdgeConnector ensures that no value is present for EdgeConnector, not even an explicit nil
### GetPages

`func (o *ConnectorCustomPagesRequest) GetPages() []PageRequest`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ConnectorCustomPagesRequest) GetPagesOk() (*[]PageRequest, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ConnectorCustomPagesRequest) SetPages(v []PageRequest)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


