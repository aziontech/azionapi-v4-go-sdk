# \ApplicationsDeviceGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceGroup**](ApplicationsDeviceGroupsAPI.md#CreateDeviceGroup) | **Post** /workspace/applications/{application_id}/device_groups | Create an Applications Device Group
[**DeleteDeviceGroups**](ApplicationsDeviceGroupsAPI.md#DeleteDeviceGroups) | **Delete** /workspace/applications/{application_id}/device_groups/{device_group_id} | Delete an Applications Device Group
[**ListDeviceGroups**](ApplicationsDeviceGroupsAPI.md#ListDeviceGroups) | **Get** /workspace/applications/{application_id}/device_groups | List Applications Device Groups
[**PartialUpdateDeviceGroup**](ApplicationsDeviceGroupsAPI.md#PartialUpdateDeviceGroup) | **Patch** /workspace/applications/{application_id}/device_groups/{device_group_id} | Partially update an Applications Device Group
[**RetrieveDeviceGroup**](ApplicationsDeviceGroupsAPI.md#RetrieveDeviceGroup) | **Get** /workspace/applications/{application_id}/device_groups/{device_group_id} | Retrieve details of a Device Group
[**UpdateDeviceGroup**](ApplicationsDeviceGroupsAPI.md#UpdateDeviceGroup) | **Put** /workspace/applications/{application_id}/device_groups/{device_group_id} | Update an Applications Device Group



## CreateDeviceGroup

> ResponseApplicationDeviceGroups CreateDeviceGroup(ctx, applicationId).ApplicationDeviceGroupsRequest(applicationDeviceGroupsRequest).Execute()

Create an Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	applicationDeviceGroupsRequest := *openapiclient.NewApplicationDeviceGroupsRequest("Name_example", "UserAgent_example") // ApplicationDeviceGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsDeviceGroupsAPI.CreateDeviceGroup(context.Background(), applicationId).ApplicationDeviceGroupsRequest(applicationDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsDeviceGroupsAPI.CreateDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceGroup`: ResponseApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsDeviceGroupsAPI.CreateDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDeviceGroupsRequest** | [**ApplicationDeviceGroupsRequest**](ApplicationDeviceGroupsRequest.md) |  | 

### Return type

[**ResponseApplicationDeviceGroups**](ResponseApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceGroups

> ResponseAsyncDeleteApplicationDeviceGroups DeleteDeviceGroups(ctx, applicationId, deviceGroupId).Execute()

Delete an Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	deviceGroupId := int64(789) // int64 | A unique integer value identifying the device group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsDeviceGroupsAPI.DeleteDeviceGroups(context.Background(), applicationId, deviceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsDeviceGroupsAPI.DeleteDeviceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeviceGroups`: ResponseAsyncDeleteApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsDeviceGroupsAPI.DeleteDeviceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**deviceGroupId** | **int64** | A unique integer value identifying the device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseAsyncDeleteApplicationDeviceGroups**](ResponseAsyncDeleteApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceGroups

> PaginatedApplicationDeviceGroupsList ListDeviceGroups(ctx, applicationId).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserAgent(userAgent).Execute()

List Applications Device Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := int64(789) // int64 | Filter by id (accepts comma-separated values). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id, user_agent) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)
	userAgent := "userAgent_example" // string | Filter by user agent (case-insensitive, partial match). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsDeviceGroupsAPI.ListDeviceGroups(context.Background(), applicationId).Fields(fields).Id(id).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserAgent(userAgent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsDeviceGroupsAPI.ListDeviceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceGroups`: PaginatedApplicationDeviceGroupsList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsDeviceGroupsAPI.ListDeviceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by id (accepts comma-separated values). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id, user_agent) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 
 **userAgent** | **string** | Filter by user agent (case-insensitive, partial match). | 

### Return type

[**PaginatedApplicationDeviceGroupsList**](PaginatedApplicationDeviceGroupsList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDeviceGroup

> ResponseApplicationDeviceGroups PartialUpdateDeviceGroup(ctx, applicationId, deviceGroupId).PatchedApplicationDeviceGroupsRequest(patchedApplicationDeviceGroupsRequest).Execute()

Partially update an Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	deviceGroupId := int64(789) // int64 | A unique integer value identifying the device group.
	patchedApplicationDeviceGroupsRequest := *openapiclient.NewPatchedApplicationDeviceGroupsRequest() // PatchedApplicationDeviceGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsDeviceGroupsAPI.PartialUpdateDeviceGroup(context.Background(), applicationId, deviceGroupId).PatchedApplicationDeviceGroupsRequest(patchedApplicationDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsDeviceGroupsAPI.PartialUpdateDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDeviceGroup`: ResponseApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsDeviceGroupsAPI.PartialUpdateDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**deviceGroupId** | **int64** | A unique integer value identifying the device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedApplicationDeviceGroupsRequest** | [**PatchedApplicationDeviceGroupsRequest**](PatchedApplicationDeviceGroupsRequest.md) |  | 

### Return type

[**ResponseApplicationDeviceGroups**](ResponseApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDeviceGroup

> ResponseRetrieveApplicationDeviceGroups RetrieveDeviceGroup(ctx, applicationId, deviceGroupId).Fields(fields).Execute()

Retrieve details of a Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	deviceGroupId := int64(789) // int64 | A unique integer value identifying the device group.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsDeviceGroupsAPI.RetrieveDeviceGroup(context.Background(), applicationId, deviceGroupId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsDeviceGroupsAPI.RetrieveDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDeviceGroup`: ResponseRetrieveApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsDeviceGroupsAPI.RetrieveDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**deviceGroupId** | **int64** | A unique integer value identifying the device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveApplicationDeviceGroups**](ResponseRetrieveApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceGroup

> ResponseApplicationDeviceGroups UpdateDeviceGroup(ctx, applicationId, deviceGroupId).ApplicationDeviceGroupsRequest(applicationDeviceGroupsRequest).Execute()

Update an Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	deviceGroupId := int64(789) // int64 | A unique integer value identifying the device group.
	applicationDeviceGroupsRequest := *openapiclient.NewApplicationDeviceGroupsRequest("Name_example", "UserAgent_example") // ApplicationDeviceGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsDeviceGroupsAPI.UpdateDeviceGroup(context.Background(), applicationId, deviceGroupId).ApplicationDeviceGroupsRequest(applicationDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsDeviceGroupsAPI.UpdateDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceGroup`: ResponseApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsDeviceGroupsAPI.UpdateDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**deviceGroupId** | **int64** | A unique integer value identifying the device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationDeviceGroupsRequest** | [**ApplicationDeviceGroupsRequest**](ApplicationDeviceGroupsRequest.md) |  | 

### Return type

[**ResponseApplicationDeviceGroups**](ResponseApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

