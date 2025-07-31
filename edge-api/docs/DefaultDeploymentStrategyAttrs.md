# DefaultDeploymentStrategyAttrs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeApplication** | **int64** |  | 
**EdgeFirewall** | Pointer to **NullableInt64** |  | [optional] 
**CustomPage** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewDefaultDeploymentStrategyAttrs

`func NewDefaultDeploymentStrategyAttrs(edgeApplication int64, ) *DefaultDeploymentStrategyAttrs`

NewDefaultDeploymentStrategyAttrs instantiates a new DefaultDeploymentStrategyAttrs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultDeploymentStrategyAttrsWithDefaults

`func NewDefaultDeploymentStrategyAttrsWithDefaults() *DefaultDeploymentStrategyAttrs`

NewDefaultDeploymentStrategyAttrsWithDefaults instantiates a new DefaultDeploymentStrategyAttrs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeApplication

`func (o *DefaultDeploymentStrategyAttrs) GetEdgeApplication() int64`

GetEdgeApplication returns the EdgeApplication field if non-nil, zero value otherwise.

### GetEdgeApplicationOk

`func (o *DefaultDeploymentStrategyAttrs) GetEdgeApplicationOk() (*int64, bool)`

GetEdgeApplicationOk returns a tuple with the EdgeApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplication

`func (o *DefaultDeploymentStrategyAttrs) SetEdgeApplication(v int64)`

SetEdgeApplication sets EdgeApplication field to given value.


### GetEdgeFirewall

`func (o *DefaultDeploymentStrategyAttrs) GetEdgeFirewall() int64`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *DefaultDeploymentStrategyAttrs) GetEdgeFirewallOk() (*int64, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *DefaultDeploymentStrategyAttrs) SetEdgeFirewall(v int64)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *DefaultDeploymentStrategyAttrs) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### SetEdgeFirewallNil

`func (o *DefaultDeploymentStrategyAttrs) SetEdgeFirewallNil(b bool)`

 SetEdgeFirewallNil sets the value for EdgeFirewall to be an explicit nil

### UnsetEdgeFirewall
`func (o *DefaultDeploymentStrategyAttrs) UnsetEdgeFirewall()`

UnsetEdgeFirewall ensures that no value is present for EdgeFirewall, not even an explicit nil
### GetCustomPage

`func (o *DefaultDeploymentStrategyAttrs) GetCustomPage() int64`

GetCustomPage returns the CustomPage field if non-nil, zero value otherwise.

### GetCustomPageOk

`func (o *DefaultDeploymentStrategyAttrs) GetCustomPageOk() (*int64, bool)`

GetCustomPageOk returns a tuple with the CustomPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPage

`func (o *DefaultDeploymentStrategyAttrs) SetCustomPage(v int64)`

SetCustomPage sets CustomPage field to given value.

### HasCustomPage

`func (o *DefaultDeploymentStrategyAttrs) HasCustomPage() bool`

HasCustomPage returns a boolean if a field has been set.

### SetCustomPageNil

`func (o *DefaultDeploymentStrategyAttrs) SetCustomPageNil(b bool)`

 SetCustomPageNil sets the value for CustomPage to be an explicit nil

### UnsetCustomPage
`func (o *DefaultDeploymentStrategyAttrs) UnsetCustomPage()`

UnsetCustomPage ensures that no value is present for CustomPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


