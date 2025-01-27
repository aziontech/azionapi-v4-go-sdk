/*
edge-api

Testing WorkloadsAPIService

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

func Test_edge_WorkloadsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkloadsAPIService CreateWorkload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkloadsAPI.CreateWorkload(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkloadsAPIService DestroyWorkload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.WorkloadsAPI.DestroyWorkload(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkloadsAPIService ListWorkloads", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkloadsAPI.ListWorkloads(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkloadsAPIService PartialUpdateWorkload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.WorkloadsAPI.PartialUpdateWorkload(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkloadsAPIService RetrieveWorkload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.WorkloadsAPI.RetrieveWorkload(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkloadsAPIService UpdateWorkload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalId string

		resp, httpRes, err := apiClient.WorkloadsAPI.UpdateWorkload(context.Background(), globalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
