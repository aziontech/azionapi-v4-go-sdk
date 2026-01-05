# \FirewallsFunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirewallFunction**](FirewallsFunctionAPI.md#CreateFirewallFunction) | **Post** /workspace/firewalls/{firewall_id}/functions | Create an Firewall Function
[**DeleteFirewallFunction**](FirewallsFunctionAPI.md#DeleteFirewallFunction) | **Delete** /workspace/firewalls/{firewall_id}/functions/{function_id} | Delete an Firewall Function
[**ListFirewallFunction**](FirewallsFunctionAPI.md#ListFirewallFunction) | **Get** /workspace/firewalls/{firewall_id}/functions | List Firewall Function
[**PartialUpdateFirewallFunction**](FirewallsFunctionAPI.md#PartialUpdateFirewallFunction) | **Patch** /workspace/firewalls/{firewall_id}/functions/{function_id} | Partially update an Firewall Function
[**RetrieveFirewallFunction**](FirewallsFunctionAPI.md#RetrieveFirewallFunction) | **Get** /workspace/firewalls/{firewall_id}/functions/{function_id} | Retrieve details of an Firewall Function
[**UpdateFirewallFunction**](FirewallsFunctionAPI.md#UpdateFirewallFunction) | **Put** /workspace/firewalls/{firewall_id}/functions/{function_id} | Update an Firewall Function



## CreateFirewallFunction

> ResponseFirewallFunctionInstance CreateFirewallFunction(ctx, firewallId).FirewallFunctionInstanceRequest(firewallFunctionInstanceRequest).Execute()

Create an Firewall Function



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
	firewallId := int64(789) // int64 | A unique integer value identifying the firewall.
	firewallFunctionInstanceRequest := *openapiclient.NewFirewallFunctionInstanceRequest("Name_example", int64(123)) // FirewallFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsFunctionAPI.CreateFirewallFunction(context.Background(), firewallId).FirewallFunctionInstanceRequest(firewallFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsFunctionAPI.CreateFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewallFunction`: ResponseFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `FirewallsFunctionAPI.CreateFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int64** | A unique integer value identifying the firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallFunctionInstanceRequest** | [**FirewallFunctionInstanceRequest**](FirewallFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseFirewallFunctionInstance**](ResponseFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallFunction

> ResponseDeleteFirewallFunctionInstance DeleteFirewallFunction(ctx, firewallId, functionId).Execute()

Delete an Firewall Function



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
	firewallId := int64(789) // int64 | A unique integer value identifying the firewall.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsFunctionAPI.DeleteFirewallFunction(context.Background(), firewallId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsFunctionAPI.DeleteFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallFunction`: ResponseDeleteFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `FirewallsFunctionAPI.DeleteFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int64** | A unique integer value identifying the firewall. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteFirewallFunctionInstance**](ResponseDeleteFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirewallFunction

> PaginatedFirewallFunctionInstanceList ListFirewallFunction(ctx, firewallId).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Firewall Function



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	firewallId := int64(789) // int64 | A unique integer value identifying the firewall.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := int64(789) // int64 | Filter by id (accepts comma-separated values). (optional)
	lastEditor := "lastEditor_example" // string | Filter by last editor (case-insensitive, partial match). (optional)
	lastModifiedGte := time.Now() // time.Time | Filter by last modified date (greater than or equal). (optional)
	lastModifiedLte := time.Now() // time.Time | Filter by last modified date (less than or equal). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, azion_form, function, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsFunctionAPI.ListFirewallFunction(context.Background(), firewallId).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsFunctionAPI.ListFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewallFunction`: PaginatedFirewallFunctionInstanceList
	fmt.Fprintf(os.Stdout, "Response from `FirewallsFunctionAPI.ListFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int64** | A unique integer value identifying the firewall. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by id (accepts comma-separated values). | 
 **lastEditor** | **string** | Filter by last editor (case-insensitive, partial match). | 
 **lastModifiedGte** | **time.Time** | Filter by last modified date (greater than or equal). | 
 **lastModifiedLte** | **time.Time** | Filter by last modified date (less than or equal). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, azion_form, function, active) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedFirewallFunctionInstanceList**](PaginatedFirewallFunctionInstanceList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateFirewallFunction

> ResponseFirewallFunctionInstance PartialUpdateFirewallFunction(ctx, firewallId, functionId).PatchedFirewallFunctionInstanceRequest(patchedFirewallFunctionInstanceRequest).Execute()

Partially update an Firewall Function



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
	firewallId := int64(789) // int64 | A unique integer value identifying the firewall.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.
	patchedFirewallFunctionInstanceRequest := *openapiclient.NewPatchedFirewallFunctionInstanceRequest() // PatchedFirewallFunctionInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsFunctionAPI.PartialUpdateFirewallFunction(context.Background(), firewallId, functionId).PatchedFirewallFunctionInstanceRequest(patchedFirewallFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsFunctionAPI.PartialUpdateFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateFirewallFunction`: ResponseFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `FirewallsFunctionAPI.PartialUpdateFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int64** | A unique integer value identifying the firewall. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedFirewallFunctionInstanceRequest** | [**PatchedFirewallFunctionInstanceRequest**](PatchedFirewallFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseFirewallFunctionInstance**](ResponseFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirewallFunction

> ResponseRetrieveFirewallFunctionInstance RetrieveFirewallFunction(ctx, firewallId, functionId).Fields(fields).Execute()

Retrieve details of an Firewall Function



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
	firewallId := int64(789) // int64 | A unique integer value identifying the firewall.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsFunctionAPI.RetrieveFirewallFunction(context.Background(), firewallId, functionId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsFunctionAPI.RetrieveFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFirewallFunction`: ResponseRetrieveFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `FirewallsFunctionAPI.RetrieveFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int64** | A unique integer value identifying the firewall. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveFirewallFunctionInstance**](ResponseRetrieveFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirewallFunction

> ResponseFirewallFunctionInstance UpdateFirewallFunction(ctx, firewallId, functionId).FirewallFunctionInstanceRequest(firewallFunctionInstanceRequest).Execute()

Update an Firewall Function



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
	firewallId := int64(789) // int64 | A unique integer value identifying the firewall.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.
	firewallFunctionInstanceRequest := *openapiclient.NewFirewallFunctionInstanceRequest("Name_example", int64(123)) // FirewallFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsFunctionAPI.UpdateFirewallFunction(context.Background(), firewallId, functionId).FirewallFunctionInstanceRequest(firewallFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsFunctionAPI.UpdateFirewallFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirewallFunction`: ResponseFirewallFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `FirewallsFunctionAPI.UpdateFirewallFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int64** | A unique integer value identifying the firewall. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallFunctionInstanceRequest** | [**FirewallFunctionInstanceRequest**](FirewallFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseFirewallFunctionInstance**](ResponseFirewallFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

