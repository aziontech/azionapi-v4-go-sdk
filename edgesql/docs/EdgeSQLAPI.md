# \EdgeSQLAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](EdgeSQLAPI.md#CreateDatabase) | **Post** /edge_sql/databases | Create a database
[**DestroyDatabase**](EdgeSQLAPI.md#DestroyDatabase) | **Delete** /edge_sql/databases/{id} | Destroy a database
[**ExecuteQuery**](EdgeSQLAPI.md#ExecuteQuery) | **Post** /edge_sql/databases/{id}/query | Execute a query into a database
[**ListDatabases**](EdgeSQLAPI.md#ListDatabases) | **Get** /edge_sql/databases | List databases
[**RetriveDatabase**](EdgeSQLAPI.md#RetriveDatabase) | **Get** /edge_sql/databases/{id} | Retrieve details from a database



## CreateDatabase

> Database CreateDatabase(ctx).DatabaseRequest(databaseRequest).Execute()

Create a database



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	databaseRequest := *openapiclient.NewDatabaseRequest("Name_example") // DatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeSQLAPI.CreateDatabase(context.Background()).DatabaseRequest(databaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSQLAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: Database
	fmt.Fprintf(os.Stdout, "Response from `EdgeSQLAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseRequest** | [**DatabaseRequest**](DatabaseRequest.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDatabase

> ResponseDeleteDatabase DestroyDatabase(ctx, id).Execute()

Destroy a database



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeSQLAPI.DestroyDatabase(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSQLAPI.DestroyDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyDatabase`: ResponseDeleteDatabase
	fmt.Fprintf(os.Stdout, "Response from `EdgeSQLAPI.DestroyDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteDatabase**](ResponseDeleteDatabase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteQuery

> Database ExecuteQuery(ctx, id).SQLStatementsRequest(sQLStatementsRequest).Execute()

Execute a query into a database



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	id := int32(56) // int32 | 
	sQLStatementsRequest := *openapiclient.NewSQLStatementsRequest([]string{"Statements_example"}) // SQLStatementsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeSQLAPI.ExecuteQuery(context.Background(), id).SQLStatementsRequest(sQLStatementsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSQLAPI.ExecuteQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteQuery`: Database
	fmt.Fprintf(os.Stdout, "Response from `EdgeSQLAPI.ExecuteQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sQLStatementsRequest** | [**SQLStatementsRequest**](SQLStatementsRequest.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> PaginatedResponseListDatabaseList ListDatabases(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List databases



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeSQLAPI.ListDatabases(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSQLAPI.ListDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDatabases`: PaginatedResponseListDatabaseList
	fmt.Fprintf(os.Stdout, "Response from `EdgeSQLAPI.ListDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListDatabaseList**](PaginatedResponseListDatabaseList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetriveDatabase

> ResponseRetrieveDatabase RetriveDatabase(ctx, id).Fields(fields).Execute()

Retrieve details from a database



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aziontech/azionapi-v4-go-sdk"
)

func main() {
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeSQLAPI.RetriveDatabase(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSQLAPI.RetriveDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetriveDatabase`: ResponseRetrieveDatabase
	fmt.Fprintf(os.Stdout, "Response from `EdgeSQLAPI.RetriveDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetriveDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveDatabase**](ResponseRetrieveDatabase.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

