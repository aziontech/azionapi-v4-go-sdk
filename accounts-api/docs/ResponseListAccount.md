# ResponseListAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Active** | **bool** |  | [readonly] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ParentId** | **int32** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Info** | **map[string]map[string]interface{}** |  | [readonly] 
**Status** | [**StatusEnum**](StatusEnum.md) |  | [readonly] 
**Reason** | [**ReasonEnum**](ReasonEnum.md) |  | [readonly] 
**CurrencyIsoCode** | [**CurrencyIsoCodeEnum**](CurrencyIsoCodeEnum.md) |  | 
**TermsOfServiceUrl** | Pointer to **string** |  | [optional] [default to "https://www.azion.com/pt-br/documentacao/contratos/tds/"]
**WorkspaceId** | **string** |  | [readonly] 

## Methods

### NewResponseListAccount

`func NewResponseListAccount(id int32, name string, active bool, lastEditor string, lastModified time.Time, parentId int32, created time.Time, info map[string]map[string]interface{}, status StatusEnum, reason ReasonEnum, currencyIsoCode CurrencyIsoCodeEnum, workspaceId string, ) *ResponseListAccount`

NewResponseListAccount instantiates a new ResponseListAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListAccountWithDefaults

`func NewResponseListAccountWithDefaults() *ResponseListAccount`

NewResponseListAccountWithDefaults instantiates a new ResponseListAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListAccount) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListAccount) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ResponseListAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListAccount) SetActive(v bool)`

SetActive sets Active field to given value.


### GetLastEditor

`func (o *ResponseListAccount) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListAccount) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListAccount) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListAccount) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListAccount) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListAccount) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetParentId

`func (o *ResponseListAccount) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ResponseListAccount) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ResponseListAccount) SetParentId(v int32)`

SetParentId sets ParentId field to given value.


### GetCreated

`func (o *ResponseListAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResponseListAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResponseListAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetInfo

`func (o *ResponseListAccount) GetInfo() map[string]map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ResponseListAccount) GetInfoOk() (*map[string]map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ResponseListAccount) SetInfo(v map[string]map[string]interface{})`

SetInfo sets Info field to given value.


### GetStatus

`func (o *ResponseListAccount) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseListAccount) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseListAccount) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetReason

`func (o *ResponseListAccount) GetReason() ReasonEnum`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseListAccount) GetReasonOk() (*ReasonEnum, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseListAccount) SetReason(v ReasonEnum)`

SetReason sets Reason field to given value.


### GetCurrencyIsoCode

`func (o *ResponseListAccount) GetCurrencyIsoCode() CurrencyIsoCodeEnum`

GetCurrencyIsoCode returns the CurrencyIsoCode field if non-nil, zero value otherwise.

### GetCurrencyIsoCodeOk

`func (o *ResponseListAccount) GetCurrencyIsoCodeOk() (*CurrencyIsoCodeEnum, bool)`

GetCurrencyIsoCodeOk returns a tuple with the CurrencyIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIsoCode

`func (o *ResponseListAccount) SetCurrencyIsoCode(v CurrencyIsoCodeEnum)`

SetCurrencyIsoCode sets CurrencyIsoCode field to given value.


### GetTermsOfServiceUrl

`func (o *ResponseListAccount) GetTermsOfServiceUrl() string`

GetTermsOfServiceUrl returns the TermsOfServiceUrl field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlOk

`func (o *ResponseListAccount) GetTermsOfServiceUrlOk() (*string, bool)`

GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrl

`func (o *ResponseListAccount) SetTermsOfServiceUrl(v string)`

SetTermsOfServiceUrl sets TermsOfServiceUrl field to given value.

### HasTermsOfServiceUrl

`func (o *ResponseListAccount) HasTermsOfServiceUrl() bool`

HasTermsOfServiceUrl returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *ResponseListAccount) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ResponseListAccount) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ResponseListAccount) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


