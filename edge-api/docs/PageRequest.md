# PageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | * &#x60;default&#x60; - default * &#x60;400&#x60; - Bad Request * &#x60;401&#x60; - Unauthorized * &#x60;403&#x60; - Forbidden * &#x60;404&#x60; - Not Found * &#x60;405&#x60; - Method Not Allowed * &#x60;406&#x60; - Not Acceptable * &#x60;408&#x60; - Request Timeout * &#x60;409&#x60; - Conflict * &#x60;410&#x60; - Gone * &#x60;411&#x60; - Length Required * &#x60;414&#x60; - URI Too Long * &#x60;415&#x60; - Unsupported Media Type * &#x60;416&#x60; - Range Not Satisfiable * &#x60;426&#x60; - Upgrade Required * &#x60;429&#x60; - Too Many Requests * &#x60;431&#x60; - Request Header Fields Too Large * &#x60;500&#x60; - Internal Server Error * &#x60;501&#x60; - Not Implemented * &#x60;502&#x60; - Bad Gateway * &#x60;503&#x60; - Service Unavailable * &#x60;504&#x60; - Gateway Timeout * &#x60;505&#x60; - HTTP Version Not Supported | 
**Page** | [**PageConnectorRequest**](PageConnectorRequest.md) |  | 

## Methods

### NewPageRequest

`func NewPageRequest(code string, page PageConnectorRequest, ) *PageRequest`

NewPageRequest instantiates a new PageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRequestWithDefaults

`func NewPageRequestWithDefaults() *PageRequest`

NewPageRequestWithDefaults instantiates a new PageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PageRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PageRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PageRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetPage

`func (o *PageRequest) GetPage() PageConnectorRequest`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageRequest) GetPageOk() (*PageConnectorRequest, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageRequest) SetPage(v PageConnectorRequest)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


