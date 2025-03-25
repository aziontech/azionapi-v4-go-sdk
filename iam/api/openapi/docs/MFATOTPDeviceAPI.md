# \MFATOTPDeviceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTotpDevice**](MFATOTPDeviceAPI.md#CreateTotpDevice) | **Post** /iam/mfa/totp | Create a TOTP device



## CreateTotpDevice

> ResponseTOTPDevice CreateTotpDevice(ctx).Execute()

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
	resp, r, err := apiClient.MFATOTPDeviceAPI.CreateTotpDevice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFATOTPDeviceAPI.CreateTotpDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTotpDevice`: ResponseTOTPDevice
	fmt.Fprintf(os.Stdout, "Response from `MFATOTPDeviceAPI.CreateTotpDevice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTotpDeviceRequest struct via the builder pattern


### Return type

[**ResponseTOTPDevice**](ResponseTOTPDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

