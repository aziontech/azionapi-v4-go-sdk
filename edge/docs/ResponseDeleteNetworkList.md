# ResponseDeleteNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateEnum**](StateEnum.md) |  | 
**Data** | [**NullableNetworkList**](NetworkList.md) |  | 

## Methods

### NewResponseDeleteNetworkList

`func NewResponseDeleteNetworkList(state StateEnum, data NullableNetworkList, ) *ResponseDeleteNetworkList`

NewResponseDeleteNetworkList instantiates a new ResponseDeleteNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteNetworkListWithDefaults

`func NewResponseDeleteNetworkListWithDefaults() *ResponseDeleteNetworkList`

NewResponseDeleteNetworkListWithDefaults instantiates a new ResponseDeleteNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteNetworkList) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteNetworkList) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteNetworkList) SetState(v StateEnum)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteNetworkList) GetData() NetworkList`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteNetworkList) GetDataOk() (*NetworkList, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteNetworkList) SetData(v NetworkList)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteNetworkList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteNetworkList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


