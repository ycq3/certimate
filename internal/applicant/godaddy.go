package applicant

import (
	"encoding/json"
	"github.com/usual2970/certimate/internal/domain"
	"os"

	godaddyProvider "github.com/go-acme/lego/v4/providers/dns/godaddy"
)

type godaddy struct {
	option *ApplyOption
}

func NewGodaddy(option *ApplyOption) Applicant {
	return &godaddy{
		option: option,
	}
}

func (a *godaddy) Apply() (*Certificate, error) {

	access := &domain.GodaddyAccess{}
	json.Unmarshal([]byte(a.option.Access), access)

	os.Setenv("GODADDY_API_KEY", access.ApiKey)
	os.Setenv("GODADDY_API_SECRET", access.ApiSecret)

	dnsProvider, err := godaddyProvider.NewDNSProvider()
	if err != nil {
		return nil, err
	}

	return apply(a.option, dnsProvider)
}
