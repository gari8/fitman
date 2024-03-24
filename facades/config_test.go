package facades

import (
	"github.com/gari8/fitman/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunConfig(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Response modules.ConfigResponse
		Err      error
	}{
		{
			Name: "should be no error",
			Response: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			Err: nil,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setupMock(t, func(mockApiClient *MockApiClient, mockFSClient *MockFSClient) {
				mockFSClient.EXPECT().ReadConfig().Return(test.Response, test.Err).AnyTimes()
			})
			err := RunConfig(nil, nil)
			assert.NoError(t, err)
		})
	}
}
