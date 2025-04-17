package cloudflare

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/cloudflare"
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
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, zone_token = %s",
			p.Provider.APIToken,
			p.Provider.ZoneToken,
		)
	}
}

func TestZoneAndAPITokens(t *testing.T) {
	fmt.Println("Testing separated Zone and DNS tokens... ")
	zone_token := "foo"
	api_token := "bar"
	config := fmt.Sprintf(`
	cloudflare {
		zone_token %s
		api_token %s
	}`, zone_token, api_token)

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

	expected = api_token
	actual = p.Provider.APIToken
	if expected != actual {
		t.Errorf("Expected APIToken to be '%s' but got '%s'", expected, actual)
	}
}

func TestPartialConfig(t *testing.T) {
	fmt.Println("Testing partial config fails to parse... ")
	zone_token := "bar"
	config := fmt.Sprintf(`
	cloudflare {
		zone_token %s
	}`, zone_token)

	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err == nil {
		t.Errorf(
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, zone_token = %s",
			p.Provider.APIToken,
			p.Provider.ZoneToken,
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
			"UnmarshalCaddyfile should have provided an error, but none was received. api_token = %s, zone_token = %s",
			p.Provider.APIToken,
			p.Provider.ZoneToken,
		)
	}
}

func TestInvalidTokens(t *testing.T) {
	badTokens := []string{
		"",
		" ",
		"{env.VAR}",
		`"Sqqty8-Vn0iOP29rvqYgwKz_xqGQ4y5JhuVL1-qU"`,
		"Sqqty8-Vn0iOP29rvqYgwKz_xqGQ4y5JhuVL1-qU-extra-characters-that-are-way-too-long",
		"abcdef",
	}

	for _, badToken := range badTokens {
		p := Provider{&cloudflare.Provider{APIToken: badToken}}
		if err := p.Provision(caddy.Context{}); err == nil {
			t.Errorf(
				"Expected token '%s' to fail provisioning, but it was accepted",
				badToken,
			)
		}
	}
}

func TestValidToken(t *testing.T) {
	goodToken := "Sqqty8-Vn0iOP29rvqYgwKz_xqGQ4y5JhuVL1-qU"
	config := fmt.Sprintf(`cloudflare %s`, goodToken)
	dispenser := caddyfile.NewTestDispenser(config)
	p := Provider{&cloudflare.Provider{}}

	err := p.UnmarshalCaddyfile(dispenser)
	if err != nil {
		t.Errorf(
			"Expected valid token '%s', but validation failed: %v",
			goodToken,
			err,
		)
	}
}
