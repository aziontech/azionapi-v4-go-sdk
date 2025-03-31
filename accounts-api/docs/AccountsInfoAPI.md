# \AccountsInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAccountInfoDetails**](AccountsInfoAPI.md#RetrieveAccountInfoDetails) | **Get** /account/accounts/{id}/info | Retrieve account information details
[**UpdateAccountInfoDetails**](AccountsInfoAPI.md#UpdateAccountInfoDetails) | **Put** /account/accounts/{id}/info | Update account information details



## RetrieveAccountInfoDetails

> ResponseRetrieveAccountInfo RetrieveAccountInfoDetails(ctx, id).Fields(fields).Execute()

Retrieve account information details



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
	id := int64(789) // int64 | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsInfoAPI.RetrieveAccountInfoDetails(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsInfoAPI.RetrieveAccountInfoDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAccountInfoDetails`: ResponseRetrieveAccountInfo
	fmt.Fprintf(os.Stdout, "Response from `AccountsInfoAPI.RetrieveAccountInfoDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAccountInfoDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveAccountInfo**](ResponseRetrieveAccountInfo.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountInfoDetails

> ResponseAccountInfo UpdateAccountInfoDetails(ctx, id).UpdateAccountInfoDetailsRequest(updateAccountInfoDetailsRequest).Execute()

Update account information details



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
	id := int64(789) // int64 | 
	updateAccountInfoDetailsRequest := *openapiclient.NewUpdateAccountInfoDetailsRequest() // UpdateAccountInfoDetailsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsInfoAPI.UpdateAccountInfoDetails(context.Background(), id).UpdateAccountInfoDetailsRequest(updateAccountInfoDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsInfoAPI.UpdateAccountInfoDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountInfoDetails`: ResponseAccountInfo
	fmt.Fprintf(os.Stdout, "Response from `AccountsInfoAPI.UpdateAccountInfoDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountInfoDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAccountInfoDetailsRequest** | [**UpdateAccountInfoDetailsRequest**](UpdateAccountInfoDetailsRequest.md) |  | 

### Return type

[**ResponseAccountInfo**](ResponseAccountInfo.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

