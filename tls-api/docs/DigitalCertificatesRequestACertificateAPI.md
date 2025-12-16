# \DigitalCertificatesRequestACertificateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestCertificate**](DigitalCertificatesRequestACertificateAPI.md#RequestCertificate) | **Post** /workspace/tls/certificates/request | Request a certificate



## RequestCertificate

> ResponseCertificate RequestCertificate(ctx).CertificateRequest(certificateRequest).Execute()

Request a certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	certificateRequest := *openapiclient.NewCertificateRequest(int64(123), "Name_example", "Issuer_example", []string{"SubjectName_example"}, "Validity_example", false, "Status_example", "StatusDetail_example", "Csr_example", "Challenge_example", "Authority_example", "ProductVersion_example", "LastEditor_example", time.Now(), time.Now(), "CommonName_example") // CertificateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesRequestACertificateAPI.RequestCertificate(context.Background()).CertificateRequest(certificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesRequestACertificateAPI.RequestCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestCertificate`: ResponseCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesRequestACertificateAPI.RequestCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateRequest** | [**CertificateRequest**](CertificateRequest.md) |  | 

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

