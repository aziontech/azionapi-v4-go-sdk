# \EdgeApplicationsErrorResponsesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListErrorResponses**](EdgeApplicationsErrorResponsesAPI.md#ListErrorResponses) | **Get** /edge_application/applications/{edgeApplicationId}/error_responses | List Error Responses for a Specific Edge Application
[**PartialUpdateErrorResponse**](EdgeApplicationsErrorResponsesAPI.md#PartialUpdateErrorResponse) | **Patch** /edge_application/applications/{edgeApplicationId}/error_responses/{id} | Partially Update an Error Response for a Specific Edge Application
[**RetrieveErrorResponse**](EdgeApplicationsErrorResponsesAPI.md#RetrieveErrorResponse) | **Get** /edge_application/applications/{edgeApplicationId}/error_responses/{id} | Retrieve Details of an Error Response for a Specific Edge Application
[**UpdateErrorResponse**](EdgeApplicationsErrorResponsesAPI.md#UpdateErrorResponse) | **Put** /edge_application/applications/{edgeApplicationId}/error_responses/{id} | Update an Error Response for a Specific Edge Application



## ListErrorResponses

> PaginatedResponseListErrorResponsesList ListErrorResponses(ctx, edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Error Responses for a Specific Edge Application



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: edge_application_id, origin_id, error_responses, name) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsErrorResponsesAPI.ListErrorResponses(context.Background(), edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsErrorResponsesAPI.ListErrorResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListErrorResponses`: PaginatedResponseListErrorResponsesList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsErrorResponsesAPI.ListErrorResponses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListErrorResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: edge_application_id, origin_id, error_responses, name) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListErrorResponsesList**](PaginatedResponseListErrorResponsesList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateErrorResponse

> ResponseErrorResponses PartialUpdateErrorResponse(ctx, edgeApplicationId, id).PatchedErrorResponsesRequest(patchedErrorResponsesRequest).Execute()

Partially Update an Error Response for a Specific Edge Application



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	patchedErrorResponsesRequest := *openapiclient.NewPatchedErrorResponsesRequest() // PatchedErrorResponsesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsErrorResponsesAPI.PartialUpdateErrorResponse(context.Background(), edgeApplicationId, id).PatchedErrorResponsesRequest(patchedErrorResponsesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsErrorResponsesAPI.PartialUpdateErrorResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateErrorResponse`: ResponseErrorResponses
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsErrorResponsesAPI.PartialUpdateErrorResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateErrorResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedErrorResponsesRequest** | [**PatchedErrorResponsesRequest**](PatchedErrorResponsesRequest.md) |  | 

### Return type

[**ResponseErrorResponses**](ResponseErrorResponses.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveErrorResponse

> ResponseRetrieveErrorResponses RetrieveErrorResponse(ctx, edgeApplicationId, id).Fields(fields).Execute()

Retrieve Details of an Error Response for a Specific Edge Application



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsErrorResponsesAPI.RetrieveErrorResponse(context.Background(), edgeApplicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsErrorResponsesAPI.RetrieveErrorResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveErrorResponse`: ResponseRetrieveErrorResponses
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsErrorResponsesAPI.RetrieveErrorResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveErrorResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveErrorResponses**](ResponseRetrieveErrorResponses.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateErrorResponse

> ResponseErrorResponses UpdateErrorResponse(ctx, edgeApplicationId, id).ErrorResponsesRequest(errorResponsesRequest).Execute()

Update an Error Response for a Specific Edge Application



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	errorResponsesRequest := *openapiclient.NewErrorResponsesRequest([]openapiclient.NestedErrorResponseRequest{*openapiclient.NewNestedErrorResponseRequest(int64(123), int64(123))}) // ErrorResponsesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsErrorResponsesAPI.UpdateErrorResponse(context.Background(), edgeApplicationId, id).ErrorResponsesRequest(errorResponsesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsErrorResponsesAPI.UpdateErrorResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateErrorResponse`: ResponseErrorResponses
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsErrorResponsesAPI.UpdateErrorResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateErrorResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **errorResponsesRequest** | [**ErrorResponsesRequest**](ErrorResponsesRequest.md) |  | 

### Return type

[**ResponseErrorResponses**](ResponseErrorResponses.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

