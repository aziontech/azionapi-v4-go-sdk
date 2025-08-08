# \EdgeApplicationsResponseRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationApiApplicationsResponseRulesCreate**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesCreate) | **Post** /edge_application/applications/{application_id}/response_rules | Create an Edge Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesDestroy**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesDestroy) | **Delete** /edge_application/applications/{application_id}/response_rules/{id} | Destroy an Edge Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesList**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesList) | **Get** /edge_application/applications/{application_id}/response_rules | List Edge Application Response Rules
[**EdgeApplicationApiApplicationsResponseRulesOrderUpdate**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesOrderUpdate) | **Put** /edge_application/applications/{application_id}/response_rules/order | Ordering Edge Application Response Rules
[**EdgeApplicationApiApplicationsResponseRulesPartialUpdate**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesPartialUpdate) | **Patch** /edge_application/applications/{application_id}/response_rules/{id} | Partially update an Edge Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesRetrieve**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesRetrieve) | **Get** /edge_application/applications/{application_id}/response_rules/{id} | Retrieve details of an Edge Application Response Rule
[**EdgeApplicationApiApplicationsResponseRulesUpdate**](EdgeApplicationsResponseRulesAPI.md#EdgeApplicationApiApplicationsResponseRulesUpdate) | **Put** /edge_application/applications/{application_id}/response_rules/{id} | Update an Edge Application Response Rule



## EdgeApplicationApiApplicationsResponseRulesCreate

> ResponseEdgeApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesCreate(ctx, applicationId).EdgeApplicationResponsePhaseRuleEngineRequest(edgeApplicationResponsePhaseRuleEngineRequest).Execute()

Create an Edge Application Response Rule



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
	edgeApplicationResponsePhaseRuleEngineRequest := *openapiclient.NewEdgeApplicationResponsePhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest{openapiclient.EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest{EdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest: openapiclient.NewEdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest("Type_example", *openapiclient.NewEdgeApplicationRuleEngineCaptureMatchGroupsAttributes("Subject_example", "Regex_example", "CapturedArray_example"))}}) // EdgeApplicationResponsePhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesCreate(context.Background(), applicationId).EdgeApplicationResponsePhaseRuleEngineRequest(edgeApplicationResponsePhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesCreate`: ResponseEdgeApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesCreate`: %v\n", resp)
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

 **edgeApplicationResponsePhaseRuleEngineRequest** | [**EdgeApplicationResponsePhaseRuleEngineRequest**](EdgeApplicationResponsePhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationResponsePhaseRuleEngine**](ResponseEdgeApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesDestroy

> ResponseDeleteEdgeApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesDestroy(ctx, applicationId, id).Execute()

Destroy an Edge Application Response Rule



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
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesDestroy(context.Background(), applicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesDestroy`: ResponseDeleteEdgeApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesDestroy`: %v\n", resp)
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

[**ResponseDeleteEdgeApplicationResponsePhaseRuleEngine**](ResponseDeleteEdgeApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesList

> PaginatedEdgeApplicationResponsePhaseRuleEngineList EdgeApplicationApiApplicationsResponseRulesList(ctx, applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Application Response Rules



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
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesList(context.Background(), applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesList`: PaginatedEdgeApplicationResponsePhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesList`: %v\n", resp)
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

[**PaginatedEdgeApplicationResponsePhaseRuleEngineList**](PaginatedEdgeApplicationResponsePhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesOrderUpdate

> PaginatedEdgeApplicationResponsePhaseRuleEngineList EdgeApplicationApiApplicationsResponseRulesOrderUpdate(ctx, applicationId).EdgeApplicationResponsePhaseRuleEngineOrderRequest(edgeApplicationResponsePhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

Ordering Edge Application Response Rules



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
	edgeApplicationResponsePhaseRuleEngineOrderRequest := *openapiclient.NewEdgeApplicationResponsePhaseRuleEngineOrderRequest([]int64{int64(123)}) // EdgeApplicationResponsePhaseRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesOrderUpdate(context.Background(), applicationId).EdgeApplicationResponsePhaseRuleEngineOrderRequest(edgeApplicationResponsePhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesOrderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesOrderUpdate`: PaginatedEdgeApplicationResponsePhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesOrderUpdate`: %v\n", resp)
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

 **edgeApplicationResponsePhaseRuleEngineOrderRequest** | [**EdgeApplicationResponsePhaseRuleEngineOrderRequest**](EdgeApplicationResponsePhaseRuleEngineOrderRequest.md) |  | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeApplicationResponsePhaseRuleEngineList**](PaginatedEdgeApplicationResponsePhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesPartialUpdate

> ResponseEdgeApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesPartialUpdate(ctx, applicationId, id).PatchedEdgeApplicationResponsePhaseRuleEngineRequest(patchedEdgeApplicationResponsePhaseRuleEngineRequest).Execute()

Partially update an Edge Application Response Rule



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
	patchedEdgeApplicationResponsePhaseRuleEngineRequest := *openapiclient.NewPatchedEdgeApplicationResponsePhaseRuleEngineRequest() // PatchedEdgeApplicationResponsePhaseRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesPartialUpdate(context.Background(), applicationId, id).PatchedEdgeApplicationResponsePhaseRuleEngineRequest(patchedEdgeApplicationResponsePhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesPartialUpdate`: ResponseEdgeApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesPartialUpdate`: %v\n", resp)
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


 **patchedEdgeApplicationResponsePhaseRuleEngineRequest** | [**PatchedEdgeApplicationResponsePhaseRuleEngineRequest**](PatchedEdgeApplicationResponsePhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationResponsePhaseRuleEngine**](ResponseEdgeApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesRetrieve

> ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesRetrieve(ctx, applicationId, id).Fields(fields).Execute()

Retrieve details of an Edge Application Response Rule



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
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesRetrieve(context.Background(), applicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesRetrieve`: ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesRetrieve`: %v\n", resp)
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

[**ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine**](ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsResponseRulesUpdate

> ResponseEdgeApplicationResponsePhaseRuleEngine EdgeApplicationApiApplicationsResponseRulesUpdate(ctx, applicationId, id).EdgeApplicationResponsePhaseRuleEngineRequest(edgeApplicationResponsePhaseRuleEngineRequest).Execute()

Update an Edge Application Response Rule



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
	edgeApplicationResponsePhaseRuleEngineRequest := *openapiclient.NewEdgeApplicationResponsePhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest{openapiclient.EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest{EdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest: openapiclient.NewEdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest("Type_example", *openapiclient.NewEdgeApplicationRuleEngineCaptureMatchGroupsAttributes("Subject_example", "Regex_example", "CapturedArray_example"))}}) // EdgeApplicationResponsePhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesUpdate(context.Background(), applicationId, id).EdgeApplicationResponsePhaseRuleEngineRequest(edgeApplicationResponsePhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsResponseRulesUpdate`: ResponseEdgeApplicationResponsePhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsResponseRulesAPI.EdgeApplicationApiApplicationsResponseRulesUpdate`: %v\n", resp)
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


 **edgeApplicationResponsePhaseRuleEngineRequest** | [**EdgeApplicationResponsePhaseRuleEngineRequest**](EdgeApplicationResponsePhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationResponsePhaseRuleEngine**](ResponseEdgeApplicationResponsePhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

