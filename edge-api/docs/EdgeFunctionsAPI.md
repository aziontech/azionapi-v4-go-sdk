# \EdgeFunctionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeFunction**](EdgeFunctionsAPI.md#CreateEdgeFunction) | **Post** /edge_functions/functions | Create an Edge Function
[**DestroyEdgeFunction**](EdgeFunctionsAPI.md#DestroyEdgeFunction) | **Delete** /edge_functions/functions/{id} | Destroy an Edge Function
[**ListEdgeFunctions**](EdgeFunctionsAPI.md#ListEdgeFunctions) | **Get** /edge_functions/functions | List Edge Functions
[**PartialUpdateEdgeFunction**](EdgeFunctionsAPI.md#PartialUpdateEdgeFunction) | **Patch** /edge_functions/functions/{id} | Partially update an Edge Function
[**RetrieveEdgeFunction**](EdgeFunctionsAPI.md#RetrieveEdgeFunction) | **Get** /edge_functions/functions/{id} | Retrieve details of an Edge Function
[**UpdateEdgeFunction**](EdgeFunctionsAPI.md#UpdateEdgeFunction) | **Put** /edge_functions/functions/{id} | Update an Edge Function



## CreateEdgeFunction

> ResponseEdgeFunctionsDoc CreateEdgeFunction(ctx).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()

Create an Edge Function



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
	edgeFunctionsRequest := *openapiclient.NewEdgeFunctionsRequest("Name_example", "Code_example") // EdgeFunctionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.CreateEdgeFunction(context.Background()).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.CreateEdgeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeFunction`: ResponseEdgeFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.CreateEdgeFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeFunctionsRequest** | [**EdgeFunctionsRequest**](EdgeFunctionsRequest.md) |  | 

### Return type

[**ResponseEdgeFunctionsDoc**](ResponseEdgeFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeFunction

> ResponseDeleteEdgeFunctionsDoc DestroyEdgeFunction(ctx, id).Execute()

Destroy an Edge Function



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.DestroyEdgeFunction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.DestroyEdgeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeFunction`: ResponseDeleteEdgeFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.DestroyEdgeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteEdgeFunctionsDoc**](ResponseDeleteEdgeFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeFunctions

> PaginatedEdgeFunctionsList ListEdgeFunctions(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Functions



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, language, json_args, runtime_environment, active, last_editor, last_modified, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.ListEdgeFunctions(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.ListEdgeFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeFunctions`: PaginatedEdgeFunctionsList
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.ListEdgeFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, language, json_args, runtime_environment, active, last_editor, last_modified, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeFunctionsList**](PaginatedEdgeFunctionsList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeFunction

> ResponseEdgeFunctionsDoc PartialUpdateEdgeFunction(ctx, id).PatchedEdgeFunctionsRequest(patchedEdgeFunctionsRequest).Execute()

Partially update an Edge Function



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
	id := "id_example" // string | 
	patchedEdgeFunctionsRequest := *openapiclient.NewPatchedEdgeFunctionsRequest() // PatchedEdgeFunctionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.PartialUpdateEdgeFunction(context.Background(), id).PatchedEdgeFunctionsRequest(patchedEdgeFunctionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.PartialUpdateEdgeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeFunction`: ResponseEdgeFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.PartialUpdateEdgeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEdgeFunctionsRequest** | [**PatchedEdgeFunctionsRequest**](PatchedEdgeFunctionsRequest.md) |  | 

### Return type

[**ResponseEdgeFunctionsDoc**](ResponseEdgeFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeFunction

> ResponseRetrieveEdgeFunctionsDoc RetrieveEdgeFunction(ctx, id).Fields(fields).Execute()

Retrieve details of an Edge Function



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
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.RetrieveEdgeFunction(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.RetrieveEdgeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeFunction`: ResponseRetrieveEdgeFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.RetrieveEdgeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeFunctionsDoc**](ResponseRetrieveEdgeFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeFunction

> ResponseEdgeFunctionsDoc UpdateEdgeFunction(ctx, id).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()

Update an Edge Function



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
	id := "id_example" // string | 
	edgeFunctionsRequest := *openapiclient.NewEdgeFunctionsRequest("Name_example", "Code_example") // EdgeFunctionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.UpdateEdgeFunction(context.Background(), id).EdgeFunctionsRequest(edgeFunctionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.UpdateEdgeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeFunction`: ResponseEdgeFunctionsDoc
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.UpdateEdgeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFunctionsRequest** | [**EdgeFunctionsRequest**](EdgeFunctionsRequest.md) |  | 

### Return type

[**ResponseEdgeFunctionsDoc**](ResponseEdgeFunctionsDoc.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

