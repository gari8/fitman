package facades

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./mock.go

import (
	"context"
	"fmt"
	"github.com/gari8/fitman/modules"
)

const (
	paramsContextName = "params"
)

type (
	Params struct {
		OnlyIdToken bool
	}
	ApiClient interface {
		Refresh() (modules.RefreshResponse, error)
		GetTokenInfo() (modules.TokenInfo, error)
		SetApiClient(apiKey, refreshToken string)
	}
	FSClient interface {
		ReadConfig() (modules.ConfigResponse, error)
		SetConfig(requests []modules.ConfigRequest, override bool) error
		Dialogue(override bool) (modules.DialogueResponse, error)
		Confirm() error
	}
)

func SetParams(parent context.Context, params *Params) context.Context {
	return context.WithValue(parent, paramsContextName, params)
}

func GetParams(ctx context.Context) (*Params, error) {
	params, ok := ctx.Value(paramsContextName).(*Params)
	if !ok {
		return nil, fmt.Errorf("params is not found")
	}
	return params, nil
}

var (
	apiClient ApiClient
	fsClient  FSClient
)

func init() {
	fsClient = modules.NewFS()
	apiClient = modules.NewApiClient(modules.NewHttpClient())
}
