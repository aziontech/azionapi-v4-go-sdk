# \ApplicationsRequestRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationApiApplicationsRequestRulesCreate**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesCreate) | **Post** /edge_application/applications/{application_id}/request_rules | Create an Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesDestroy**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesDestroy) | **Delete** /edge_application/applications/{application_id}/request_rules/{id} | Destroy an Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesList**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesList) | **Get** /edge_application/applications/{application_id}/request_rules | List Application Request Rules
[**EdgeApplicationApiApplicationsRequestRulesOrderUpdate**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesOrderUpdate) | **Put** /edge_application/applications/{application_id}/request_rules/order | Ordering Application Request Rules
[**EdgeApplicationApiApplicationsRequestRulesPartialUpdate**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesPartialUpdate) | **Patch** /edge_application/applications/{application_id}/request_rules/{id} | Partially update an Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesRetrieve**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesRetrieve) | **Get** /edge_application/applications/{application_id}/request_rules/{id} | Retrieve details of an Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesUpdate**](ApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesUpdate) | **Put** /edge_application/applications/{application_id}/request_rules/{id} | Update an Application Request Rule



## EdgeApplicationApiApplicationsRequestRulesCreate

> ResponseApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesCreate(ctx, applicationId).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()

Create an Application Request Rule



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
	applicationId := "applicationId_example" // string | 
	applicationRequestPhaseRuleEngineRequest := *openapiclient.NewApplicationRequestPhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest: openapiclient.NewApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest("Type_example", *openapiclient.NewApplicationRuleEngineAddHeaderAttributesRequest("Value_example"))}}) // ApplicationRequestPhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesCreate(context.Background(), applicationId).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesCreate`: ResponseApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationRequestPhaseRuleEngineRequest** | [**ApplicationRequestPhaseRuleEngineRequest**](ApplicationRequestPhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseApplicationRequestPhaseRuleEngine**](ResponseApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesDestroy

> ResponseDeleteApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesDestroy(ctx, applicationId, id).Execute()

Destroy an Application Request Rule



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
	applicationId := "applicationId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesDestroy(context.Background(), applicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesDestroy`: ResponseDeleteApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteApplicationRequestPhaseRuleEngine**](ResponseDeleteApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesList

> PaginatedApplicationRequestPhaseRuleEngineList EdgeApplicationApiApplicationsRequestRulesList(ctx, applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Application Request Rules



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
	applicationId := "applicationId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, active, description, order, criteria, last_editor, last_modified, behaviors) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesList(context.Background(), applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesList`: PaginatedApplicationRequestPhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, active, description, order, criteria, last_editor, last_modified, behaviors) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedApplicationRequestPhaseRuleEngineList**](PaginatedApplicationRequestPhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesOrderUpdate

> PaginatedApplicationRequestPhaseRuleEngineList EdgeApplicationApiApplicationsRequestRulesOrderUpdate(ctx, applicationId).ApplicationRequestPhaseRuleEngineOrderRequest(applicationRequestPhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

Ordering Application Request Rules



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
	applicationId := "applicationId_example" // string | 
	applicationRequestPhaseRuleEngineOrderRequest := *openapiclient.NewApplicationRequestPhaseRuleEngineOrderRequest([]int64{int64(123)}) // ApplicationRequestPhaseRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesOrderUpdate(context.Background(), applicationId).ApplicationRequestPhaseRuleEngineOrderRequest(applicationRequestPhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesOrderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesOrderUpdate`: PaginatedApplicationRequestPhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesOrderUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesOrderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationRequestPhaseRuleEngineOrderRequest** | [**ApplicationRequestPhaseRuleEngineOrderRequest**](ApplicationRequestPhaseRuleEngineOrderRequest.md) |  | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedApplicationRequestPhaseRuleEngineList**](PaginatedApplicationRequestPhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesPartialUpdate

> ResponseApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesPartialUpdate(ctx, applicationId, id).PatchedApplicationRequestPhaseRuleEngineRequest(patchedApplicationRequestPhaseRuleEngineRequest).Execute()

Partially update an Application Request Rule



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
	applicationId := "applicationId_example" // string | 
	id := "id_example" // string | 
	patchedApplicationRequestPhaseRuleEngineRequest := *openapiclient.NewPatchedApplicationRequestPhaseRuleEngineRequest() // PatchedApplicationRequestPhaseRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesPartialUpdate(context.Background(), applicationId, id).PatchedApplicationRequestPhaseRuleEngineRequest(patchedApplicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesPartialUpdate`: ResponseApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedApplicationRequestPhaseRuleEngineRequest** | [**PatchedApplicationRequestPhaseRuleEngineRequest**](PatchedApplicationRequestPhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseApplicationRequestPhaseRuleEngine**](ResponseApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesRetrieve

> ResponseRetrieveApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesRetrieve(ctx, applicationId, id).Fields(fields).Execute()

Retrieve details of an Application Request Rule



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
	applicationId := "applicationId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesRetrieve(context.Background(), applicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesRetrieve`: ResponseRetrieveApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveApplicationRequestPhaseRuleEngine**](ResponseRetrieveApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesUpdate

> ResponseApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesUpdate(ctx, applicationId, id).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()

Update an Application Request Rule



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
	applicationId := "applicationId_example" // string | 
	id := "id_example" // string | 
	applicationRequestPhaseRuleEngineRequest := *openapiclient.NewApplicationRequestPhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest: openapiclient.NewApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest("Type_example", *openapiclient.NewApplicationRuleEngineAddHeaderAttributesRequest("Value_example"))}}) // ApplicationRequestPhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesUpdate(context.Background(), applicationId, id).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesUpdate`: ResponseApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsRequestRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationRequestPhaseRuleEngineRequest** | [**ApplicationRequestPhaseRuleEngineRequest**](ApplicationRequestPhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseApplicationRequestPhaseRuleEngine**](ResponseApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

