# ResellerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CurrencyIsoCode** | **string** | * &#x60;USD&#x60; - USD * &#x60;BRL&#x60; - BRL | 
**TermsOfServiceUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewResellerRequest

`func NewResellerRequest(name string, currencyIsoCode string, ) *ResellerRequest`

NewResellerRequest instantiates a new ResellerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResellerRequestWithDefaults

`func NewResellerRequestWithDefaults() *ResellerRequest`

NewResellerRequestWithDefaults instantiates a new ResellerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResellerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResellerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResellerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCurrencyIsoCode

`func (o *ResellerRequest) GetCurrencyIsoCode() string`

GetCurrencyIsoCode returns the CurrencyIsoCode field if non-nil, zero value otherwise.

### GetCurrencyIsoCodeOk

`func (o *ResellerRequest) GetCurrencyIsoCodeOk() (*string, bool)`

GetCurrencyIsoCodeOk returns a tuple with the CurrencyIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode

`func (o *ResellerRequest) SetCurrencyIsoCode(v string)`

SetCurrencyIsoCode sets CurrencyIsoCode field to given value.


### GetTermsOfServiceUrl

`func (o *ResellerRequest) GetTermsOfServiceUrl() string`

GetTermsOfServiceUrl returns the TermsOfServiceUrl field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlOk

`func (o *ResellerRequest) GetTermsOfServiceUrlOk() (*string, bool)`

GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrl

`func (o *ResellerRequest) SetTermsOfServiceUrl(v string)`

SetTermsOfServiceUrl sets TermsOfServiceUrl field to given value.

### HasTermsOfServiceUrl

`func (o *ResellerRequest) HasTermsOfServiceUrl() bool`

HasTermsOfServiceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


