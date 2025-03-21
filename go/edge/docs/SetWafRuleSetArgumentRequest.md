# SetWafRuleSetArgumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Mode** | **string** | * &#x60;learning&#x60; - learning * &#x60;blocking&#x60; - blocking | 

## Methods

### NewSetWafRuleSetArgumentRequest

`func NewSetWafRuleSetArgumentRequest(id int64, mode string, ) *SetWafRuleSetArgumentRequest`

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

`func (o *SetWafRuleSetArgumentRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SetWafRuleSetArgumentRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SetWafRuleSetArgumentRequest) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


