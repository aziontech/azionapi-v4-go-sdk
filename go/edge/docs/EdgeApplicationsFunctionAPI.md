# \EdgeApplicationsFunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeFirewallFunctionInstance**](EdgeApplicationsFunctionAPI.md#CreateEdgeFirewallFunctionInstance) | **Post** /edge_application/applications/{edgeApplicationId}/functions | Create an Edge Application Function Instance
[**DestroyEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#DestroyEdgeApplicationFunctionInstance) | **Delete** /edge_application/applications/{edgeApplicationId}/functions/{id} | Destroy an Edge Application Function Instance
[**ListEdgeApplicationFunctionInstances**](EdgeApplicationsFunctionAPI.md#ListEdgeApplicationFunctionInstances) | **Get** /edge_application/applications/{edgeApplicationId}/functions | List Function Instances
[**PartialUpdateEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#PartialUpdateEdgeApplicationFunctionInstance) | **Patch** /edge_application/applications/{edgeApplicationId}/functions/{id} | Partially update an Edge Application Function Instance
[**RetrieveFunctionInstance**](EdgeApplicationsFunctionAPI.md#RetrieveFunctionInstance) | **Get** /edge_application/applications/{edgeApplicationId}/functions/{id} | Retrieve details of an Edge Application Function Instance
[**UpdateEdgeApplicationFunctionInstance**](EdgeApplicationsFunctionAPI.md#UpdateEdgeApplicationFunctionInstance) | **Put** /edge_application/applications/{edgeApplicationId}/functions/{id} | Update an Edge Application Function Instance



## CreateEdgeFirewallFunctionInstance

> ResponseEdgeApplicationFunctionInstance CreateEdgeFirewallFunctionInstance(ctx, edgeApplicationId).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	edgeApplicationFunctionInstanceRequest := *openapiclient.NewEdgeApplicationFunctionInstanceRequest("Name_example", interface{}(123), int64(123)) // EdgeApplicationFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.CreateEdgeFirewallFunctionInstance(context.Background(), edgeApplicationId).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.CreateEdgeFirewallFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeFirewallFunctionInstance`: ResponseEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.CreateEdgeFirewallFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFirewallFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeApplicationFunctionInstanceRequest** | [**EdgeApplicationFunctionInstanceRequest**](EdgeApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationFunctionInstance**](ResponseEdgeApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeApplicationFunctionInstance

> ResponseDeleteEdgeApplicationFunctionInstance DestroyEdgeApplicationFunctionInstance(ctx, edgeApplicationId, id).Execute()

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.DestroyEdgeApplicationFunctionInstance(context.Background(), edgeApplicationId, id).Execute()
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
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

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

> PaginatedResponseListEdgeApplicationFunctionInstanceList ListEdgeApplicationFunctionInstances(ctx, edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, json_args, edge_function, active, last_editor, last_modified) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances(context.Background(), edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeApplicationFunctionInstances`: PaginatedResponseListEdgeApplicationFunctionInstanceList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeApplicationFunctionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, json_args, edge_function, active, last_editor, last_modified) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListEdgeApplicationFunctionInstanceList**](PaginatedResponseListEdgeApplicationFunctionInstanceList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeApplicationFunctionInstance

> ResponseEdgeApplicationFunctionInstance PartialUpdateEdgeApplicationFunctionInstance(ctx, edgeApplicationId, id).PatchedEdgeApplicationFunctionInstanceRequest(patchedEdgeApplicationFunctionInstanceRequest).Execute()

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	patchedEdgeApplicationFunctionInstanceRequest := *openapiclient.NewPatchedEdgeApplicationFunctionInstanceRequest() // PatchedEdgeApplicationFunctionInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.PartialUpdateEdgeApplicationFunctionInstance(context.Background(), edgeApplicationId, id).PatchedEdgeApplicationFunctionInstanceRequest(patchedEdgeApplicationFunctionInstanceRequest).Execute()
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
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

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

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFunctionInstance

> ResponseRetrieveEdgeApplicationFunctionInstance RetrieveFunctionInstance(ctx, edgeApplicationId, id).Fields(fields).Execute()

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.RetrieveFunctionInstance(context.Background(), edgeApplicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsFunctionAPI.RetrieveFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFunctionInstance`: ResponseRetrieveEdgeApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsFunctionAPI.RetrieveFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFunctionInstanceRequest struct via the builder pattern


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

> ResponseEdgeApplicationFunctionInstance UpdateEdgeApplicationFunctionInstance(ctx, edgeApplicationId, id).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	edgeApplicationFunctionInstanceRequest := *openapiclient.NewEdgeApplicationFunctionInstanceRequest("Name_example", interface{}(123), int64(123)) // EdgeApplicationFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsFunctionAPI.UpdateEdgeApplicationFunctionInstance(context.Background(), edgeApplicationId, id).EdgeApplicationFunctionInstanceRequest(edgeApplicationFunctionInstanceRequest).Execute()
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
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

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

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

