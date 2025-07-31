# WAFExceptionSpecificConditionOnValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | * &#x60;specific_body_form_field_value&#x60; - specific_body_form_field_value * &#x60;specific_http_header_value&#x60; - specific_http_header_value * &#x60;specific_query_string_value&#x60; - specific_query_string_value | 
**Value** | **NullableString** |  | 

## Methods

### NewWAFExceptionSpecificConditionOnValue

`func NewWAFExceptionSpecificConditionOnValue(match string, value NullableString, ) *WAFExceptionSpecificConditionOnValue`

NewWAFExceptionSpecificConditionOnValue instantiates a new WAFExceptionSpecificConditionOnValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFExceptionSpecificConditionOnValueWithDefaults

`func NewWAFExceptionSpecificConditionOnValueWithDefaults() *WAFExceptionSpecificConditionOnValue`

NewWAFExceptionSpecificConditionOnValueWithDefaults instantiates a new WAFExceptionSpecificConditionOnValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *WAFExceptionSpecificConditionOnValue) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *WAFExceptionSpecificConditionOnValue) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *WAFExceptionSpecificConditionOnValue) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetValue

`func (o *WAFExceptionSpecificConditionOnValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WAFExceptionSpecificConditionOnValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WAFExceptionSpecificConditionOnValue) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *WAFExceptionSpecificConditionOnValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WAFExceptionSpecificConditionOnValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


