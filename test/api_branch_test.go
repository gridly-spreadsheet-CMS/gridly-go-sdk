/*
Gridly API

Testing BranchApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gridly

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    gridly "./openapi"
)

func Test_gridly_BranchApiService(t *testing.T) {

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)

    t.Run("Test BranchApiService Create", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BranchApi.Create(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BranchApiService CreateDiffCheck", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BranchApi.CreateDiffCheck(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BranchApiService Delete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var branchId string

        resp, httpRes, err := apiClient.BranchApi.Delete(context.Background(), branchId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BranchApiService Get", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var branchId string

        resp, httpRes, err := apiClient.BranchApi.Get(context.Background(), branchId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BranchApiService GetDiffCheck", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var taskId string

        resp, httpRes, err := apiClient.BranchApi.GetDiffCheck(context.Background(), taskId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BranchApiService List", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BranchApi.List(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BranchApiService Merge", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var branchId string

        resp, httpRes, err := apiClient.BranchApi.Merge(context.Background(), branchId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
