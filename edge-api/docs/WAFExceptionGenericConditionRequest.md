# WAFExceptionGenericConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | * &#x60;any_http_header_name&#x60; - any_http_header_name * &#x60;any_http_header_value&#x60; - any_http_header_value * &#x60;any_query_string_name&#x60; - any_query_string_name * &#x60;any_query_string_value&#x60; - any_query_string_value * &#x60;any_url&#x60; - any_url * &#x60;body_form_field_name&#x60; - body_form_field_name * &#x60;body_form_field_value&#x60; - body_form_field_value * &#x60;file_extension&#x60; - file_extension * &#x60;raw_body&#x60; - raw_body | 

## Methods

### NewWAFExceptionGenericConditionRequest

`func NewWAFExceptionGenericConditionRequest(match string, ) *WAFExceptionGenericConditionRequest`

NewWAFExceptionGenericConditionRequest instantiates a new WAFExceptionGenericConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFExceptionGenericConditionRequestWithDefaults

`func NewWAFExceptionGenericConditionRequestWithDefaults() *WAFExceptionGenericConditionRequest`

NewWAFExceptionGenericConditionRequestWithDefaults instantiates a new WAFExceptionGenericConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *WAFExceptionGenericConditionRequest) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *WAFExceptionGenericConditionRequest) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *WAFExceptionGenericConditionRequest) SetMatch(v string)`

SetMatch sets Match field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


