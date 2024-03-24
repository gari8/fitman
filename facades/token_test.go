package facades

import (
	"context"
	"github.com/gari8/fitman/modules"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunGet(t *testing.T) {
	for _, test := range []struct {
		Name string
		*Params
		modules.ConfigResponse
		ConfigErr error
		modules.RefreshResponse
		RefreshErr error
	}{
		{
			Name:   "should be no error",
			Params: &Params{OnlyIdToken: true},
			ConfigResponse: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			ConfigErr: nil,
			RefreshResponse: modules.RefreshResponse{
				AccessToken:  "test",
				ExpiresIn:    3600,
				TokenType:    "test",
				RefreshToken: "test",
				IdToken:      "test",
				UserId:       "test",
				ProjectId:    "test",
			},
			RefreshErr: nil,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setupMock(t, func(mockApiClient *MockApiClient, mockFSClient *MockFSClient) {
				mockFSClient.EXPECT().ReadConfig().Return(test.ConfigResponse, test.ConfigErr).AnyTimes()
				mockApiClient.EXPECT().SetApiClient(gomock.Any(), gomock.Any()).Times(1)
				mockApiClient.EXPECT().Refresh().Return(test.RefreshResponse, test.RefreshErr).AnyTimes()
			})
			cmd := &cobra.Command{}
			cmd.SetContext(SetParams(context.Background(), test.Params))
			err := RunGet(cmd, nil)
			assert.NoError(t, err)
		})
	}
}

func TestRunInit(t *testing.T) {
	for _, test := range []struct {
		Name string
		*Params
		modules.RefreshResponse
		RefreshErr error
		modules.DialogueResponse
		DialogueErr  error
		TokenInfoErr error
		SetConfigErr error
	}{
		{
			Name:   "should be no error",
			Params: &Params{OnlyIdToken: true},
			RefreshResponse: modules.RefreshResponse{
				AccessToken:  "test",
				ExpiresIn:    3600,
				TokenType:    "test",
				RefreshToken: "test",
				IdToken:      "test",
				UserId:       "test",
				ProjectId:    "test",
			},
			RefreshErr: nil,
			DialogueResponse: modules.DialogueResponse{
				ApiKey:       "test",
				RefreshToken: "test",
			},
			DialogueErr:  nil,
			TokenInfoErr: nil,
			SetConfigErr: nil,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setupMock(t, func(mockApiClient *MockApiClient, mockFSClient *MockFSClient) {
				mockFSClient.EXPECT().Dialogue(gomock.Any()).Return(test.DialogueResponse, test.DialogueErr).AnyTimes()
				mockFSClient.EXPECT().SetConfig(gomock.Any(), gomock.Any()).Return(test.SetConfigErr).AnyTimes()
				mockApiClient.EXPECT().SetApiClient(gomock.Any(), gomock.Any()).Times(1)
				mockApiClient.EXPECT().GetTokenInfo().Return(modules.TokenInfo{}, test.TokenInfoErr).AnyTimes()
				mockApiClient.EXPECT().Refresh().Return(test.RefreshResponse, test.RefreshErr).AnyTimes()
			})
			cmd := &cobra.Command{}
			cmd.SetContext(SetParams(context.Background(), test.Params))
			err := RunInit(cmd, nil)
			assert.NoError(t, err)
		})
	}
}

func TestRunAdd(t *testing.T) {
	for _, test := range []struct {
		Name string
		Fail bool
		*Params
		modules.ConfigResponse
		ConfigErr error
		modules.RefreshResponse
		RefreshErr error
		modules.DialogueResponse
		DialogueErr  error
		TokenInfoErr error
		SetConfigErr error
		Args         []string
	}{
		{
			Name:   "should be no error",
			Fail:   false,
			Params: &Params{OnlyIdToken: true},
			ConfigResponse: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			ConfigErr: nil,
			RefreshResponse: modules.RefreshResponse{
				AccessToken:  "test",
				ExpiresIn:    3600,
				TokenType:    "test",
				RefreshToken: "test",
				IdToken:      "test",
				UserId:       "test",
				ProjectId:    "test",
			},
			RefreshErr: nil,
			DialogueResponse: modules.DialogueResponse{
				ApiKey:       "test",
				RefreshToken: "test",
			},
			DialogueErr:  nil,
			TokenInfoErr: nil,
			SetConfigErr: nil,
			Args:         []string{"test1"},
		},
		{
			Name:   "should be error / default profile is already exists",
			Fail:   true,
			Params: &Params{OnlyIdToken: true},
			ConfigResponse: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			ConfigErr: nil,
			RefreshResponse: modules.RefreshResponse{
				AccessToken:  "test",
				ExpiresIn:    3600,
				TokenType:    "test",
				RefreshToken: "test",
				IdToken:      "test",
				UserId:       "test",
				ProjectId:    "test",
			},
			RefreshErr: nil,
			DialogueResponse: modules.DialogueResponse{
				ApiKey:       "test",
				RefreshToken: "test",
			},
			DialogueErr:  nil,
			TokenInfoErr: nil,
			SetConfigErr: nil,
			Args:         nil,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setupMock(t, func(mockApiClient *MockApiClient, mockFSClient *MockFSClient) {
				mockFSClient.EXPECT().Dialogue(gomock.Any()).Return(test.DialogueResponse, test.DialogueErr).AnyTimes()
				mockFSClient.EXPECT().ReadConfig().Return(test.ConfigResponse, test.ConfigErr).AnyTimes()
				mockFSClient.EXPECT().SetConfig(gomock.Any(), gomock.Any()).Return(test.SetConfigErr).AnyTimes()
				mockApiClient.EXPECT().SetApiClient(gomock.Any(), gomock.Any()).Times(1)
				mockApiClient.EXPECT().GetTokenInfo().Return(modules.TokenInfo{}, test.TokenInfoErr).AnyTimes()
				mockApiClient.EXPECT().Refresh().Return(test.RefreshResponse, test.RefreshErr).AnyTimes()
			})
			cmd := &cobra.Command{}
			cmd.SetContext(SetParams(context.Background(), test.Params))
			err := RunAdd(cmd, test.Args)
			if test.Fail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRunList(t *testing.T) {
	for _, test := range []struct {
		Name string
		Fail bool
		*Params
		modules.ConfigResponse
		ConfigErr error
	}{
		{
			Name:   "should be no error",
			Fail:   true,
			Params: &Params{OnlyIdToken: true},
			ConfigResponse: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			ConfigErr: nil,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setupMock(t, func(mockApiClient *MockApiClient, mockFSClient *MockFSClient) {
				mockFSClient.EXPECT().ReadConfig().Return(test.ConfigResponse, test.ConfigErr).AnyTimes()
			})
			cmd := &cobra.Command{}
			cmd.SetContext(SetParams(context.Background(), test.Params))
			err := RunList(cmd, nil)
			assert.NoError(t, err)
		})
	}
}

func TestRunDelete(t *testing.T) {
	for _, test := range []struct {
		Name string
		Fail bool
		*Params
		modules.ConfigResponse
		ConfigErr    error
		SetConfigErr error
		Args         []string
		ConfirmErr   error
	}{
		{
			Name:   "should be no error",
			Fail:   false,
			Params: &Params{OnlyIdToken: true},
			ConfigResponse: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			ConfigErr:    nil,
			SetConfigErr: nil,
			Args:         []string{"default"},
			ConfirmErr:   nil,
		},
		{
			Name:   "should be error / invalid profile",
			Fail:   true,
			Params: &Params{OnlyIdToken: true},
			ConfigResponse: modules.ConfigResponse{
				"default": {
					ApiKey:       "test",
					RefreshToken: "test",
				},
			},
			ConfigErr:    nil,
			SetConfigErr: nil,
			Args:         []string{"invalid"},
			ConfirmErr:   nil,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			setupMock(t, func(mockApiClient *MockApiClient, mockFSClient *MockFSClient) {
				mockFSClient.EXPECT().ReadConfig().Return(test.ConfigResponse, test.ConfigErr).AnyTimes()
				mockFSClient.EXPECT().SetConfig(gomock.Any(), gomock.Any()).Return(test.SetConfigErr).AnyTimes()
				mockFSClient.EXPECT().Confirm().Return(test.ConfirmErr).AnyTimes()
			})
			cmd := &cobra.Command{}
			cmd.SetContext(SetParams(context.Background(), test.Params))
			err := RunDelete(cmd, test.Args)
			if test.Fail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
