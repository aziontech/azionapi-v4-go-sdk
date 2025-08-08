# \EdgeApplicationsFunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#CreateEdgeApplicationFunctionInstance) | **Post** /edge_application/applications/{application_id}/functions | Create an Edge Application Function Instance
[**DestroyEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#DestroyEdgeApplicationFunctionInstance) | **Delete** /edge_application/applications/{application_id}/functions/{function_id} | Destroy an Edge Application Function Instance
[**ListEdgeApplicationFunctionInstances**](EdgeApplicationsFunctionAPI.md#ListEdgeApplicationFunctionInstances) | **Get** /edge_application/applications/{application_id}/functions | List Function Instances
[**PartialUpdateEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#PartialUpdateEdgeApplicationFunctionInstance) | **Patch** /edge_application/applications/{application_id}/functions/{function_id} | Partially update an Edge Application Function Instance
[**RetrieveEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#RetrieveEdgeApplicationFunctionInstance) | **Get** /edge_application/applications/{application_id}/functions/{function_id} | Retrieve details of an Edge Application Function Instance
[**UpdateEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#UpdateEdgeApplicationFunctionInstance) | **Put** /edge_application/applications/{application_id}/functions/{function_id} | Update an Edge Application Function Instance



## CreateEdgeApplicationFunctionInstance

> ResponseEdgeApplicationFunctionInstance CreateEdgeApplicationFunctionInstance(ctx, applicationId).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()

Create an Edge Application Function Instance



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
	applicationId := "applicationId_example" // string | 
	edgeApplicationFunctionInstanceRequest := *openapiclient.NewEdgeApplicationFunctionInstanceRequest("Name_example", int64(123)) // EdgeApplicationFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.CreateEdgeApplicationFunctionInstance(context.Background(), applicationId).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.CreateEdgeApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeApplicationFunctionInstance`: ResponseEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.CreateEdgeApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeApplicationFunctionInstanceRequest** | [**EdgeApplicationFunctionInstanceRequest**](EdgeApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationFunctionInstance**](ResponseEdgeApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeApplicationFunctionInstance

> ResponseDeleteEdgeApplicationFunctionInstance DestroyEdgeApplicationFunctionInstance(ctx, applicationId, functionId).Execute()

Destroy an Edge Application Function Instance



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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.DestroyEdgeApplicationFunctionInstance(context.Background(), applicationId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.DestroyEdgeApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeApplicationFunctionInstance`: ResponseDeleteEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.DestroyEdgeApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteEdgeApplicationFunctionInstance**](ResponseDeleteEdgeApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeApplicationFunctionInstances

> PaginatedEdgeApplicationFunctionInstanceList ListEdgeApplicationFunctionInstances(ctx, applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Function Instances



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
	applicationId := "applicationId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, edge_function, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances(context.Background(), applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeApplicationFunctionInstances`: PaginatedEdgeApplicationFunctionInstanceList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeApplicationFunctionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, edge_function, active) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeApplicationFunctionInstanceList**](PaginatedEdgeApplicationFunctionInstanceList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeApplicationFunctionInstance

> ResponseEdgeApplicationFunctionInstance PartialUpdateEdgeApplicationFunctionInstance(ctx, applicationId, functionId).PatchedEdgeApplicationFunctionInstanceRequest(patchedEdgeApplicationFunctionInstanceRequest).Execute()

Partially update an Edge Application Function Instance



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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 
	patchedEdgeApplicationFunctionInstanceRequest := *openapiclient.NewPatchedEdgeApplicationFunctionInstanceRequest() // PatchedEdgeApplicationFunctionInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.PartialUpdateEdgeApplicationFunctionInstance(context.Background(), applicationId, functionId).PatchedEdgeApplicationFunctionInstanceRequest(patchedEdgeApplicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.PartialUpdateEdgeApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeApplicationFunctionInstance`: ResponseEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.PartialUpdateEdgeApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedEdgeApplicationFunctionInstanceRequest** | [**PatchedEdgeApplicationFunctionInstanceRequest**](PatchedEdgeApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationFunctionInstance**](ResponseEdgeApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeApplicationFunctionInstance

> ResponseRetrieveEdgeApplicationFunctionInstance RetrieveEdgeApplicationFunctionInstance(ctx, applicationId, functionId).Fields(fields).Execute()

Retrieve details of an Edge Application Function Instance



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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.RetrieveEdgeApplicationFunctionInstance(context.Background(), applicationId, functionId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.RetrieveEdgeApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeApplicationFunctionInstance`: ResponseRetrieveEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.RetrieveEdgeApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeApplicationFunctionInstance**](ResponseRetrieveEdgeApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeApplicationFunctionInstance

> ResponseEdgeApplicationFunctionInstance UpdateEdgeApplicationFunctionInstance(ctx, applicationId, functionId).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()

Update an Edge Application Function Instance



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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 
	edgeApplicationFunctionInstanceRequest := *openapiclient.NewEdgeApplicationFunctionInstanceRequest("Name_example", int64(123)) // EdgeApplicationFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.UpdateEdgeApplicationFunctionInstance(context.Background(), applicationId, functionId).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.UpdateEdgeApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeApplicationFunctionInstance`: ResponseEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.UpdateEdgeApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeApplicationFunctionInstanceRequest** | [**EdgeApplicationFunctionInstanceRequest**](EdgeApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationFunctionInstance**](ResponseEdgeApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

