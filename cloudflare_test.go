package cloudflare

import (
	"fmt"
	"testing"

	cloudflare "github.com/aliask/libdns-cloudflare"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func TestSingleArg(t *testing.T) {
	fmt.Println("Testing single string argument (APIToken)... ")
	api_token := "abc123"
	config := fmt.Sprintf("cloudflare %s", api_token)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err != nil {
		t.Errorf("UnmarshalCaddyfile failed with %v", err)
		return
	}

	expected := api_token
	actual := p.Provider.APIToken
	if expected != actual {
		t.Errorf("Expected APIToken to be '%s' but got '%s'", expected, actual)
	}
}

func TestAPITokenInBlock(t *testing.T) {
	fmt.Println("Testing APIToken provided in block... ")
	api_token := "abc123"
	config := fmt.Sprintf(`cloudflare {
		api_token %s
	}`, api_token)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err != nil {
		t.Errorf("UnmarshalCaddyfile failed with %v", err)
		return
	}

	expected := api_token
	actual := p.Provider.APIToken
	if expected != actual {
		t.Errorf("Expected APIToken to be '%s' but got '%s'", expected, actual)
	}
}

func TestEmptyConfig(t *testing.T) {
	fmt.Println("Testing empty config fails to parse... ")
	config := "cloudflare"

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, zone_token = %s, dns_token = %s",
			p.Provider.APIToken,
			p.Provider.ZoneToken,
			p.Provider.DNSToken,
		)
	}
}

func TestZoneAndDNSTokens(t *testing.T) {
	fmt.Println("Testing separated Zone and DNS tokens... ")
	zone_token := "foo"
	dns_token := "bar"
	config := fmt.Sprintf(`
	cloudflare {
		zone_token %s
		dns_token %s
	}`, zone_token, dns_token)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err != nil {
		t.Errorf("UnmarshalCaddyfile failed with %v", err)
		return
	}

	expected := zone_token
	actual := p.Provider.ZoneToken
	if expected != actual {
		t.Errorf("Expected ZoneToken to be '%s' but got '%s'", expected, actual)
	}

	expected = dns_token
	actual = p.Provider.DNSToken
	if expected != actual {
		t.Errorf("Expected DNSToken to be '%s' but got '%s'", expected, actual)
	}

	expected = ""
	actual = p.Provider.APIToken
	if expected != actual {
		t.Errorf("Expected APIToken to be '%s' but got '%s'", expected, actual)
	}
}

func TestPartialConfig(t *testing.T) {
	fmt.Println("Testing partial config fails to parse... ")
	dns_token := "bar"
	config := fmt.Sprintf(`
	cloudflare {
		zone_token
		dns_token %s
	}`, dns_token)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, zone_token = %s, dns_token = %s",
			p.Provider.APIToken,
			p.Provider.ZoneToken,
			p.Provider.DNSToken,
		)
	}
}

func TestTooManyArgs(t *testing.T) {
	fmt.Println("Testing too many args... ")
	api_token := "foo"
	config := fmt.Sprintf("cloudflare %s with more", api_token)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, zone_token = %s, dns_token = %s",
			p.Provider.APIToken,
			p.Provider.ZoneToken,
			p.Provider.DNSToken,
		)
	}
}
