# \EdgeApplicationsRequestRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationApiApplicationsRequestRulesCreate**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesCreate) | **Post** /edge_application/applications/{application_id}/request_rules | Create an Edge Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesDestroy**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesDestroy) | **Delete** /edge_application/applications/{application_id}/request_rules/{id} | Destroy an Edge Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesList**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesList) | **Get** /edge_application/applications/{application_id}/request_rules | List Edge Application Request Rules
[**EdgeApplicationApiApplicationsRequestRulesOrderUpdate**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesOrderUpdate) | **Put** /edge_application/applications/{application_id}/request_rules/order | Ordering Edge Application Request Rules
[**EdgeApplicationApiApplicationsRequestRulesPartialUpdate**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesPartialUpdate) | **Patch** /edge_application/applications/{application_id}/request_rules/{id} | Partially update an Edge Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesRetrieve**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesRetrieve) | **Get** /edge_application/applications/{application_id}/request_rules/{id} | Retrieve details of an Edge Application Request Rule
[**EdgeApplicationApiApplicationsRequestRulesUpdate**](EdgeApplicationsRequestRulesAPI.md#EdgeApplicationApiApplicationsRequestRulesUpdate) | **Put** /edge_application/applications/{application_id}/request_rules/{id} | Update an Edge Application Request Rule



## EdgeApplicationApiApplicationsRequestRulesCreate

> ResponseEdgeApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesCreate(ctx, applicationId).EdgeApplicationRequestPhaseRuleEngineRequest(edgeApplicationRequestPhaseRuleEngineRequest).Execute()

Create an Edge Application Request Rule



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
	edgeApplicationRequestPhaseRuleEngineRequest := *openapiclient.NewEdgeApplicationRequestPhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest{openapiclient.EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest{EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineAddHeaderRequest: openapiclient.NewEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineAddHeaderRequest("Type_example", *openapiclient.NewEdgeApplicationRuleEngineAddHeaderAttributesRequest("Value_example"))}}) // EdgeApplicationRequestPhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesCreate(context.Background(), applicationId).EdgeApplicationRequestPhaseRuleEngineRequest(edgeApplicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesCreate`: ResponseEdgeApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesCreate`: %v\n", resp)
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

 **edgeApplicationRequestPhaseRuleEngineRequest** | [**EdgeApplicationRequestPhaseRuleEngineRequest**](EdgeApplicationRequestPhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationRequestPhaseRuleEngine**](ResponseEdgeApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesDestroy

> ResponseDeleteEdgeApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesDestroy(ctx, applicationId, id).Execute()

Destroy an Edge Application Request Rule



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
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesDestroy(context.Background(), applicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesDestroy`: ResponseDeleteEdgeApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesDestroy`: %v\n", resp)
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

[**ResponseDeleteEdgeApplicationRequestPhaseRuleEngine**](ResponseDeleteEdgeApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesList

> PaginatedEdgeApplicationRequestPhaseRuleEngineList EdgeApplicationApiApplicationsRequestRulesList(ctx, applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Application Request Rules



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
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesList(context.Background(), applicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesList`: PaginatedEdgeApplicationRequestPhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesList`: %v\n", resp)
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

[**PaginatedEdgeApplicationRequestPhaseRuleEngineList**](PaginatedEdgeApplicationRequestPhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesOrderUpdate

> PaginatedEdgeApplicationRequestPhaseRuleEngineList EdgeApplicationApiApplicationsRequestRulesOrderUpdate(ctx, applicationId).EdgeApplicationRequestPhaseRuleEngineOrderRequest(edgeApplicationRequestPhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

Ordering Edge Application Request Rules



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
	edgeApplicationRequestPhaseRuleEngineOrderRequest := *openapiclient.NewEdgeApplicationRequestPhaseRuleEngineOrderRequest([]int64{int64(123)}) // EdgeApplicationRequestPhaseRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesOrderUpdate(context.Background(), applicationId).EdgeApplicationRequestPhaseRuleEngineOrderRequest(edgeApplicationRequestPhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesOrderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesOrderUpdate`: PaginatedEdgeApplicationRequestPhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesOrderUpdate`: %v\n", resp)
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

 **edgeApplicationRequestPhaseRuleEngineOrderRequest** | [**EdgeApplicationRequestPhaseRuleEngineOrderRequest**](EdgeApplicationRequestPhaseRuleEngineOrderRequest.md) |  | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeApplicationRequestPhaseRuleEngineList**](PaginatedEdgeApplicationRequestPhaseRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesPartialUpdate

> ResponseEdgeApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesPartialUpdate(ctx, applicationId, id).PatchedEdgeApplicationRequestPhaseRuleEngineRequest(patchedEdgeApplicationRequestPhaseRuleEngineRequest).Execute()

Partially update an Edge Application Request Rule



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
	patchedEdgeApplicationRequestPhaseRuleEngineRequest := *openapiclient.NewPatchedEdgeApplicationRequestPhaseRuleEngineRequest() // PatchedEdgeApplicationRequestPhaseRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesPartialUpdate(context.Background(), applicationId, id).PatchedEdgeApplicationRequestPhaseRuleEngineRequest(patchedEdgeApplicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesPartialUpdate`: ResponseEdgeApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesPartialUpdate`: %v\n", resp)
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


 **patchedEdgeApplicationRequestPhaseRuleEngineRequest** | [**PatchedEdgeApplicationRequestPhaseRuleEngineRequest**](PatchedEdgeApplicationRequestPhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationRequestPhaseRuleEngine**](ResponseEdgeApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesRetrieve

> ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesRetrieve(ctx, applicationId, id).Fields(fields).Execute()

Retrieve details of an Edge Application Request Rule



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
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesRetrieve(context.Background(), applicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesRetrieve`: ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesRetrieve`: %v\n", resp)
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

[**ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine**](ResponseRetrieveEdgeApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationApiApplicationsRequestRulesUpdate

> ResponseEdgeApplicationRequestPhaseRuleEngine EdgeApplicationApiApplicationsRequestRulesUpdate(ctx, applicationId, id).EdgeApplicationRequestPhaseRuleEngineRequest(edgeApplicationRequestPhaseRuleEngineRequest).Execute()

Update an Edge Application Request Rule



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
	edgeApplicationRequestPhaseRuleEngineRequest := *openapiclient.NewEdgeApplicationRequestPhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest{openapiclient.EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest{EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineAddHeaderRequest: openapiclient.NewEdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineAddHeaderRequest("Type_example", *openapiclient.NewEdgeApplicationRuleEngineAddHeaderAttributesRequest("Value_example"))}}) // EdgeApplicationRequestPhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesUpdate(context.Background(), applicationId, id).EdgeApplicationRequestPhaseRuleEngineRequest(edgeApplicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationApiApplicationsRequestRulesUpdate`: ResponseEdgeApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRequestRulesAPI.EdgeApplicationApiApplicationsRequestRulesUpdate`: %v\n", resp)
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


 **edgeApplicationRequestPhaseRuleEngineRequest** | [**EdgeApplicationRequestPhaseRuleEngineRequest**](EdgeApplicationRequestPhaseRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationRequestPhaseRuleEngine**](ResponseEdgeApplicationRequestPhaseRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

