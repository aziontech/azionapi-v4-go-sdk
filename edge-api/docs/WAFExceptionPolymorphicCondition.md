# WAFExceptionPolymorphicCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | * &#x60;specific_body_form_field_value&#x60; - specific_body_form_field_value * &#x60;specific_http_header_value&#x60; - specific_http_header_value * &#x60;specific_query_string_value&#x60; - specific_query_string_value | 
**Name** | **NullableString** |  | 
**Value** | **NullableString** |  | 

## Methods

### NewWAFExceptionPolymorphicCondition

`func NewWAFExceptionPolymorphicCondition(match string, name NullableString, value NullableString, ) *WAFExceptionPolymorphicCondition`

NewWAFExceptionPolymorphicCondition instantiates a new WAFExceptionPolymorphicCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFExceptionPolymorphicConditionWithDefaults

`func NewWAFExceptionPolymorphicConditionWithDefaults() *WAFExceptionPolymorphicCondition`

NewWAFExceptionPolymorphicConditionWithDefaults instantiates a new WAFExceptionPolymorphicCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *WAFExceptionPolymorphicCondition) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *WAFExceptionPolymorphicCondition) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *WAFExceptionPolymorphicCondition) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetName

`func (o *WAFExceptionPolymorphicCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WAFExceptionPolymorphicCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WAFExceptionPolymorphicCondition) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *WAFExceptionPolymorphicCondition) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WAFExceptionPolymorphicCondition) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *WAFExceptionPolymorphicCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WAFExceptionPolymorphicCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WAFExceptionPolymorphicCondition) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *WAFExceptionPolymorphicCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WAFExceptionPolymorphicCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


