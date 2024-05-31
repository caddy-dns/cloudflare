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
	cloudflare "github.com/aliask/libdns-cloudflare"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

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
	p.Provider.DNSToken = caddy.NewReplacer().ReplaceAll(p.Provider.DNSToken, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Three syntaxes supported:
//
// Seperate Zone/DNS tokens
//
//	cloudflare {
//	  zone_token <zone_token>   // Zone read access - all zones
//	  dns_token <dns_token>     // Zone DNS write access - scoped to applicable Zone(s)
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
			case "dns_token":
				if d.NextArg() {
					p.Provider.DNSToken = d.Val()
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
	if p.Provider.DNSToken != "" || p.Provider.ZoneToken != "" {
		if p.Provider.ZoneToken == "" {
			return d.Err("dns_token provided but no zone_token found")
		}
		if p.Provider.DNSToken == "" {
			return d.Err("zone_token provided but no dns_token found")
		}
	} else {
		if p.Provider.APIToken == "" {
			return d.Err("missing API tokens")
		}
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
