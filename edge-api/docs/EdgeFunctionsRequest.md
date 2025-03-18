# EdgeFunctionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**LanguageEnum** | Pointer to **string** | * &#x60;javascript&#x60; - JavaScript * &#x60;lua&#x60; - Lua | [optional] 
**Code** | **string** |  | 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**InitiatorTypeEnum** | Pointer to **string** | * &#x60;edge_application&#x60; - Edge Application * &#x60;edge_firewall&#x60; - Edge Firewall | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdgeFunctionsRequest

`func NewEdgeFunctionsRequest(name string, code string, ) *EdgeFunctionsRequest`

NewEdgeFunctionsRequest instantiates a new EdgeFunctionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionsRequestWithDefaults

`func NewEdgeFunctionsRequestWithDefaults() *EdgeFunctionsRequest`

NewEdgeFunctionsRequestWithDefaults instantiates a new EdgeFunctionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFunctionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFunctionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFunctionsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLanguageEnum

`func (o *EdgeFunctionsRequest) GetLanguageEnum() string`

GetLanguageEnum returns the LanguageEnum field if non-nil, zero value otherwise.

### GetLanguageEnumOk

`func (o *EdgeFunctionsRequest) GetLanguageEnumOk() (*string, bool)`

GetLanguageEnumOk returns a tuple with the LanguageEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageEnum

`func (o *EdgeFunctionsRequest) SetLanguageEnum(v string)`

SetLanguageEnum sets LanguageEnum field to given value.

### HasLanguageEnum

`func (o *EdgeFunctionsRequest) HasLanguageEnum() bool`

HasLanguageEnum returns a boolean if a field has been set.

### GetCode

`func (o *EdgeFunctionsRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EdgeFunctionsRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EdgeFunctionsRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetJsonArgs

`func (o *EdgeFunctionsRequest) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *EdgeFunctionsRequest) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *EdgeFunctionsRequest) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *EdgeFunctionsRequest) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *EdgeFunctionsRequest) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *EdgeFunctionsRequest) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetInitiatorTypeEnum

`func (o *EdgeFunctionsRequest) GetInitiatorTypeEnum() string`

GetInitiatorTypeEnum returns the InitiatorTypeEnum field if non-nil, zero value otherwise.

### GetInitiatorTypeEnumOk

`func (o *EdgeFunctionsRequest) GetInitiatorTypeEnumOk() (*string, bool)`

GetInitiatorTypeEnumOk returns a tuple with the InitiatorTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorTypeEnum

`func (o *EdgeFunctionsRequest) SetInitiatorTypeEnum(v string)`

SetInitiatorTypeEnum sets InitiatorTypeEnum field to given value.

### HasInitiatorTypeEnum

`func (o *EdgeFunctionsRequest) HasInitiatorTypeEnum() bool`

HasInitiatorTypeEnum returns a boolean if a field has been set.

### GetActive

`func (o *EdgeFunctionsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFunctionsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFunctionsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFunctionsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


