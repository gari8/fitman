package facades

import (
	"context"
	"fmt"
)

const (
	paramsContextName = "params"
)

type Params struct {
	Verbose bool
}

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
