# \EdgeApplicationsCacheSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCacheSetting**](EdgeApplicationsCacheSettingsAPI.md#CreateCacheSetting) | **Post** /edge_application/applications/{edgeApplicationId}/cache_settings | Create an Edge Applications Cache Setting
[**DestroyCacheSetting**](EdgeApplicationsCacheSettingsAPI.md#DestroyCacheSetting) | **Delete** /edge_application/applications/{edgeApplicationId}/cache_settings/{id} | Destroy an Edge Applications Cache Setting
[**ListCacheSettings**](EdgeApplicationsCacheSettingsAPI.md#ListCacheSettings) | **Get** /edge_application/applications/{edgeApplicationId}/cache_settings | List all Edge Applications Cache Settings
[**PartialUpdateCacheSetting**](EdgeApplicationsCacheSettingsAPI.md#PartialUpdateCacheSetting) | **Patch** /edge_application/applications/{edgeApplicationId}/cache_settings/{id} | Partially update an Edge Applications Cache Setting
[**RetrieveCacheSetting**](EdgeApplicationsCacheSettingsAPI.md#RetrieveCacheSetting) | **Get** /edge_application/applications/{edgeApplicationId}/cache_settings/{id} | Retrieve details of an Edge Applications Cache Setting
[**UpdateCacheSetting**](EdgeApplicationsCacheSettingsAPI.md#UpdateCacheSetting) | **Put** /edge_application/applications/{edgeApplicationId}/cache_settings/{id} | Update an Edge Applications Cache Setting



## CreateCacheSetting

> ResponseCacheSetting CreateCacheSetting(ctx, edgeApplicationId).CacheSettingRequest(cacheSettingRequest).Execute()

Create an Edge Applications Cache Setting



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	cacheSettingRequest := *openapiclient.NewCacheSettingRequest("Name_example", *openapiclient.NewBrowserCacheModuleRequest(*openapiclient.NewBrowserCacheModuleBehaviorEnum(), int32(123)), *openapiclient.NewEdgeCacheModuleRequest(*openapiclient.NewEdgeCacheModuleBehaviorEnum(), int32(123), false, false, false, false), *openapiclient.NewApplicationControlsModuleRequest(*openapiclient.NewCacheByQueryStringEnum(), []string{"QueryStringFields_example"}, false, *openapiclient.NewCacheByCookiesEnum(), []string{"CookieNames_example"}, *openapiclient.NewAdaptiveDeliveryActionEnum(), []int32{int32(123)}), *openapiclient.NewSliceControlsModuleRequest(false, false, false)) // CacheSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.CreateCacheSetting(context.Background(), edgeApplicationId).CacheSettingRequest(cacheSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.CreateCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCacheSetting`: ResponseCacheSetting
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.CreateCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cacheSettingRequest** | [**CacheSettingRequest**](CacheSettingRequest.md) |  | 

### Return type

[**ResponseCacheSetting**](ResponseCacheSetting.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCacheSetting

> ResponseDeleteCacheSetting DestroyCacheSetting(ctx, edgeApplicationId, id).Execute()

Destroy an Edge Applications Cache Setting



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.DestroyCacheSetting(context.Background(), edgeApplicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.DestroyCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyCacheSetting`: ResponseDeleteCacheSetting
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.DestroyCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteCacheSetting**](ResponseDeleteCacheSetting.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCacheSettings

> PaginatedResponseListCacheSettingList ListCacheSettings(ctx, edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List all Edge Applications Cache Settings



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name) (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.ListCacheSettings(context.Background(), edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.ListCacheSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCacheSettings`: PaginatedResponseListCacheSettingList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.ListCacheSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCacheSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name) | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListCacheSettingList**](PaginatedResponseListCacheSettingList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateCacheSetting

> ResponseCacheSetting PartialUpdateCacheSetting(ctx, edgeApplicationId, id).PatchedCacheSettingRequest(patchedCacheSettingRequest).Execute()

Partially update an Edge Applications Cache Setting



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	patchedCacheSettingRequest := *openapiclient.NewPatchedCacheSettingRequest() // PatchedCacheSettingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.PartialUpdateCacheSetting(context.Background(), edgeApplicationId, id).PatchedCacheSettingRequest(patchedCacheSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.PartialUpdateCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateCacheSetting`: ResponseCacheSetting
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.PartialUpdateCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedCacheSettingRequest** | [**PatchedCacheSettingRequest**](PatchedCacheSettingRequest.md) |  | 

### Return type

[**ResponseCacheSetting**](ResponseCacheSetting.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCacheSetting

> ResponseRetrieveCacheSetting RetrieveCacheSetting(ctx, edgeApplicationId, id).Fields(fields).Execute()

Retrieve details of an Edge Applications Cache Setting



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.RetrieveCacheSetting(context.Background(), edgeApplicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.RetrieveCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCacheSetting`: ResponseRetrieveCacheSetting
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.RetrieveCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveCacheSetting**](ResponseRetrieveCacheSetting.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCacheSetting

> ResponseCacheSetting UpdateCacheSetting(ctx, edgeApplicationId, id).CacheSettingRequest(cacheSettingRequest).Execute()

Update an Edge Applications Cache Setting



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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	id := "id_example" // string | 
	cacheSettingRequest := *openapiclient.NewCacheSettingRequest("Name_example", *openapiclient.NewBrowserCacheModuleRequest(*openapiclient.NewBrowserCacheModuleBehaviorEnum(), int32(123)), *openapiclient.NewEdgeCacheModuleRequest(*openapiclient.NewEdgeCacheModuleBehaviorEnum(), int32(123), false, false, false, false), *openapiclient.NewApplicationControlsModuleRequest(*openapiclient.NewCacheByQueryStringEnum(), []string{"QueryStringFields_example"}, false, *openapiclient.NewCacheByCookiesEnum(), []string{"CookieNames_example"}, *openapiclient.NewAdaptiveDeliveryActionEnum(), []int32{int32(123)}), *openapiclient.NewSliceControlsModuleRequest(false, false, false)) // CacheSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.UpdateCacheSetting(context.Background(), edgeApplicationId, id).CacheSettingRequest(cacheSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.UpdateCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCacheSetting`: ResponseCacheSetting
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.UpdateCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cacheSettingRequest** | [**CacheSettingRequest**](CacheSettingRequest.md) |  | 

### Return type

[**ResponseCacheSetting**](ResponseCacheSetting.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

