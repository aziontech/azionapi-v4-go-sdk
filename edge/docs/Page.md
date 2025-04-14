# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | * &#x60;default&#x60; - default * &#x60;400&#x60; - Bad Request * &#x60;401&#x60; - Unauthorized * &#x60;403&#x60; - Forbidden * &#x60;404&#x60; - Not Found * &#x60;405&#x60; - Method Not Allowed * &#x60;406&#x60; - Not Acceptable * &#x60;408&#x60; - Request Timeout * &#x60;409&#x60; - Conflict * &#x60;410&#x60; - Gone * &#x60;411&#x60; - Length Required * &#x60;414&#x60; - URI Too Long * &#x60;415&#x60; - Unsupported Media Type * &#x60;416&#x60; - Range Not Satisfiable * &#x60;426&#x60; - Upgrade Required * &#x60;429&#x60; - Too Many Requests * &#x60;431&#x60; - Request Header Fields Too Large * &#x60;500&#x60; - Internal Server Error * &#x60;501&#x60; - Not Implemented * &#x60;502&#x60; - Bad Gateway * &#x60;503&#x60; - Service Unavailable * &#x60;504&#x60; - Gateway Timeout * &#x60;505&#x60; - HTTP Version Not Supported | 
**Ttl** | Pointer to **int64** |  | [optional] 
**Uri** | Pointer to **NullableString** |  | [optional] 
**CustomStatusCode** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewPage

`func NewPage(code string, ) *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Page) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Page) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Page) SetCode(v string)`

SetCode sets Code field to given value.


### GetTtl

`func (o *Page) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Page) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Page) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Page) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUri

`func (o *Page) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Page) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Page) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Page) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *Page) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *Page) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetCustomStatusCode

`func (o *Page) GetCustomStatusCode() int64`

GetCustomStatusCode returns the CustomStatusCode field if non-nil, zero value otherwise.

### GetCustomStatusCodeOk

`func (o *Page) GetCustomStatusCodeOk() (*int64, bool)`

GetCustomStatusCodeOk returns a tuple with the CustomStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusCode

`func (o *Page) SetCustomStatusCode(v int64)`

SetCustomStatusCode sets CustomStatusCode field to given value.

### HasCustomStatusCode

`func (o *Page) HasCustomStatusCode() bool`

HasCustomStatusCode returns a boolean if a field has been set.

### SetCustomStatusCodeNil

`func (o *Page) SetCustomStatusCodeNil(b bool)`

 SetCustomStatusCodeNil sets the value for CustomStatusCode to be an explicit nil

### UnsetCustomStatusCode
`func (o *Page) UnsetCustomStatusCode()`

UnsetCustomStatusCode ensures that no value is present for CustomStatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


