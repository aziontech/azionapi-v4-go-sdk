# \DataStreamStreamsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataStream**](DataStreamStreamsAPI.md#CreateDataStream) | **Post** /data_stream/streams | Create a Data Stream
[**DestroyDataStream**](DataStreamStreamsAPI.md#DestroyDataStream) | **Delete** /data_stream/streams/{id} | Destroy a Data Stream
[**ListDataStreams**](DataStreamStreamsAPI.md#ListDataStreams) | **Get** /data_stream/streams | List Data Streams
[**PartialUpdateDataStream**](DataStreamStreamsAPI.md#PartialUpdateDataStream) | **Patch** /data_stream/streams/{id} | Partially update a Data Stream
[**RetrieveDataStream**](DataStreamStreamsAPI.md#RetrieveDataStream) | **Get** /data_stream/streams/{id} | Retrieve details of a Data Stream
[**UpdateDataStream**](DataStreamStreamsAPI.md#UpdateDataStream) | **Put** /data_stream/streams/{id} | Update a Data Stream



## CreateDataStream

> ResponseDataStream CreateDataStream(ctx).DataStreamRequest(dataStreamRequest).Execute()

Create a Data Stream



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
	dataStreamRequest := *openapiclient.NewDataStreamRequest("Name_example", []openapiclient.InputPolymorphicInputDataSourceAttributesRequest{*openapiclient.NewInputPolymorphicInputDataSourceAttributesRequest("Type_example", *openapiclient.NewInputDataSourceRequest("DataSource_example"))}, []openapiclient.TransformPolymorphicRequest{openapiclient.TransformPolymorphicRequest{TransformPolymorphicTransformFilterWorkloadsAttributesRequest: openapiclient.NewTransformPolymorphicTransformFilterWorkloadsAttributesRequest("Type_example", *openapiclient.NewTransformFilterWorkloadsRequest([]int64{int64(123)}))}}, []openapiclient.OutputRequest{*openapiclient.NewOutputRequest("Type_example", openapiclient.OutputPolymorphicRequest{AWSKinesisFirehoseEndpointRequest: openapiclient.NewAWSKinesisFirehoseEndpointRequest("AccessKey_example", "StreamName_example", "Region_example", "SecretKey_example")})}) // DataStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataStreamStreamsAPI.CreateDataStream(context.Background()).DataStreamRequest(dataStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamStreamsAPI.CreateDataStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataStream`: ResponseDataStream
	fmt.Fprintf(os.Stdout, "Response from `DataStreamStreamsAPI.CreateDataStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataStreamRequest** | [**DataStreamRequest**](DataStreamRequest.md) |  | 

### Return type

[**ResponseDataStream**](ResponseDataStream.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDataStream

> ResponseDeleteDataStream DestroyDataStream(ctx, id).Execute()

Destroy a Data Stream



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
	resp, r, err := apiClient.DataStreamStreamsAPI.DestroyDataStream(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamStreamsAPI.DestroyDataStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyDataStream`: ResponseDeleteDataStream
	fmt.Fprintf(os.Stdout, "Response from `DataStreamStreamsAPI.DestroyDataStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDataStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteDataStream**](ResponseDeleteDataStream.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataStreams

> PaginatedResponseListDataStreamList ListDataStreams(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Data Streams



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
	resp, r, err := apiClient.DataStreamStreamsAPI.ListDataStreams(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamStreamsAPI.ListDataStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataStreams`: PaginatedResponseListDataStreamList
	fmt.Fprintf(os.Stdout, "Response from `DataStreamStreamsAPI.ListDataStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListDataStreamList**](PaginatedResponseListDataStreamList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDataStream

> ResponseDataStream PartialUpdateDataStream(ctx, id).PatchedDataStreamRequest(patchedDataStreamRequest).Execute()

Partially update a Data Stream



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
	patchedDataStreamRequest := *openapiclient.NewPatchedDataStreamRequest() // PatchedDataStreamRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataStreamStreamsAPI.PartialUpdateDataStream(context.Background(), id).PatchedDataStreamRequest(patchedDataStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamStreamsAPI.PartialUpdateDataStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDataStream`: ResponseDataStream
	fmt.Fprintf(os.Stdout, "Response from `DataStreamStreamsAPI.PartialUpdateDataStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDataStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDataStreamRequest** | [**PatchedDataStreamRequest**](PatchedDataStreamRequest.md) |  | 

### Return type

[**ResponseDataStream**](ResponseDataStream.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDataStream

> ResponseRetrieveDataStream RetrieveDataStream(ctx, id).Fields(fields).Execute()

Retrieve details of a Data Stream



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
	resp, r, err := apiClient.DataStreamStreamsAPI.RetrieveDataStream(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamStreamsAPI.RetrieveDataStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDataStream`: ResponseRetrieveDataStream
	fmt.Fprintf(os.Stdout, "Response from `DataStreamStreamsAPI.RetrieveDataStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDataStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveDataStream**](ResponseRetrieveDataStream.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataStream

> ResponseDataStream UpdateDataStream(ctx, id).DataStreamRequest(dataStreamRequest).Execute()

Update a Data Stream



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
	dataStreamRequest := *openapiclient.NewDataStreamRequest("Name_example", []openapiclient.InputPolymorphicInputDataSourceAttributesRequest{*openapiclient.NewInputPolymorphicInputDataSourceAttributesRequest("Type_example", *openapiclient.NewInputDataSourceRequest("DataSource_example"))}, []openapiclient.TransformPolymorphicRequest{openapiclient.TransformPolymorphicRequest{TransformPolymorphicTransformFilterWorkloadsAttributesRequest: openapiclient.NewTransformPolymorphicTransformFilterWorkloadsAttributesRequest("Type_example", *openapiclient.NewTransformFilterWorkloadsRequest([]int64{int64(123)}))}}, []openapiclient.OutputRequest{*openapiclient.NewOutputRequest("Type_example", openapiclient.OutputPolymorphicRequest{AWSKinesisFirehoseEndpointRequest: openapiclient.NewAWSKinesisFirehoseEndpointRequest("AccessKey_example", "StreamName_example", "Region_example", "SecretKey_example")})}) // DataStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataStreamStreamsAPI.UpdateDataStream(context.Background(), id).DataStreamRequest(dataStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamStreamsAPI.UpdateDataStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataStream`: ResponseDataStream
	fmt.Fprintf(os.Stdout, "Response from `DataStreamStreamsAPI.UpdateDataStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataStreamRequest** | [**DataStreamRequest**](DataStreamRequest.md) |  | 

### Return type

[**ResponseDataStream**](ResponseDataStream.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

