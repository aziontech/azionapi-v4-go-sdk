# SplunkEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**ApiKey** | **string** |  | 
**Type** | **string** | Type identifier for this endpoint (splunk) | 

## Methods

### NewSplunkEndpoint

`func NewSplunkEndpoint(url string, apiKey string, type_ string, ) *SplunkEndpoint`

NewSplunkEndpoint instantiates a new SplunkEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplunkEndpointWithDefaults

`func NewSplunkEndpointWithDefaults() *SplunkEndpoint`

NewSplunkEndpointWithDefaults instantiates a new SplunkEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SplunkEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SplunkEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SplunkEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetApiKey

`func (o *SplunkEndpoint) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SplunkEndpoint) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SplunkEndpoint) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetType

`func (o *SplunkEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SplunkEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SplunkEndpoint) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


