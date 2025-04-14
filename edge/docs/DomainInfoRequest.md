# DomainInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** |  | [optional] 
**AllowAccess** | **bool** |  | 

## Methods

### NewDomainInfoRequest

`func NewDomainInfoRequest(allowAccess bool, ) *DomainInfoRequest`

NewDomainInfoRequest instantiates a new DomainInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInfoRequestWithDefaults

`func NewDomainInfoRequestWithDefaults() *DomainInfoRequest`

NewDomainInfoRequestWithDefaults instantiates a new DomainInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainInfoRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainInfoRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainInfoRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainInfoRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *DomainInfoRequest) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *DomainInfoRequest) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetAllowAccess

`func (o *DomainInfoRequest) GetAllowAccess() bool`

GetAllowAccess returns the AllowAccess field if non-nil, zero value otherwise.

### GetAllowAccessOk

`func (o *DomainInfoRequest) GetAllowAccessOk() (*bool, bool)`

GetAllowAccessOk returns a tuple with the AllowAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccess

`func (o *DomainInfoRequest) SetAllowAccess(v bool)`

SetAllowAccess sets AllowAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


