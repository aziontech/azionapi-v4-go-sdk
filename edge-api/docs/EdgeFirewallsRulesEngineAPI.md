# \EdgeFirewallsRulesEngineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeFirewallRule**](EdgeFirewallsRulesEngineAPI.md#CreateEdgeFirewallRule) | **Post** /edge_firewall/firewalls/{edge_firewall_id}/request_rules | Create an Edge Firewall Rule
[**DestroyEdgeFirewallRule**](EdgeFirewallsRulesEngineAPI.md#DestroyEdgeFirewallRule) | **Delete** /edge_firewall/firewalls/{edge_firewall_id}/request_rules/{id} | Destroy an Edge Firewall Rule
[**ListEdgeFirewallRules**](EdgeFirewallsRulesEngineAPI.md#ListEdgeFirewallRules) | **Get** /edge_firewall/firewalls/{edge_firewall_id}/request_rules | List Edge Firewall Rules
[**OrderEdgeFirewallRules**](EdgeFirewallsRulesEngineAPI.md#OrderEdgeFirewallRules) | **Put** /edge_firewall/firewalls/{edge_firewall_id}/request_rules/order | Ordering Edge Firewall Rules
[**PartialUpdateEdgeFirewallRule**](EdgeFirewallsRulesEngineAPI.md#PartialUpdateEdgeFirewallRule) | **Patch** /edge_firewall/firewalls/{edge_firewall_id}/request_rules/{id} | Partially update an Edge Firewall Rule
[**RetrieveEdgeFirewallRule**](EdgeFirewallsRulesEngineAPI.md#RetrieveEdgeFirewallRule) | **Get** /edge_firewall/firewalls/{edge_firewall_id}/request_rules/{id} | Retrieve details of an Edge Firewall Rule
[**UpdateEdgeFirewallRule**](EdgeFirewallsRulesEngineAPI.md#UpdateEdgeFirewallRule) | **Put** /edge_firewall/firewalls/{edge_firewall_id}/request_rules/{id} | Update an Edge Firewall Rule



## CreateEdgeFirewallRule

> ResponseEdgeFirewallRuleEngine CreateEdgeFirewallRule(ctx, edgeFirewallId).EdgeFirewallRuleEngineRequest(edgeFirewallRuleEngineRequest).Execute()

Create an Edge Firewall Rule



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	edgeFirewallRuleEngineRequest := *openapiclient.NewEdgeFirewallRuleEngineRequest("Name_example", [][]EdgeFirewallCriterionFieldRequest{[]openapiclient.EdgeFirewallCriterionFieldRequest{*openapiclient.NewEdgeFirewallCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.EdgeFirewallBehaviorsRequest{openapiclient.EdgeFirewallBehaviorsRequest{EdgeFirewallBehaviorsEdgeFirewallBehaviorNoArgumentsRequest: openapiclient.NewEdgeFirewallBehaviorsEdgeFirewallBehaviorNoArgumentsRequest("Type_example")}}) // EdgeFirewallRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.CreateEdgeFirewallRule(context.Background(), edgeFirewallId).EdgeFirewallRuleEngineRequest(edgeFirewallRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.CreateEdgeFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEdgeFirewallRule`: ResponseEdgeFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.CreateEdgeFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEdgeFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFirewallRuleEngineRequest** | [**EdgeFirewallRuleEngineRequest**](EdgeFirewallRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeFirewallRuleEngine**](ResponseEdgeFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyEdgeFirewallRule

> ResponseDeleteEdgeFirewallRuleEngine DestroyEdgeFirewallRule(ctx, edgeFirewallId, id).Execute()

Destroy an Edge Firewall Rule



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.DestroyEdgeFirewallRule(context.Background(), edgeFirewallId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.DestroyEdgeFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DestroyEdgeFirewallRule`: ResponseDeleteEdgeFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.DestroyEdgeFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyEdgeFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteEdgeFirewallRuleEngine**](ResponseDeleteEdgeFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEdgeFirewallRules

> PaginatedEdgeFirewallRuleEngineList ListEdgeFirewallRules(ctx, edgeFirewallId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Edge Firewall Rules



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, last_editor, last_modified, active, description, order, behaviors, criteria) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.ListEdgeFirewallRules(context.Background(), edgeFirewallId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.ListEdgeFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEdgeFirewallRules`: PaginatedEdgeFirewallRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.ListEdgeFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEdgeFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, last_editor, last_modified, active, description, order, behaviors, criteria) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeFirewallRuleEngineList**](PaginatedEdgeFirewallRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderEdgeFirewallRules

> PaginatedEdgeFirewallRuleEngineList OrderEdgeFirewallRules(ctx, edgeFirewallId).EdgeFirewallRuleEngineOrderRequest(edgeFirewallRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

Ordering Edge Firewall Rules



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	edgeFirewallRuleEngineOrderRequest := *openapiclient.NewEdgeFirewallRuleEngineOrderRequest([]int64{int64(123)}) // EdgeFirewallRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.OrderEdgeFirewallRules(context.Background(), edgeFirewallId).EdgeFirewallRuleEngineOrderRequest(edgeFirewallRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.OrderEdgeFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderEdgeFirewallRules`: PaginatedEdgeFirewallRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.OrderEdgeFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderEdgeFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFirewallRuleEngineOrderRequest** | [**EdgeFirewallRuleEngineOrderRequest**](EdgeFirewallRuleEngineOrderRequest.md) |  | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEdgeFirewallRuleEngineList**](PaginatedEdgeFirewallRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateEdgeFirewallRule

> ResponseEdgeFirewallRuleEngine PartialUpdateEdgeFirewallRule(ctx, edgeFirewallId, id).PatchedEdgeFirewallRuleEngineRequest(patchedEdgeFirewallRuleEngineRequest).Execute()

Partially update an Edge Firewall Rule



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 
	patchedEdgeFirewallRuleEngineRequest := *openapiclient.NewPatchedEdgeFirewallRuleEngineRequest() // PatchedEdgeFirewallRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.PartialUpdateEdgeFirewallRule(context.Background(), edgeFirewallId, id).PatchedEdgeFirewallRuleEngineRequest(patchedEdgeFirewallRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.PartialUpdateEdgeFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateEdgeFirewallRule`: ResponseEdgeFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.PartialUpdateEdgeFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateEdgeFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedEdgeFirewallRuleEngineRequest** | [**PatchedEdgeFirewallRuleEngineRequest**](PatchedEdgeFirewallRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeFirewallRuleEngine**](ResponseEdgeFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveEdgeFirewallRule

> ResponseRetrieveEdgeFirewallRuleEngine RetrieveEdgeFirewallRule(ctx, edgeFirewallId, id).Fields(fields).Execute()

Retrieve details of an Edge Firewall Rule



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.RetrieveEdgeFirewallRule(context.Background(), edgeFirewallId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.RetrieveEdgeFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveEdgeFirewallRule`: ResponseRetrieveEdgeFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.RetrieveEdgeFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveEdgeFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveEdgeFirewallRuleEngine**](ResponseRetrieveEdgeFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEdgeFirewallRule

> ResponseEdgeFirewallRuleEngine UpdateEdgeFirewallRule(ctx, edgeFirewallId, id).EdgeFirewallRuleEngineRequest(edgeFirewallRuleEngineRequest).Execute()

Update an Edge Firewall Rule



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
	edgeFirewallId := "edgeFirewallId_example" // string | 
	id := "id_example" // string | 
	edgeFirewallRuleEngineRequest := *openapiclient.NewEdgeFirewallRuleEngineRequest("Name_example", [][]EdgeFirewallCriterionFieldRequest{[]openapiclient.EdgeFirewallCriterionFieldRequest{*openapiclient.NewEdgeFirewallCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.EdgeFirewallBehaviorsRequest{openapiclient.EdgeFirewallBehaviorsRequest{EdgeFirewallBehaviorsEdgeFirewallBehaviorNoArgumentsRequest: openapiclient.NewEdgeFirewallBehaviorsEdgeFirewallBehaviorNoArgumentsRequest("Type_example")}}) // EdgeFirewallRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFirewallsRulesEngineAPI.UpdateEdgeFirewallRule(context.Background(), edgeFirewallId, id).EdgeFirewallRuleEngineRequest(edgeFirewallRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFirewallsRulesEngineAPI.UpdateEdgeFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEdgeFirewallRule`: ResponseEdgeFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `EdgeFirewallsRulesEngineAPI.UpdateEdgeFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEdgeFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **edgeFirewallRuleEngineRequest** | [**EdgeFirewallRuleEngineRequest**](EdgeFirewallRuleEngineRequest.md) |  | 

### Return type

[**ResponseEdgeFirewallRuleEngine**](ResponseEdgeFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

