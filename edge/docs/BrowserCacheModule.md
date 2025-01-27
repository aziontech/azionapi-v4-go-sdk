# BrowserCacheModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | [**BrowserCacheModuleBehaviorEnum**](BrowserCacheModuleBehaviorEnum.md) |  | 
**MaxAge** | **int32** |  | 

## Methods

### NewBrowserCacheModule

`func NewBrowserCacheModule(behavior BrowserCacheModuleBehaviorEnum, maxAge int32, ) *BrowserCacheModule`

NewBrowserCacheModule instantiates a new BrowserCacheModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserCacheModuleWithDefaults

`func NewBrowserCacheModuleWithDefaults() *BrowserCacheModule`

NewBrowserCacheModuleWithDefaults instantiates a new BrowserCacheModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *BrowserCacheModule) GetBehavior() BrowserCacheModuleBehaviorEnum`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *BrowserCacheModule) GetBehaviorOk() (*BrowserCacheModuleBehaviorEnum, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *BrowserCacheModule) SetBehavior(v BrowserCacheModuleBehaviorEnum)`

SetBehavior sets Behavior field to given value.


### GetMaxAge

`func (o *BrowserCacheModule) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *BrowserCacheModule) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *BrowserCacheModule) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


