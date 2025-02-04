# ResponseDeleteAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateEnum**](StateEnum.md) |  | 
**Data** | [**NullableAccount**](Account.md) |  | 

## Methods

### NewResponseDeleteAccount

`func NewResponseDeleteAccount(state StateEnum, data NullableAccount, ) *ResponseDeleteAccount`

NewResponseDeleteAccount instantiates a new ResponseDeleteAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteAccountWithDefaults

`func NewResponseDeleteAccountWithDefaults() *ResponseDeleteAccount`

NewResponseDeleteAccountWithDefaults instantiates a new ResponseDeleteAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteAccount) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteAccount) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteAccount) SetState(v StateEnum)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteAccount) GetData() Account`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteAccount) GetDataOk() (*Account, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteAccount) SetData(v Account)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteAccount) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteAccount) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


