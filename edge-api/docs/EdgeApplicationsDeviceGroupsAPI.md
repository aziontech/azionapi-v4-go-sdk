# \EdgeApplicationsDeviceGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceGroup**](EdgeApplicationsDeviceGroupsAPI.md#CreateDeviceGroup) | **Post** /edge_application/applications/{edgeApplicationId}/device_groups | Create an Edge Applications Device Group
[**DestroyDeviceGroup**](EdgeApplicationsDeviceGroupsAPI.md#DestroyDeviceGroup) | **Delete** /edge_application/applications/{edgeApplicationId}/device_groups/{id} | Destroy an Edge Applications Device Group
[**ListDeviceGroups**](EdgeApplicationsDeviceGroupsAPI.md#ListDeviceGroups) | **Get** /edge_application/applications/{edgeApplicationId}/device_groups | List Edge Applications Device Groups
[**PartialUpdateDeviceGroup**](EdgeApplicationsDeviceGroupsAPI.md#PartialUpdateDeviceGroup) | **Patch** /edge_application/applications/{edgeApplicationId}/device_groups/{id} | Partially update an Edge Applications Device Group
[**RetrieveDeviceGroup**](EdgeApplicationsDeviceGroupsAPI.md#RetrieveDeviceGroup) | **Get** /edge_application/applications/{edgeApplicationId}/device_groups/{id} | Retrieve details of a Device Group
[**UpdateDeviceGroup**](EdgeApplicationsDeviceGroupsAPI.md#UpdateDeviceGroup) | **Put** /edge_application/applications/{edgeApplicationId}/device_groups/{id} | Update an Edge Applications Device Group



## CreateDeviceGroup

> ResponseEdgeApplicationDeviceGroups CreateDeviceGroup(ctx, edgeApplicationId).EdgeApplicationDeviceGroupsRequest(edgeApplicationDeviceGroupsRequest).Execute()

Create an Edge Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	edgeApplicationId := "edgeApplicationId_example" // string | 
	edgeApplicationDeviceGroupsRequest := *openapiclient.NewEdgeApplicationDeviceGroupsRequest("Name_example", "UserAgent_example") // EdgeApplicationDeviceGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.CreateDeviceGroup(context.Background(), edgeApplicationId).EdgeApplicationDeviceGroupsRequest(edgeApplicationDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.CreateDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceGroup`: ResponseEdgeApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.CreateDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeApplicationDeviceGroupsRequest** | [**EdgeApplicationDeviceGroupsRequest**](EdgeApplicationDeviceGroupsRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationDeviceGroups**](ResponseEdgeApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDeviceGroup

> ResponseDeleteEdgeApplicationDeviceGroups DestroyDeviceGroup(ctx, edgeApplicationId, id).Execute()

Destroy an Edge Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.DestroyDeviceGroup(context.Background(), edgeApplicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.DestroyDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyDeviceGroup`: ResponseDeleteEdgeApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.DestroyDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteEdgeApplicationDeviceGroups**](ResponseDeleteEdgeApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceGroups

> PaginatedResponseListEdgeApplicationDeviceGroupsList ListDeviceGroups(ctx, edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Applications Device Groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	edgeApplicationId := "edgeApplicationId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id, user_agent) (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.ListDeviceGroups(context.Background(), edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.ListDeviceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceGroups`: PaginatedResponseListEdgeApplicationDeviceGroupsList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.ListDeviceGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id, user_agent) | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListEdgeApplicationDeviceGroupsList**](PaginatedResponseListEdgeApplicationDeviceGroupsList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDeviceGroup

> ResponseEdgeApplicationDeviceGroups PartialUpdateDeviceGroup(ctx, edgeApplicationId, id).PatchedEdgeApplicationDeviceGroupsRequest(patchedEdgeApplicationDeviceGroupsRequest).Execute()

Partially update an Edge Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	patchedEdgeApplicationDeviceGroupsRequest := *openapiclient.NewPatchedEdgeApplicationDeviceGroupsRequest() // PatchedEdgeApplicationDeviceGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.PartialUpdateDeviceGroup(context.Background(), edgeApplicationId, id).PatchedEdgeApplicationDeviceGroupsRequest(patchedEdgeApplicationDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.PartialUpdateDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDeviceGroup`: ResponseEdgeApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.PartialUpdateDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedEdgeApplicationDeviceGroupsRequest** | [**PatchedEdgeApplicationDeviceGroupsRequest**](PatchedEdgeApplicationDeviceGroupsRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationDeviceGroups**](ResponseEdgeApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDeviceGroup

> ResponseRetrieveEdgeApplicationDeviceGroups RetrieveDeviceGroup(ctx, edgeApplicationId, id).Fields(fields).Execute()

Retrieve details of a Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.RetrieveDeviceGroup(context.Background(), edgeApplicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.RetrieveDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDeviceGroup`: ResponseRetrieveEdgeApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.RetrieveDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeApplicationDeviceGroups**](ResponseRetrieveEdgeApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceGroup

> ResponseEdgeApplicationDeviceGroups UpdateDeviceGroup(ctx, edgeApplicationId, id).EdgeApplicationDeviceGroupsRequest(edgeApplicationDeviceGroupsRequest).Execute()

Update an Edge Applications Device Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	edgeApplicationDeviceGroupsRequest := *openapiclient.NewEdgeApplicationDeviceGroupsRequest("Name_example", "UserAgent_example") // EdgeApplicationDeviceGroupsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.UpdateDeviceGroup(context.Background(), edgeApplicationId, id).EdgeApplicationDeviceGroupsRequest(edgeApplicationDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.UpdateDeviceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceGroup`: ResponseEdgeApplicationDeviceGroups
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.UpdateDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeApplicationDeviceGroupsRequest** | [**EdgeApplicationDeviceGroupsRequest**](EdgeApplicationDeviceGroupsRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationDeviceGroups**](ResponseEdgeApplicationDeviceGroups.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

