# CacheSettingsTieredCacheModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | **string** | * &#x60;override&#x60; - override | 
**MaxAge** | **int64** |  | 
**Topology** | Pointer to **string** | * &#x60;near-edge&#x60; - near-edge * &#x60;br-east-1&#x60; - br-east-1 * &#x60;us-east-1&#x60; - us-east-1 | [optional] 

## Methods

### NewCacheSettingsTieredCacheModule

`func NewCacheSettingsTieredCacheModule(behavior string, maxAge int64, ) *CacheSettingsTieredCacheModule`

NewCacheSettingsTieredCacheModule instantiates a new CacheSettingsTieredCacheModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsTieredCacheModuleWithDefaults

`func NewCacheSettingsTieredCacheModuleWithDefaults() *CacheSettingsTieredCacheModule`

NewCacheSettingsTieredCacheModuleWithDefaults instantiates a new CacheSettingsTieredCacheModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *CacheSettingsTieredCacheModule) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *CacheSettingsTieredCacheModule) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *CacheSettingsTieredCacheModule) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.


### GetMaxAge

`func (o *CacheSettingsTieredCacheModule) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CacheSettingsTieredCacheModule) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CacheSettingsTieredCacheModule) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.


### GetTopology

`func (o *CacheSettingsTieredCacheModule) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *CacheSettingsTieredCacheModule) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *CacheSettingsTieredCacheModule) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *CacheSettingsTieredCacheModule) HasTopology() bool`

HasTopology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


