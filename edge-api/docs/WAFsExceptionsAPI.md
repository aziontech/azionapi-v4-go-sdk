# \WAFsExceptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWAFException**](WAFsExceptionsAPI.md#CreateWAFException) | **Post** /edge_firewall/wafs/{waf_id}/exceptions | Create an Exception for a Web Application Firewall (WAF)
[**DestroyWAFException**](WAFsExceptionsAPI.md#DestroyWAFException) | **Delete** /edge_firewall/wafs/{waf_id}/exceptions/{exception_id} | Destroy an Exception from a Web Application Firewall (WAF)
[**ListWAFExceptions**](WAFsExceptionsAPI.md#ListWAFExceptions) | **Get** /edge_firewall/wafs/{waf_id}/exceptions | List Exceptions for a Web Application Firewall (WAF)
[**PartialUpdateWAFException**](WAFsExceptionsAPI.md#PartialUpdateWAFException) | **Patch** /edge_firewall/wafs/{waf_id}/exceptions/{exception_id} | Partially update an Exception for a Web Application Firewall (WAF)
[**RetrieveWAFException**](WAFsExceptionsAPI.md#RetrieveWAFException) | **Get** /edge_firewall/wafs/{waf_id}/exceptions/{exception_id} | Retrieve details of an Exception from a Web Application Firewall (WAF)
[**UpdateWAFException**](WAFsExceptionsAPI.md#UpdateWAFException) | **Put** /edge_firewall/wafs/{waf_id}/exceptions/{exception_id} | Update an Exception for a Web Application Firewall (WAF)



## CreateWAFException

> ResponseWAFRule CreateWAFException(ctx, wafId).WAFRuleRequest(wAFRuleRequest).Execute()

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
	wafId := "wafId_example" // string | 
	wAFRuleRequest := *openapiclient.NewWAFRuleRequest("Name_example", []openapiclient.WAFExceptionPolymorphicConditionRequest{openapiclient.WAFExceptionPolymorphicConditionRequest{WAFExceptionGenericConditionRequest: openapiclient.NewWAFExceptionGenericConditionRequest("Match_example")}}) // WAFRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.CreateWAFException(context.Background(), wafId).WAFRuleRequest(wAFRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.CreateWAFException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWAFException`: ResponseWAFRule
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.CreateWAFException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWAFExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wAFRuleRequest** | [**WAFRuleRequest**](WAFRuleRequest.md) |  | 

### Return type

[**ResponseWAFRule**](ResponseWAFRule.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWAFException

> ResponseDeleteWAFRule DestroyWAFException(ctx, exceptionId, wafId).Execute()

Destroy an Exception from a Web Application Firewall (WAF)



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
	exceptionId := "exceptionId_example" // string | 
	wafId := "wafId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.DestroyWAFException(context.Background(), exceptionId, wafId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.DestroyWAFException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyWAFException`: ResponseDeleteWAFRule
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.DestroyWAFException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **string** |  | 
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWAFExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteWAFRule**](ResponseDeleteWAFRule.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWAFExceptions

> PaginatedWAFRuleList ListWAFExceptions(ctx, wafId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Exceptions for a Web Application Firewall (WAF)



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
	wafId := "wafId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: rule_id, name, path, conditions, operator, active, last_editor, last_modified) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.ListWAFExceptions(context.Background(), wafId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.ListWAFExceptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWAFExceptions`: PaginatedWAFRuleList
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.ListWAFExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWAFExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: rule_id, name, path, conditions, operator, active, last_editor, last_modified) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

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


## PartialUpdateWAFException

> ResponseWAFRule PartialUpdateWAFException(ctx, exceptionId, wafId).PatchedWAFRuleRequest(patchedWAFRuleRequest).Execute()

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
	exceptionId := "exceptionId_example" // string | 
	wafId := "wafId_example" // string | 
	patchedWAFRuleRequest := *openapiclient.NewPatchedWAFRuleRequest() // PatchedWAFRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.PartialUpdateWAFException(context.Background(), exceptionId, wafId).PatchedWAFRuleRequest(patchedWAFRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.PartialUpdateWAFException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateWAFException`: ResponseWAFRule
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.PartialUpdateWAFException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **string** |  | 
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWAFExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedWAFRuleRequest** | [**PatchedWAFRuleRequest**](PatchedWAFRuleRequest.md) |  | 

### Return type

[**ResponseWAFRule**](ResponseWAFRule.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWAFException

> ResponseRetrieveWAFRule RetrieveWAFException(ctx, exceptionId, wafId).Fields(fields).Execute()

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
	exceptionId := "exceptionId_example" // string | 
	wafId := "wafId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.RetrieveWAFException(context.Background(), exceptionId, wafId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.RetrieveWAFException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWAFException`: ResponseRetrieveWAFRule
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.RetrieveWAFException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **string** |  | 
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWAFExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveWAFRule**](ResponseRetrieveWAFRule.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWAFException

> ResponseWAFRule UpdateWAFException(ctx, exceptionId, wafId).WAFRuleRequest(wAFRuleRequest).Execute()

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
	exceptionId := "exceptionId_example" // string | 
	wafId := "wafId_example" // string | 
	wAFRuleRequest := *openapiclient.NewWAFRuleRequest("Name_example", []openapiclient.WAFExceptionPolymorphicConditionRequest{openapiclient.WAFExceptionPolymorphicConditionRequest{WAFExceptionGenericConditionRequest: openapiclient.NewWAFExceptionGenericConditionRequest("Match_example")}}) // WAFRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsExceptionsAPI.UpdateWAFException(context.Background(), exceptionId, wafId).WAFRuleRequest(wAFRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsExceptionsAPI.UpdateWAFException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWAFException`: ResponseWAFRule
	fmt.Fprintf(os.Stdout, "Response from `WAFsExceptionsAPI.UpdateWAFException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exceptionId** | **string** |  | 
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWAFExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wAFRuleRequest** | [**WAFRuleRequest**](WAFRuleRequest.md) |  | 

### Return type

[**ResponseWAFRule**](ResponseWAFRule.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

