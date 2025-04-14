# CustomPagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ConnectorCustomPages** | [**ConnectorCustomPagesRequest**](ConnectorCustomPagesRequest.md) |  | 

## Methods

### NewCustomPagesRequest

`func NewCustomPagesRequest(name string, connectorCustomPages ConnectorCustomPagesRequest, ) *CustomPagesRequest`

NewCustomPagesRequest instantiates a new CustomPagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPagesRequestWithDefaults

`func NewCustomPagesRequestWithDefaults() *CustomPagesRequest`

NewCustomPagesRequestWithDefaults instantiates a new CustomPagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomPagesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPagesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPagesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *CustomPagesRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomPagesRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomPagesRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CustomPagesRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConnectorCustomPages

`func (o *CustomPagesRequest) GetConnectorCustomPages() ConnectorCustomPagesRequest`

GetConnectorCustomPages returns the ConnectorCustomPages field if non-nil, zero value otherwise.

### GetConnectorCustomPagesOk

`func (o *CustomPagesRequest) GetConnectorCustomPagesOk() (*ConnectorCustomPagesRequest, bool)`

GetConnectorCustomPagesOk returns a tuple with the ConnectorCustomPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCustomPages

`func (o *CustomPagesRequest) SetConnectorCustomPages(v ConnectorCustomPagesRequest)`

SetConnectorCustomPages sets ConnectorCustomPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


