# WAFExceptionPolymorphicConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | * &#x60;specific_body_form_field_value&#x60; - specific_body_form_field_value * &#x60;specific_http_header_value&#x60; - specific_http_header_value * &#x60;specific_query_string_value&#x60; - specific_query_string_value | 
**Name** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewWAFExceptionPolymorphicConditionRequest

`func NewWAFExceptionPolymorphicConditionRequest(match string, name string, value string, ) *WAFExceptionPolymorphicConditionRequest`

NewWAFExceptionPolymorphicConditionRequest instantiates a new WAFExceptionPolymorphicConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFExceptionPolymorphicConditionRequestWithDefaults

`func NewWAFExceptionPolymorphicConditionRequestWithDefaults() *WAFExceptionPolymorphicConditionRequest`

NewWAFExceptionPolymorphicConditionRequestWithDefaults instantiates a new WAFExceptionPolymorphicConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *WAFExceptionPolymorphicConditionRequest) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *WAFExceptionPolymorphicConditionRequest) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *WAFExceptionPolymorphicConditionRequest) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetName

`func (o *WAFExceptionPolymorphicConditionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WAFExceptionPolymorphicConditionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WAFExceptionPolymorphicConditionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *WAFExceptionPolymorphicConditionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WAFExceptionPolymorphicConditionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WAFExceptionPolymorphicConditionRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


