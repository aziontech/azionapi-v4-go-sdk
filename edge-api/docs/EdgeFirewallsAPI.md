# \EdgeFirewallsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneEdgeFirewall**](EdgeFirewallsAPI.md#CloneEdgeFirewall) | **Post** /edge_firewall/firewalls/{id}/clone | Clone an Edge Firewall
[**CreateEdgeFirewall**](EdgeFirewallsAPI.md#CreateEdgeFirewall) | **Post** /edge_firewall/firewalls | Create an Edge Firewall
[**DestroyEdgeFirewall**](EdgeFirewallsAPI.md#DestroyEdgeFirewall) | **Delete** /edge_firewall/firewalls/{id} | Destroy an Edge Firewall
[**ListEdgeFirewalls**](EdgeFirewallsAPI.md#ListEdgeFirewalls) | **Get** /edge_firewall/firewalls | List Edge Firewalls
[**PartialUpdateEdgeFirewall**](EdgeFirewallsAPI.md#PartialUpdateEdgeFirewall) | **Patch** /edge_firewall/firewalls/{id} | Partially update an Edge Firewall
[**RetrieveEdgeFirewall**](EdgeFirewallsAPI.md#RetrieveEdgeFirewall) | **Get** /edge_firewall/firewalls/{id} | Retrieve details from an Edge Firewall
[**UpdateEdgeFirewall**](EdgeFirewallsAPI.md#UpdateEdgeFirewall) | **Put** /edge_firewall/firewalls/{id} | Update an Edge Firewall



## CloneEdgeFirewall

> ResponseEdgeFirewall CloneEdgeFirewall(ctx, id).CloneEdgeFirewallRequest(cloneEdgeFirewallRequest).Execute()

Clone an Edge Firewall



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
	cloneEdgeFirewallRequest := *openapiclient.NewCloneEdgeFirewallRequest("Name_example") // CloneEdgeFirewallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsAPI.CloneEdgeFirewall(context.Background(), id).CloneEdgeFirewallRequest(cloneEdgeFirewallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.CloneEdgeFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneEdgeFirewall`: ResponseEdgeFirewall
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.CloneEdgeFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneEdgeFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneEdgeFirewallRequest** | [**CloneEdgeFirewallRequest**](CloneEdgeFirewallRequest.md) |  | 

### Return type

[**ResponseEdgeFirewall**](ResponseEdgeFirewall.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEdgeFirewall

> ResponseEdgeFirewall CreateEdgeFirewall(ctx).EdgeFirewallRequest(edgeFirewallRequest).Execute()

Create an Edge Firewall



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
	edgeFirewallRequest := *openapiclient.NewEdgeFirewallRequest("Name_example") // EdgeFirewallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsAPI.CreateEdgeFirewall(context.Background()).EdgeFirewallRequest(edgeFirewallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.CreateEdgeFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeFirewall`: ResponseEdgeFirewall
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.CreateEdgeFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeFirewallRequest** | [**EdgeFirewallRequest**](EdgeFirewallRequest.md) |  | 

### Return type

[**ResponseEdgeFirewall**](ResponseEdgeFirewall.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeFirewall

> ResponseDeleteEdgeFirewall DestroyEdgeFirewall(ctx, id).Execute()

Destroy an Edge Firewall



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
	resp, r, err := apiClient.EdgeFirewallsAPI.DestroyEdgeFirewall(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.DestroyEdgeFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeFirewall`: ResponseDeleteEdgeFirewall
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.DestroyEdgeFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteEdgeFirewall**](ResponseDeleteEdgeFirewall.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeFirewalls

> PaginatedEdgeFirewallList ListEdgeFirewalls(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Firewalls



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id, debug, active, last_editor, last_modified, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsAPI.ListEdgeFirewalls(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.ListEdgeFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeFirewalls`: PaginatedEdgeFirewallList
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.ListEdgeFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id, debug, active, last_editor, last_modified, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeFirewallList**](PaginatedEdgeFirewallList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeFirewall

> ResponseEdgeFirewall PartialUpdateEdgeFirewall(ctx, id).PatchedEdgeFirewallRequest(patchedEdgeFirewallRequest).Execute()

Partially update an Edge Firewall



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
	patchedEdgeFirewallRequest := *openapiclient.NewPatchedEdgeFirewallRequest() // PatchedEdgeFirewallRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsAPI.PartialUpdateEdgeFirewall(context.Background(), id).PatchedEdgeFirewallRequest(patchedEdgeFirewallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.PartialUpdateEdgeFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeFirewall`: ResponseEdgeFirewall
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.PartialUpdateEdgeFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEdgeFirewallRequest** | [**PatchedEdgeFirewallRequest**](PatchedEdgeFirewallRequest.md) |  | 

### Return type

[**ResponseEdgeFirewall**](ResponseEdgeFirewall.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeFirewall

> ResponseRetrieveEdgeFirewall RetrieveEdgeFirewall(ctx, id).Fields(fields).Execute()

Retrieve details from an Edge Firewall



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
	resp, r, err := apiClient.EdgeFirewallsAPI.RetrieveEdgeFirewall(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.RetrieveEdgeFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeFirewall`: ResponseRetrieveEdgeFirewall
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.RetrieveEdgeFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeFirewall**](ResponseRetrieveEdgeFirewall.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeFirewall

> ResponseEdgeFirewall UpdateEdgeFirewall(ctx, id).EdgeFirewallRequest(edgeFirewallRequest).Execute()

Update an Edge Firewall



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
	edgeFirewallRequest := *openapiclient.NewEdgeFirewallRequest("Name_example") // EdgeFirewallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsAPI.UpdateEdgeFirewall(context.Background(), id).EdgeFirewallRequest(edgeFirewallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsAPI.UpdateEdgeFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeFirewall`: ResponseEdgeFirewall
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsAPI.UpdateEdgeFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFirewallRequest** | [**EdgeFirewallRequest**](EdgeFirewallRequest.md) |  | 

### Return type

[**ResponseEdgeFirewall**](ResponseEdgeFirewall.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

