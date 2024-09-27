package applicant

import (
	"encoding/json"
	"github.com/usual2970/certimate/app/domain"
	"os"

	cf "github.com/go-acme/lego/v4/providers/dns/cloudflare"
)

type cloudflare struct {
	option *ApplyOption
}

func NewCloudflare(option *ApplyOption) Applicant {
	return &cloudflare{
		option: option,
	}
}

func (c *cloudflare) Apply() (*Certificate, error) {
	access := &domain.CloudflareAccess{}
	json.Unmarshal([]byte(c.option.Access), access)

	os.Setenv("CLOUDFLARE_DNS_API_TOKEN", access.DnsApiToken)

	provider, err := cf.NewDNSProvider()
	if err != nil {
		return nil, err
	}

	return apply(c.option, provider)
}
