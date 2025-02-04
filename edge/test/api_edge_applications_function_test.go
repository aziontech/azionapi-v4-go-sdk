/*
edge-api

Testing EdgeApplicationsFunctionAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package edge

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_edge_EdgeApplicationsFunctionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsFunctionAPIService CreateEdgeFirewallFunctionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string

		resp, httpRes, err := apiClient.EdgeApplicationsFunctionAPI.CreateEdgeFirewallFunctionInstance(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsFunctionAPIService DestroyEdgeApplicationFunctionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsFunctionAPI.DestroyEdgeApplicationFunctionInstance(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsFunctionAPIService ListEdgeApplicationFunctionInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string

		resp, httpRes, err := apiClient.EdgeApplicationsFunctionAPI.ListEdgeApplicationFunctionInstances(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsFunctionAPIService PartialUpdateEdgeApplicationFunctionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsFunctionAPI.PartialUpdateEdgeApplicationFunctionInstance(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsFunctionAPIService RetrieveFunctionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsFunctionAPI.RetrieveFunctionInstance(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsFunctionAPIService UpdateEdgeApplicationFunctionInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsFunctionAPI.UpdateEdgeApplicationFunctionInstance(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
