# ResponseDeleteWAF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableWAF**](WAF.md) |  | 

## Methods

### NewResponseDeleteWAF

`func NewResponseDeleteWAF(state string, data NullableWAF, ) *ResponseDeleteWAF`

NewResponseDeleteWAF instantiates a new ResponseDeleteWAF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteWAFWithDefaults

`func NewResponseDeleteWAFWithDefaults() *ResponseDeleteWAF`

NewResponseDeleteWAFWithDefaults instantiates a new ResponseDeleteWAF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteWAF) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteWAF) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteWAF) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteWAF) GetData() WAF`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteWAF) GetDataOk() (*WAF, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteWAF) SetData(v WAF)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteWAF) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteWAF) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


