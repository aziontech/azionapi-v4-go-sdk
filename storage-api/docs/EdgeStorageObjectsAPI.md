# \EdgeStorageObjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectKey**](EdgeStorageObjectsAPI.md#CreateObjectKey) | **Post** /edge_storage/buckets/{bucketName}/objects/{objectKey} | Create new object key.
[**DeleteObjectKey**](EdgeStorageObjectsAPI.md#DeleteObjectKey) | **Delete** /edge_storage/buckets/{bucketName}/objects/{objectKey} | Delete object key
[**DownloadObject**](EdgeStorageObjectsAPI.md#DownloadObject) | **Get** /edge_storage/buckets/{bucketName}/objects/{objectKey} | Download object
[**ListObjectKeys**](EdgeStorageObjectsAPI.md#ListObjectKeys) | **Get** /edge_storage/buckets/{bucketName}/objects | List buckets objects
[**UpdateObjectKey**](EdgeStorageObjectsAPI.md#UpdateObjectKey) | **Put** /edge_storage/buckets/{bucketName}/objects/{objectKey} | Update the object key.



## CreateObjectKey

> SuccessObjectOperation CreateObjectKey(ctx, bucketName, objectKey).ContentType(contentType).Body(body).Execute()

Create new object key.



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 
	contentType := "contentType_example" // string | The content type of the file (Example: application/octet-stream). (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageObjectsAPI.CreateObjectKey(context.Background(), bucketName, objectKey).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageObjectsAPI.CreateObjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectKey`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageObjectsAPI.CreateObjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** | The content type of the file (Example: application/octet-stream). | 
 **body** | ***os.File** |  | 

### Return type

[**SuccessObjectOperation**](SuccessObjectOperation.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectKey

> ResponseDeleteBucketObject DeleteObjectKey(ctx, bucketName, objectKey).Execute()

Delete object key



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageObjectsAPI.DeleteObjectKey(context.Background(), bucketName, objectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageObjectsAPI.DeleteObjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectKey`: ResponseDeleteBucketObject
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageObjectsAPI.DeleteObjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteBucketObject**](ResponseDeleteBucketObject.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadObject

> *os.File DownloadObject(ctx, bucketName, objectKey).Fields(fields).Execute()

Download object



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageObjectsAPI.DownloadObject(context.Background(), bucketName, objectKey).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageObjectsAPI.DownloadObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadObject`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageObjectsAPI.DownloadObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectKeys

> []ResponseBucketObject ListObjectKeys(ctx, bucketName).ContinuationToken(continuationToken).Fields(fields).MaxObjectCount(maxObjectCount).Execute()

List buckets objects



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
	bucketName := "bucketName_example" // string | 
	continuationToken := "continuationToken_example" // string | A continuation token for the next page of records. (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	maxObjectCount := int64(789) // int64 | Number of results to be returned on the page. Limited to 1000 objects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageObjectsAPI.ListObjectKeys(context.Background(), bucketName).ContinuationToken(continuationToken).Fields(fields).MaxObjectCount(maxObjectCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageObjectsAPI.ListObjectKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListObjectKeys`: []ResponseBucketObject
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageObjectsAPI.ListObjectKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **string** | A continuation token for the next page of records. | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **maxObjectCount** | **int64** | Number of results to be returned on the page. Limited to 1000 objects. | 

### Return type

[**[]ResponseBucketObject**](ResponseBucketObject.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectKey

> SuccessObjectOperation UpdateObjectKey(ctx, bucketName, objectKey).ContentType(contentType).Body(body).Execute()

Update the object key.



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 
	contentType := "contentType_example" // string | The content type of the file (Example: application/octet-stream). (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeStorageObjectsAPI.UpdateObjectKey(context.Background(), bucketName, objectKey).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeStorageObjectsAPI.UpdateObjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectKey`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `EdgeStorageObjectsAPI.UpdateObjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** | The content type of the file (Example: application/octet-stream). | 
 **body** | ***os.File** |  | 

### Return type

[**SuccessObjectOperation**](SuccessObjectOperation.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

