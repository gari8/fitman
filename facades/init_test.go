package facades

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getApiClientMock(t *testing.T) *MockApiClient {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockApiClient(ctrl)
}

func getFSClientMock(t *testing.T) *MockFSClient {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockFSClient(ctrl)
}

func setupMock(t *testing.T, setup func(mockApiClient *MockApiClient, mockFSClient *MockFSClient)) {
	mockApiClient, mockFSClient := getApiClientMock(t), getFSClientMock(t)
	setup(mockApiClient, mockFSClient)
	apiClient = mockApiClient
	fsClient = mockFSClient
}

func TestSetParams(t *testing.T) {
	for _, test := range []struct {
		Name string
		*Params
		context.Context
	}{
		{
			Name:    "OnlyIdToken is true / empty context",
			Context: context.Background(),
			Params:  &Params{OnlyIdToken: true},
		},
		{
			Name:    "OnlyIdToken is false / context is registered",
			Context: context.WithValue(context.Background(), paramsContextName, "empty"),
			Params:  &Params{OnlyIdToken: false},
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			ctx := SetParams(test.Context, test.Params)
			p, ok := ctx.Value(paramsContextName).(*Params)
			assert.Equal(t, ok, true)
			assert.Equal(t, test.Params, p)
		})
	}
}

func TestGetParams(t *testing.T) {
	for _, test := range []struct {
		Name string
		*Params
	}{
		{
			Name:   "empty context",
			Params: nil,
		},
		{
			Name:   "context is registered",
			Params: &Params{OnlyIdToken: false},
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), paramsContextName, test.Params)
			params, err := GetParams(ctx)
			assert.NoError(t, err)
			assert.Equal(t, test.Params, params)
		})
	}
}
