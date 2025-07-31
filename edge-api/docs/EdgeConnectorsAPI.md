# \EdgeConnectorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeConnector**](EdgeConnectorsAPI.md#CreateEdgeConnector) | **Post** /edge_connector/connectors | Create an Edge Connector
[**DestroyEdgeConnector**](EdgeConnectorsAPI.md#DestroyEdgeConnector) | **Delete** /edge_connector/connectors/{id} | Destroy an Edge Connector
[**ListEdgeConnectors**](EdgeConnectorsAPI.md#ListEdgeConnectors) | **Get** /edge_connector/connectors | List Edge Connectors
[**PartialUpdateEdgeConnector**](EdgeConnectorsAPI.md#PartialUpdateEdgeConnector) | **Patch** /edge_connector/connectors/{id} | Partially update an Edge Connector
[**RetrieveEdgeConnector**](EdgeConnectorsAPI.md#RetrieveEdgeConnector) | **Get** /edge_connector/connectors/{id} | Retrieve details of an Edge Connector
[**UpdateEdgeConnector**](EdgeConnectorsAPI.md#UpdateEdgeConnector) | **Put** /edge_connector/connectors/{id} | Update an Edge Connector



## CreateEdgeConnector

> ResponseEdgeConnectorPolymorphic CreateEdgeConnector(ctx).EdgeConnectorPolymorphicRequest(edgeConnectorPolymorphicRequest).Execute()

Create an Edge Connector



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
	edgeConnectorPolymorphicRequest := openapiclient.EdgeConnectorPolymorphicRequest{EdgeConnectorHTTPRequest: openapiclient.NewEdgeConnectorHTTPRequest("Name_example", "Type_example", *openapiclient.NewEdgeConnectorHTTPAttributesRequest([]openapiclient.AddressRequest{*openapiclient.NewAddressRequest("Address_example")}))} // EdgeConnectorPolymorphicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeConnectorsAPI.CreateEdgeConnector(context.Background()).EdgeConnectorPolymorphicRequest(edgeConnectorPolymorphicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeConnectorsAPI.CreateEdgeConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeConnector`: ResponseEdgeConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `EdgeConnectorsAPI.CreateEdgeConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeConnectorPolymorphicRequest** | [**EdgeConnectorPolymorphicRequest**](EdgeConnectorPolymorphicRequest.md) |  | 

### Return type

[**ResponseEdgeConnectorPolymorphic**](ResponseEdgeConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeConnector

> ResponseDeleteEdgeConnectorPolymorphic DestroyEdgeConnector(ctx, id).Execute()

Destroy an Edge Connector



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
	resp, r, err := apiClient.EdgeConnectorsAPI.DestroyEdgeConnector(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeConnectorsAPI.DestroyEdgeConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeConnector`: ResponseDeleteEdgeConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `EdgeConnectorsAPI.DestroyEdgeConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteEdgeConnectorPolymorphic**](ResponseDeleteEdgeConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeConnectors

> PaginatedEdgeConnectorPolymorphicList ListEdgeConnectors(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Connectors



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: ) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeConnectorsAPI.ListEdgeConnectors(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeConnectorsAPI.ListEdgeConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeConnectors`: PaginatedEdgeConnectorPolymorphicList
	fmt.Fprintf(os.Stdout, "Response from `EdgeConnectorsAPI.ListEdgeConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: ) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeConnectorPolymorphicList**](PaginatedEdgeConnectorPolymorphicList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeConnector

> ResponseEdgeConnectorPolymorphic PartialUpdateEdgeConnector(ctx, id).PatchedEdgeConnectorPolymorphicRequest(patchedEdgeConnectorPolymorphicRequest).Execute()

Partially update an Edge Connector



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
	patchedEdgeConnectorPolymorphicRequest := openapiclient.PatchedEdgeConnectorPolymorphicRequest{PatchedEdgeConnectorHTTPRequest: openapiclient.NewPatchedEdgeConnectorHTTPRequest()} // PatchedEdgeConnectorPolymorphicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeConnectorsAPI.PartialUpdateEdgeConnector(context.Background(), id).PatchedEdgeConnectorPolymorphicRequest(patchedEdgeConnectorPolymorphicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeConnectorsAPI.PartialUpdateEdgeConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeConnector`: ResponseEdgeConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `EdgeConnectorsAPI.PartialUpdateEdgeConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEdgeConnectorPolymorphicRequest** | [**PatchedEdgeConnectorPolymorphicRequest**](PatchedEdgeConnectorPolymorphicRequest.md) |  | 

### Return type

[**ResponseEdgeConnectorPolymorphic**](ResponseEdgeConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeConnector

> ResponseRetrieveEdgeConnectorPolymorphic RetrieveEdgeConnector(ctx, id).Fields(fields).Execute()

Retrieve details of an Edge Connector



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
	resp, r, err := apiClient.EdgeConnectorsAPI.RetrieveEdgeConnector(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeConnectorsAPI.RetrieveEdgeConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeConnector`: ResponseRetrieveEdgeConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `EdgeConnectorsAPI.RetrieveEdgeConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeConnectorPolymorphic**](ResponseRetrieveEdgeConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeConnector

> ResponseEdgeConnectorPolymorphic UpdateEdgeConnector(ctx, id).EdgeConnectorPolymorphicRequest(edgeConnectorPolymorphicRequest).Execute()

Update an Edge Connector



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
	edgeConnectorPolymorphicRequest := openapiclient.EdgeConnectorPolymorphicRequest{EdgeConnectorHTTPRequest: openapiclient.NewEdgeConnectorHTTPRequest("Name_example", "Type_example", *openapiclient.NewEdgeConnectorHTTPAttributesRequest([]openapiclient.AddressRequest{*openapiclient.NewAddressRequest("Address_example")}))} // EdgeConnectorPolymorphicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeConnectorsAPI.UpdateEdgeConnector(context.Background(), id).EdgeConnectorPolymorphicRequest(edgeConnectorPolymorphicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeConnectorsAPI.UpdateEdgeConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeConnector`: ResponseEdgeConnectorPolymorphic
	fmt.Fprintf(os.Stdout, "Response from `EdgeConnectorsAPI.UpdateEdgeConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeConnectorPolymorphicRequest** | [**EdgeConnectorPolymorphicRequest**](EdgeConnectorPolymorphicRequest.md) |  | 

### Return type

[**ResponseEdgeConnectorPolymorphic**](ResponseEdgeConnectorPolymorphic.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

