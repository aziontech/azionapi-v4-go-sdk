# \DigitalCertificatesCertificateSigningRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateSigningRequestCSR**](DigitalCertificatesCertificateSigningRequestsAPI.md#CreateCertificateSigningRequestCSR) | **Post** /digital_certificates/csr | Create a certificate signing request (CSR)



## CreateCertificateSigningRequestCSR

> ResponseCertificate CreateCertificateSigningRequestCSR(ctx).CertificateSigningRequestRequest(certificateSigningRequestRequest).Execute()

Create a certificate signing request (CSR)



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
	certificateSigningRequestRequest := *openapiclient.NewCertificateSigningRequestRequest("Name_example", "CommonName_example", "Country_example", "State_example", "Locality_example", "Organization_example", "OrganizationUnity_example", "Email_example") // CertificateSigningRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificateSigningRequestsAPI.CreateCertificateSigningRequestCSR(context.Background()).CertificateSigningRequestRequest(certificateSigningRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateSigningRequestsAPI.CreateCertificateSigningRequestCSR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertificateSigningRequestCSR`: ResponseCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateSigningRequestsAPI.CreateCertificateSigningRequestCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateSigningRequestCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateSigningRequestRequest** | [**CertificateSigningRequestRequest**](CertificateSigningRequestRequest.md) |  | 

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

