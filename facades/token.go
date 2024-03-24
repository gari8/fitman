package facades

import (
	"encoding/json"
	"fmt"
	"github.com/gari8/fitman/modules"
	"github.com/spf13/cobra"
)

const defaultProfile = "default"

func RunGet(cmd *cobra.Command, args []string) error {
	profile := defaultProfile
	if len(args) > 0 {
		profile = args[0]
	}
	conf, err := fsClient.ReadConfig()
	if err != nil {
		return err
	}
	if !conf.Contains(profile) {
		return fmt.Errorf("invalid profile")
	}
	apiClient.SetApiClient(conf[profile].ApiKey, conf[profile].RefreshToken)
	refreshToken, err := apiClient.Refresh()
	if err != nil {
		return err
	}
	params, err := GetParams(cmd.Context())
	if err != nil {
		return err
	}
	if !params.OnlyIdToken {
		b, err := json.Marshal(refreshToken)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}
	fmt.Println(refreshToken.IdToken)
	return nil
}

func RunInit(cmd *cobra.Command, args []string) error {
	profile := defaultProfile
	if len(args) > 0 {
		profile = args[0]
	}
	resp, err := fsClient.Dialogue(true)
	if err != nil {
		return err
	}
	apiClient.SetApiClient(resp.ApiKey, resp.RefreshToken)

	if _, err := apiClient.GetTokenInfo(); err != nil {
		return err
	}
	refreshResp, err := apiClient.Refresh()
	if err != nil {
		return err
	}
	if err := fsClient.SetConfig([]modules.ConfigRequest{
		{
			Profile:      profile,
			ApiKey:       resp.ApiKey,
			RefreshToken: refreshResp.RefreshToken,
		},
	}, true); err != nil {
		return err
	}
	params, err := GetParams(cmd.Context())
	if err != nil {
		return err
	}
	if !params.OnlyIdToken {
		b, err := json.Marshal(refreshResp)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}
	fmt.Println(refreshResp.IdToken)
	return nil
}

func RunAdd(cmd *cobra.Command, args []string) error {
	profile := defaultProfile
	if len(args) > 0 {
		profile = args[0]
	}
	confResp, err := fsClient.ReadConfig()
	if err != nil {
		return err
	}
	if confResp.Contains(profile) {
		return fmt.Errorf("%s already exists", profile)
	}
	resp, err := fsClient.Dialogue(false)
	if err != nil {
		return err
	}
	apiClient.SetApiClient(resp.ApiKey, resp.RefreshToken)
	if _, err := apiClient.GetTokenInfo(); err != nil {
		return err
	}
	refreshResp, err := apiClient.Refresh()
	if err != nil {
		return err
	}
	if err := fsClient.SetConfig([]modules.ConfigRequest{
		{
			Profile:      profile,
			ApiKey:       resp.ApiKey,
			RefreshToken: refreshResp.RefreshToken,
		},
	}, false); err != nil {
		return err
	}
	params, err := GetParams(cmd.Context())
	if err != nil {
		return err
	}
	if !params.OnlyIdToken {
		b, err := json.Marshal(refreshResp)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}
	fmt.Println(refreshResp.IdToken)
	return nil
}

func RunList(cmd *cobra.Command, args []string) error {
	conf, err := fsClient.ReadConfig()
	if err != nil {
		return err
	}
	b, err := json.Marshal(conf.Keys())
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func RunDelete(cmd *cobra.Command, args []string) error {
	profile := defaultProfile
	if len(args) > 0 {
		profile = args[0]
	}
	conf, err := fsClient.ReadConfig()
	if err != nil {
		return err
	}
	if !conf.Contains(profile) {
		return fmt.Errorf("invalid profile")
	}
	if err := fsClient.Confirm(); err != nil {
		return err
	}
	delete(conf, profile)
	var requests []modules.ConfigRequest
	for k, v := range conf {
		requests = append(requests, modules.ConfigRequest{
			Profile:      k,
			ApiKey:       v.ApiKey,
			RefreshToken: v.RefreshToken,
		})
	}
	if err := fsClient.SetConfig(requests, true); err != nil {
		return err
	}
	fmt.Printf("...%s profile is deleted\n", profile)
	return nil
}
