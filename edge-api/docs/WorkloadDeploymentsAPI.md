# \WorkloadDeploymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkloadDeployment**](WorkloadDeploymentsAPI.md#CreateWorkloadDeployment) | **Post** /workspace/workloads/{workload_id}/deployments | Create a Workload Deployment
[**DeleteWorkloadDeployment**](WorkloadDeploymentsAPI.md#DeleteWorkloadDeployment) | **Delete** /workspace/workloads/{workload_id}/deployments/{deployment_id} | Delete a Workload Deployment
[**ListWorkloadDeployments**](WorkloadDeploymentsAPI.md#ListWorkloadDeployments) | **Get** /workspace/workloads/{workload_id}/deployments | List Workload Deployments
[**PartialUpdateWorkloadDeployment**](WorkloadDeploymentsAPI.md#PartialUpdateWorkloadDeployment) | **Patch** /workspace/workloads/{workload_id}/deployments/{deployment_id} | Partially update a Workload Deployment
[**RetrieveWorkloadDeployment**](WorkloadDeploymentsAPI.md#RetrieveWorkloadDeployment) | **Get** /workspace/workloads/{workload_id}/deployments/{deployment_id} | Retrieve details of a Workload Deployment
[**UpdateWorkloadDeployment**](WorkloadDeploymentsAPI.md#UpdateWorkloadDeployment) | **Put** /workspace/workloads/{workload_id}/deployments/{deployment_id} | Update a Workload Deployment



## CreateWorkloadDeployment

> ResponseWorkloadDeployment CreateWorkloadDeployment(ctx, workloadId).WorkloadDeploymentRequest(workloadDeploymentRequest).Execute()

Create a Workload Deployment



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
	workloadDeploymentRequest := *openapiclient.NewWorkloadDeploymentRequest("Name_example", *openapiclient.NewDeploymentStrategyDefaultDeploymentStrategyRequest("Type_example", *openapiclient.NewDefaultDeploymentStrategyAttrsRequest(int64(123)))) // WorkloadDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.CreateWorkloadDeployment(context.Background(), workloadId).WorkloadDeploymentRequest(workloadDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDeploymentsAPI.CreateWorkloadDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkloadDeployment`: ResponseWorkloadDeployment
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.CreateWorkloadDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workloadDeploymentRequest** | [**WorkloadDeploymentRequest**](WorkloadDeploymentRequest.md) |  | 

### Return type

[**ResponseWorkloadDeployment**](ResponseWorkloadDeployment.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkloadDeployment

> ResponseDeleteWorkloadDeployment DeleteWorkloadDeployment(ctx, deploymentId, workloadId).Execute()

Delete a Workload Deployment



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
	deploymentId := int64(789) // int64 | 
	workloadId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.DeleteWorkloadDeployment(context.Background(), deploymentId, workloadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadDeploymentsAPI.DeleteWorkloadDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkloadDeployment`: ResponseDeleteWorkloadDeployment
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.DeleteWorkloadDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** |  | 
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkloadDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteWorkloadDeployment**](ResponseDeleteWorkloadDeployment.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkloadDeployments

> PaginatedWorkloadDeploymentList ListWorkloadDeployments(ctx, workloadId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Workload Deployments



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, active, last_editor, last_modified, current) (optional)
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
	// response from `ListWorkloadDeployments`: PaginatedWorkloadDeploymentList
	fmt.Fprintf(os.Stdout, "Response from `WorkloadDeploymentsAPI.ListWorkloadDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkloadDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, active, last_editor, last_modified, current) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedWorkloadDeploymentList**](PaginatedWorkloadDeploymentList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWorkloadDeployment

> ResponseWorkloadDeployment PartialUpdateWorkloadDeployment(ctx, deploymentId, workloadId).PatchedWorkloadDeploymentRequest(patchedWorkloadDeploymentRequest).Execute()

Partially update a Workload Deployment



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
	deploymentId := int64(789) // int64 | 
	workloadId := int64(789) // int64 | 
	patchedWorkloadDeploymentRequest := *openapiclient.NewPatchedWorkloadDeploymentRequest() // PatchedWorkloadDeploymentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.PartialUpdateWorkloadDeployment(context.Background(), deploymentId, workloadId).PatchedWorkloadDeploymentRequest(patchedWorkloadDeploymentRequest).Execute()
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
**deploymentId** | **int64** |  | 
**workloadId** | **int64** |  | 

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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWorkloadDeployment

> ResponseRetrieveWorkloadDeployment RetrieveWorkloadDeployment(ctx, deploymentId, workloadId).Fields(fields).Execute()

Retrieve details of a Workload Deployment



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
	deploymentId := int64(789) // int64 | 
	workloadId := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.RetrieveWorkloadDeployment(context.Background(), deploymentId, workloadId).Fields(fields).Execute()
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
**deploymentId** | **int64** |  | 
**workloadId** | **int64** |  | 

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

> ResponseWorkloadDeployment UpdateWorkloadDeployment(ctx, deploymentId, workloadId).WorkloadDeploymentRequest(workloadDeploymentRequest).Execute()

Update a Workload Deployment



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
	deploymentId := int64(789) // int64 | 
	workloadId := int64(789) // int64 | 
	workloadDeploymentRequest := *openapiclient.NewWorkloadDeploymentRequest("Name_example", *openapiclient.NewDeploymentStrategyDefaultDeploymentStrategyRequest("Type_example", *openapiclient.NewDefaultDeploymentStrategyAttrsRequest(int64(123)))) // WorkloadDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadDeploymentsAPI.UpdateWorkloadDeployment(context.Background(), deploymentId, workloadId).WorkloadDeploymentRequest(workloadDeploymentRequest).Execute()
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
**deploymentId** | **int64** |  | 
**workloadId** | **int64** |  | 

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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

