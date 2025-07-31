# CacheVaryByCookiesModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to **string** | * &#x60;ignore&#x60; - ignore * &#x60;all&#x60; - all * &#x60;allowlist&#x60; - allowlist * &#x60;denylist&#x60; - denylist | [optional] 
**CookieNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCacheVaryByCookiesModule

`func NewCacheVaryByCookiesModule() *CacheVaryByCookiesModule`

NewCacheVaryByCookiesModule instantiates a new CacheVaryByCookiesModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheVaryByCookiesModuleWithDefaults

`func NewCacheVaryByCookiesModuleWithDefaults() *CacheVaryByCookiesModule`

NewCacheVaryByCookiesModuleWithDefaults instantiates a new CacheVaryByCookiesModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *CacheVaryByCookiesModule) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *CacheVaryByCookiesModule) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *CacheVaryByCookiesModule) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *CacheVaryByCookiesModule) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetCookieNames

`func (o *CacheVaryByCookiesModule) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *CacheVaryByCookiesModule) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *CacheVaryByCookiesModule) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.

### HasCookieNames

`func (o *CacheVaryByCookiesModule) HasCookieNames() bool`

HasCookieNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


