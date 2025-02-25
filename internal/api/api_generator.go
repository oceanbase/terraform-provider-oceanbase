/*
 * Copyright 2025 OceanBase
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"net/http"
	"os"

	"github.com/icholy/digest"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

const (
	HeaderProjectId string = "X-Ob-Project-Id"
)

const (
	SiteMock   string = "MOCK"
	SiteChina  string = "CHINA"
	SiteGlobal string = "GLOBAL"
)

const (
	APIHostMock   string = "127.0.0.1:8080"
	APIHostChina  string = "api-cloud-cn.oceanbase.com"
	APIHostGlobal string = "g-api-cloud.oceanbase.com"
)

const (
	SchemeHttps string = "https"
	SchemeHttp  string = "http"
)

const (
	EnvAPIHostChina  string = "TF_OCEANBASE_API_HOST_CHINA"
	EnvAPIHostGlobal string = "TF_OCEANBASE_API_HOST_GLOBAL"
)

type Generator struct {
	Site      string
	AccessKey string
	SecretKey string
}

func createDigestClient(username, password string) *http.Client {
	t := &digest.Transport{
		Username: username,
		Password: password,
	}
	return &http.Client{Transport: t}
}

func getAPIServerSchemeHost(site string) (string, string) {
	switch site {
	case SiteMock:
		return SchemeHttp, APIHostMock
	case SiteChina:
		apiHostChina := APIHostChina
		envAPIHostChina := os.Getenv(EnvAPIHostChina)
		if envAPIHostChina != "" {
			apiHostChina = envAPIHostChina
		}
		return SchemeHttps, apiHostChina
	case SiteGlobal:
		apiHostGlobal := APIHostGlobal
		envAPIHostGlobal := os.Getenv(EnvAPIHostGlobal)
		if envAPIHostGlobal != "" {
			apiHostGlobal = envAPIHostGlobal
		}
		return SchemeHttps, apiHostGlobal
	default:
		return "", ""
	}
}

func (g *Generator) GetClient() *obcloudsdk.APIClient {
	cfg := obcloudsdk.NewConfiguration()
	cfg.Scheme, cfg.Host = getAPIServerSchemeHost(g.Site)
	cfg.HTTPClient = createDigestClient(g.AccessKey, g.SecretKey)
	return obcloudsdk.NewAPIClient(cfg)
}
