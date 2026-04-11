package cloudflare

import (
	"fmt"
	"strings"
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
		// cfut_/cfat_ with suffix >256 (too long for new; too long for legacy)
		"cfut_" + strings.Repeat("a", 257),
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

func TestValidCloudflareTokenFormats(t *testing.T) {
	validAPITokens := []string{
		// legacy
		"Sqqty8-Vn0iOP29rvqYgwKz_xqGQ4y5JhuVL1-qU",
		// user token (cfut_)
		"cfut_6YNSv6zUiWehDxZfdh8vNf3orQJzA6rrcDFJhql65e203820",
		// account token (cfat_)
		"cfat_3Sak5MsJJAsXNHCEmRufjbkbrdcxrhZFofmAe1pN435b71a9",
		// minimum-length new-format suffix (32 chars)
		"cfut_12345678901234567890123456789012",
		"cfat_12345678901234567890123456789012",
		// same shape as cfut_/cfat_ but 35–50 chars total → accepted as legacy
		"cfut_123456789012345678901234567890",
		"cfxt_12345678901234567890123456789012",
	}

	for _, tok := range validAPITokens {
		t.Run(tok[:min(12, len(tok))]+"...", func(t *testing.T) {
			p := Provider{&cloudflare.Provider{APIToken: tok}}
			if err := p.Provision(caddy.Context{}); err != nil {
				t.Errorf("Provision: %v", err)
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestProvisionErrorRedactsToken(t *testing.T) {
	// A token that contains characters outside [A-Za-z0-9_-] so it fails
	// both regexes, but is long enough (>12 chars) to exercise redaction.
	badToken := "this_is_not_a_valid_token!!"
	p := Provider{&cloudflare.Provider{APIToken: badToken}}

	err := p.Provision(caddy.Context{})
	if err == nil {
		t.Fatal("expected Provision to fail for an invalid token")
	}

	errMsg := err.Error()

	// The full token must NOT appear anywhere in the error message.
	if strings.Contains(errMsg, badToken) {
		t.Errorf("error message contains the full unredacted token: %s", errMsg)
	}

	// The redacted form should be present (first 8 + last 4 visible).
	redacted := redactToken(badToken)
	if !strings.Contains(errMsg, redacted) {
		t.Errorf("error message does not contain the redacted token %q: %s", redacted, errMsg)
	}
}

func TestRedactToken(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{"short token (<= 12 chars)", "abc", "***"},
		{"exactly 12 chars", "123456789012", "************"},
		{"13 chars shows first 8 + last 4", "1234567890123", "12345678*0123"},
		{"legacy token", "Sqqty8-Vn0iOP29rvqYgwKz_xqGQ4y5JhuVL1-qU", "Sqqty8-V****************************1-qU"},
		{"cfat_ token", "cfat_THIS_IS_A_FAKE_TOKEN_FOR_TESTING", "cfat_THI*************************TING"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := redactToken(tt.token)
			if got != tt.expected {
				t.Errorf("redactToken(%q) = %q, want %q", tt.token, got, tt.expected)
			}
			// Regardless of expected value, the full token must not survive redaction
			// (unless the token is <= 12 chars, in which case it's fully masked).
			if len(tt.token) > 12 && strings.Contains(got, tt.token) {
				t.Errorf("redactToken(%q) still contains the full token", tt.token)
			}
		})
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
