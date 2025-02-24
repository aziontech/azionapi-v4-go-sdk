# ResponseListDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**ClientId** | **string** |  | [readonly] 
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**DeletedAt** | **NullableTime** |  | [readonly] 
**IsActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewResponseListDatabase

`func NewResponseListDatabase(id int32, name string, clientId string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, ) *ResponseListDatabase`

NewResponseListDatabase instantiates a new ResponseListDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListDatabaseWithDefaults

`func NewResponseListDatabaseWithDefaults() *ResponseListDatabase`

NewResponseListDatabaseWithDefaults instantiates a new ResponseListDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListDatabase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListDatabase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListDatabase) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListDatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListDatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListDatabase) SetName(v string)`

SetName sets Name field to given value.


### GetClientId

`func (o *ResponseListDatabase) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ResponseListDatabase) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ResponseListDatabase) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetStatus

`func (o *ResponseListDatabase) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseListDatabase) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseListDatabase) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseListDatabase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseListDatabase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseListDatabase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseListDatabase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResponseListDatabase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResponseListDatabase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResponseListDatabase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *ResponseListDatabase) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ResponseListDatabase) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ResponseListDatabase) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *ResponseListDatabase) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ResponseListDatabase) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetIsActive

`func (o *ResponseListDatabase) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ResponseListDatabase) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ResponseListDatabase) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ResponseListDatabase) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


