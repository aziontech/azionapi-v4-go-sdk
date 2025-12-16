# \WorkloadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkload**](WorkloadsAPI.md#CreateWorkload) | **Post** /workspace/workloads | Create an Workload
[**DeleteWorkload**](WorkloadsAPI.md#DeleteWorkload) | **Delete** /workspace/workloads/{workload_id} | Delete an Workload
[**ListWorkloads**](WorkloadsAPI.md#ListWorkloads) | **Get** /workspace/workloads | List Workloads
[**PartialUpdateWorkload**](WorkloadsAPI.md#PartialUpdateWorkload) | **Patch** /workspace/workloads/{workload_id} | Partially update an Workload
[**RetrieveWorkload**](WorkloadsAPI.md#RetrieveWorkload) | **Get** /workspace/workloads/{workload_id} | Retrieve details of an Workload
[**UpdateWorkload**](WorkloadsAPI.md#UpdateWorkload) | **Put** /workspace/workloads/{workload_id} | Update an Workload



## CreateWorkload

> ResponseWorkload CreateWorkload(ctx).WorkloadRequest(workloadRequest).Execute()

Create an Workload



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
	workloadRequest := *openapiclient.NewWorkloadRequest("Name_example") // WorkloadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.CreateWorkload(context.Background()).WorkloadRequest(workloadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.CreateWorkload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkload`: ResponseWorkload
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.CreateWorkload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workloadRequest** | [**WorkloadRequest**](WorkloadRequest.md) |  | 

### Return type

[**ResponseWorkload**](ResponseWorkload.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkload

> ResponseAsyncDeleteWorkload DeleteWorkload(ctx, workloadId).Execute()

Delete an Workload



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
	workloadId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.DeleteWorkload(context.Background(), workloadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.DeleteWorkload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkload`: ResponseAsyncDeleteWorkload
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.DeleteWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseAsyncDeleteWorkload**](ResponseAsyncDeleteWorkload.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkloads

> PaginatedWorkloadList ListWorkloads(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Workloads



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, last_editor, last_modified, active, workload_domain_allow_access, workload_domain, infrastructure, domains, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.ListWorkloads(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.ListWorkloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkloads`: PaginatedWorkloadList
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.ListWorkloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, last_editor, last_modified, active, workload_domain_allow_access, workload_domain, infrastructure, domains, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedWorkloadList**](PaginatedWorkloadList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWorkload

> ResponseWorkload PartialUpdateWorkload(ctx, workloadId).PatchedWorkloadRequest(patchedWorkloadRequest).Execute()

Partially update an Workload



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
	workloadId := int64(789) // int64 | 
	patchedWorkloadRequest := *openapiclient.NewPatchedWorkloadRequest() // PatchedWorkloadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.PartialUpdateWorkload(context.Background(), workloadId).PatchedWorkloadRequest(patchedWorkloadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.PartialUpdateWorkload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateWorkload`: ResponseWorkload
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.PartialUpdateWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWorkloadRequest** | [**PatchedWorkloadRequest**](PatchedWorkloadRequest.md) |  | 

### Return type

[**ResponseWorkload**](ResponseWorkload.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWorkload

> ResponseRetrieveWorkload RetrieveWorkload(ctx, workloadId).Fields(fields).Execute()

Retrieve details of an Workload



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
	workloadId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.RetrieveWorkload(context.Background(), workloadId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.RetrieveWorkload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWorkload`: ResponseRetrieveWorkload
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.RetrieveWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveWorkload**](ResponseRetrieveWorkload.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkload

> ResponseWorkload UpdateWorkload(ctx, workloadId).WorkloadRequest(workloadRequest).Execute()

Update an Workload



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
	workloadId := int64(789) // int64 | 
	workloadRequest := *openapiclient.NewWorkloadRequest("Name_example") // WorkloadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.UpdateWorkload(context.Background(), workloadId).WorkloadRequest(workloadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.UpdateWorkload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkload`: ResponseWorkload
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.UpdateWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workloadRequest** | [**WorkloadRequest**](WorkloadRequest.md) |  | 

### Return type

[**ResponseWorkload**](ResponseWorkload.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

