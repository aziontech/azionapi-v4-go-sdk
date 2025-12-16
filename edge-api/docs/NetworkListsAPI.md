# \NetworkListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkList**](NetworkListsAPI.md#CreateNetworkList) | **Post** /workspace/network_lists | Create a Network List
[**DeleteNetworkList**](NetworkListsAPI.md#DeleteNetworkList) | **Delete** /workspace/network_lists/{network_list_id} | Delete a Network List
[**ListNetworkLists**](NetworkListsAPI.md#ListNetworkLists) | **Get** /workspace/network_lists | List Network Lists
[**PartialUpdateNetworkList**](NetworkListsAPI.md#PartialUpdateNetworkList) | **Patch** /workspace/network_lists/{network_list_id} | Partially update a Network List
[**RetrieveNetworkList**](NetworkListsAPI.md#RetrieveNetworkList) | **Get** /workspace/network_lists/{network_list_id} | Retrieve details of a Network List
[**UpdateNetworkList**](NetworkListsAPI.md#UpdateNetworkList) | **Put** /workspace/network_lists/{network_list_id} | Update a Network List



## CreateNetworkList

> ResponseNetworkListDetail CreateNetworkList(ctx).NetworkListDetailRequest(networkListDetailRequest).Execute()

Create a Network List



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
	networkListDetailRequest := *openapiclient.NewNetworkListDetailRequest("Name_example", "Type_example", []string{"Items_example"}) // NetworkListDetailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.CreateNetworkList(context.Background()).NetworkListDetailRequest(networkListDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.CreateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkList`: ResponseNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.CreateNetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkListDetailRequest** | [**NetworkListDetailRequest**](NetworkListDetailRequest.md) |  | 

### Return type

[**ResponseNetworkListDetail**](ResponseNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkList

> ResponseAsyncDeleteNetworkListDetail DeleteNetworkList(ctx, networkListId).Execute()

Delete a Network List



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
	networkListId := int64(789) // int64 | A unique integer value identifying the network list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.DeleteNetworkList(context.Background(), networkListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.DeleteNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkList`: ResponseAsyncDeleteNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.DeleteNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkListId** | **int64** | A unique integer value identifying the network list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseAsyncDeleteNetworkListDetail**](ResponseAsyncDeleteNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkLists

> PaginatedNetworkListList ListNetworkLists(ctx).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).ListTypeIn(listTypeIn).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Network Lists



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := int64(789) // int64 | Filter by id (accepts comma-separated values). (optional)
	lastEditor := "lastEditor_example" // string | Filter by last editor (case-insensitive, partial match). (optional)
	lastModifiedGte := time.Now() // time.Time | Filter by last modified date (greater than or equal). (optional)
	lastModifiedLte := time.Now() // time.Time | Filter by last modified date (less than or equal). (optional)
	listTypeIn := "listTypeIn_example" // string | Filter by list type (accepts comma-separated values). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, type, last_editor, last_modified, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.ListNetworkLists(context.Background()).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).ListTypeIn(listTypeIn).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.ListNetworkLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkLists`: PaginatedNetworkListList
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.ListNetworkLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by id (accepts comma-separated values). | 
 **lastEditor** | **string** | Filter by last editor (case-insensitive, partial match). | 
 **lastModifiedGte** | **time.Time** | Filter by last modified date (greater than or equal). | 
 **lastModifiedLte** | **time.Time** | Filter by last modified date (less than or equal). | 
 **listTypeIn** | **string** | Filter by list type (accepts comma-separated values). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, type, last_editor, last_modified, active) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedNetworkListList**](PaginatedNetworkListList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateNetworkList

> ResponseNetworkListDetail PartialUpdateNetworkList(ctx, networkListId).PatchedNetworkListDetailRequest(patchedNetworkListDetailRequest).Execute()

Partially update a Network List



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
	networkListId := int64(789) // int64 | A unique integer value identifying the network list.
	patchedNetworkListDetailRequest := *openapiclient.NewPatchedNetworkListDetailRequest() // PatchedNetworkListDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.PartialUpdateNetworkList(context.Background(), networkListId).PatchedNetworkListDetailRequest(patchedNetworkListDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.PartialUpdateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateNetworkList`: ResponseNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.PartialUpdateNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkListId** | **int64** | A unique integer value identifying the network list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNetworkListDetailRequest** | [**PatchedNetworkListDetailRequest**](PatchedNetworkListDetailRequest.md) |  | 

### Return type

[**ResponseNetworkListDetail**](ResponseNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNetworkList

> ResponseRetrieveNetworkListDetail RetrieveNetworkList(ctx, networkListId).Fields(fields).Ipv4(ipv4).Ipv6(ipv6).Execute()

Retrieve details of a Network List



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
	networkListId := int64(789) // int64 | A unique integer value identifying the network list.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ipv4 := true // bool | Filter by IPv4. Only applicable for network lists of type 'ip_cidr'. (optional)
	ipv6 := true // bool | Filter by IPv6. Only applicable for network lists of type 'ip_cidr'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.RetrieveNetworkList(context.Background(), networkListId).Fields(fields).Ipv4(ipv4).Ipv6(ipv6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.RetrieveNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveNetworkList`: ResponseRetrieveNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.RetrieveNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkListId** | **int64** | A unique integer value identifying the network list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ipv4** | **bool** | Filter by IPv4. Only applicable for network lists of type &#39;ip_cidr&#39;. | 
 **ipv6** | **bool** | Filter by IPv6. Only applicable for network lists of type &#39;ip_cidr&#39;. | 

### Return type

[**ResponseRetrieveNetworkListDetail**](ResponseRetrieveNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkList

> ResponseNetworkListDetail UpdateNetworkList(ctx, networkListId).NetworkListDetailRequest(networkListDetailRequest).Execute()

Update a Network List



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
	networkListId := int64(789) // int64 | A unique integer value identifying the network list.
	networkListDetailRequest := *openapiclient.NewNetworkListDetailRequest("Name_example", "Type_example", []string{"Items_example"}) // NetworkListDetailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkListsAPI.UpdateNetworkList(context.Background(), networkListId).NetworkListDetailRequest(networkListDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkListsAPI.UpdateNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkList`: ResponseNetworkListDetail
	fmt.Fprintf(os.Stdout, "Response from `NetworkListsAPI.UpdateNetworkList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkListId** | **int64** | A unique integer value identifying the network list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkListDetailRequest** | [**NetworkListDetailRequest**](NetworkListDetailRequest.md) |  | 

### Return type

[**ResponseNetworkListDetail**](ResponseNetworkListDetail.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

