/*
edge-api

Testing EdgeApplicationsErrorResponsesAPIService

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

func Test_edge_EdgeApplicationsErrorResponsesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsErrorResponsesAPIService ListErrorResponses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string

		resp, httpRes, err := apiClient.EdgeApplicationsErrorResponsesAPI.ListErrorResponses(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsErrorResponsesAPIService PartialUpdateErrorResponse", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsErrorResponsesAPI.PartialUpdateErrorResponse(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsErrorResponsesAPIService RetrieveErrorResponse", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsErrorResponsesAPI.RetrieveErrorResponse(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsErrorResponsesAPIService UpdateErrorResponse", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsErrorResponsesAPI.UpdateErrorResponse(context.Background(), edgeApplicationId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
