# \EdgeFirewallsFunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeFirewallFunction**](EdgeFirewallsFunctionAPI.md#CreateEdgeFirewallFunction) | **Post** /edge_firewall/firewalls/{edge_firewall_id}/functions | Create an Edge Firewall Function
[**DestroyEdgeFirewallFunction**](EdgeFirewallsFunctionAPI.md#DestroyEdgeFirewallFunction) | **Delete** /edge_firewall/firewalls/{edge_firewall_id}/functions/{id} | Destroy an Edge Firewall Function
[**ListEdgeFirewallFunction**](EdgeFirewallsFunctionAPI.md#ListEdgeFirewallFunction) | **Get** /edge_firewall/firewalls/{edge_firewall_id}/functions | List Edge Firewall Function
[**PartialUpdateEdgeFirewallFunction**](EdgeFirewallsFunctionAPI.md#PartialUpdateEdgeFirewallFunction) | **Patch** /edge_firewall/firewalls/{edge_firewall_id}/functions/{id} | Partially update an Edge Firewall Function
[**RetrieveEdgeFirewallFunction**](EdgeFirewallsFunctionAPI.md#RetrieveEdgeFirewallFunction) | **Get** /edge_firewall/firewalls/{edge_firewall_id}/functions/{id} | Retrieve details of an Edge Firewall Function
[**UpdateEdgeFirewallFunction**](EdgeFirewallsFunctionAPI.md#UpdateEdgeFirewallFunction) | **Put** /edge_firewall/firewalls/{edge_firewall_id}/functions/{id} | Update an Edge Firewall Function



## CreateEdgeFirewallFunction

> ResponseEdgeFirewallFunctionInstance CreateEdgeFirewallFunction(ctx, edgeFirewallId).EdgeFirewallFunctionInstanceRequest(edgeFirewallFunctionInstanceRequest).Execute()

Create an Edge Firewall Function



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	edgeFirewallFunctionInstanceRequest := *openapiclient.NewEdgeFirewallFunctionInstanceRequest("Name_example", int64(123)) // EdgeFirewallFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsFunctionAPI.CreateEdgeFirewallFunction(context.Background(), edgeFirewallId).EdgeFirewallFunctionInstanceRequest(edgeFirewallFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsFunctionAPI.CreateEdgeFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeFirewallFunction`: ResponseEdgeFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsFunctionAPI.CreateEdgeFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFirewallFunctionInstanceRequest** | [**EdgeFirewallFunctionInstanceRequest**](EdgeFirewallFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeFirewallFunctionInstance**](ResponseEdgeFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeFirewallFunction

> ResponseDeleteEdgeFirewallFunctionInstance DestroyEdgeFirewallFunction(ctx, edgeFirewallId, id).Execute()

Destroy an Edge Firewall Function



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsFunctionAPI.DestroyEdgeFirewallFunction(context.Background(), edgeFirewallId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsFunctionAPI.DestroyEdgeFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeFirewallFunction`: ResponseDeleteEdgeFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsFunctionAPI.DestroyEdgeFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteEdgeFirewallFunctionInstance**](ResponseDeleteEdgeFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeFirewallFunction

> PaginatedEdgeFirewallFunctionInstanceList ListEdgeFirewallFunction(ctx, edgeFirewallId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Firewall Function



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, edge_function, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsFunctionAPI.ListEdgeFirewallFunction(context.Background(), edgeFirewallId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsFunctionAPI.ListEdgeFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeFirewallFunction`: PaginatedEdgeFirewallFunctionInstanceList
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsFunctionAPI.ListEdgeFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, edge_function, active) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeFirewallFunctionInstanceList**](PaginatedEdgeFirewallFunctionInstanceList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeFirewallFunction

> ResponseEdgeFirewallFunctionInstance PartialUpdateEdgeFirewallFunction(ctx, edgeFirewallId, id).PatchedEdgeFirewallFunctionInstanceRequest(patchedEdgeFirewallFunctionInstanceRequest).Execute()

Partially update an Edge Firewall Function



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 
	patchedEdgeFirewallFunctionInstanceRequest := *openapiclient.NewPatchedEdgeFirewallFunctionInstanceRequest() // PatchedEdgeFirewallFunctionInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsFunctionAPI.PartialUpdateEdgeFirewallFunction(context.Background(), edgeFirewallId, id).PatchedEdgeFirewallFunctionInstanceRequest(patchedEdgeFirewallFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsFunctionAPI.PartialUpdateEdgeFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeFirewallFunction`: ResponseEdgeFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsFunctionAPI.PartialUpdateEdgeFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedEdgeFirewallFunctionInstanceRequest** | [**PatchedEdgeFirewallFunctionInstanceRequest**](PatchedEdgeFirewallFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeFirewallFunctionInstance**](ResponseEdgeFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeFirewallFunction

> ResponseRetrieveEdgeFirewallFunctionInstance RetrieveEdgeFirewallFunction(ctx, edgeFirewallId, id).Fields(fields).Execute()

Retrieve details of an Edge Firewall Function



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsFunctionAPI.RetrieveEdgeFirewallFunction(context.Background(), edgeFirewallId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsFunctionAPI.RetrieveEdgeFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeFirewallFunction`: ResponseRetrieveEdgeFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsFunctionAPI.RetrieveEdgeFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeFirewallFunctionInstance**](ResponseRetrieveEdgeFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeFirewallFunction

> ResponseEdgeFirewallFunctionInstance UpdateEdgeFirewallFunction(ctx, edgeFirewallId, id).EdgeFirewallFunctionInstanceRequest(edgeFirewallFunctionInstanceRequest).Execute()

Update an Edge Firewall Function



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 
	edgeFirewallFunctionInstanceRequest := *openapiclient.NewEdgeFirewallFunctionInstanceRequest("Name_example", int64(123)) // EdgeFirewallFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsFunctionAPI.UpdateEdgeFirewallFunction(context.Background(), edgeFirewallId, id).EdgeFirewallFunctionInstanceRequest(edgeFirewallFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsFunctionAPI.UpdateEdgeFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeFirewallFunction`: ResponseEdgeFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsFunctionAPI.UpdateEdgeFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeFirewallFunctionInstanceRequest** | [**EdgeFirewallFunctionInstanceRequest**](EdgeFirewallFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseEdgeFirewallFunctionInstance**](ResponseEdgeFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

