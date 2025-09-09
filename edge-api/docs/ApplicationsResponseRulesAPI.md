# \ApplicationsResponseRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationApiApplicationsResponseRulesCreate**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesCreate) | **Post** /edge_application/applications/{application_id}/response_rules | Create an Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesDestroy**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesDestroy) | **Delete** /edge_application/applications/{application_id}/response_rules/{id} | Destroy an Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesList**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesList) | **Get** /edge_application/applications/{application_id}/response_rules | List Application Response Rules
[**EdgeApplicationApiApplicationsResponseRulesOrderUpdate**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesOrderUpdate) | **Put** /edge_application/applications/{application_id}/response_rules/order | Ordering Application Response Rules
[**EdgeApplicationApiApplicationsResponseRulesPartialUpdate**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesPartialUpdate) | **Patch** /edge_application/applications/{application_id}/response_rules/{id} | Partially update an Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesRetrieve**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesRetrieve) | **Get** /edge_application/applications/{application_id}/response_rules/{id} | Retrieve details of an Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesUpdate**](ApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesUpdate) | **Put** /edge_application/applications/{application_id}/response_rules/{id} | Update an Application Response Rule



## EdgeApplicationApiApplicationsResponseRulesCreate

> ResponseApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesCreate(ctx, applicationId).ApplicationResponsePhaseRuleEngineRequest(applicationResponsePhaseRuleEngineRequest).Execute()

Create an Application Response Rule



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
	applicationResponsePhaseRuleEngineRequest := *openapiclient.NewApplicationResponsePhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ApplicationRuleEngineResponsePhaseBehaviorsRequest{openapiclient.ApplicationRuleEngineResponsePhaseBehaviorsRequest{ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineAddResponseHeaderRequest: openapiclient.NewApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineAddResponseHeaderRequest("Type_example", *openapiclient.NewApplicationRuleEngineAddResponseHeaderAttributesRequest("Value_example"))}}) // ApplicationResponsePhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesCreate(context.Background(), applicationId).ApplicationResponsePhaseRuleEngineRequest(applicationResponsePhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesCreate`: ResponseApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationResponsePhaseRuleEngineRequest** | [**ApplicationResponsePhaseRuleEngineRequest**](ApplicationResponsePhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseApplicationResponsePhaseRuleEngine**](ResponseApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesDestroy

> ResponseDeleteApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesDestroy(ctx, applicationId, id).Execute()

Destroy an Application Response Rule



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
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesDestroy(context.Background(), applicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesDestroy`: ResponseDeleteApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteApplicationResponsePhaseRuleEngine**](ResponseDeleteApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesList

> PaginatedApplicationResponsePhaseRuleEngineList EdgeApplicationApiApplicationsResponseRulesList(ctx, applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Application Response Rules



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
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesList(context.Background(), applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesList`: PaginatedApplicationResponsePhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, active, description, order, criteria, last_editor, last_modified, behaviors) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedApplicationResponsePhaseRuleEngineList**](PaginatedApplicationResponsePhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesOrderUpdate

> PaginatedApplicationResponsePhaseRuleEngineList EdgeApplicationApiApplicationsResponseRulesOrderUpdate(ctx, applicationId).ApplicationResponsePhaseRuleEngineOrderRequest(applicationResponsePhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

Ordering Application Response Rules



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
	applicationResponsePhaseRuleEngineOrderRequest := *openapiclient.NewApplicationResponsePhaseRuleEngineOrderRequest([]int64{int64(123)}) // ApplicationResponsePhaseRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesOrderUpdate(context.Background(), applicationId).ApplicationResponsePhaseRuleEngineOrderRequest(applicationResponsePhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesOrderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesOrderUpdate`: PaginatedApplicationResponsePhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesOrderUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesOrderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationResponsePhaseRuleEngineOrderRequest** | [**ApplicationResponsePhaseRuleEngineOrderRequest**](ApplicationResponsePhaseRuleEngineOrderRequest.md) |  | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedApplicationResponsePhaseRuleEngineList**](PaginatedApplicationResponsePhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesPartialUpdate

> ResponseApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesPartialUpdate(ctx, applicationId, id).PatchedApplicationResponsePhaseRuleEngineRequest(patchedApplicationResponsePhaseRuleEngineRequest).Execute()

Partially update an Application Response Rule



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
	patchedApplicationResponsePhaseRuleEngineRequest := *openapiclient.NewPatchedApplicationResponsePhaseRuleEngineRequest() // PatchedApplicationResponsePhaseRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesPartialUpdate(context.Background(), applicationId, id).PatchedApplicationResponsePhaseRuleEngineRequest(patchedApplicationResponsePhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesPartialUpdate`: ResponseApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedApplicationResponsePhaseRuleEngineRequest** | [**PatchedApplicationResponsePhaseRuleEngineRequest**](PatchedApplicationResponsePhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseApplicationResponsePhaseRuleEngine**](ResponseApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesRetrieve

> ResponseRetrieveApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesRetrieve(ctx, applicationId, id).Fields(fields).Execute()

Retrieve details of an Application Response Rule



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
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesRetrieve(context.Background(), applicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesRetrieve`: ResponseRetrieveApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesRetrieveRequest struct via the builder pattern


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


## EdgeApplicationApiApplicationsResponseRulesUpdate

> ResponseApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesUpdate(ctx, applicationId, id).ApplicationResponsePhaseRuleEngineRequest(applicationResponsePhaseRuleEngineRequest).Execute()

Update an Application Response Rule



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
	applicationResponsePhaseRuleEngineRequest := *openapiclient.NewApplicationResponsePhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ApplicationRuleEngineResponsePhaseBehaviorsRequest{openapiclient.ApplicationRuleEngineResponsePhaseBehaviorsRequest{ApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineAddResponseHeaderRequest: openapiclient.NewApplicationRuleEngineResponsePhaseBehaviorsApplicationRuleEngineAddResponseHeaderRequest("Type_example", *openapiclient.NewApplicationRuleEngineAddResponseHeaderAttributesRequest("Value_example"))}}) // ApplicationResponsePhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesUpdate(context.Background(), applicationId, id).ApplicationResponsePhaseRuleEngineRequest(applicationResponsePhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesUpdate`: ResponseApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationApiApplicationsResponseRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationResponsePhaseRuleEngineRequest** | [**ApplicationResponsePhaseRuleEngineRequest**](ApplicationResponsePhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseApplicationResponsePhaseRuleEngine**](ResponseApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

