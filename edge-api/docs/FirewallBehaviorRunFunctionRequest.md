# FirewallBehaviorRunFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**FirewallBehaviorRunFunctionAttributesRequest**](FirewallBehaviorRunFunctionAttributesRequest.md) |  | 

## Methods

### NewFirewallBehaviorRunFunctionRequest

`func NewFirewallBehaviorRunFunctionRequest(type_ string, attributes FirewallBehaviorRunFunctionAttributesRequest, ) *FirewallBehaviorRunFunctionRequest`

NewFirewallBehaviorRunFunctionRequest instantiates a new FirewallBehaviorRunFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorRunFunctionRequestWithDefaults

`func NewFirewallBehaviorRunFunctionRequestWithDefaults() *FirewallBehaviorRunFunctionRequest`

NewFirewallBehaviorRunFunctionRequestWithDefaults instantiates a new FirewallBehaviorRunFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorRunFunctionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorRunFunctionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorRunFunctionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorRunFunctionRequest) GetAttributes() FirewallBehaviorRunFunctionAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorRunFunctionRequest) GetAttributesOk() (*FirewallBehaviorRunFunctionAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorRunFunctionRequest) SetAttributes(v FirewallBehaviorRunFunctionAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


