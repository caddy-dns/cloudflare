package cloudflare

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/cloudflare"
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

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// cloudflare [<api_token>] {
//     api_token <api_token>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	repl := caddy.NewReplacer()
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = repl.ReplaceAll(d.Val(), "")
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				p.Provider.APIToken = repl.ReplaceAll(d.Val(), "")
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guard
var _ caddyfile.Unmarshaler = (*Provider)(nil)
