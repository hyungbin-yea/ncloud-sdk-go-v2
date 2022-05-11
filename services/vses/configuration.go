/*
 * api
 *
 * Search Engine Service(VPC) Open API<br/>https://vpcsearchengine.apigw.ntruss.com/api/v1
 *
 * API version: 2022-04-21T09:07:51Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vses

import (
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"os"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

func NewConfiguration(region string, apiKeys ...*ncloud.APIKey) *ncloud.Configuration {
	cfg := &ncloud.Configuration{
		BasePath:      "https://vpcsearchengine.apigw.ntruss.com/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "vses//go",
	}
	if len(apiKeys) > 0 {
		cfg.APIKey = apiKeys[0]
	}
	cfg.InitCredentials()

	var ncloudApiGw string
	if os.Getenv("NCLOUD_API_GW") == "" {
		ncloudApiGw = "https://vpcsearchengine.apigw.ntruss.com"
	} else {
		ncloudApiGw = os.Getenv("NCLOUD_API_GW")
	}

	ncloudApiGw = strings.Replace(ncloudApiGw, "https://ncloud.", "https://vpcsearchengine.", 1)
	ncloudApiGw = strings.Replace(ncloudApiGw, "https://fin-ncloud.", "https://vpcsearchengine.", 1)

	return cfg
}
