# \StorageObjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectKey**](StorageObjectsAPI.md#CreateObjectKey) | **Post** /workspace/storage/buckets/{bucket_name}/objects/{object_key} | Create new object key.
[**DeleteObjectKey**](StorageObjectsAPI.md#DeleteObjectKey) | **Delete** /workspace/storage/buckets/{bucket_name}/objects/{object_key} | Delete object key
[**DownloadObject**](StorageObjectsAPI.md#DownloadObject) | **Get** /workspace/storage/buckets/{bucket_name}/objects/{object_key} | Download object
[**ListObjects**](StorageObjectsAPI.md#ListObjects) | **Get** /workspace/storage/buckets/{bucket_name}/objects | List objects from bucket
[**UpdateObjectKey**](StorageObjectsAPI.md#UpdateObjectKey) | **Put** /workspace/storage/buckets/{bucket_name}/objects/{object_key} | Update the object key.



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
	bucketName := "bucketName_example" // string | The name of the bucket
	objectKey := "objectKey_example" // string | The key/path of the object within the bucket
	contentType := "contentType_example" // string | The MIME type of the object being uploaded (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageObjectsAPI.CreateObjectKey(context.Background(), bucketName, objectKey).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectsAPI.CreateObjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateObjectKey`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `StorageObjectsAPI.CreateObjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 
**objectKey** | **string** | The key/path of the object within the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** | The MIME type of the object being uploaded | 
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

> ResponseAsyncDeleteBucketObject DeleteObjectKey(ctx, bucketName, objectKey).Execute()

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
	bucketName := "bucketName_example" // string | The name of the bucket
	objectKey := "objectKey_example" // string | The key/path of the object within the bucket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageObjectsAPI.DeleteObjectKey(context.Background(), bucketName, objectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectsAPI.DeleteObjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectKey`: ResponseAsyncDeleteBucketObject
	fmt.Fprintf(os.Stdout, "Response from `StorageObjectsAPI.DeleteObjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 
**objectKey** | **string** | The key/path of the object within the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseAsyncDeleteBucketObject**](ResponseAsyncDeleteBucketObject.md)

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
	bucketName := "bucketName_example" // string | The name of the bucket
	objectKey := "objectKey_example" // string | The key/path of the object within the bucket
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageObjectsAPI.DownloadObject(context.Background(), bucketName, objectKey).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectsAPI.DownloadObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadObject`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StorageObjectsAPI.DownloadObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 
**objectKey** | **string** | The key/path of the object within the bucket | 

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


## ListObjects

> ResponseBucketObject ListObjects(ctx, bucketName).AllLevels(allLevels).ContinuationToken(continuationToken).Fields(fields).MaxObjectCount(maxObjectCount).Prefix(prefix).Execute()

List objects from bucket



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
	bucketName := "bucketName_example" // string | The name of the bucket
	allLevels := true // bool | If true, lists objects recursively. If false, lists only the first level using Delimiter='/' (default: true). (optional)
	continuationToken := "continuationToken_example" // string | A continuation token for the next page of records. (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	maxObjectCount := int64(789) // int64 | Number of results to be returned on the page. Limited to 1000 objects. (optional)
	prefix := "prefix_example" // string | Filter objects by key prefix. If empty, lists from the bucket root (default: empty). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageObjectsAPI.ListObjects(context.Background(), bucketName).AllLevels(allLevels).ContinuationToken(continuationToken).Fields(fields).MaxObjectCount(maxObjectCount).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectsAPI.ListObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListObjects`: ResponseBucketObject
	fmt.Fprintf(os.Stdout, "Response from `StorageObjectsAPI.ListObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allLevels** | **bool** | If true, lists objects recursively. If false, lists only the first level using Delimiter&#x3D;&#39;/&#39; (default: true). | 
 **continuationToken** | **string** | A continuation token for the next page of records. | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **maxObjectCount** | **int64** | Number of results to be returned on the page. Limited to 1000 objects. | 
 **prefix** | **string** | Filter objects by key prefix. If empty, lists from the bucket root (default: empty). | 

### Return type

[**ResponseBucketObject**](ResponseBucketObject.md)

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
	bucketName := "bucketName_example" // string | The name of the bucket
	objectKey := "objectKey_example" // string | The key/path of the object within the bucket
	contentType := "contentType_example" // string | The MIME type of the object being uploaded (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageObjectsAPI.UpdateObjectKey(context.Background(), bucketName, objectKey).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageObjectsAPI.UpdateObjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateObjectKey`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `StorageObjectsAPI.UpdateObjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** | The name of the bucket | 
**objectKey** | **string** | The key/path of the object within the bucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** | The MIME type of the object being uploaded | 
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

