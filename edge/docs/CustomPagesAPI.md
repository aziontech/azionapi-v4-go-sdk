# \CustomPagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomPage**](CustomPagesAPI.md#CreateCustomPage) | **Post** /workspace/custom_pages | Create a Custom Page
[**DestroyCustomPage**](CustomPagesAPI.md#DestroyCustomPage) | **Delete** /workspace/custom_pages/{id} | Destroy a Custom Page
[**ListCustomPages**](CustomPagesAPI.md#ListCustomPages) | **Get** /workspace/custom_pages | List Custom Pages
[**PartialUpdateCustomPage**](CustomPagesAPI.md#PartialUpdateCustomPage) | **Patch** /workspace/custom_pages/{id} | Partially update a Custom Page
[**RetrieveCustomPage**](CustomPagesAPI.md#RetrieveCustomPage) | **Get** /workspace/custom_pages/{id} | Retrieve details of a Custom Page
[**UpdateCustomPage**](CustomPagesAPI.md#UpdateCustomPage) | **Put** /workspace/custom_pages/{id} | Update a Custom Page



## CreateCustomPage

> ResponseCustomPages CreateCustomPage(ctx).CustomPagesRequest(customPagesRequest).Execute()

Create a Custom Page



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
	customPagesRequest := *openapiclient.NewCustomPagesRequest("Name_example", *openapiclient.NewConnectorCustomPagesRequest([]openapiclient.PageRequest{*openapiclient.NewPageRequest("Code_example")})) // CustomPagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPagesAPI.CreateCustomPage(context.Background()).CustomPagesRequest(customPagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.CreateCustomPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomPage`: ResponseCustomPages
	fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.CreateCustomPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customPagesRequest** | [**CustomPagesRequest**](CustomPagesRequest.md) |  | 

### Return type

[**ResponseCustomPages**](ResponseCustomPages.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCustomPage

> ResponseDeleteCustomPages DestroyCustomPage(ctx, id).Execute()

Destroy a Custom Page



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
	resp, r, err := apiClient.CustomPagesAPI.DestroyCustomPage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.DestroyCustomPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyCustomPage`: ResponseDeleteCustomPages
	fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.DestroyCustomPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyCustomPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteCustomPages**](ResponseDeleteCustomPages.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomPages

> PaginatedResponseListCustomPagesList ListCustomPages(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Custom Pages



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, last_editor, last_modified, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPagesAPI.ListCustomPages(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ListCustomPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomPages`: PaginatedResponseListCustomPagesList
	fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ListCustomPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, last_editor, last_modified, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListCustomPagesList**](PaginatedResponseListCustomPagesList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateCustomPage

> ResponseCustomPages PartialUpdateCustomPage(ctx, id).PatchedCustomPagesRequest(patchedCustomPagesRequest).Execute()

Partially update a Custom Page



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
	patchedCustomPagesRequest := *openapiclient.NewPatchedCustomPagesRequest() // PatchedCustomPagesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPagesAPI.PartialUpdateCustomPage(context.Background(), id).PatchedCustomPagesRequest(patchedCustomPagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.PartialUpdateCustomPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateCustomPage`: ResponseCustomPages
	fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.PartialUpdateCustomPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCustomPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCustomPagesRequest** | [**PatchedCustomPagesRequest**](PatchedCustomPagesRequest.md) |  | 

### Return type

[**ResponseCustomPages**](ResponseCustomPages.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCustomPage

> ResponseRetrieveCustomPages RetrieveCustomPage(ctx, id).Fields(fields).Execute()

Retrieve details of a Custom Page



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
	resp, r, err := apiClient.CustomPagesAPI.RetrieveCustomPage(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.RetrieveCustomPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCustomPage`: ResponseRetrieveCustomPages
	fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.RetrieveCustomPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCustomPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveCustomPages**](ResponseRetrieveCustomPages.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomPage

> ResponseCustomPages UpdateCustomPage(ctx, id).CustomPagesRequest(customPagesRequest).Execute()

Update a Custom Page



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
	customPagesRequest := *openapiclient.NewCustomPagesRequest("Name_example", *openapiclient.NewConnectorCustomPagesRequest([]openapiclient.PageRequest{*openapiclient.NewPageRequest("Code_example")})) // CustomPagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPagesAPI.UpdateCustomPage(context.Background(), id).CustomPagesRequest(customPagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.UpdateCustomPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomPage`: ResponseCustomPages
	fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.UpdateCustomPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customPagesRequest** | [**CustomPagesRequest**](CustomPagesRequest.md) |  | 

### Return type

[**ResponseCustomPages**](ResponseCustomPages.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

