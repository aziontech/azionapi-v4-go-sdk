# \AuthRefreshAccessTokenAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthUserRefreshToken**](AuthRefreshAccessTokenAPI.md#AuthUserRefreshToken) | **Post** /auth/token | Refresh user JWT access token



## AuthUserRefreshToken

> ResponseToken AuthUserRefreshToken(ctx).TokenRequest(tokenRequest).Execute()

Refresh user JWT access token



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
	tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthRefreshAccessTokenAPI.AuthUserRefreshToken(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthRefreshAccessTokenAPI.AuthUserRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthUserRefreshToken`: ResponseToken
	fmt.Fprintf(os.Stdout, "Response from `AuthRefreshAccessTokenAPI.AuthUserRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthUserRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**ResponseToken**](ResponseToken.md)

### Authorization

[JWT Refresh Authentication](../README.md#JWT Refresh Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

