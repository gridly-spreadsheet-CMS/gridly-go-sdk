/*
Gridly API

Testing DatabaseApiService

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

func Test_gridly_DatabaseApiService(t *testing.T) {

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)

    t.Run("Test DatabaseApiService Create", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DatabaseApi.Create(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseApiService Delete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var dbId string

        resp, httpRes, err := apiClient.DatabaseApi.Delete(context.Background(), dbId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseApiService Duplicate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var dbId string

        resp, httpRes, err := apiClient.DatabaseApi.Duplicate(context.Background(), dbId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseApiService Get", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var dbId string

        resp, httpRes, err := apiClient.DatabaseApi.Get(context.Background(), dbId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseApiService List", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DatabaseApi.List(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DatabaseApiService Update", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var dbId string

        resp, httpRes, err := apiClient.DatabaseApi.Update(context.Background(), dbId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}