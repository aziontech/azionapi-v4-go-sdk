# \AccountsLoggedInAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLoggedAccount**](AccountsLoggedInAPI.md#DeleteLoggedAccount) | **Delete** /account/account | Delete logged account
[**PartialUpdateLoggedAccount**](AccountsLoggedInAPI.md#PartialUpdateLoggedAccount) | **Patch** /account/account | Partially update logged account details
[**RetrieveLoggedAccountDetails**](AccountsLoggedInAPI.md#RetrieveLoggedAccountDetails) | **Get** /account/account | Retrieve logged account details
[**UpdateLoggedAccount**](AccountsLoggedInAPI.md#UpdateLoggedAccount) | **Put** /account/account | Update logged account details



## DeleteLoggedAccount

> ResponseDeleteAccount DeleteLoggedAccount(ctx).Execute()

Delete logged account



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsLoggedInAPI.DeleteLoggedAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsLoggedInAPI.DeleteLoggedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLoggedAccount`: ResponseDeleteAccount
	fmt.Fprintf(os.Stdout, "Response from `AccountsLoggedInAPI.DeleteLoggedAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoggedAccountRequest struct via the builder pattern


### Return type

[**ResponseDeleteAccount**](ResponseDeleteAccount.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateLoggedAccount

> ResponseAccount PartialUpdateLoggedAccount(ctx).PatchedAccountRequest(patchedAccountRequest).Execute()

Partially update logged account details



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
	patchedAccountRequest := openapiclient.PatchedAccountRequest{PatchedBrandRequest: openapiclient.NewPatchedBrandRequest()} // PatchedAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsLoggedInAPI.PartialUpdateLoggedAccount(context.Background()).PatchedAccountRequest(patchedAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsLoggedInAPI.PartialUpdateLoggedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateLoggedAccount`: ResponseAccount
	fmt.Fprintf(os.Stdout, "Response from `AccountsLoggedInAPI.PartialUpdateLoggedAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateLoggedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedAccountRequest** | [**PatchedAccountRequest**](PatchedAccountRequest.md) |  | 

### Return type

[**ResponseAccount**](ResponseAccount.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveLoggedAccountDetails

> ResponseRetrieveAccount RetrieveLoggedAccountDetails(ctx).Fields(fields).Execute()

Retrieve logged account details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsLoggedInAPI.RetrieveLoggedAccountDetails(context.Background()).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsLoggedInAPI.RetrieveLoggedAccountDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveLoggedAccountDetails`: ResponseRetrieveAccount
	fmt.Fprintf(os.Stdout, "Response from `AccountsLoggedInAPI.RetrieveLoggedAccountDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveLoggedAccountDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveAccount**](ResponseRetrieveAccount.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoggedAccount

> ResponseAccount UpdateLoggedAccount(ctx).AccountRequest(accountRequest).Execute()

Update logged account details



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
	accountRequest := openapiclient.AccountRequest{BrandRequest: openapiclient.NewBrandRequest("Name_example")} // AccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsLoggedInAPI.UpdateLoggedAccount(context.Background()).AccountRequest(accountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsLoggedInAPI.UpdateLoggedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoggedAccount`: ResponseAccount
	fmt.Fprintf(os.Stdout, "Response from `AccountsLoggedInAPI.UpdateLoggedAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoggedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountRequest** | [**AccountRequest**](AccountRequest.md) |  | 

### Return type

[**ResponseAccount**](ResponseAccount.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

