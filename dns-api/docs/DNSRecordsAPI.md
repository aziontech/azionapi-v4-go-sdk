# \DNSRecordsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsRecord**](DNSRecordsAPI.md#CreateDnsRecord) | **Post** /edge_dns/zones/{zoneId}/records | Create a DNS Record
[**DestroyDnsRecord**](DNSRecordsAPI.md#DestroyDnsRecord) | **Delete** /edge_dns/zones/{zoneId}/records/{recordId} | Destroy a DNS Record
[**ListDnsRecords**](DNSRecordsAPI.md#ListDnsRecords) | **Get** /edge_dns/zones/{zoneId}/records | List DNS Records
[**PartialUpdateDnsRecord**](DNSRecordsAPI.md#PartialUpdateDnsRecord) | **Patch** /edge_dns/zones/{zoneId}/records/{recordId} | Partially update a DNS Record
[**RetrieveDnsRecord**](DNSRecordsAPI.md#RetrieveDnsRecord) | **Get** /edge_dns/zones/{zoneId}/records/{recordId} | Retrieve details of a DNS Record
[**UpdateDnsRecord**](DNSRecordsAPI.md#UpdateDnsRecord) | **Put** /edge_dns/zones/{zoneId}/records/{recordId} | Update a DNS Record



## CreateDnsRecord

> ResponseRecord CreateDnsRecord(ctx, zoneId).RecordRequest(recordRequest).Execute()

Create a DNS Record



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
	zoneId := "zoneId_example" // string | 
	recordRequest := *openapiclient.NewRecordRequest("Name_example", "Type_example", []string{"Rdata_example"}) // RecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordsAPI.CreateDnsRecord(context.Background(), zoneId).RecordRequest(recordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsAPI.CreateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDnsRecord`: ResponseRecord
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordsAPI.CreateDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordRequest** | [**RecordRequest**](RecordRequest.md) |  | 

### Return type

[**ResponseRecord**](ResponseRecord.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDnsRecord

> ResponseDeleteRecord DestroyDnsRecord(ctx, recordId, zoneId).Execute()

Destroy a DNS Record



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
	recordId := "recordId_example" // string | 
	zoneId := "zoneId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordsAPI.DestroyDnsRecord(context.Background(), recordId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsAPI.DestroyDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyDnsRecord`: ResponseDeleteRecord
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordsAPI.DestroyDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteRecord**](ResponseDeleteRecord.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDnsRecords

> PaginatedRecordList ListDnsRecords(ctx, zoneId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List DNS Records



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
	zoneId := "zoneId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, weight) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordsAPI.ListDnsRecords(context.Background(), zoneId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsAPI.ListDnsRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDnsRecords`: PaginatedRecordList
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordsAPI.ListDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, weight) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRecordList**](PaginatedRecordList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateDnsRecord

> ResponseRecord PartialUpdateDnsRecord(ctx, recordId, zoneId).PatchedRecordRequest(patchedRecordRequest).Execute()

Partially update a DNS Record



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
	recordId := "recordId_example" // string | 
	zoneId := "zoneId_example" // string | 
	patchedRecordRequest := *openapiclient.NewPatchedRecordRequest() // PatchedRecordRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordsAPI.PartialUpdateDnsRecord(context.Background(), recordId, zoneId).PatchedRecordRequest(patchedRecordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsAPI.PartialUpdateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateDnsRecord`: ResponseRecord
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordsAPI.PartialUpdateDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedRecordRequest** | [**PatchedRecordRequest**](PatchedRecordRequest.md) |  | 

### Return type

[**ResponseRecord**](ResponseRecord.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDnsRecord

> ResponseRetrieveRecord RetrieveDnsRecord(ctx, recordId, zoneId).Fields(fields).Execute()

Retrieve details of a DNS Record



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
	recordId := "recordId_example" // string | 
	zoneId := "zoneId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordsAPI.RetrieveDnsRecord(context.Background(), recordId, zoneId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsAPI.RetrieveDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDnsRecord`: ResponseRetrieveRecord
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordsAPI.RetrieveDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveRecord**](ResponseRetrieveRecord.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDnsRecord

> ResponseRecord UpdateDnsRecord(ctx, recordId, zoneId).RecordRequest(recordRequest).Execute()

Update a DNS Record



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
	recordId := "recordId_example" // string | 
	zoneId := "zoneId_example" // string | 
	recordRequest := *openapiclient.NewRecordRequest("Name_example", "Type_example", []string{"Rdata_example"}) // RecordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSRecordsAPI.UpdateDnsRecord(context.Background(), recordId, zoneId).RecordRequest(recordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsAPI.UpdateDnsRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDnsRecord`: ResponseRecord
	fmt.Fprintf(os.Stdout, "Response from `DNSRecordsAPI.UpdateDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordRequest** | [**RecordRequest**](RecordRequest.md) |  | 

### Return type

[**ResponseRecord**](ResponseRecord.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

