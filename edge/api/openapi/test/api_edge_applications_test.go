/*
edge-api

Testing EdgeApplicationsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package edge/api/openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_edge/api/openapi_EdgeApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsAPIService CloneEdgeApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.CloneEdgeApplication(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsAPIService CreateEdgeApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.CreateEdgeApplication(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsAPIService DestroyEdgeApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.DestroyEdgeApplication(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsAPIService ListEdgeApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.ListEdgeApplications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsAPIService PartialUpdateEdgeApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.PartialUpdateEdgeApplication(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsAPIService RetrieveEdgeApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.RetrieveEdgeApplication(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsAPIService UpdateEdgeApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.EdgeApplicationsAPI.UpdateEdgeApplication(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
