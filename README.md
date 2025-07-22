Cloudflare module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Cloudflare accounts.

## Caddy module name

```
dns.providers.cloudflare
```

## Configuration

This module gives the user two ways of configuring API tokens.

1. Separate Zone and DNS Tokens (Deprecated)
	- **Zone Token:** `Zone.Zone:Read` permission for the domain(s) you're managing with Caddy
	- **DNS Token:** `Zone.DNS:Edit` permission for the domain(s) you're managing with Caddy 
2. Single API Token (Recommended)
	- **API Token:** `Zone.Zone:Read` and `Zone.DNS:Edit` permissions for the domain(s) you're managing with Caddy 

**Note:** Deprecated separate tokens support is only there for backward compatibility and might be removed in a future version of this module.

### JSON Example

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "cloudflare",
				"api_token": "{env.CF_API_TOKEN}"
			}
		}
	}
}
```

### Caddyfile Examples

#### Dual-key approach

```Caddyfile
tls {
	dns cloudflare {
		zone_token {env.CF_ZONE_TOKEN}
		api_token {env.CF_API_TOKEN}
	}
}
```

#### Single-key approach

```Caddyfile
tls {
	dns cloudflare {env.CF_API_TOKEN}
}
```

You can replace the `{env.CF_*}` placeholders with the actual auth token if you prefer to put it directly in your config instead of an environment variable, however it is less secure.


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/cloudflare) for important information about credentials.

**NOTE**: If migrating from Caddy v1, you will need to change from using a Cloudflare API Key to a scoped API Token. Please see link above for more information.

## Troubleshooting

### Error: `Invalid request headers`

If providing your API token via an ENV var which is accidentally not set/available when running Caddy, you'll receive this error from Cloudflare.

Double check that Caddy has access to a valid CF API token.

### Error: `timed out waiting for record to fully propagate`

Some environments may have trouble querying the `_acme-challenge` TXT record from Cloudflare. Verify in the Cloudflare dashboard that the temporary record is being created.

If the record does exist, your DNS resolver may be caching an earlier response before the record is valid. You can instead configure Caddy to use an alternative DNS resolver such as [Cloudflare's official `1.1.1.1`](https://www.cloudflare.com/en-gb/learning/dns/what-is-1.1.1.1/).

Add a custom `resolver` to the [`tls` directive](https://caddyserver.com/docs/caddyfile/directives/tls):

```
tls {
  dns cloudflare {env.CF_API_TOKEN}
  resolvers 1.1.1.1
}
```

Or with Caddy JSON to the `acme` module: [`challenges.dns.provider.resolvers: ["1.1.1.1"]`](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/challenges/dns/resolvers/).
