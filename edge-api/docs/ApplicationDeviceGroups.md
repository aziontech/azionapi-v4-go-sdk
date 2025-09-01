# ApplicationDeviceGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**UserAgent** | **string** | Enter a valid regular expression pattern to identify user agents. | 

## Methods

### NewApplicationDeviceGroups

`func NewApplicationDeviceGroups(id int64, name string, userAgent string, ) *ApplicationDeviceGroups`

NewApplicationDeviceGroups instantiates a new ApplicationDeviceGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDeviceGroupsWithDefaults

`func NewApplicationDeviceGroupsWithDefaults() *ApplicationDeviceGroups`

NewApplicationDeviceGroupsWithDefaults instantiates a new ApplicationDeviceGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationDeviceGroups) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationDeviceGroups) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationDeviceGroups) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationDeviceGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationDeviceGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationDeviceGroups) SetName(v string)`

SetName sets Name field to given value.


### GetUserAgent

`func (o *ApplicationDeviceGroups) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ApplicationDeviceGroups) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ApplicationDeviceGroups) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


