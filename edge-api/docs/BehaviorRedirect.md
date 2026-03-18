# BehaviorRedirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 | 
**Attributes** | [**BehaviorStringAttributes**](BehaviorStringAttributes.md) |  | 

## Methods

### NewBehaviorRedirect

`func NewBehaviorRedirect(type_ string, attributes BehaviorStringAttributes, ) *BehaviorRedirect`

NewBehaviorRedirect instantiates a new BehaviorRedirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRedirectWithDefaults

`func NewBehaviorRedirectWithDefaults() *BehaviorRedirect`

NewBehaviorRedirectWithDefaults instantiates a new BehaviorRedirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorRedirect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorRedirect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorRedirect) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorRedirect) GetAttributes() BehaviorStringAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorRedirect) GetAttributesOk() (*BehaviorStringAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorRedirect) SetAttributes(v BehaviorStringAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


