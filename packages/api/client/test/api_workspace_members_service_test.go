/*
@open-sauced/api.opensauced.pizza

Testing WorkspaceMembersServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"testing"

	openapiclient "github.com/open-sauced/go-api/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_client_WorkspaceMembersServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkspaceMembersServiceAPIService DeleteOneWorkspaceMemberForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var memberId string

		resp, httpRes, err := apiClient.WorkspaceMembersServiceAPI.DeleteOneWorkspaceMemberForUser(context.Background(), id, memberId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspaceMembersServiceAPIService DeleteWorkspaceMembersForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.WorkspaceMembersServiceAPI.DeleteWorkspaceMembersForUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspaceMembersServiceAPIService GetWorkspaceMembersForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.WorkspaceMembersServiceAPI.GetWorkspaceMembersForUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspaceMembersServiceAPIService UpdateOneWorkspaceMemberForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var memberId string

		resp, httpRes, err := apiClient.WorkspaceMembersServiceAPI.UpdateOneWorkspaceMemberForUser(context.Background(), id, memberId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspaceMembersServiceAPIService UpdateWorkspaceMembersForUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.WorkspaceMembersServiceAPI.UpdateWorkspaceMembersForUser(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
