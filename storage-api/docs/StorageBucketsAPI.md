# \StorageBucketsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](StorageBucketsAPI.md#CreateBucket) | **Post** /edge_storage/buckets | Create a new bucket
[**DeleteBucket**](StorageBucketsAPI.md#DeleteBucket) | **Delete** /edge_storage/buckets/{name} | Delete a bucket
[**ListBuckets**](StorageBucketsAPI.md#ListBuckets) | **Get** /edge_storage/buckets | List buckets
[**RetrieveBucket**](StorageBucketsAPI.md#RetrieveBucket) | **Get** /edge_storage/buckets/{name} | Retrieve details from a bucket
[**UpdateBucket**](StorageBucketsAPI.md#UpdateBucket) | **Patch** /edge_storage/buckets/{name} | Update bucket info



## CreateBucket

> ResponseBucketCreate CreateBucket(ctx).BucketCreateRequest(bucketCreateRequest).Execute()

Create a new bucket



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
	bucketCreateRequest := *openapiclient.NewBucketCreateRequest("Name_example", "EdgeAccess_example") // BucketCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.CreateBucket(context.Background()).BucketCreateRequest(bucketCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.CreateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBucket`: ResponseBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.CreateBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketCreateRequest** | [**BucketCreateRequest**](BucketCreateRequest.md) |  | 

### Return type

[**ResponseBucketCreate**](ResponseBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> ResponseDeleteBucketCreate DeleteBucket(ctx, name).Execute()

Delete a bucket



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.DeleteBucket(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.DeleteBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBucket`: ResponseDeleteBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.DeleteBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteBucketCreate**](ResponseDeleteBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> PaginatedBucketList ListBuckets(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List buckets



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
	resp, r, err := apiClient.StorageBucketsAPI.ListBuckets(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.ListBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuckets`: PaginatedBucketList
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.ListBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedBucketList**](PaginatedBucketList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBucket

> ResponseRetrieveBucket RetrieveBucket(ctx, name).Fields(fields).Execute()

Retrieve details from a bucket



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
	name := "name_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.RetrieveBucket(context.Background(), name).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.RetrieveBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBucket`: ResponseRetrieveBucket
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.RetrieveBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveBucket**](ResponseRetrieveBucket.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> ResponseBucketCreate UpdateBucket(ctx, name).PatchedBucketRequest(patchedBucketRequest).Execute()

Update bucket info



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
	name := "name_example" // string | 
	patchedBucketRequest := *openapiclient.NewPatchedBucketRequest() // PatchedBucketRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageBucketsAPI.UpdateBucket(context.Background(), name).PatchedBucketRequest(patchedBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageBucketsAPI.UpdateBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBucket`: ResponseBucketCreate
	fmt.Fprintf(os.Stdout, "Response from `StorageBucketsAPI.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedBucketRequest** | [**PatchedBucketRequest**](PatchedBucketRequest.md) |  | 

### Return type

[**ResponseBucketCreate**](ResponseBucketCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

