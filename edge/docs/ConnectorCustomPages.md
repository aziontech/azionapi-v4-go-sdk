# ConnectorCustomPages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeConnector** | Pointer to **NullableInt64** |  | [optional] 
**Pages** | [**[]Page**](Page.md) |  | 

## Methods

### NewConnectorCustomPages

`func NewConnectorCustomPages(pages []Page, ) *ConnectorCustomPages`

NewConnectorCustomPages instantiates a new ConnectorCustomPages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorCustomPagesWithDefaults

`func NewConnectorCustomPagesWithDefaults() *ConnectorCustomPages`

NewConnectorCustomPagesWithDefaults instantiates a new ConnectorCustomPages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeConnector

`func (o *ConnectorCustomPages) GetEdgeConnector() int64`

GetEdgeConnector returns the EdgeConnector field if non-nil, zero value otherwise.

### GetEdgeConnectorOk

`func (o *ConnectorCustomPages) GetEdgeConnectorOk() (*int64, bool)`

GetEdgeConnectorOk returns a tuple with the EdgeConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeConnector

`func (o *ConnectorCustomPages) SetEdgeConnector(v int64)`

SetEdgeConnector sets EdgeConnector field to given value.

### HasEdgeConnector

`func (o *ConnectorCustomPages) HasEdgeConnector() bool`

HasEdgeConnector returns a boolean if a field has been set.

### SetEdgeConnectorNil

`func (o *ConnectorCustomPages) SetEdgeConnectorNil(b bool)`

 SetEdgeConnectorNil sets the value for EdgeConnector to be an explicit nil

### UnsetEdgeConnector
`func (o *ConnectorCustomPages) UnsetEdgeConnector()`

UnsetEdgeConnector ensures that no value is present for EdgeConnector, not even an explicit nil
### GetPages

`func (o *ConnectorCustomPages) GetPages() []Page`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ConnectorCustomPages) GetPagesOk() (*[]Page, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ConnectorCustomPages) SetPages(v []Page)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


