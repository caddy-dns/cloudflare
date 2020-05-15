Cloudflare module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Cloudflare accounts.

## Caddy module name

```
dns.providers.cloudflare
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```
{
	"module": "acme",
	"dns": {
		"provider": {
			"name": "cloudflare",
			"api_token": "{env.CLOUDFLARE_API_TOKEN}"
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns cloudflare {env.CLOUDFLARE_API_TOKEN}
}
```

You can replace `{env.CLOUDFLARE_API_TOKEN}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/cloudflare) for important information about credentials.
