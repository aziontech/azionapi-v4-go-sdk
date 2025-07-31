# \AuthRevokeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthUserRevoke**](AuthRevokeAPI.md#AuthUserRevoke) | **Post** /auth/revoke | Revoke user JWT refresh token



## AuthUserRevoke

> StateExecutedResponse AuthUserRevoke(ctx).Execute()

Revoke user JWT refresh token



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
	resp, r, err := apiClient.AuthRevokeAPI.AuthUserRevoke(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthRevokeAPI.AuthUserRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthUserRevoke`: StateExecutedResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthRevokeAPI.AuthUserRevoke`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthUserRevokeRequest struct via the builder pattern


### Return type

[**StateExecutedResponse**](StateExecutedResponse.md)

### Authorization

[JWT Refresh Authentication](../README.md#JWT Refresh Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

