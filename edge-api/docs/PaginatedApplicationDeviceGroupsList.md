# PaginatedApplicationDeviceGroupsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Total number of items | [optional] 
**TotalPages** | Pointer to **int64** | Total number of pages | [optional] 
**Page** | Pointer to **int64** | Current page number | [optional] 
**PageSize** | Pointer to **int64** | Number of items per page | [optional] 
**Results** | Pointer to [**[]ApplicationDeviceGroups**](ApplicationDeviceGroups.md) |  | [optional] 

## Methods

### NewPaginatedApplicationDeviceGroupsList

`func NewPaginatedApplicationDeviceGroupsList() *PaginatedApplicationDeviceGroupsList`

NewPaginatedApplicationDeviceGroupsList instantiates a new PaginatedApplicationDeviceGroupsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedApplicationDeviceGroupsListWithDefaults

`func NewPaginatedApplicationDeviceGroupsListWithDefaults() *PaginatedApplicationDeviceGroupsList`

NewPaginatedApplicationDeviceGroupsListWithDefaults instantiates a new PaginatedApplicationDeviceGroupsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedApplicationDeviceGroupsList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedApplicationDeviceGroupsList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedApplicationDeviceGroupsList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedApplicationDeviceGroupsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *PaginatedApplicationDeviceGroupsList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginatedApplicationDeviceGroupsList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginatedApplicationDeviceGroupsList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginatedApplicationDeviceGroupsList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetPage

`func (o *PaginatedApplicationDeviceGroupsList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedApplicationDeviceGroupsList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedApplicationDeviceGroupsList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedApplicationDeviceGroupsList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PaginatedApplicationDeviceGroupsList) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginatedApplicationDeviceGroupsList) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginatedApplicationDeviceGroupsList) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PaginatedApplicationDeviceGroupsList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedApplicationDeviceGroupsList) GetResults() []ApplicationDeviceGroups`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedApplicationDeviceGroupsList) GetResultsOk() (*[]ApplicationDeviceGroups, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedApplicationDeviceGroupsList) SetResults(v []ApplicationDeviceGroups)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedApplicationDeviceGroupsList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


