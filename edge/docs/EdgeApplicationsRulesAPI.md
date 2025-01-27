# \EdgeApplicationsRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeApplicationRule**](EdgeApplicationsRulesAPI.md#CreateEdgeApplicationRule) | **Post** /edge_application/applications/{edgeApplicationId}/rules | Create an Edge Application Rule
[**DestroyEdgeApplicationRule**](EdgeApplicationsRulesAPI.md#DestroyEdgeApplicationRule) | **Delete** /edge_application/applications/{edgeApplicationId}/rules/{id} | Destroy an Edge Application Rule
[**ListEdgeApplicationRule**](EdgeApplicationsRulesAPI.md#ListEdgeApplicationRule) | **Get** /edge_application/applications/{edgeApplicationId}/rules | List Edge Application Rules
[**OrderEdgeApplicationRules**](EdgeApplicationsRulesAPI.md#OrderEdgeApplicationRules) | **Put** /edge_application/applications/{edgeApplicationId}/rules/order | Ordering Edge Application Rules
[**PartialUpdateEdgeApplicationRule**](EdgeApplicationsRulesAPI.md#PartialUpdateEdgeApplicationRule) | **Patch** /edge_application/applications/{edgeApplicationId}/rules/{id} | Partially update an Edge Application Rule
[**RetrieveEdgeApplicationRule**](EdgeApplicationsRulesAPI.md#RetrieveEdgeApplicationRule) | **Get** /edge_application/applications/{edgeApplicationId}/rules/{id} | Retrieve details of an Edge Application Rule
[**UpdateEdgeApplicationRule**](EdgeApplicationsRulesAPI.md#UpdateEdgeApplicationRule) | **Put** /edge_application/applications/{edgeApplicationId}/rules/{id} | Update an Edge Application Rule



## CreateEdgeApplicationRule

> ResponseEdgeApplicationRuleEngine CreateEdgeApplicationRule(ctx, edgeApplicationId).EdgeApplicationRuleEngineRequest(edgeApplicationRuleEngineRequest).Execute()

Create an Edge Application Rule



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
	edgeApplicationRuleEngineRequest := *openapiclient.NewEdgeApplicationRuleEngineRequest("Name_example", *openapiclient.NewPhaseEnum(), []openapiclient.EdgeApplicationBehaviorFieldRequest{*openapiclient.NewEdgeApplicationBehaviorFieldRequest(*openapiclient.NewEdgeApplicationBehaviorNameField())}, [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest(*openapiclient.NewEdgeApplicationCriterionFieldVariableEnum(), *openapiclient.NewConditionalEnum(), *openapiclient.NewOperatorEnum())}}) // EdgeApplicationRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.CreateEdgeApplicationRule(context.Background(), edgeApplicationId).EdgeApplicationRuleEngineRequest(edgeApplicationRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.CreateEdgeApplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeApplicationRule`: ResponseEdgeApplicationRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.CreateEdgeApplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeApplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeApplicationRuleEngineRequest** | [**EdgeApplicationRuleEngineRequest**](EdgeApplicationRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationRuleEngine**](ResponseEdgeApplicationRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeApplicationRule

> ResponseDeleteEdgeApplicationRuleEngine DestroyEdgeApplicationRule(ctx, edgeApplicationId, id).Execute()

Destroy an Edge Application Rule



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
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.DestroyEdgeApplicationRule(context.Background(), edgeApplicationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.DestroyEdgeApplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeApplicationRule`: ResponseDeleteEdgeApplicationRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.DestroyEdgeApplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeApplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteEdgeApplicationRuleEngine**](ResponseDeleteEdgeApplicationRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeApplicationRule

> PaginatedResponseListEdgeApplicationRuleEngineList ListEdgeApplicationRule(ctx, edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Application Rules



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, phase, active, description, order, behaviors, criteria, last_editor, last_modified) (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.ListEdgeApplicationRule(context.Background(), edgeApplicationId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.ListEdgeApplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeApplicationRule`: PaginatedResponseListEdgeApplicationRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.ListEdgeApplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeApplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, phase, active, description, order, behaviors, criteria, last_editor, last_modified) | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedResponseListEdgeApplicationRuleEngineList**](PaginatedResponseListEdgeApplicationRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderEdgeApplicationRules

> EdgeApplicationRuleEngine OrderEdgeApplicationRules(ctx, edgeApplicationId).EdgeApplicationRuleEngineOrderRequest(edgeApplicationRuleEngineOrderRequest).Execute()

Ordering Edge Application Rules



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
	edgeApplicationRuleEngineOrderRequest := *openapiclient.NewEdgeApplicationRuleEngineOrderRequest([]int64{int64(123)}) // EdgeApplicationRuleEngineOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.OrderEdgeApplicationRules(context.Background(), edgeApplicationId).EdgeApplicationRuleEngineOrderRequest(edgeApplicationRuleEngineOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.OrderEdgeApplicationRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderEdgeApplicationRules`: EdgeApplicationRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.OrderEdgeApplicationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderEdgeApplicationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeApplicationRuleEngineOrderRequest** | [**EdgeApplicationRuleEngineOrderRequest**](EdgeApplicationRuleEngineOrderRequest.md) |  | 

### Return type

[**EdgeApplicationRuleEngine**](EdgeApplicationRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeApplicationRule

> ResponseEdgeApplicationRuleEngine PartialUpdateEdgeApplicationRule(ctx, edgeApplicationId, id).PatchedEdgeApplicationRuleEngineRequest(patchedEdgeApplicationRuleEngineRequest).Execute()

Partially update an Edge Application Rule



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
	patchedEdgeApplicationRuleEngineRequest := *openapiclient.NewPatchedEdgeApplicationRuleEngineRequest() // PatchedEdgeApplicationRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.PartialUpdateEdgeApplicationRule(context.Background(), edgeApplicationId, id).PatchedEdgeApplicationRuleEngineRequest(patchedEdgeApplicationRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.PartialUpdateEdgeApplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeApplicationRule`: ResponseEdgeApplicationRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.PartialUpdateEdgeApplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeApplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedEdgeApplicationRuleEngineRequest** | [**PatchedEdgeApplicationRuleEngineRequest**](PatchedEdgeApplicationRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationRuleEngine**](ResponseEdgeApplicationRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeApplicationRule

> ResponseRetrieveEdgeApplicationRuleEngine RetrieveEdgeApplicationRule(ctx, edgeApplicationId, id).Fields(fields).Execute()

Retrieve details of an Edge Application Rule



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
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.RetrieveEdgeApplicationRule(context.Background(), edgeApplicationId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.RetrieveEdgeApplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeApplicationRule`: ResponseRetrieveEdgeApplicationRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.RetrieveEdgeApplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeApplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeApplicationRuleEngine**](ResponseRetrieveEdgeApplicationRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeApplicationRule

> ResponseEdgeApplicationRuleEngine UpdateEdgeApplicationRule(ctx, edgeApplicationId, id).EdgeApplicationRuleEngineRequest(edgeApplicationRuleEngineRequest).Execute()

Update an Edge Application Rule



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
	edgeApplicationRuleEngineRequest := *openapiclient.NewEdgeApplicationRuleEngineRequest("Name_example", *openapiclient.NewPhaseEnum(), []openapiclient.EdgeApplicationBehaviorFieldRequest{*openapiclient.NewEdgeApplicationBehaviorFieldRequest(*openapiclient.NewEdgeApplicationBehaviorNameField())}, [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest(*openapiclient.NewEdgeApplicationCriterionFieldVariableEnum(), *openapiclient.NewConditionalEnum(), *openapiclient.NewOperatorEnum())}}) // EdgeApplicationRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsRulesAPI.UpdateEdgeApplicationRule(context.Background(), edgeApplicationId, id).EdgeApplicationRuleEngineRequest(edgeApplicationRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesAPI.UpdateEdgeApplicationRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeApplicationRule`: ResponseEdgeApplicationRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesAPI.UpdateEdgeApplicationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeApplicationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeApplicationRuleEngineRequest** | [**EdgeApplicationRuleEngineRequest**](EdgeApplicationRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeApplicationRuleEngine**](ResponseEdgeApplicationRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

