# ResponseAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateEnum**](StateEnum.md) |  | 
**Data** | [**AccountInfo**](AccountInfo.md) |  | 

## Methods

### NewResponseAccountInfo

`func NewResponseAccountInfo(state StateEnum, data AccountInfo, ) *ResponseAccountInfo`

NewResponseAccountInfo instantiates a new ResponseAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAccountInfoWithDefaults

`func NewResponseAccountInfoWithDefaults() *ResponseAccountInfo`

NewResponseAccountInfoWithDefaults instantiates a new ResponseAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAccountInfo) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAccountInfo) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAccountInfo) SetState(v StateEnum)`

SetState sets State field to given value.


### GetData

`func (o *ResponseAccountInfo) GetData() AccountInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAccountInfo) GetDataOk() (*AccountInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAccountInfo) SetData(v AccountInfo)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


