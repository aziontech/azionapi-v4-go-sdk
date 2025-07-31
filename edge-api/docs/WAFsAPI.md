# \WAFsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneWAF**](WAFsAPI.md#CloneWAF) | **Post** /edge_firewall/wafs/{waf_id}/clone | Clone a Web Application Firewall (WAF)
[**CreateWAF**](WAFsAPI.md#CreateWAF) | **Post** /edge_firewall/wafs | Create a Web Application Firewall (WAF)
[**DestroyWAF**](WAFsAPI.md#DestroyWAF) | **Delete** /edge_firewall/wafs/{waf_id} | Destroy a Web Application Firewall (WAF)
[**ListWAFs**](WAFsAPI.md#ListWAFs) | **Get** /edge_firewall/wafs | List Web Application Firewalls (WAFs)
[**PartialUpdateWAF**](WAFsAPI.md#PartialUpdateWAF) | **Patch** /edge_firewall/wafs/{waf_id} | Partially update a Web Application Firewall (WAF)
[**RetrieveWAF**](WAFsAPI.md#RetrieveWAF) | **Get** /edge_firewall/wafs/{waf_id} | Retrieve details from a Web Application Firewall (WAF)
[**UpdateWAF**](WAFsAPI.md#UpdateWAF) | **Put** /edge_firewall/wafs/{waf_id} | Update a Web Application Firewall (WAF)



## CloneWAF

> ResponseWAF CloneWAF(ctx, wafId).CloneWAFRequest(cloneWAFRequest).Execute()

Clone a Web Application Firewall (WAF)



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
	cloneWAFRequest := *openapiclient.NewCloneWAFRequest("Name_example") // CloneWAFRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.CloneWAF(context.Background(), wafId).CloneWAFRequest(cloneWAFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.CloneWAF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneWAF`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.CloneWAF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneWAFRequest** | [**CloneWAFRequest**](CloneWAFRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWAF

> ResponseWAF CreateWAF(ctx).WAFRequest(wAFRequest).Execute()

Create a Web Application Firewall (WAF)



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
	wAFRequest := *openapiclient.NewWAFRequest("Name_example") // WAFRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.CreateWAF(context.Background()).WAFRequest(wAFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.CreateWAF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWAF`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.CreateWAF`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wAFRequest** | [**WAFRequest**](WAFRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyWAF

> ResponseDeleteWAF DestroyWAF(ctx, wafId).Execute()

Destroy a Web Application Firewall (WAF)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.DestroyWAF(context.Background(), wafId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.DestroyWAF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyWAF`: ResponseDeleteWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.DestroyWAF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteWAF**](ResponseDeleteWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWAFs

> PaginatedWAFList ListWAFs(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Web Application Firewalls (WAFs)



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id, active, last_editor, last_modified, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.ListWAFs(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.ListWAFs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWAFs`: PaginatedWAFList
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.ListWAFs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWAFsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id, active, last_editor, last_modified, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedWAFList**](PaginatedWAFList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateWAF

> ResponseWAF PartialUpdateWAF(ctx, wafId).PatchedWAFRequest(patchedWAFRequest).Execute()

Partially update a Web Application Firewall (WAF)



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
	patchedWAFRequest := *openapiclient.NewPatchedWAFRequest() // PatchedWAFRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.PartialUpdateWAF(context.Background(), wafId).PatchedWAFRequest(patchedWAFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.PartialUpdateWAF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateWAF`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.PartialUpdateWAF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWAFRequest** | [**PatchedWAFRequest**](PatchedWAFRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWAF

> ResponseRetrieveWAF RetrieveWAF(ctx, wafId).Fields(fields).Execute()

Retrieve details from a Web Application Firewall (WAF)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.RetrieveWAF(context.Background(), wafId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.RetrieveWAF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWAF`: ResponseRetrieveWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.RetrieveWAF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveWAF**](ResponseRetrieveWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWAF

> ResponseWAF UpdateWAF(ctx, wafId).WAFRequest(wAFRequest).Execute()

Update a Web Application Firewall (WAF)



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
	wAFRequest := *openapiclient.NewWAFRequest("Name_example") // WAFRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WAFsAPI.UpdateWAF(context.Background(), wafId).WAFRequest(wAFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WAFsAPI.UpdateWAF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWAF`: ResponseWAF
	fmt.Fprintf(os.Stdout, "Response from `WAFsAPI.UpdateWAF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wAFRequest** | [**WAFRequest**](WAFRequest.md) |  | 

### Return type

[**ResponseWAF**](ResponseWAF.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

