# \DigitalCertificatesCertificatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificate**](DigitalCertificatesCertificatesAPI.md#CreateCertificate) | **Post** /digital_certificates/certificates | Create a certificate
[**DestroyCertificate**](DigitalCertificatesCertificatesAPI.md#DestroyCertificate) | **Delete** /digital_certificates/certificates/{id} | Destroy a certificate
[**ListCertificates**](DigitalCertificatesCertificatesAPI.md#ListCertificates) | **Get** /digital_certificates/certificates | List certificates
[**PartialUpdateCertificate**](DigitalCertificatesCertificatesAPI.md#PartialUpdateCertificate) | **Patch** /digital_certificates/certificates/{id} | Partially update a certificate
[**RetriveCertificate**](DigitalCertificatesCertificatesAPI.md#RetriveCertificate) | **Get** /digital_certificates/certificates/{id} | Retrieve details from a certificate
[**UpdateCertificate**](DigitalCertificatesCertificatesAPI.md#UpdateCertificate) | **Put** /digital_certificates/certificates/{id} | Update a certificate



## CreateCertificate

> ResponseCertificate CreateCertificate(ctx).CertificateRequest(certificateRequest).Execute()

Create a certificate



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
	certificateRequest := *openapiclient.NewCertificateRequest("Name_example") // CertificateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificatesAPI.CreateCertificate(context.Background()).CertificateRequest(certificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificatesAPI.CreateCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertificate`: ResponseCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificatesAPI.CreateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateRequest struct via the builder pattern


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


## DestroyCertificate

> ResponseDeleteCertificate DestroyCertificate(ctx, id).Execute()

Destroy a certificate



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
	resp, r, err := apiClient.DigitalCertificatesCertificatesAPI.DestroyCertificate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificatesAPI.DestroyCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyCertificate`: ResponseDeleteCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificatesAPI.DestroyCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteCertificate**](ResponseDeleteCertificate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificates

> PaginatedCertificateList ListCertificates(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List certificates



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, certificate, issuer, validity, subject_name, type, managed, status, status_detail, csr, key_algorithm, challenge, authority, active, product_version, last_editor, last_modified, renewed_at) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificatesAPI.ListCertificates(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificatesAPI.ListCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertificates`: PaginatedCertificateList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificatesAPI.ListCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, certificate, issuer, validity, subject_name, type, managed, status, status_detail, csr, key_algorithm, challenge, authority, active, product_version, last_editor, last_modified, renewed_at) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedCertificateList**](PaginatedCertificateList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateCertificate

> ResponseAsyncCertificate PartialUpdateCertificate(ctx, id).PatchedCertificateRequest(patchedCertificateRequest).Execute()

Partially update a certificate



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
	patchedCertificateRequest := *openapiclient.NewPatchedCertificateRequest() // PatchedCertificateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificatesAPI.PartialUpdateCertificate(context.Background(), id).PatchedCertificateRequest(patchedCertificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificatesAPI.PartialUpdateCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateCertificate`: ResponseAsyncCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificatesAPI.PartialUpdateCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCertificateRequest** | [**PatchedCertificateRequest**](PatchedCertificateRequest.md) |  | 

### Return type

[**ResponseAsyncCertificate**](ResponseAsyncCertificate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetriveCertificate

> ResponseRetrieveCertificate RetriveCertificate(ctx, id).Fields(fields).Execute()

Retrieve details from a certificate



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
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificatesAPI.RetriveCertificate(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificatesAPI.RetriveCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetriveCertificate`: ResponseRetrieveCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificatesAPI.RetriveCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetriveCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveCertificate**](ResponseRetrieveCertificate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> ResponseAsyncCertificate UpdateCertificate(ctx, id).CertificateRequest(certificateRequest).Execute()

Update a certificate



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
	certificateRequest := *openapiclient.NewCertificateRequest("Name_example") // CertificateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificatesAPI.UpdateCertificate(context.Background(), id).CertificateRequest(certificateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificatesAPI.UpdateCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCertificate`: ResponseAsyncCertificate
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificatesAPI.UpdateCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateRequest** | [**CertificateRequest**](CertificateRequest.md) |  | 

### Return type

[**ResponseAsyncCertificate**](ResponseAsyncCertificate.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

