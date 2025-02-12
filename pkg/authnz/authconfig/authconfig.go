package authconfig

import (
	"github.com/place1/wg-access-server/pkg/authnz/authruntime"
)

type AuthConfig struct {
	OIDC   *OIDCConfig      `yaml:"oidc"`
	Gitlab *GitlabConfig    `yaml:"gitlab"`
	Basic  *BasicAuthConfig `yaml:"basic"`
	Proxy  *ProxyAuthConfig `yaml:"proxy"`
}

func (c *AuthConfig) IsEnabled() bool {
	return c.OIDC != nil || c.Gitlab != nil || c.Basic != nil || c.Proxy != nil
}

func (c *AuthConfig) Providers() []*authruntime.Provider {
	providers := []*authruntime.Provider{}

	if c.OIDC != nil {
		providers = append(providers, c.OIDC.Provider())
	}

	if c.Gitlab != nil {
		providers = append(providers, c.Gitlab.Provider())
	}

	if c.Basic != nil {
		providers = append(providers, c.Basic.Provider())
	}

	if c.Proxy != nil {
		providers = append(providers, c.Proxy.Provider())
	}

	return providers
}
