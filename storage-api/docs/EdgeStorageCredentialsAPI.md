# \EdgeStorageCredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](EdgeStorageCredentialsAPI.md#CreateCredential) | **Post** /edge_storage/s3-credentials | Create a new credential
[**DeleteCredential**](EdgeStorageCredentialsAPI.md#DeleteCredential) | **Delete** /edge_storage/s3-credentials/{accessKey} | Delete a Credential
[**ListCredentials**](EdgeStorageCredentialsAPI.md#ListCredentials) | **Get** /edge_storage/s3-credentials | List credentials
[**RetriveCredential**](EdgeStorageCredentialsAPI.md#RetriveCredential) | **Get** /edge_storage/s3-credentials/{accessKey} | Retrieve details from a credential



## CreateCredential

> ResponseCredential CreateCredential(ctx).CredentialCreateRequest(credentialCreateRequest).Execute()

Create a new credential



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
	credentialCreateRequest := *openapiclient.NewCredentialCreateRequest("Name_example", []string{"Capabilities_example"}) // CredentialCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageCredentialsAPI.CreateCredential(context.Background()).CredentialCreateRequest(credentialCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageCredentialsAPI.CreateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCredential`: ResponseCredential
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageCredentialsAPI.CreateCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialCreateRequest** | [**CredentialCreateRequest**](CredentialCreateRequest.md) |  | 

### Return type

[**ResponseCredential**](ResponseCredential.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredential

> ResponseDeleteCredential DeleteCredential(ctx, accessKey).Execute()

Delete a Credential



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
	accessKey := "accessKey_example" // string | The credential access key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageCredentialsAPI.DeleteCredential(context.Background(), accessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageCredentialsAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCredential`: ResponseDeleteCredential
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageCredentialsAPI.DeleteCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** | The credential access key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteCredential**](ResponseDeleteCredential.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentials

> PaginatedResponseListCredentialList ListCredentials(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List credentials



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageCredentialsAPI.ListCredentials(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageCredentialsAPI.ListCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCredentials`: PaginatedResponseListCredentialList
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageCredentialsAPI.ListCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListCredentialList**](PaginatedResponseListCredentialList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetriveCredential

> ResponseRetrieveCredential RetriveCredential(ctx, accessKey).Fields(fields).Execute()

Retrieve details from a credential



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
	accessKey := "accessKey_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageCredentialsAPI.RetriveCredential(context.Background(), accessKey).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageCredentialsAPI.RetriveCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetriveCredential`: ResponseRetrieveCredential
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageCredentialsAPI.RetriveCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetriveCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveCredential**](ResponseRetrieveCredential.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

