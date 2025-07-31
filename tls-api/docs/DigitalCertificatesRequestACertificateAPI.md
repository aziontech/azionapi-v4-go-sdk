# \DigitalCertificatesRequestACertificateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestACertificate**](DigitalCertificatesRequestACertificateAPI.md#RequestACertificate) | **Post** /digital_certificates/certificates/request | Request a certificate



## RequestACertificate

> ResponseCertificate RequestACertificate(ctx).CertificateRequestRequest(certificateRequestRequest).Execute()

Request a certificate



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
	certificateRequestRequest := *openapiclient.NewCertificateRequestRequest("Name_example", "Challenge_example", "Authority_example", "CommonName_example") // CertificateRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesRequestACertificateAPI.RequestACertificate(context.Background()).CertificateRequestRequest(certificateRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesRequestACertificateAPI.RequestACertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestACertificate`: ResponseCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesRequestACertificateAPI.RequestACertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestACertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateRequestRequest** | [**CertificateRequestRequest**](CertificateRequestRequest.md) |  | 

### Return type

[**ResponseCertificate**](ResponseCertificate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

