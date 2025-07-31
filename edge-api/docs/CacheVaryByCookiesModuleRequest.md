# CacheVaryByCookiesModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to **string** | * &#x60;ignore&#x60; - ignore * &#x60;all&#x60; - all * &#x60;allowlist&#x60; - allowlist * &#x60;denylist&#x60; - denylist | [optional] 
**CookieNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCacheVaryByCookiesModuleRequest

`func NewCacheVaryByCookiesModuleRequest() *CacheVaryByCookiesModuleRequest`

NewCacheVaryByCookiesModuleRequest instantiates a new CacheVaryByCookiesModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheVaryByCookiesModuleRequestWithDefaults

`func NewCacheVaryByCookiesModuleRequestWithDefaults() *CacheVaryByCookiesModuleRequest`

NewCacheVaryByCookiesModuleRequestWithDefaults instantiates a new CacheVaryByCookiesModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *CacheVaryByCookiesModuleRequest) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *CacheVaryByCookiesModuleRequest) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *CacheVaryByCookiesModuleRequest) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *CacheVaryByCookiesModuleRequest) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetCookieNames

`func (o *CacheVaryByCookiesModuleRequest) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *CacheVaryByCookiesModuleRequest) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *CacheVaryByCookiesModuleRequest) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.

### HasCookieNames

`func (o *CacheVaryByCookiesModuleRequest) HasCookieNames() bool`

HasCookieNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


