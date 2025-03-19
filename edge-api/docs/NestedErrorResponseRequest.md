# NestedErrorResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** | * &#x60;400&#x60; - 400: Bad Request * &#x60;401&#x60; - 401: Unauthorized * &#x60;403&#x60; - 403: Forbidden * &#x60;404&#x60; - 404: Not Found * &#x60;405&#x60; - 405: Method Not Allowed * &#x60;406&#x60; - 406: Not Acceptable * &#x60;408&#x60; - 408: Request Timeout * &#x60;409&#x60; - 409: Conflict * &#x60;410&#x60; - 410: Gone * &#x60;411&#x60; - 411: Length Required * &#x60;414&#x60; - 414: URI Too Long * &#x60;415&#x60; - 415: Unsupported Media Type * &#x60;416&#x60; - 416: Requested Range Not Satisfiable * &#x60;426&#x60; - 426: Upgrade Required * &#x60;429&#x60; - 429: Too Many Requests * &#x60;431&#x60; - 431: Request Header Fields Too Large * &#x60;500&#x60; - 500: Internal Server Error * &#x60;501&#x60; - 501: Not Implemented * &#x60;502&#x60; - 502: Bad Gateway * &#x60;503&#x60; - 503: Service Unavailable * &#x60;504&#x60; - 504: Gateway Timeout * &#x60;505&#x60; - 505: HTTP Version Not Supported * &#x60;any&#x60; - any | 
**Timeout** | **int64** |  | 
**Uri** | Pointer to **NullableString** |  | [optional] 
**CustomStatusCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNestedErrorResponseRequest

`func NewNestedErrorResponseRequest(code int64, timeout int64, ) *NestedErrorResponseRequest`

NewNestedErrorResponseRequest instantiates a new NestedErrorResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedErrorResponseRequestWithDefaults

`func NewNestedErrorResponseRequestWithDefaults() *NestedErrorResponseRequest`

NewNestedErrorResponseRequestWithDefaults instantiates a new NestedErrorResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NestedErrorResponseRequest) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NestedErrorResponseRequest) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NestedErrorResponseRequest) SetCode(v int64)`

SetCode sets Code field to given value.


### GetTimeout

`func (o *NestedErrorResponseRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NestedErrorResponseRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NestedErrorResponseRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.


### GetUri

`func (o *NestedErrorResponseRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NestedErrorResponseRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NestedErrorResponseRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NestedErrorResponseRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *NestedErrorResponseRequest) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *NestedErrorResponseRequest) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetCustomStatusCode

`func (o *NestedErrorResponseRequest) GetCustomStatusCode() string`

GetCustomStatusCode returns the CustomStatusCode field if non-nil, zero value otherwise.

### GetCustomStatusCodeOk

`func (o *NestedErrorResponseRequest) GetCustomStatusCodeOk() (*string, bool)`

GetCustomStatusCodeOk returns a tuple with the CustomStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusCode

`func (o *NestedErrorResponseRequest) SetCustomStatusCode(v string)`

SetCustomStatusCode sets CustomStatusCode field to given value.

### HasCustomStatusCode

`func (o *NestedErrorResponseRequest) HasCustomStatusCode() bool`

HasCustomStatusCode returns a boolean if a field has been set.

### SetCustomStatusCodeNil

`func (o *NestedErrorResponseRequest) SetCustomStatusCodeNil(b bool)`

 SetCustomStatusCodeNil sets the value for CustomStatusCode to be an explicit nil

### UnsetCustomStatusCode
`func (o *NestedErrorResponseRequest) UnsetCustomStatusCode()`

UnsetCustomStatusCode ensures that no value is present for CustomStatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


