# \ApplicationsFunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationFunctionInstance**](ApplicationsFunctionAPI.md#CreateApplicationFunctionInstance) | **Post** /workspace/applications/{application_id}/functions | Create an Application Function Instance
[**DeleteApplicationFunctionInstance**](ApplicationsFunctionAPI.md#DeleteApplicationFunctionInstance) | **Delete** /workspace/applications/{application_id}/functions/{function_id} | Delete an Application Function Instance
[**ListApplicationFunctionInstances**](ApplicationsFunctionAPI.md#ListApplicationFunctionInstances) | **Get** /workspace/applications/{application_id}/functions | List Function Instances
[**PartialUpdateApplicationFunctionInstance**](ApplicationsFunctionAPI.md#PartialUpdateApplicationFunctionInstance) | **Patch** /workspace/applications/{application_id}/functions/{function_id} | Partially update an Application Function Instance
[**RetrieveApplicationFunctionInstance**](ApplicationsFunctionAPI.md#RetrieveApplicationFunctionInstance) | **Get** /workspace/applications/{application_id}/functions/{function_id} | Retrieve details of an Application Function Instance
[**UpdateApplicationFunctionInstance**](ApplicationsFunctionAPI.md#UpdateApplicationFunctionInstance) | **Put** /workspace/applications/{application_id}/functions/{function_id} | Update an Edge Application Function Instance



## CreateApplicationFunctionInstance

> ResponseApplicationFunctionInstance CreateApplicationFunctionInstance(ctx, applicationId).ApplicationFunctionInstanceRequest(applicationFunctionInstanceRequest).Execute()

Create an Application Function Instance



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	applicationFunctionInstanceRequest := *openapiclient.NewApplicationFunctionInstanceRequest("Name_example", int64(123)) // ApplicationFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.CreateApplicationFunctionInstance(context.Background(), applicationId).ApplicationFunctionInstanceRequest(applicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.CreateApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationFunctionInstance`: ResponseApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.CreateApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationFunctionInstanceRequest** | [**ApplicationFunctionInstanceRequest**](ApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseApplicationFunctionInstance**](ResponseApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationFunctionInstance

> ResponseAsyncDeleteApplicationFunctionInstance DeleteApplicationFunctionInstance(ctx, applicationId, functionId).Execute()

Delete an Application Function Instance



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.DeleteApplicationFunctionInstance(context.Background(), applicationId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.DeleteApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationFunctionInstance`: ResponseAsyncDeleteApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.DeleteApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseAsyncDeleteApplicationFunctionInstance**](ResponseAsyncDeleteApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationFunctionInstances

> PaginatedApplicationFunctionInstanceList ListApplicationFunctionInstances(ctx, applicationId).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Function Instances



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
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
	resp, r, err := apiClient.ApplicationsFunctionAPI.ListApplicationFunctionInstances(context.Background(), applicationId).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.ListApplicationFunctionInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationFunctionInstances`: PaginatedApplicationFunctionInstanceList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.ListApplicationFunctionInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationFunctionInstancesRequest struct via the builder pattern


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

[**PaginatedApplicationFunctionInstanceList**](PaginatedApplicationFunctionInstanceList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateApplicationFunctionInstance

> ResponseApplicationFunctionInstance PartialUpdateApplicationFunctionInstance(ctx, applicationId, functionId).PatchedApplicationFunctionInstanceRequest(patchedApplicationFunctionInstanceRequest).Execute()

Partially update an Application Function Instance



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.
	patchedApplicationFunctionInstanceRequest := *openapiclient.NewPatchedApplicationFunctionInstanceRequest() // PatchedApplicationFunctionInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.PartialUpdateApplicationFunctionInstance(context.Background(), applicationId, functionId).PatchedApplicationFunctionInstanceRequest(patchedApplicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.PartialUpdateApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateApplicationFunctionInstance`: ResponseApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.PartialUpdateApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedApplicationFunctionInstanceRequest** | [**PatchedApplicationFunctionInstanceRequest**](PatchedApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseApplicationFunctionInstance**](ResponseApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveApplicationFunctionInstance

> ResponseRetrieveApplicationFunctionInstance RetrieveApplicationFunctionInstance(ctx, applicationId, functionId).Fields(fields).Execute()

Retrieve details of an Application Function Instance



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.RetrieveApplicationFunctionInstance(context.Background(), applicationId, functionId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.RetrieveApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveApplicationFunctionInstance`: ResponseRetrieveApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.RetrieveApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveApplicationFunctionInstance**](ResponseRetrieveApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationFunctionInstance

> ResponseApplicationFunctionInstance UpdateApplicationFunctionInstance(ctx, applicationId, functionId).ApplicationFunctionInstanceRequest(applicationFunctionInstanceRequest).Execute()

Update an Edge Application Function Instance



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	functionId := int64(789) // int64 | A unique integer value identifying the function instance.
	applicationFunctionInstanceRequest := *openapiclient.NewApplicationFunctionInstanceRequest("Name_example", int64(123)) // ApplicationFunctionInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.UpdateApplicationFunctionInstance(context.Background(), applicationId, functionId).ApplicationFunctionInstanceRequest(applicationFunctionInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.UpdateApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationFunctionInstance`: ResponseApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.UpdateApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**functionId** | **int64** | A unique integer value identifying the function instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationFunctionInstanceRequest** | [**ApplicationFunctionInstanceRequest**](ApplicationFunctionInstanceRequest.md) |  | 

### Return type

[**ResponseApplicationFunctionInstance**](ResponseApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

