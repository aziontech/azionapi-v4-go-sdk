# \DigitalCertificatesCertificateRevocationListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateRevocationListsCRL**](DigitalCertificatesCertificateRevocationListsAPI.md#CreateCertificateRevocationListsCRL) | **Post** /digital_certificates/crls | Create a certificate revocation lists (CRL)
[**DestroyCertificateRevocationListsCRL**](DigitalCertificatesCertificateRevocationListsAPI.md#DestroyCertificateRevocationListsCRL) | **Delete** /digital_certificates/crls/{id} | Destroy a certificate revocation lists (CRL)
[**ListCertificateRevocationListsCRL**](DigitalCertificatesCertificateRevocationListsAPI.md#ListCertificateRevocationListsCRL) | **Get** /digital_certificates/crls | List certificate revocation lists (CRL)
[**PartialUpdateCertificateRevocationListsCRL**](DigitalCertificatesCertificateRevocationListsAPI.md#PartialUpdateCertificateRevocationListsCRL) | **Patch** /digital_certificates/crls/{id} | Update a certificate revocation lists (CRL)
[**RetriveCertificateRevocationListsCRL**](DigitalCertificatesCertificateRevocationListsAPI.md#RetriveCertificateRevocationListsCRL) | **Get** /digital_certificates/crls/{id} | Retrieve details from a certificate revocation lists (CRL)
[**UpdateCertificateRevocationListsCRL**](DigitalCertificatesCertificateRevocationListsAPI.md#UpdateCertificateRevocationListsCRL) | **Put** /digital_certificates/crls/{id} | Update a certificate revocation lists (CRL)



## CreateCertificateRevocationListsCRL

> ResponseCertificateRevocationList CreateCertificateRevocationListsCRL(ctx).CertificateRevocationListRequest(certificateRevocationListRequest).Execute()

Create a certificate revocation lists (CRL)



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
	certificateRevocationListRequest := *openapiclient.NewCertificateRevocationListRequest("Name_example", "Crl_example") // CertificateRevocationListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificateRevocationListsAPI.CreateCertificateRevocationListsCRL(context.Background()).CertificateRevocationListRequest(certificateRevocationListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateRevocationListsAPI.CreateCertificateRevocationListsCRL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertificateRevocationListsCRL`: ResponseCertificateRevocationList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateRevocationListsAPI.CreateCertificateRevocationListsCRL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateRevocationListsCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateRevocationListRequest** | [**CertificateRevocationListRequest**](CertificateRevocationListRequest.md) |  | 

### Return type

[**ResponseCertificateRevocationList**](ResponseCertificateRevocationList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCertificateRevocationListsCRL

> ResponseDeleteCertificateRevocationList DestroyCertificateRevocationListsCRL(ctx, id).Execute()

Destroy a certificate revocation lists (CRL)



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
	resp, r, err := apiClient.DigitalCertificatesCertificateRevocationListsAPI.DestroyCertificateRevocationListsCRL(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateRevocationListsAPI.DestroyCertificateRevocationListsCRL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyCertificateRevocationListsCRL`: ResponseDeleteCertificateRevocationList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateRevocationListsAPI.DestroyCertificateRevocationListsCRL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyCertificateRevocationListsCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteCertificateRevocationList**](ResponseDeleteCertificateRevocationList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificateRevocationListsCRL

> PaginatedCertificateRevocationListList ListCertificateRevocationListsCRL(ctx).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List certificate revocation lists (CRL)



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificateRevocationListsAPI.ListCertificateRevocationListsCRL(context.Background()).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateRevocationListsAPI.ListCertificateRevocationListsCRL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertificateRevocationListsCRL`: PaginatedCertificateRevocationListList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateRevocationListsAPI.ListCertificateRevocationListsCRL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertificateRevocationListsCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedCertificateRevocationListList**](PaginatedCertificateRevocationListList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateCertificateRevocationListsCRL

> ResponseCertificateRevocationList PartialUpdateCertificateRevocationListsCRL(ctx, id).PatchedCertificateRevocationListRequest(patchedCertificateRevocationListRequest).Execute()

Update a certificate revocation lists (CRL)



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
	patchedCertificateRevocationListRequest := *openapiclient.NewPatchedCertificateRevocationListRequest() // PatchedCertificateRevocationListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificateRevocationListsAPI.PartialUpdateCertificateRevocationListsCRL(context.Background(), id).PatchedCertificateRevocationListRequest(patchedCertificateRevocationListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateRevocationListsAPI.PartialUpdateCertificateRevocationListsCRL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateCertificateRevocationListsCRL`: ResponseCertificateRevocationList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateRevocationListsAPI.PartialUpdateCertificateRevocationListsCRL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCertificateRevocationListsCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCertificateRevocationListRequest** | [**PatchedCertificateRevocationListRequest**](PatchedCertificateRevocationListRequest.md) |  | 

### Return type

[**ResponseCertificateRevocationList**](ResponseCertificateRevocationList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetriveCertificateRevocationListsCRL

> ResponseRetrieveCertificateRevocationList RetriveCertificateRevocationListsCRL(ctx, id).Fields(fields).Execute()

Retrieve details from a certificate revocation lists (CRL)



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
	resp, r, err := apiClient.DigitalCertificatesCertificateRevocationListsAPI.RetriveCertificateRevocationListsCRL(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateRevocationListsAPI.RetriveCertificateRevocationListsCRL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetriveCertificateRevocationListsCRL`: ResponseRetrieveCertificateRevocationList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateRevocationListsAPI.RetriveCertificateRevocationListsCRL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetriveCertificateRevocationListsCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveCertificateRevocationList**](ResponseRetrieveCertificateRevocationList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateRevocationListsCRL

> ResponseCertificateRevocationList UpdateCertificateRevocationListsCRL(ctx, id).CertificateRevocationListRequest(certificateRevocationListRequest).Execute()

Update a certificate revocation lists (CRL)



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
	certificateRevocationListRequest := *openapiclient.NewCertificateRevocationListRequest("Name_example", "Crl_example") // CertificateRevocationListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalCertificatesCertificateRevocationListsAPI.UpdateCertificateRevocationListsCRL(context.Background(), id).CertificateRevocationListRequest(certificateRevocationListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalCertificatesCertificateRevocationListsAPI.UpdateCertificateRevocationListsCRL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCertificateRevocationListsCRL`: ResponseCertificateRevocationList
	fmt.Fprintf(os.Stdout, "Response from `DigitalCertificatesCertificateRevocationListsAPI.UpdateCertificateRevocationListsCRL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateRevocationListsCRLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateRevocationListRequest** | [**CertificateRevocationListRequest**](CertificateRevocationListRequest.md) |  | 

### Return type

[**ResponseCertificateRevocationList**](ResponseCertificateRevocationList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

