# \WAFsExceptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWafException**](WAFsExceptionsAPI.md#CreateWafException) | **Post** /workspace/wafs/{waf_id}/exceptions | Create an Exception for a Web Application Firewall (WAF)
[**DeleteWafException**](WAFsExceptionsAPI.md#DeleteWafException) | **Delete** /workspace/wafs/{waf_id}/exceptions/{exception_id} | Delete an Exception from a Web Application Firewall (WAF)
[**ListWafExceptions**](WAFsExceptionsAPI.md#ListWafExceptions) | **Get** /workspace/wafs/{waf_id}/exceptions | List Exceptions for a Web Application Firewall (WAF)
[**PartialUpdateWafException**](WAFsExceptionsAPI.md#PartialUpdateWafException) | **Patch** /workspace/wafs/{waf_id}/exceptions/{exception_id} | Partially update an Exception for a Web Application Firewall (WAF)
[**RetrieveWafException**](WAFsExceptionsAPI.md#RetrieveWafException) | **Get** /workspace/wafs/{waf_id}/exceptions/{exception_id} | Retrieve details of an Exception from a Web Application Firewall (WAF)
[**UpdateWafException**](WAFsExceptionsAPI.md#UpdateWafException) | **Put** /workspace/wafs/{waf_id}/exceptions/{exception_id} | Update an Exception for a Web Application Firewall (WAF)



## CreateWafException

> WAFRuleResponse CreateWafException(ctx, wafId).WAFRuleRequest(wAFRuleRequest).Execute()

Create an Exception for a Web Application Firewall (WAF)



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
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	wAFRuleRequest := *openapiclient.NewWAFRuleRequest("Name_example", []openapiclient.WAFExceptionPolymorphicConditionRequest{openapiclient.WAFExceptionPolymorphicConditionRequest{WAFExceptionGenericConditionRequest: openapiclient.NewWAFExceptionGenericConditionRequest("Match_example")}}) // WAFRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.CreateWafException(context.Background(), wafId).WAFRuleRequest(wAFRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.CreateWafException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWafException`: WAFRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.CreateWafException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWafExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wAFRuleRequest** | [**WAFRuleRequest**](WAFRuleRequest.md) |  | 

### Return type

[**WAFRuleResponse**](WAFRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWafException

> DeleteResponse DeleteWafException(ctx, exceptionId, wafId).Execute()

Delete an Exception from a Web Application Firewall (WAF)



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
	exceptionId := int64(789) // int64 | A unique integer value identifying the WAF exception.
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.DeleteWafException(context.Background(), exceptionId, wafId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.DeleteWafException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWafException`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.DeleteWafException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **int64** | A unique integer value identifying the WAF exception. | 
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWafExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWafExceptions

> PaginatedWAFRuleList ListWafExceptions(ctx, wafId).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Description(description).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Ordering(ordering).Page(page).PageSize(pageSize).Path(path).Search(search).Execute()

List Exceptions for a Web Application Firewall (WAF)



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
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	createdAtGte := time.Now() // time.Time | Filter by creation date (greater than or equal). (optional)
	createdAtLte := time.Now() // time.Time | Filter by creation date (less than or equal). (optional)
	description := "description_example" // string | Filter by description (case-insensitive, partial match). (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := int64(789) // int64 | Filter by id (accepts comma-separated values). (optional)
	lastEditor := "lastEditor_example" // string | Filter by last editor (case-insensitive, partial match). (optional)
	lastModifiedGte := time.Now() // time.Time | Filter by last modified date (greater than or equal). (optional)
	lastModifiedLte := time.Now() // time.Time | Filter by last modified date (less than or equal). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, description, path, last_editor, last_modified, created_at) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	path := "path_example" // string | Filter by path (case-insensitive, partial match). (optional)
	search := "search_example" // string | A search term to filter results. Searches across the following fields: description, path, last_editor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.ListWafExceptions(context.Background(), wafId).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Description(description).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Ordering(ordering).Page(page).PageSize(pageSize).Path(path).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.ListWafExceptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWafExceptions`: PaginatedWAFRuleList
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.ListWafExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWafExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdAtGte** | **time.Time** | Filter by creation date (greater than or equal). | 
 **createdAtLte** | **time.Time** | Filter by creation date (less than or equal). | 
 **description** | **string** | Filter by description (case-insensitive, partial match). | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by id (accepts comma-separated values). | 
 **lastEditor** | **string** | Filter by last editor (case-insensitive, partial match). | 
 **lastModifiedGte** | **time.Time** | Filter by last modified date (greater than or equal). | 
 **lastModifiedLte** | **time.Time** | Filter by last modified date (less than or equal). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, description, path, last_editor, last_modified, created_at) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **path** | **string** | Filter by path (case-insensitive, partial match). | 
 **search** | **string** | A search term to filter results. Searches across the following fields: description, path, last_editor. | 

### Return type

[**PaginatedWAFRuleList**](PaginatedWAFRuleList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWafException

> WAFRuleResponse PartialUpdateWafException(ctx, exceptionId, wafId).PatchedWAFRuleRequest(patchedWAFRuleRequest).Execute()

Partially update an Exception for a Web Application Firewall (WAF)



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
	exceptionId := int64(789) // int64 | A unique integer value identifying the WAF exception.
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	patchedWAFRuleRequest := *openapiclient.NewPatchedWAFRuleRequest() // PatchedWAFRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.PartialUpdateWafException(context.Background(), exceptionId, wafId).PatchedWAFRuleRequest(patchedWAFRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.PartialUpdateWafException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateWafException`: WAFRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.PartialUpdateWafException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **int64** | A unique integer value identifying the WAF exception. | 
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWafExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedWAFRuleRequest** | [**PatchedWAFRuleRequest**](PatchedWAFRuleRequest.md) |  | 

### Return type

[**WAFRuleResponse**](WAFRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWafException

> WAFRuleResponse RetrieveWafException(ctx, exceptionId, wafId).Fields(fields).Execute()

Retrieve details of an Exception from a Web Application Firewall (WAF)



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
	exceptionId := int64(789) // int64 | A unique integer value identifying the WAF exception.
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.RetrieveWafException(context.Background(), exceptionId, wafId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.RetrieveWafException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWafException`: WAFRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.RetrieveWafException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **int64** | A unique integer value identifying the WAF exception. | 
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWafExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**WAFRuleResponse**](WAFRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWafException

> WAFRuleResponse UpdateWafException(ctx, exceptionId, wafId).WAFRuleRequest(wAFRuleRequest).Execute()

Update an Exception for a Web Application Firewall (WAF)



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
	exceptionId := int64(789) // int64 | A unique integer value identifying the WAF exception.
	wafId := int64(789) // int64 | A unique integer value identifying the WAF.
	wAFRuleRequest := *openapiclient.NewWAFRuleRequest("Name_example", []openapiclient.WAFExceptionPolymorphicConditionRequest{openapiclient.WAFExceptionPolymorphicConditionRequest{WAFExceptionGenericConditionRequest: openapiclient.NewWAFExceptionGenericConditionRequest("Match_example")}}) // WAFRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.UpdateWafException(context.Background(), exceptionId, wafId).WAFRuleRequest(wAFRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.UpdateWafException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWafException`: WAFRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.UpdateWafException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **int64** | A unique integer value identifying the WAF exception. | 
**wafId** | **int64** | A unique integer value identifying the WAF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWafExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wAFRuleRequest** | [**WAFRuleRequest**](WAFRuleRequest.md) |  | 

### Return type

[**WAFRuleResponse**](WAFRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

