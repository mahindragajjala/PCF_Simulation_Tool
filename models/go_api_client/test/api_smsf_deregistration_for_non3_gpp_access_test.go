/*
Nudm_UECM

Testing SMSFDeregistrationForNon3GPPAccessAPIService

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

func Test_openapi_SMSFDeregistrationForNon3GPPAccessAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SMSFDeregistrationForNon3GPPAccessAPIService Non3GppSmsfDeregistration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ueId string

		httpRes, err := apiClient.SMSFDeregistrationForNon3GPPAccessAPI.Non3GppSmsfDeregistration(context.Background(), ueId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
