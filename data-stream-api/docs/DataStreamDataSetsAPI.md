# \DataStreamDataSetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataSet**](DataStreamDataSetsAPI.md#CreateDataSet) | **Post** /data_stream/data_sets | Create a Data Set
[**DestroyDataSet**](DataStreamDataSetsAPI.md#DestroyDataSet) | **Delete** /data_stream/data_sets/{id} | Destroy a Data Set
[**ListDataSets**](DataStreamDataSetsAPI.md#ListDataSets) | **Get** /data_stream/data_sets | List Data Sets
[**PartialUpdateDataSet**](DataStreamDataSetsAPI.md#PartialUpdateDataSet) | **Patch** /data_stream/data_sets/{id} | Partially update a Data Set
[**RetrieveDataSet**](DataStreamDataSetsAPI.md#RetrieveDataSet) | **Get** /data_stream/data_sets/{id} | Retrieve details of a Data Set
[**UpdateDataSet**](DataStreamDataSetsAPI.md#UpdateDataSet) | **Put** /data_stream/data_sets/{id} | Update a Data Set



## CreateDataSet

> ResponseDataSet CreateDataSet(ctx).DataSetRequest(dataSetRequest).Execute()

Create a Data Set



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
	dataSetRequest := *openapiclient.NewDataSetRequest("Name_example", "DataSet_example") // DataSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataStreamDataSetsAPI.CreateDataSet(context.Background()).DataSetRequest(dataSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamDataSetsAPI.CreateDataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataSet`: ResponseDataSet
	fmt.Fprintf(os.Stdout, "Response from `DataStreamDataSetsAPI.CreateDataSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSetRequest** | [**DataSetRequest**](DataSetRequest.md) |  | 

### Return type

[**ResponseDataSet**](ResponseDataSet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDataSet

> ResponseDeleteDataSet DestroyDataSet(ctx, id).Execute()

Destroy a Data Set



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
	resp, r, err := apiClient.DataStreamDataSetsAPI.DestroyDataSet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamDataSetsAPI.DestroyDataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyDataSet`: ResponseDeleteDataSet
	fmt.Fprintf(os.Stdout, "Response from `DataStreamDataSetsAPI.DestroyDataSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteDataSet**](ResponseDeleteDataSet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSets

> PaginatedResponseListDataSetList ListDataSets(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Data Sets



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
	resp, r, err := apiClient.DataStreamDataSetsAPI.ListDataSets(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamDataSetsAPI.ListDataSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDataSets`: PaginatedResponseListDataSetList
	fmt.Fprintf(os.Stdout, "Response from `DataStreamDataSetsAPI.ListDataSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListDataSetList**](PaginatedResponseListDataSetList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDataSet

> ResponseDataSet PartialUpdateDataSet(ctx, id).PatchedDataSetRequest(patchedDataSetRequest).Execute()

Partially update a Data Set



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
	patchedDataSetRequest := *openapiclient.NewPatchedDataSetRequest() // PatchedDataSetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataStreamDataSetsAPI.PartialUpdateDataSet(context.Background(), id).PatchedDataSetRequest(patchedDataSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamDataSetsAPI.PartialUpdateDataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDataSet`: ResponseDataSet
	fmt.Fprintf(os.Stdout, "Response from `DataStreamDataSetsAPI.PartialUpdateDataSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDataSetRequest** | [**PatchedDataSetRequest**](PatchedDataSetRequest.md) |  | 

### Return type

[**ResponseDataSet**](ResponseDataSet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDataSet

> ResponseRetrieveDataSet RetrieveDataSet(ctx, id).Fields(fields).Execute()

Retrieve details of a Data Set



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
	resp, r, err := apiClient.DataStreamDataSetsAPI.RetrieveDataSet(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamDataSetsAPI.RetrieveDataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDataSet`: ResponseRetrieveDataSet
	fmt.Fprintf(os.Stdout, "Response from `DataStreamDataSetsAPI.RetrieveDataSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveDataSet**](ResponseRetrieveDataSet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataSet

> ResponseDataSet UpdateDataSet(ctx, id).DataSetRequest(dataSetRequest).Execute()

Update a Data Set



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
	dataSetRequest := *openapiclient.NewDataSetRequest("Name_example", "DataSet_example") // DataSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataStreamDataSetsAPI.UpdateDataSet(context.Background(), id).DataSetRequest(dataSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStreamDataSetsAPI.UpdateDataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataSet`: ResponseDataSet
	fmt.Fprintf(os.Stdout, "Response from `DataStreamDataSetsAPI.UpdateDataSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSetRequest** | [**DataSetRequest**](DataSetRequest.md) |  | 

### Return type

[**ResponseDataSet**](ResponseDataSet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

