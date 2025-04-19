// Copyright 2020 Matthew Holt
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudflare

import (
	"fmt"
	"regexp"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/cloudflare"
)

// cloudflareTokenRegexp matches Cloudflare tokens consisting of 35 to 50 alphanumeric characters, dashes, or underscores.
var cloudflareTokenRegexp = regexp.MustCompile(`^[A-Za-z0-9_-]{35,50}$`)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *cloudflare.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.cloudflare",
		New: func() caddy.Module { return &Provider{new(cloudflare.Provider)} },
	}
}

// Before using the provider config, resolve placeholders in the API token(s).
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIToken = caddy.NewReplacer().ReplaceAll(p.Provider.APIToken, "")
	p.Provider.ZoneToken = caddy.NewReplacer().ReplaceAll(p.Provider.ZoneToken, "")
	if !validCloudflareToken(p.Provider.APIToken) {
		return fmt.Errorf("API token '%s' appears invalid; ensure it's correctly entered and not wrapped in braces nor quotes", p.Provider.APIToken)
	}
	return nil
}

// validCloudflareToken validates if the provided token matches the expected Cloudflare token format using a regular expression.
func validCloudflareToken(token string) bool {
	return cloudflareTokenRegexp.MatchString(token)
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Three syntaxes supported:
//
// Seperate Zone/DNS tokens
//
//	cloudflare {
//	  api_token <api_token>     // Zone DNS write access - scoped to applicable Zone(s)
//	  zone_token <zone_token>   // Zone read access - all zones
//	}
//
//	Single API Token
//
//	cloudflare <api_token>      // Zone read access and Zone DNS write for all zones
//
//	Single API Token, alternative syntax
//
//	cloudflare {
//	  api_token <api_token>     // Zone read access and Zone DNS write for all zones
//	}
//
// Expansion of placeholders in the API token is left to the JSON config caddy.Provisioner (above).
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	d.Next() // consume directive name

	if d.NextArg() {
		p.Provider.APIToken = d.Val()
	} else {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if d.NextArg() {
					p.Provider.APIToken = d.Val()
				} else {
					return d.ArgErr()
				}
			case "zone_token":
				if d.NextArg() {
					p.Provider.ZoneToken = d.Val()
				} else {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if d.NextArg() {
		return d.Errf("unexpected argument '%s'", d.Val())
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
