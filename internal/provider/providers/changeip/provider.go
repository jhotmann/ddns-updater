package changeip

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/netip"
	"net/url"

	"github.com/qdm12/ddns-updater/internal/models"
	"github.com/qdm12/ddns-updater/internal/provider/constants"
	"github.com/qdm12/ddns-updater/internal/provider/errors"
	"github.com/qdm12/ddns-updater/internal/provider/headers"
	"github.com/qdm12/ddns-updater/internal/provider/utils"
	"github.com/qdm12/ddns-updater/pkg/publicip/ipversion"
)

type Provider struct {
	domain        string
	host          string
	ipVersion     ipversion.IPVersion
	ipv6Suffix    netip.Prefix
	username      string
	password      string
	useProviderIP bool
}

type settings struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	UseProviderIP bool   `json:"provider_ip"`
}

func New(data json.RawMessage, domain, host string,
	ipVersion ipversion.IPVersion, ipv6Suffix netip.Prefix) (
	p *Provider, err error) {
	var providerSpecificSettings settings
	err = json.Unmarshal(data, &providerSpecificSettings)
	if err != nil {
		return nil, fmt.Errorf("json decoding provider specific settings: %w", err)
	}
	err = validateSettings(domain, host, providerSpecificSettings)
	if err != nil {
		return nil, fmt.Errorf("validating settings: %w", err)
	}
	return &Provider{
		domain:        domain,
		host:          host,
		ipVersion:     ipVersion,
		ipv6Suffix:    ipv6Suffix,
		username:      providerSpecificSettings.Username,
		password:      providerSpecificSettings.Password,
		useProviderIP: providerSpecificSettings.UseProviderIP,
	}, nil
}

func validateSettings(domain, host string, settings settings) error {
	switch {
	case domain == "":
		return fmt.Errorf("%w", errors.ErrDomainNotSet)
	case host == "":
		return fmt.Errorf("%w", errors.ErrHostNotSet)
	case settings.Username == "":
		return fmt.Errorf("%w", errors.ErrUsernameNotSet)
	case settings.Password == "":
		return fmt.Errorf("%w", errors.ErrPasswordNotSet)
	}
	return nil
}

func (p *Provider) String() string {
	return utils.ToString(p.domain, p.host, constants.Changeip, p.ipVersion)
}

func (p *Provider) Domain() string {
	return p.domain
}

func (p *Provider) Host() string {
	return p.host
}

func (p *Provider) IPVersion() ipversion.IPVersion {
	return p.ipVersion
}

func (p *Provider) IPv6Suffix() netip.Prefix {
	return p.ipv6Suffix
}

func (p *Provider) Proxied() bool {
	return false
}

func (p *Provider) BuildDomainName() string {
	return utils.BuildDomainName(p.host, p.domain)
}

func (p *Provider) HTML() models.HTMLRow {
	return models.HTMLRow{
		Domain:    fmt.Sprintf("<a href=\"http://%s\">%s</a>", p.BuildDomainName(), p.BuildDomainName()),
		Host:      p.Host(),
		Provider:  "<a href=\"https://www.changeip.com\">changeip.com</a>",
		IPVersion: p.ipVersion.String(),
	}
}

func (p *Provider) Update(ctx context.Context, client *http.Client, ip netip.Addr) (newIP netip.Addr, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   "nic.ChangeIP.com",
		Path:   "/nic/update",
	}
	values := url.Values{}
	values.Set("hostname", utils.BuildURLQueryHostname(p.host, p.domain))
	useProviderIP := p.useProviderIP && (ip.Is4() || !p.ipv6Suffix.IsValid())
	if !useProviderIP {
		values.Set("ip", ip.String())
	}
	u.RawQuery = values.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("creating http request: %w", err)
	}
	request.SetBasicAuth(p.username, p.password)
	headers.SetUserAgent(request)

	response, err := client.Do(request)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("doing http request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		return netip.Addr{}, fmt.Errorf("%w: %d: %s", errors.ErrHTTPStatusNotValid,
			response.StatusCode, utils.BodyToSingleLine(response.Body))
	}

	err = response.Body.Close()
	if err != nil {
		return netip.Addr{}, fmt.Errorf("closing response body: %w", err)
	}

	return ip, nil
}
