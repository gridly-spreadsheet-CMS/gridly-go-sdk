/*
Gridly API

Testing CdnApiService

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

func Test_gridly_CdnApiService(t *testing.T) {

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)

    t.Run("Test CdnApiService List", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CdnApi.List(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CdnApiService Publish", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var cdnId string

        resp, httpRes, err := apiClient.CdnApi.Publish(context.Background(), cdnId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CdnApiService UnPublish", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var cdnId string

        resp, httpRes, err := apiClient.CdnApi.UnPublish(context.Background(), cdnId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
