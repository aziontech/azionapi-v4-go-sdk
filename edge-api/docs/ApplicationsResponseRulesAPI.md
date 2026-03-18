# \ApplicationsResponseRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationResponseRule**](ApplicationsResponseRulesAPI.md#CreateApplicationResponseRule) | **Post** /workspace/applications/{application_id}/response_rules | Create an Application Response Rule
[**DeleteApplicationResponseRule**](ApplicationsResponseRulesAPI.md#DeleteApplicationResponseRule) | **Delete** /workspace/applications/{application_id}/response_rules/{response_rule_id} | Delete an Application Response Rule
[**ListApplicationResponseRules**](ApplicationsResponseRulesAPI.md#ListApplicationResponseRules) | **Get** /workspace/applications/{application_id}/response_rules | List Application Response Rules
[**PartialUpdateApplicationResponseRule**](ApplicationsResponseRulesAPI.md#PartialUpdateApplicationResponseRule) | **Patch** /workspace/applications/{application_id}/response_rules/{response_rule_id} | Partially update an Application Response Rule
[**RetrieveApplicationResponseRule**](ApplicationsResponseRulesAPI.md#RetrieveApplicationResponseRule) | **Get** /workspace/applications/{application_id}/response_rules/{response_rule_id} | Retrieve details of an Application Response Rule
[**UpdateApplicationResponseRule**](ApplicationsResponseRulesAPI.md#UpdateApplicationResponseRule) | **Put** /workspace/applications/{application_id}/response_rules/{response_rule_id} | Update an Application Response Rule
[**UpdateApplicationResponseRulesOrder**](ApplicationsResponseRulesAPI.md#UpdateApplicationResponseRulesOrder) | **Put** /workspace/applications/{application_id}/response_rules/order | Ordering Application Response Rules



## CreateApplicationResponseRule

> ResponsePhaseRuleResponse CreateApplicationResponseRule(ctx, applicationId).ResponsePhaseRuleRequest(responsePhaseRuleRequest).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	responsePhaseRuleRequest := *openapiclient.NewResponsePhaseRuleRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ResponsePhaseBehaviorRequest{openapiclient.ResponsePhaseBehaviorRequest{BehaviorArgs: openapiclient.NewBehaviorArgs("Type_example", *openapiclient.NewBehaviorArgsAttributes(openapiclient.BehaviorArgsAttributes_value{Int64: new(int64)}))}}) // ResponsePhaseRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.CreateApplicationResponseRule(context.Background(), applicationId).ResponsePhaseRuleRequest(responsePhaseRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.CreateApplicationResponseRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationResponseRule`: ResponsePhaseRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.CreateApplicationResponseRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationResponseRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **responsePhaseRuleRequest** | [**ResponsePhaseRuleRequest**](ResponsePhaseRuleRequest.md) |  | 

### Return type

[**ResponsePhaseRuleResponse**](ResponsePhaseRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationResponseRule

> DeleteResponse DeleteApplicationResponseRule(ctx, applicationId, responseRuleId).Execute()

Delete an Application Response Rule



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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	responseRuleId := int64(789) // int64 | A unique integer value identifying the response rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.DeleteApplicationResponseRule(context.Background(), applicationId, responseRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.DeleteApplicationResponseRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationResponseRule`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.DeleteApplicationResponseRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**responseRuleId** | **int64** | A unique integer value identifying the response rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationResponseRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationResponseRules

> PaginatedResponsePhaseRuleList ListApplicationResponseRules(ctx, applicationId).Description(description).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).OrderGte(orderGte).OrderLte(orderLte).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Application Response Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	description := "description_example" // string | Filter by description (case-insensitive, partial match). (optional)
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)
	id := int64(789) // int64 | Filter by id (accepts comma-separated values). (optional)
	lastEditor := "lastEditor_example" // string | Filter by last editor (case-insensitive, partial match). (optional)
	lastModifiedGte := time.Now() // time.Time | Filter by last modified date (greater than or equal). (optional)
	lastModifiedLte := time.Now() // time.Time | Filter by last modified date (less than or equal). (optional)
	name := "name_example" // string | Filter by name (case-insensitive, partial match). (optional)
	orderGte := int64(789) // int64 | Filter by order (greater than or equal). (optional)
	orderLte := int64(789) // int64 | Filter by order (less than or equal). (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: id, name, description, last_editor, last_modified, order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term to filter results. Searches across the following fields: name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.ListApplicationResponseRules(context.Background(), applicationId).Description(description).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).OrderGte(orderGte).OrderLte(orderLte).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.ListApplicationResponseRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationResponseRules`: PaginatedResponsePhaseRuleList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.ListApplicationResponseRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationResponseRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** | Filter by description (case-insensitive, partial match). | 
 **fields** | **string** | Comma-separated list of field names to include in the response. | 
 **id** | **int64** | Filter by id (accepts comma-separated values). | 
 **lastEditor** | **string** | Filter by last editor (case-insensitive, partial match). | 
 **lastModifiedGte** | **time.Time** | Filter by last modified date (greater than or equal). | 
 **lastModifiedLte** | **time.Time** | Filter by last modified date (less than or equal). | 
 **name** | **string** | Filter by name (case-insensitive, partial match). | 
 **orderGte** | **int64** | Filter by order (greater than or equal). | 
 **orderLte** | **int64** | Filter by order (less than or equal). | 
 **ordering** | **string** | Which field to use when ordering the results. (Valid fields: id, name, description, last_editor, last_modified, order) | 
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | A numeric value that indicates the number of items per page. | 
 **search** | **string** | A search term to filter results. Searches across the following fields: name. | 

### Return type

[**PaginatedResponsePhaseRuleList**](PaginatedResponsePhaseRuleList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateApplicationResponseRule

> ResponsePhaseRuleResponse PartialUpdateApplicationResponseRule(ctx, applicationId, responseRuleId).PatchedResponsePhaseRuleRequest(patchedResponsePhaseRuleRequest).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	responseRuleId := int64(789) // int64 | A unique integer value identifying the response rule.
	patchedResponsePhaseRuleRequest := *openapiclient.NewPatchedResponsePhaseRuleRequest() // PatchedResponsePhaseRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.PartialUpdateApplicationResponseRule(context.Background(), applicationId, responseRuleId).PatchedResponsePhaseRuleRequest(patchedResponsePhaseRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.PartialUpdateApplicationResponseRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateApplicationResponseRule`: ResponsePhaseRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.PartialUpdateApplicationResponseRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**responseRuleId** | **int64** | A unique integer value identifying the response rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateApplicationResponseRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedResponsePhaseRuleRequest** | [**PatchedResponsePhaseRuleRequest**](PatchedResponsePhaseRuleRequest.md) |  | 

### Return type

[**ResponsePhaseRuleResponse**](ResponsePhaseRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveApplicationResponseRule

> ResponsePhaseRuleResponse RetrieveApplicationResponseRule(ctx, applicationId, responseRuleId).Fields(fields).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	responseRuleId := int64(789) // int64 | A unique integer value identifying the response rule.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.RetrieveApplicationResponseRule(context.Background(), applicationId, responseRuleId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.RetrieveApplicationResponseRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveApplicationResponseRule`: ResponsePhaseRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.RetrieveApplicationResponseRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**responseRuleId** | **int64** | A unique integer value identifying the response rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveApplicationResponseRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Comma-separated list of field names to include in the response. | 

### Return type

[**ResponsePhaseRuleResponse**](ResponsePhaseRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationResponseRule

> ResponsePhaseRuleResponse UpdateApplicationResponseRule(ctx, applicationId, responseRuleId).ResponsePhaseRuleRequest(responsePhaseRuleRequest).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	responseRuleId := int64(789) // int64 | A unique integer value identifying the response rule.
	responsePhaseRuleRequest := *openapiclient.NewResponsePhaseRuleRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ResponsePhaseBehaviorRequest{openapiclient.ResponsePhaseBehaviorRequest{BehaviorArgs: openapiclient.NewBehaviorArgs("Type_example", *openapiclient.NewBehaviorArgsAttributes(openapiclient.BehaviorArgsAttributes_value{Int64: new(int64)}))}}) // ResponsePhaseRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.UpdateApplicationResponseRule(context.Background(), applicationId, responseRuleId).ResponsePhaseRuleRequest(responsePhaseRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.UpdateApplicationResponseRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationResponseRule`: ResponsePhaseRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.UpdateApplicationResponseRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**responseRuleId** | **int64** | A unique integer value identifying the response rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationResponseRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **responsePhaseRuleRequest** | [**ResponsePhaseRuleRequest**](ResponsePhaseRuleRequest.md) |  | 

### Return type

[**ResponsePhaseRuleResponse**](ResponsePhaseRuleResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationResponseRulesOrder

> PaginatedResponsePhaseRuleList UpdateApplicationResponseRulesOrder(ctx, applicationId).ApplicationResponsePhaseRuleEngineOrderRequest(applicationResponsePhaseRuleEngineOrderRequest).Search(search).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	applicationResponsePhaseRuleEngineOrderRequest := *openapiclient.NewApplicationResponsePhaseRuleEngineOrderRequest([]int64{int64(123)}) // ApplicationResponsePhaseRuleEngineOrderRequest | 
	search := "search_example" // string | A search term to filter results. Searches across the following fields: name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsResponseRulesAPI.UpdateApplicationResponseRulesOrder(context.Background(), applicationId).ApplicationResponsePhaseRuleEngineOrderRequest(applicationResponsePhaseRuleEngineOrderRequest).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsResponseRulesAPI.UpdateApplicationResponseRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationResponseRulesOrder`: PaginatedResponsePhaseRuleList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsResponseRulesAPI.UpdateApplicationResponseRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationResponseRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationResponsePhaseRuleEngineOrderRequest** | [**ApplicationResponsePhaseRuleEngineOrderRequest**](ApplicationResponsePhaseRuleEngineOrderRequest.md) |  | 
 **search** | **string** | A search term to filter results. Searches across the following fields: name. | 

### Return type

[**PaginatedResponsePhaseRuleList**](PaginatedResponsePhaseRuleList.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

