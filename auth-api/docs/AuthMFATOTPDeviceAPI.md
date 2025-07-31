# \AuthMFATOTPDeviceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTotpDevice**](AuthMFATOTPDeviceAPI.md#CreateTotpDevice) | **Post** /auth/mfa/totp | Create a TOTP device
[**DestroyTotpDevice**](AuthMFATOTPDeviceAPI.md#DestroyTotpDevice) | **Delete** /auth/mfa/totp/{id} | Destroy a TOTP device
[**ListTotpDevices**](AuthMFATOTPDeviceAPI.md#ListTotpDevices) | **Get** /auth/mfa/totp | List of TOTP devices



## CreateTotpDevice

> ResponseTOTPDeviceCreate CreateTotpDevice(ctx).Execute()

Create a TOTP device



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthMFATOTPDeviceAPI.CreateTotpDevice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthMFATOTPDeviceAPI.CreateTotpDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTotpDevice`: ResponseTOTPDeviceCreate
	fmt.Fprintf(os.Stdout, "Response from `AuthMFATOTPDeviceAPI.CreateTotpDevice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTotpDeviceRequest struct via the builder pattern


### Return type

[**ResponseTOTPDeviceCreate**](ResponseTOTPDeviceCreate.md)

### Authorization

[JWT MFA Authentication](../README.md#JWT MFA Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyTotpDevice

> ResponseDeleteTOTPDeviceCreate DestroyTotpDevice(ctx, id).Execute()

Destroy a TOTP device



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
	resp, r, err := apiClient.AuthMFATOTPDeviceAPI.DestroyTotpDevice(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthMFATOTPDeviceAPI.DestroyTotpDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyTotpDevice`: ResponseDeleteTOTPDeviceCreate
	fmt.Fprintf(os.Stdout, "Response from `AuthMFATOTPDeviceAPI.DestroyTotpDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyTotpDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteTOTPDeviceCreate**](ResponseDeleteTOTPDeviceCreate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTotpDevices

> PaginatedTOTPDeviceListList ListTotpDevices(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List of TOTP devices



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, user_id, confirmed, id) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthMFATOTPDeviceAPI.ListTotpDevices(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthMFATOTPDeviceAPI.ListTotpDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTotpDevices`: PaginatedTOTPDeviceListList
	fmt.Fprintf(os.Stdout, "Response from `AuthMFATOTPDeviceAPI.ListTotpDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTotpDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, user_id, confirmed, id) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedTOTPDeviceListList**](PaginatedTOTPDeviceListList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

