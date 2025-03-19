# \WorkloadDeploymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListWorkloadDeployments**](WorkloadDeploymentsAPI.md#ListWorkloadDeployments) | **Get** /workspace/workloads/{workloadId}/deployments | List Workload Deployments
[**PartialUpdateWorkloadDeployment**](WorkloadDeploymentsAPI.md#PartialUpdateWorkloadDeployment) | **Patch** /workspace/workloads/{workloadId}/deployments/{id} | Partially update a Workload Deployment
[**RetrieveWorkloadDeployment**](WorkloadDeploymentsAPI.md#RetrieveWorkloadDeployment) | **Get** /workspace/workloads/{workloadId}/deployments/{id} | Retrieve details of a Workload Deployment
[**UpdateWorkloadDeployment**](WorkloadDeploymentsAPI.md#UpdateWorkloadDeployment) | **Put** /workspace/workloads/{workloadId}/deployments/{id} | Update a Workload Deployment



## ListWorkloadDeployments

> PaginatedResponseListWorkloadDeploymentList ListWorkloadDeployments(ctx, workloadId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Workload Deployments



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
	workloadId := "workloadId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, tag, current) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.ListWorkloadDeployments(context.Background(), workloadId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDeploymentsAPI.ListWorkloadDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkloadDeployments`: PaginatedResponseListWorkloadDeploymentList
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.ListWorkloadDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkloadDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, tag, current) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListWorkloadDeploymentList**](PaginatedResponseListWorkloadDeploymentList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWorkloadDeployment

> ResponseWorkloadDeployment PartialUpdateWorkloadDeployment(ctx, id, workloadId).PatchedWorkloadDeploymentRequest(patchedWorkloadDeploymentRequest).Execute()

Partially update a Workload Deployment



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
	id := "id_example" // string | 
	workloadId := "workloadId_example" // string | 
	patchedWorkloadDeploymentRequest := *openapiclient.NewPatchedWorkloadDeploymentRequest() // PatchedWorkloadDeploymentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.PartialUpdateWorkloadDeployment(context.Background(), id, workloadId).PatchedWorkloadDeploymentRequest(patchedWorkloadDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDeploymentsAPI.PartialUpdateWorkloadDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateWorkloadDeployment`: ResponseWorkloadDeployment
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.PartialUpdateWorkloadDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**workloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWorkloadDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedWorkloadDeploymentRequest** | [**PatchedWorkloadDeploymentRequest**](PatchedWorkloadDeploymentRequest.md) |  | 

### Return type

[**ResponseWorkloadDeployment**](ResponseWorkloadDeployment.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWorkloadDeployment

> ResponseRetrieveWorkloadDeployment RetrieveWorkloadDeployment(ctx, id, workloadId).Fields(fields).Execute()

Retrieve details of a Workload Deployment



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
	id := "id_example" // string | 
	workloadId := "workloadId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.RetrieveWorkloadDeployment(context.Background(), id, workloadId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDeploymentsAPI.RetrieveWorkloadDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWorkloadDeployment`: ResponseRetrieveWorkloadDeployment
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.RetrieveWorkloadDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**workloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWorkloadDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveWorkloadDeployment**](ResponseRetrieveWorkloadDeployment.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkloadDeployment

> ResponseWorkloadDeployment UpdateWorkloadDeployment(ctx, id, workloadId).WorkloadDeploymentRequest(workloadDeploymentRequest).Execute()

Update a Workload Deployment



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
	id := "id_example" // string | 
	workloadId := "workloadId_example" // string | 
	workloadDeploymentRequest := *openapiclient.NewWorkloadDeploymentRequest("Tag_example", *openapiclient.NewWorkloadDeploymentBindsRequest(int64(123))) // WorkloadDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.UpdateWorkloadDeployment(context.Background(), id, workloadId).WorkloadDeploymentRequest(workloadDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDeploymentsAPI.UpdateWorkloadDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkloadDeployment`: ResponseWorkloadDeployment
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.UpdateWorkloadDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**workloadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkloadDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workloadDeploymentRequest** | [**WorkloadDeploymentRequest**](WorkloadDeploymentRequest.md) |  | 

### Return type

[**ResponseWorkloadDeployment**](ResponseWorkloadDeployment.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

