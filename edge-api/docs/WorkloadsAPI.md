# \WorkloadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkload**](WorkloadsAPI.md#CreateWorkload) | **Post** /workspace/workloads | Create an Workload
[**DestroyWorkload**](WorkloadsAPI.md#DestroyWorkload) | **Delete** /workspace/workloads/{globalId} | Destroy an Workload
[**ListWorkloads**](WorkloadsAPI.md#ListWorkloads) | **Get** /workspace/workloads | List Workloads
[**PartialUpdateWorkload**](WorkloadsAPI.md#PartialUpdateWorkload) | **Patch** /workspace/workloads/{globalId} | Partially update an Workload
[**RetrieveWorkload**](WorkloadsAPI.md#RetrieveWorkload) | **Get** /workspace/workloads/{globalId} | Retrieve details of an Workload
[**UpdateWorkload**](WorkloadsAPI.md#UpdateWorkload) | **Put** /workspace/workloads/{globalId} | Update an Workload



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
	openapiclient "openapiclient "github.com/aziontech/azionapi-v4-go-sdk""
)

func main() {
	workloadRequest := *openapiclient.NewWorkloadRequest("Name_example", int64(123)) // WorkloadRequest | 

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

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWorkload

> ResponseDeleteWorkload DestroyWorkload(ctx, globalId).Execute()

Destroy an Workload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "openapiclient "github.com/aziontech/azionapi-v4-go-sdk""
)

func main() {
	globalId := "globalId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.DestroyWorkload(context.Background(), globalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.DestroyWorkload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyWorkload`: ResponseDeleteWorkload
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.DestroyWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteWorkload**](ResponseDeleteWorkload.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkloads

> PaginatedResponseListWorkloadList ListWorkloads(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Workloads



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "openapiclient "github.com/aziontech/azionapi-v4-go-sdk""
)

func main() {
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: edge_application, edge_firewall, id, name, last_editor, last_modified, active, alternate_domains, network_map, domains, product_version) (optional)
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
	// response from `ListWorkloads`: PaginatedResponseListWorkloadList
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.ListWorkloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: edge_application, edge_firewall, id, name, last_editor, last_modified, active, alternate_domains, network_map, domains, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListWorkloadList**](PaginatedResponseListWorkloadList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWorkload

> ResponseWorkload PartialUpdateWorkload(ctx, globalId).PatchedWorkloadRequest(patchedWorkloadRequest).Execute()

Partially update an Workload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "openapiclient "github.com/aziontech/azionapi-v4-go-sdk""
)

func main() {
	globalId := "globalId_example" // string | 
	patchedWorkloadRequest := *openapiclient.NewPatchedWorkloadRequest() // PatchedWorkloadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.PartialUpdateWorkload(context.Background(), globalId).PatchedWorkloadRequest(patchedWorkloadRequest).Execute()
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
**globalId** | **string** |  | 

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

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWorkload

> ResponseRetrieveWorkload RetrieveWorkload(ctx, globalId).Fields(fields).Execute()

Retrieve details of an Workload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "openapiclient "github.com/aziontech/azionapi-v4-go-sdk""
)

func main() {
	globalId := "globalId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.RetrieveWorkload(context.Background(), globalId).Fields(fields).Execute()
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
**globalId** | **string** |  | 

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

> ResponseWorkload UpdateWorkload(ctx, globalId).WorkloadRequest(workloadRequest).Execute()

Update an Workload



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "openapiclient "github.com/aziontech/azionapi-v4-go-sdk""
)

func main() {
	globalId := "globalId_example" // string | 
	workloadRequest := *openapiclient.NewWorkloadRequest("Name_example", int64(123)) // WorkloadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsAPI.UpdateWorkload(context.Background(), globalId).WorkloadRequest(workloadRequest).Execute()
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
**globalId** | **string** |  | 

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

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

