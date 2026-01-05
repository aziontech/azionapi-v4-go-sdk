# \ApplicationsRequestRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationRequestRule**](ApplicationsRequestRulesAPI.md#CreateApplicationRequestRule) | **Post** /workspace/applications/{application_id}/request_rules | Create an Application Request Rule
[**DeleteApplicationRequestRule**](ApplicationsRequestRulesAPI.md#DeleteApplicationRequestRule) | **Delete** /workspace/applications/{application_id}/request_rules/{request_rule_id} | Delete an Application Request Rule
[**ListApplicationRequestRules**](ApplicationsRequestRulesAPI.md#ListApplicationRequestRules) | **Get** /workspace/applications/{application_id}/request_rules | List Application Request Rules
[**PartialUpdateApplicationRequestRule**](ApplicationsRequestRulesAPI.md#PartialUpdateApplicationRequestRule) | **Patch** /workspace/applications/{application_id}/request_rules/{request_rule_id} | Partially update an Application Request Rule
[**RetrieveApplicationRequestRule**](ApplicationsRequestRulesAPI.md#RetrieveApplicationRequestRule) | **Get** /workspace/applications/{application_id}/request_rules/{request_rule_id} | Retrieve details of an Application Request Rule
[**UpdateApplicationRequestRule**](ApplicationsRequestRulesAPI.md#UpdateApplicationRequestRule) | **Put** /workspace/applications/{application_id}/request_rules/{request_rule_id} | Update an Application Request Rule
[**UpdateApplicationRequestRulesOrder**](ApplicationsRequestRulesAPI.md#UpdateApplicationRequestRulesOrder) | **Put** /workspace/applications/{application_id}/request_rules/order | Ordering Application Request Rules



## CreateApplicationRequestRule

> ResponseApplicationRequestPhaseRuleEngine CreateApplicationRequestRule(ctx, applicationId).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	applicationRequestPhaseRuleEngineRequest := *openapiclient.NewApplicationRequestPhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest: openapiclient.NewApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest("Type_example", *openapiclient.NewApplicationRuleEngineAddHeaderAttributesRequest("Value_example"))}}) // ApplicationRequestPhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.CreateApplicationRequestRule(context.Background(), applicationId).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.CreateApplicationRequestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationRequestRule`: ResponseApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.CreateApplicationRequestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequestRuleRequest struct via the builder pattern


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


## DeleteApplicationRequestRule

> ResponseDeleteApplicationRequestPhaseRuleEngine DeleteApplicationRequestRule(ctx, applicationId, requestRuleId).Execute()

Delete an Application Request Rule



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
	requestRuleId := int64(789) // int64 | A unique integer value identifying the request rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.DeleteApplicationRequestRule(context.Background(), applicationId, requestRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.DeleteApplicationRequestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationRequestRule`: ResponseDeleteApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.DeleteApplicationRequestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**requestRuleId** | **int64** | A unique integer value identifying the request rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequestRuleRequest struct via the builder pattern


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


## ListApplicationRequestRules

> PaginatedApplicationRequestPhaseRuleEngineList ListApplicationRequestRules(ctx, applicationId).Description(description).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).OrderGte(orderGte).OrderLte(orderLte).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

List Application Request Rules



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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: name, active, description, order, criteria, last_editor, last_modified, behaviors) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | A numeric value that indicates the number of items per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.ListApplicationRequestRules(context.Background(), applicationId).Description(description).Fields(fields).Id(id).LastEditor(lastEditor).LastModifiedGte(lastModifiedGte).LastModifiedLte(lastModifiedLte).Name(name).OrderGte(orderGte).OrderLte(orderLte).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.ListApplicationRequestRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationRequestRules`: PaginatedApplicationRequestPhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.ListApplicationRequestRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationRequestRulesRequest struct via the builder pattern


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


## PartialUpdateApplicationRequestRule

> ResponseApplicationRequestPhaseRuleEngine PartialUpdateApplicationRequestRule(ctx, applicationId, requestRuleId).PatchedApplicationRequestPhaseRuleEngineRequest(patchedApplicationRequestPhaseRuleEngineRequest).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	requestRuleId := int64(789) // int64 | A unique integer value identifying the request rule.
	patchedApplicationRequestPhaseRuleEngineRequest := *openapiclient.NewPatchedApplicationRequestPhaseRuleEngineRequest() // PatchedApplicationRequestPhaseRuleEngineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.PartialUpdateApplicationRequestRule(context.Background(), applicationId, requestRuleId).PatchedApplicationRequestPhaseRuleEngineRequest(patchedApplicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.PartialUpdateApplicationRequestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateApplicationRequestRule`: ResponseApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.PartialUpdateApplicationRequestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**requestRuleId** | **int64** | A unique integer value identifying the request rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateApplicationRequestRuleRequest struct via the builder pattern


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


## RetrieveApplicationRequestRule

> ResponseRetrieveApplicationRequestPhaseRuleEngine RetrieveApplicationRequestRule(ctx, applicationId, requestRuleId).Fields(fields).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	requestRuleId := int64(789) // int64 | A unique integer value identifying the request rule.
	fields := "fields_example" // string | Comma-separated list of field names to include in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.RetrieveApplicationRequestRule(context.Background(), applicationId, requestRuleId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.RetrieveApplicationRequestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveApplicationRequestRule`: ResponseRetrieveApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.RetrieveApplicationRequestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**requestRuleId** | **int64** | A unique integer value identifying the request rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveApplicationRequestRuleRequest struct via the builder pattern


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


## UpdateApplicationRequestRule

> ResponseApplicationRequestPhaseRuleEngine UpdateApplicationRequestRule(ctx, applicationId, requestRuleId).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	requestRuleId := int64(789) // int64 | A unique integer value identifying the request rule.
	applicationRequestPhaseRuleEngineRequest := *openapiclient.NewApplicationRequestPhaseRuleEngineRequest("Name_example", [][]EdgeApplicationCriterionFieldRequest{[]openapiclient.EdgeApplicationCriterionFieldRequest{*openapiclient.NewEdgeApplicationCriterionFieldRequest("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{openapiclient.ApplicationRuleEngineRequestPhaseBehaviorsRequest{ApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest: openapiclient.NewApplicationRuleEngineRequestPhaseBehaviorsApplicationRuleEngineAddHeaderRequest("Type_example", *openapiclient.NewApplicationRuleEngineAddHeaderAttributesRequest("Value_example"))}}) // ApplicationRequestPhaseRuleEngineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.UpdateApplicationRequestRule(context.Background(), applicationId, requestRuleId).ApplicationRequestPhaseRuleEngineRequest(applicationRequestPhaseRuleEngineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.UpdateApplicationRequestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationRequestRule`: ResponseApplicationRequestPhaseRuleEngine
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.UpdateApplicationRequestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 
**requestRuleId** | **int64** | A unique integer value identifying the request rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequestRuleRequest struct via the builder pattern


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


## UpdateApplicationRequestRulesOrder

> PaginatedApplicationRequestPhaseRuleEngineList UpdateApplicationRequestRulesOrder(ctx, applicationId).ApplicationRequestPhaseRuleEngineOrderRequest(applicationRequestPhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()

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
	applicationId := int64(789) // int64 | A unique integer value identifying the application.
	applicationRequestPhaseRuleEngineOrderRequest := *openapiclient.NewApplicationRequestPhaseRuleEngineOrderRequest([]int64{int64(123)}) // ApplicationRequestPhaseRuleEngineOrderRequest | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (Valid fields: order) (optional)
	page := int64(789) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(789) // int64 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsRequestRulesAPI.UpdateApplicationRequestRulesOrder(context.Background(), applicationId).ApplicationRequestPhaseRuleEngineOrderRequest(applicationRequestPhaseRuleEngineOrderRequest).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsRequestRulesAPI.UpdateApplicationRequestRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationRequestRulesOrder`: PaginatedApplicationRequestPhaseRuleEngineList
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsRequestRulesAPI.UpdateApplicationRequestRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | A unique integer value identifying the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequestRulesOrderRequest struct via the builder pattern


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

