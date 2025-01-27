# PatchedEdgeFunctionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to [**LanguageEnum**](LanguageEnum.md) |  | [optional] [default to javascript]
**Code** | Pointer to **string** |  | [optional] 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**InitiatorType** | Pointer to [**InitiatorTypeEnum**](InitiatorTypeEnum.md) |  | [optional] [default to edge_application]
**Active** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPatchedEdgeFunctionsRequest

`func NewPatchedEdgeFunctionsRequest() *PatchedEdgeFunctionsRequest`

NewPatchedEdgeFunctionsRequest instantiates a new PatchedEdgeFunctionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeFunctionsRequestWithDefaults

`func NewPatchedEdgeFunctionsRequestWithDefaults() *PatchedEdgeFunctionsRequest`

NewPatchedEdgeFunctionsRequestWithDefaults instantiates a new PatchedEdgeFunctionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeFunctionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeFunctionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeFunctionsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeFunctionsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *PatchedEdgeFunctionsRequest) GetLanguage() LanguageEnum`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PatchedEdgeFunctionsRequest) GetLanguageOk() (*LanguageEnum, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PatchedEdgeFunctionsRequest) SetLanguage(v LanguageEnum)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PatchedEdgeFunctionsRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCode

`func (o *PatchedEdgeFunctionsRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PatchedEdgeFunctionsRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PatchedEdgeFunctionsRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PatchedEdgeFunctionsRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetJsonArgs

`func (o *PatchedEdgeFunctionsRequest) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *PatchedEdgeFunctionsRequest) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *PatchedEdgeFunctionsRequest) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *PatchedEdgeFunctionsRequest) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *PatchedEdgeFunctionsRequest) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *PatchedEdgeFunctionsRequest) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetInitiatorType

`func (o *PatchedEdgeFunctionsRequest) GetInitiatorType() InitiatorTypeEnum`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *PatchedEdgeFunctionsRequest) GetInitiatorTypeOk() (*InitiatorTypeEnum, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *PatchedEdgeFunctionsRequest) SetInitiatorType(v InitiatorTypeEnum)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *PatchedEdgeFunctionsRequest) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeFunctionsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeFunctionsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeFunctionsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeFunctionsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


