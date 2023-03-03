package config

import (
	"github.com/aziontech/azionapi-go-sdk/domains"
	"github.com/aziontech/azionapi-go-sdk/idns"
)

type Client struct {
	domain *domains.APIClient
	idns   *idns.APIClient
}

type ApiConfig struct {
	domain *domains.Configuration
	idns   *idns.Configuration
}

type AzionApiConfig interface {
	SetAuthorizationHeader(string) bool
	SetAcceptHeader(string) bool
}

func (ac *ApiConfig) SetAuthorizationHeader(token string) bool {
	ac.domain.AddDefaultHeader("Authorization", "token "+token)
	ac.idns.AddDefaultHeader("Authorization", "token "+token)
	return true
}

func (ac *ApiConfig) SetAcceptHeader(content string) bool {
	switch content_type := content; {
	case content_type == "json":
		ac.domain.AddDefaultHeader("Accept", "application/json;version=3")
		ac.idns.AddDefaultHeader("Accept", "application/json;version=3")
	}

	return true
}

func InitConfig(config AzionApiConfig, token string, content string) bool {
	config.SetAuthorizationHeader(token)
	config.SetAcceptHeader(content)
	return true
}
