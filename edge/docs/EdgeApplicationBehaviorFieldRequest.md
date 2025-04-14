# EdgeApplicationBehaviorFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | * &#x60;deny&#x60; - deny * &#x60;run_function&#x60; - run_function * &#x60;no_content&#x60; - no_content * &#x60;deliver&#x60; - deliver * &#x60;finish_request_phase&#x60; - finish_request_phase * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 * &#x60;forward_cookies&#x60; - forward_cookies * &#x60;optimize_images&#x60; - optimize_images * &#x60;set_origin&#x60; - set_origin * &#x60;set_edge_connector&#x60; - set_edge_connector * &#x60;set_cache_policy&#x60; - set_cache_policy * &#x60;bypass_cache_phase&#x60; - bypass_cache_phase * &#x60;enable_gzip&#x60; - enable_gzip * &#x60;redirect_http_to_https&#x60; - redirect_http_to_https * &#x60;set_cookie&#x60; - set_cookie * &#x60;rewrite_request&#x60; - rewrite_request * &#x60;add_request_header&#x60; - add_request_header * &#x60;filter_request_header&#x60; - filter_request_header * &#x60;add_response_header&#x60; - add_response_header * &#x60;filter_response_header&#x60; - filter_response_header * &#x60;capture_match_groups&#x60; - capture_match_groups * &#x60;add_request_cookie&#x60; - add_request_cookie * &#x60;filter_response_cookie&#x60; - filter_response_cookie * &#x60;filter_request_cookie&#x60; - filter_request_cookie | 
**Argument** | Pointer to [**NullableEdgeApplicationBehaviorPolymorphicArgumentRequest**](EdgeApplicationBehaviorPolymorphicArgumentRequest.md) |  | [optional] 

## Methods

### NewEdgeApplicationBehaviorFieldRequest

`func NewEdgeApplicationBehaviorFieldRequest(name string, ) *EdgeApplicationBehaviorFieldRequest`

NewEdgeApplicationBehaviorFieldRequest instantiates a new EdgeApplicationBehaviorFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationBehaviorFieldRequestWithDefaults

`func NewEdgeApplicationBehaviorFieldRequestWithDefaults() *EdgeApplicationBehaviorFieldRequest`

NewEdgeApplicationBehaviorFieldRequestWithDefaults instantiates a new EdgeApplicationBehaviorFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeApplicationBehaviorFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationBehaviorFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationBehaviorFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgument

`func (o *EdgeApplicationBehaviorFieldRequest) GetArgument() EdgeApplicationBehaviorPolymorphicArgumentRequest`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeApplicationBehaviorFieldRequest) GetArgumentOk() (*EdgeApplicationBehaviorPolymorphicArgumentRequest, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeApplicationBehaviorFieldRequest) SetArgument(v EdgeApplicationBehaviorPolymorphicArgumentRequest)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeApplicationBehaviorFieldRequest) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeApplicationBehaviorFieldRequest) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeApplicationBehaviorFieldRequest) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


