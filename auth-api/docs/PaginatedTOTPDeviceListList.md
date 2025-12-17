# PaginatedTOTPDeviceListList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Total number of items | [optional] 
**TotalPages** | Pointer to **int64** | Total number of pages | [optional] 
**Page** | Pointer to **int64** | Current page number | [optional] 
**PageSize** | Pointer to **int64** | Number of items per page | [optional] 
**Results** | Pointer to [**[]TOTPDeviceList**](TOTPDeviceList.md) |  | [optional] 

## Methods

### NewPaginatedTOTPDeviceListList

`func NewPaginatedTOTPDeviceListList() *PaginatedTOTPDeviceListList`

NewPaginatedTOTPDeviceListList instantiates a new PaginatedTOTPDeviceListList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedTOTPDeviceListListWithDefaults

`func NewPaginatedTOTPDeviceListListWithDefaults() *PaginatedTOTPDeviceListList`

NewPaginatedTOTPDeviceListListWithDefaults instantiates a new PaginatedTOTPDeviceListList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedTOTPDeviceListList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedTOTPDeviceListList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedTOTPDeviceListList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedTOTPDeviceListList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *PaginatedTOTPDeviceListList) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginatedTOTPDeviceListList) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginatedTOTPDeviceListList) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PaginatedTOTPDeviceListList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetPage

`func (o *PaginatedTOTPDeviceListList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedTOTPDeviceListList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedTOTPDeviceListList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedTOTPDeviceListList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PaginatedTOTPDeviceListList) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginatedTOTPDeviceListList) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginatedTOTPDeviceListList) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PaginatedTOTPDeviceListList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedTOTPDeviceListList) GetResults() []TOTPDeviceList`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedTOTPDeviceListList) GetResultsOk() (*[]TOTPDeviceList, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedTOTPDeviceListList) SetResults(v []TOTPDeviceList)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedTOTPDeviceListList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


