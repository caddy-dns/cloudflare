package cloudflare

import (
	"github.com/caddyserver/caddy/v2"
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
		New: func() caddy.Module { return new(Provider) },
	}
}
