# \EdgeApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneEdgeApplication**](EdgeApplicationsAPI.md#CloneEdgeApplication) | **Post** /edge_application/applications/{application_id}/clone | Clone an Edge Application
[**CreateEdgeApplication**](EdgeApplicationsAPI.md#CreateEdgeApplication) | **Post** /edge_application/applications | Create an Edge Application
[**DestroyEdgeApplication**](EdgeApplicationsAPI.md#DestroyEdgeApplication) | **Delete** /edge_application/applications/{application_id} | Destroy an Edge Application
[**ListEdgeApplications**](EdgeApplicationsAPI.md#ListEdgeApplications) | **Get** /edge_application/applications | List Edge Applications
[**PartialUpdateEdgeApplication**](EdgeApplicationsAPI.md#PartialUpdateEdgeApplication) | **Patch** /edge_application/applications/{application_id} | Partially update an Edge Application
[**RetrieveEdgeApplication**](EdgeApplicationsAPI.md#RetrieveEdgeApplication) | **Get** /edge_application/applications/{application_id} | Retrieve details of an Edge Application
[**UpdateEdgeApplication**](EdgeApplicationsAPI.md#UpdateEdgeApplication) | **Put** /edge_application/applications/{application_id} | Update an Edge Application



## CloneEdgeApplication

> ResponseRetrieveEdgeApplication CloneEdgeApplication(ctx, applicationId).CloneEdgeApplicationRequest(cloneEdgeApplicationRequest).Execute()

Clone an Edge Application



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
	applicationId := "applicationId_example" // string | 
	cloneEdgeApplicationRequest := *openapiclient.NewCloneEdgeApplicationRequest("Name_example") // CloneEdgeApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.CloneEdgeApplication(context.Background(), applicationId).CloneEdgeApplicationRequest(cloneEdgeApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.CloneEdgeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneEdgeApplication`: ResponseRetrieveEdgeApplication
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.CloneEdgeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneEdgeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneEdgeApplicationRequest** | [**CloneEdgeApplicationRequest**](CloneEdgeApplicationRequest.md) |  | 

### Return type

[**ResponseRetrieveEdgeApplication**](ResponseRetrieveEdgeApplication.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEdgeApplication

> ResponseEdgeApplication CreateEdgeApplication(ctx).EdgeApplicationRequest(edgeApplicationRequest).Execute()

Create an Edge Application



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
	edgeApplicationRequest := *openapiclient.NewEdgeApplicationRequest("Name_example") // EdgeApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.CreateEdgeApplication(context.Background()).EdgeApplicationRequest(edgeApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.CreateEdgeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeApplication`: ResponseEdgeApplication
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.CreateEdgeApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeApplicationRequest** | [**EdgeApplicationRequest**](EdgeApplicationRequest.md) |  | 

### Return type

[**ResponseEdgeApplication**](ResponseEdgeApplication.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeApplication

> ResponseDeleteEdgeApplication DestroyEdgeApplication(ctx, applicationId).Execute()

Destroy an Edge Application



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
	applicationId := int64(789) // int64 | A unique integer value identifying the edge application.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.DestroyEdgeApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.DestroyEdgeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeApplication`: ResponseDeleteEdgeApplication
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.DestroyEdgeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the edge application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteEdgeApplication**](ResponseDeleteEdgeApplication.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeApplications

> PaginatedEdgeApplicationList ListEdgeApplications(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Applications



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, id, last_editor, last_modified, active, debug, product_version) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.ListEdgeApplications(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.ListEdgeApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeApplications`: PaginatedEdgeApplicationList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.ListEdgeApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, id, last_editor, last_modified, active, debug, product_version) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeApplicationList**](PaginatedEdgeApplicationList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeApplication

> ResponseEdgeApplication PartialUpdateEdgeApplication(ctx, applicationId).PatchedEdgeApplicationRequest(patchedEdgeApplicationRequest).Execute()

Partially update an Edge Application



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
	applicationId := int64(789) // int64 | A unique integer value identifying the edge application.
	patchedEdgeApplicationRequest := *openapiclient.NewPatchedEdgeApplicationRequest() // PatchedEdgeApplicationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.PartialUpdateEdgeApplication(context.Background(), applicationId).PatchedEdgeApplicationRequest(patchedEdgeApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.PartialUpdateEdgeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeApplication`: ResponseEdgeApplication
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.PartialUpdateEdgeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the edge application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEdgeApplicationRequest** | [**PatchedEdgeApplicationRequest**](PatchedEdgeApplicationRequest.md) |  | 

### Return type

[**ResponseEdgeApplication**](ResponseEdgeApplication.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeApplication

> ResponseRetrieveEdgeApplication RetrieveEdgeApplication(ctx, applicationId).Fields(fields).Execute()

Retrieve details of an Edge Application



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
	applicationId := int64(789) // int64 | A unique integer value identifying the edge application.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.RetrieveEdgeApplication(context.Background(), applicationId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.RetrieveEdgeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeApplication`: ResponseRetrieveEdgeApplication
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.RetrieveEdgeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the edge application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeApplication**](ResponseRetrieveEdgeApplication.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeApplication

> ResponseEdgeApplication UpdateEdgeApplication(ctx, applicationId).EdgeApplicationRequest(edgeApplicationRequest).Execute()

Update an Edge Application



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
	applicationId := int64(789) // int64 | A unique integer value identifying the edge application.
	edgeApplicationRequest := *openapiclient.NewEdgeApplicationRequest("Name_example") // EdgeApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsAPI.UpdateEdgeApplication(context.Background(), applicationId).EdgeApplicationRequest(edgeApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsAPI.UpdateEdgeApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeApplication`: ResponseEdgeApplication
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsAPI.UpdateEdgeApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the edge application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeApplicationRequest** | [**EdgeApplicationRequest**](EdgeApplicationRequest.md) |  | 

### Return type

[**ResponseEdgeApplication**](ResponseEdgeApplication.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

