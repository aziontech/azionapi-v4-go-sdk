# \ApplicationsFunctionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationFunctionInstance**](ApplicationsFunctionAPI.md#CreateApplicationFunctionInstance) | **Post** /edge_application/applications/{application_id}/functions | Create an Application Function Instance
[**DestroyApplicationFunctionInstance**](ApplicationsFunctionAPI.md#DestroyApplicationFunctionInstance) | **Delete** /edge_application/applications/{application_id}/functions/{function_id} | Destroy an Application Function Instance
[**ListApplicationFunctionInstances**](ApplicationsFunctionAPI.md#ListApplicationFunctionInstances) | **Get** /edge_application/applications/{application_id}/functions | List Function Instances
[**PartialUpdateApplicationFunctionInstance**](ApplicationsFunctionAPI.md#PartialUpdateApplicationFunctionInstance) | **Patch** /edge_application/applications/{application_id}/functions/{function_id} | Partially update an Application Function Instance
[**RetrieveApplicationFunctionInstance**](ApplicationsFunctionAPI.md#RetrieveApplicationFunctionInstance) | **Get** /edge_application/applications/{application_id}/functions/{function_id} | Retrieve details of an Application Function Instance
[**UpdateApplicationFunctionInstance**](ApplicationsFunctionAPI.md#UpdateApplicationFunctionInstance) | **Put** /edge_application/applications/{application_id}/functions/{function_id} | Update an Edge Application Function Instance



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
	applicationId := "applicationId_example" // string | 
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
**applicationId** | **string** |  | 

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


## DestroyApplicationFunctionInstance

> ResponseDeleteApplicationFunctionInstance DestroyApplicationFunctionInstance(ctx, applicationId, functionId).Execute()

Destroy an Application Function Instance



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
	functionId := "functionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.DestroyApplicationFunctionInstance(context.Background(), applicationId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionAPI.DestroyApplicationFunctionInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyApplicationFunctionInstance`: ResponseDeleteApplicationFunctionInstance
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionAPI.DestroyApplicationFunctionInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyApplicationFunctionInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteApplicationFunctionInstance**](ResponseDeleteApplicationFunctionInstance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationFunctionInstances

> PaginatedApplicationFunctionInstanceList ListApplicationFunctionInstances(ctx, applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Function Instances



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, last_editor, last_modified, name, args, azion_form, function, active) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsFunctionAPI.ListApplicationFunctionInstances(context.Background(), applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
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
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationFunctionInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 
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
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 
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
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

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
	applicationId := "applicationId_example" // string | 
	functionId := "functionId_example" // string | 
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
**applicationId** | **string** |  | 
**functionId** | **string** |  | 

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

