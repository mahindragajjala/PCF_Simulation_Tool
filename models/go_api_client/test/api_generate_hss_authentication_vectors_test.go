/*
NudmUEAU

Testing GenerateHSSAuthenticationVectorsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_GenerateHSSAuthenticationVectorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GenerateHSSAuthenticationVectorsAPIService GenerateAv", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var supi string
		var hssAuthType HssAuthTypeInUri

		resp, httpRes, err := apiClient.GenerateHSSAuthenticationVectorsAPI.GenerateAv(context.Background(), supi, hssAuthType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
