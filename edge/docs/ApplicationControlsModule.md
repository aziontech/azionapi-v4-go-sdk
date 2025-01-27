# ApplicationControlsModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheByQueryString** | [**CacheByQueryStringEnum**](CacheByQueryStringEnum.md) |  | 
**QueryStringFields** | **[]string** |  | 
**QueryStringSortEnabled** | **bool** |  | 
**CacheByCookies** | [**CacheByCookiesEnum**](CacheByCookiesEnum.md) |  | 
**CookieNames** | **[]string** |  | 
**AdaptiveDeliveryAction** | [**AdaptiveDeliveryActionEnum**](AdaptiveDeliveryActionEnum.md) |  | 
**DeviceGroup** | **[]int32** |  | 

## Methods

### NewApplicationControlsModule

`func NewApplicationControlsModule(cacheByQueryString CacheByQueryStringEnum, queryStringFields []string, queryStringSortEnabled bool, cacheByCookies CacheByCookiesEnum, cookieNames []string, adaptiveDeliveryAction AdaptiveDeliveryActionEnum, deviceGroup []int32, ) *ApplicationControlsModule`

NewApplicationControlsModule instantiates a new ApplicationControlsModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationControlsModuleWithDefaults

`func NewApplicationControlsModuleWithDefaults() *ApplicationControlsModule`

NewApplicationControlsModuleWithDefaults instantiates a new ApplicationControlsModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheByQueryString

`func (o *ApplicationControlsModule) GetCacheByQueryString() CacheByQueryStringEnum`

GetCacheByQueryString returns the CacheByQueryString field if non-nil, zero value otherwise.

### GetCacheByQueryStringOk

`func (o *ApplicationControlsModule) GetCacheByQueryStringOk() (*CacheByQueryStringEnum, bool)`

GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByQueryString

`func (o *ApplicationControlsModule) SetCacheByQueryString(v CacheByQueryStringEnum)`

SetCacheByQueryString sets CacheByQueryString field to given value.


### GetQueryStringFields

`func (o *ApplicationControlsModule) GetQueryStringFields() []string`

GetQueryStringFields returns the QueryStringFields field if non-nil, zero value otherwise.

### GetQueryStringFieldsOk

`func (o *ApplicationControlsModule) GetQueryStringFieldsOk() (*[]string, bool)`

GetQueryStringFieldsOk returns a tuple with the QueryStringFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStringFields

`func (o *ApplicationControlsModule) SetQueryStringFields(v []string)`

SetQueryStringFields sets QueryStringFields field to given value.


### GetQueryStringSortEnabled

`func (o *ApplicationControlsModule) GetQueryStringSortEnabled() bool`

GetQueryStringSortEnabled returns the QueryStringSortEnabled field if non-nil, zero value otherwise.

### GetQueryStringSortEnabledOk

`func (o *ApplicationControlsModule) GetQueryStringSortEnabledOk() (*bool, bool)`

GetQueryStringSortEnabledOk returns a tuple with the QueryStringSortEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStringSortEnabled

`func (o *ApplicationControlsModule) SetQueryStringSortEnabled(v bool)`

SetQueryStringSortEnabled sets QueryStringSortEnabled field to given value.


### GetCacheByCookies

`func (o *ApplicationControlsModule) GetCacheByCookies() CacheByCookiesEnum`

GetCacheByCookies returns the CacheByCookies field if non-nil, zero value otherwise.

### GetCacheByCookiesOk

`func (o *ApplicationControlsModule) GetCacheByCookiesOk() (*CacheByCookiesEnum, bool)`

GetCacheByCookiesOk returns a tuple with the CacheByCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByCookies

`func (o *ApplicationControlsModule) SetCacheByCookies(v CacheByCookiesEnum)`

SetCacheByCookies sets CacheByCookies field to given value.


### GetCookieNames

`func (o *ApplicationControlsModule) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *ApplicationControlsModule) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *ApplicationControlsModule) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.


### GetAdaptiveDeliveryAction

`func (o *ApplicationControlsModule) GetAdaptiveDeliveryAction() AdaptiveDeliveryActionEnum`

GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field if non-nil, zero value otherwise.

### GetAdaptiveDeliveryActionOk

`func (o *ApplicationControlsModule) GetAdaptiveDeliveryActionOk() (*AdaptiveDeliveryActionEnum, bool)`

GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveDeliveryAction

`func (o *ApplicationControlsModule) SetAdaptiveDeliveryAction(v AdaptiveDeliveryActionEnum)`

SetAdaptiveDeliveryAction sets AdaptiveDeliveryAction field to given value.


### GetDeviceGroup

`func (o *ApplicationControlsModule) GetDeviceGroup() []int32`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *ApplicationControlsModule) GetDeviceGroupOk() (*[]int32, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *ApplicationControlsModule) SetDeviceGroup(v []int32)`

SetDeviceGroup sets DeviceGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


