# \FirewallsRulesEngineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirewallRule**](FirewallsRulesEngineAPI.md#CreateFirewallRule) | **Post** /workspace/firewalls/{firewall_id}/request_rules | Create an Firewall Rule
[**DeleteFirewallRule**](FirewallsRulesEngineAPI.md#DeleteFirewallRule) | **Delete** /workspace/firewalls/{firewall_id}/request_rules/{id} | Delete an Firewall Rule
[**ListFirewallRules**](FirewallsRulesEngineAPI.md#ListFirewallRules) | **Get** /workspace/firewalls/{firewall_id}/request_rules | List Firewall Rules
[**OrderFirewallRules**](FirewallsRulesEngineAPI.md#OrderFirewallRules) | **Put** /workspace/firewalls/{firewall_id}/request_rules/order | Ordering Firewall Rules
[**PartialUpdateFirewallRule**](FirewallsRulesEngineAPI.md#PartialUpdateFirewallRule) | **Patch** /workspace/firewalls/{firewall_id}/request_rules/{id} | Partially update an Firewall Rule
[**RetrieveFirewallRule**](FirewallsRulesEngineAPI.md#RetrieveFirewallRule) | **Get** /workspace/firewalls/{firewall_id}/request_rules/{id} | Retrieve details of an Firewall Rule
[**UpdateFirewallRule**](FirewallsRulesEngineAPI.md#UpdateFirewallRule) | **Put** /workspace/firewalls/{firewall_id}/request_rules/{id} | Update an Firewall Rule



## CreateFirewallRule

> ResponseFirewallRuleEngine CreateFirewallRule(ctx, firewallId).FirewallRuleEngineRequest(firewallRuleEngineRequest).Execute()

Create an Firewall Rule



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
	firewallId := "firewallId_example" // string | 
	firewallRuleEngineRequest := *openapiclient.NewFirewallRuleEngineRequest("Name_example", [][]EdgeFirewallCriterionFieldRequest{[]openapiclient.EdgeFirewallCriterionFieldRequest{*openapiclient.NewEdgeFirewallCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.FirewallBehaviorsRequest{openapiclient.FirewallBehaviorsRequest{FirewallBehaviorsFirewallBehaviorNoArgumentsRequest: openapiclient.NewFirewallBehaviorsFirewallBehaviorNoArgumentsRequest("Type_example")}}) // FirewallRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.CreateFirewallRule(context.Background(), firewallId).FirewallRuleEngineRequest(firewallRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.CreateFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirewallRule`: ResponseFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.CreateFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRuleEngineRequest** | [**FirewallRuleEngineRequest**](FirewallRuleEngineRequest.md) |  | 

### Return type

[**ResponseFirewallRuleEngine**](ResponseFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRule

> ResponseDeleteFirewallRuleEngine DeleteFirewallRule(ctx, firewallId, id).Execute()

Delete an Firewall Rule



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
	firewallId := "firewallId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.DeleteFirewallRule(context.Background(), firewallId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.DeleteFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRule`: ResponseDeleteFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.DeleteFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeleteFirewallRuleEngine**](ResponseDeleteFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFirewallRules

> PaginatedFirewallRuleEngineList ListFirewallRules(ctx, firewallId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Firewall Rules



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
	firewallId := "firewallId_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, last_editor, last_modified, active, description, order, behaviors, criteria) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.ListFirewallRules(context.Background(), firewallId).Fields(fields).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.ListFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewallRules`: PaginatedFirewallRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.ListFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: name, last_editor, last_modified, active, description, order, behaviors, criteria) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedFirewallRuleEngineList**](PaginatedFirewallRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderFirewallRules

> PaginatedFirewallRuleEngineList OrderFirewallRules(ctx, firewallId).FirewallRuleEngineOrderRequest(firewallRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

Ordering Firewall Rules



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
	firewallId := "firewallId_example" // string | 
	firewallRuleEngineOrderRequest := *openapiclient.NewFirewallRuleEngineOrderRequest([]int64{int64(123)}) // FirewallRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.OrderFirewallRules(context.Background(), firewallId).FirewallRuleEngineOrderRequest(firewallRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.OrderFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderFirewallRules`: PaginatedFirewallRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.OrderFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRuleEngineOrderRequest** | [**FirewallRuleEngineOrderRequest**](FirewallRuleEngineOrderRequest.md) |  | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedFirewallRuleEngineList**](PaginatedFirewallRuleEngineList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateFirewallRule

> ResponseFirewallRuleEngine PartialUpdateFirewallRule(ctx, firewallId, id).PatchedFirewallRuleEngineRequest(patchedFirewallRuleEngineRequest).Execute()

Partially update an Firewall Rule



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
	firewallId := "firewallId_example" // string | 
	id := "id_example" // string | 
	patchedFirewallRuleEngineRequest := *openapiclient.NewPatchedFirewallRuleEngineRequest() // PatchedFirewallRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.PartialUpdateFirewallRule(context.Background(), firewallId, id).PatchedFirewallRuleEngineRequest(patchedFirewallRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.PartialUpdateFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateFirewallRule`: ResponseFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.PartialUpdateFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedFirewallRuleEngineRequest** | [**PatchedFirewallRuleEngineRequest**](PatchedFirewallRuleEngineRequest.md) |  | 

### Return type

[**ResponseFirewallRuleEngine**](ResponseFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirewallRule

> ResponseRetrieveFirewallRuleEngine RetrieveFirewallRule(ctx, firewallId, id).Fields(fields).Execute()

Retrieve details of an Firewall Rule



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
	firewallId := "firewallId_example" // string | 
	id := "id_example" // string | 
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.RetrieveFirewallRule(context.Background(), firewallId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.RetrieveFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveFirewallRule`: ResponseRetrieveFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.RetrieveFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponseRetrieveFirewallRuleEngine**](ResponseRetrieveFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirewallRule

> ResponseFirewallRuleEngine UpdateFirewallRule(ctx, firewallId, id).FirewallRuleEngineRequest(firewallRuleEngineRequest).Execute()

Update an Firewall Rule



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
	firewallId := "firewallId_example" // string | 
	id := "id_example" // string | 
	firewallRuleEngineRequest := *openapiclient.NewFirewallRuleEngineRequest("Name_example", [][]EdgeFirewallCriterionFieldRequest{[]openapiclient.EdgeFirewallCriterionFieldRequest{*openapiclient.NewEdgeFirewallCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.FirewallBehaviorsRequest{openapiclient.FirewallBehaviorsRequest{FirewallBehaviorsFirewallBehaviorNoArgumentsRequest: openapiclient.NewFirewallBehaviorsFirewallBehaviorNoArgumentsRequest("Type_example")}}) // FirewallRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallsRulesEngineAPI.UpdateFirewallRule(context.Background(), firewallId, id).FirewallRuleEngineRequest(firewallRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallsRulesEngineAPI.UpdateFirewallRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirewallRule`: ResponseFirewallRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `FirewallsRulesEngineAPI.UpdateFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallRuleEngineRequest** | [**FirewallRuleEngineRequest**](FirewallRuleEngineRequest.md) |  | 

### Return type

[**ResponseFirewallRuleEngine**](ResponseFirewallRuleEngine.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

