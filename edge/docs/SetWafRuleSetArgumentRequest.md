# SetWafRuleSetArgumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Mode** | [**ModeEnum**](ModeEnum.md) |  | 

## Methods

### NewSetWafRuleSetArgumentRequest

`func NewSetWafRuleSetArgumentRequest(id int64, mode ModeEnum, ) *SetWafRuleSetArgumentRequest`

NewSetWafRuleSetArgumentRequest instantiates a new SetWafRuleSetArgumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWafRuleSetArgumentRequestWithDefaults

`func NewSetWafRuleSetArgumentRequestWithDefaults() *SetWafRuleSetArgumentRequest`

NewSetWafRuleSetArgumentRequestWithDefaults instantiates a new SetWafRuleSetArgumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetWafRuleSetArgumentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetWafRuleSetArgumentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetWafRuleSetArgumentRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetMode

`func (o *SetWafRuleSetArgumentRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SetWafRuleSetArgumentRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SetWafRuleSetArgumentRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


