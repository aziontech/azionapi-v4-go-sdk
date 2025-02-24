# DatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewDatabaseRequest

`func NewDatabaseRequest(name string, ) *DatabaseRequest`

NewDatabaseRequest instantiates a new DatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseRequestWithDefaults

`func NewDatabaseRequestWithDefaults() *DatabaseRequest`

NewDatabaseRequestWithDefaults instantiates a new DatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DatabaseRequest) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseRequest) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseRequest) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsActive

`func (o *DatabaseRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DatabaseRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DatabaseRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DatabaseRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


