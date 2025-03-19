# BrowserCacheModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | **string** | * &#x60;honor&#x60; - Honor Origin Cache Headers * &#x60;override&#x60; - Override Cache Settings * &#x60;no-cache&#x60; - No Cache | 
**MaxAge** | **int64** |  | 

## Methods

### NewBrowserCacheModuleRequest

`func NewBrowserCacheModuleRequest(behavior string, maxAge int64, ) *BrowserCacheModuleRequest`

NewBrowserCacheModuleRequest instantiates a new BrowserCacheModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserCacheModuleRequestWithDefaults

`func NewBrowserCacheModuleRequestWithDefaults() *BrowserCacheModuleRequest`

NewBrowserCacheModuleRequestWithDefaults instantiates a new BrowserCacheModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *BrowserCacheModuleRequest) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *BrowserCacheModuleRequest) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *BrowserCacheModuleRequest) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.


### GetMaxAge

`func (o *BrowserCacheModuleRequest) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *BrowserCacheModuleRequest) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *BrowserCacheModuleRequest) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


