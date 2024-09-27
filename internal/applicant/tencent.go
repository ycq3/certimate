package applicant

import (
	"encoding/json"
	"github.com/usual2970/certimate/internal/domain"
	"os"

	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
)

type tencent struct {
	option *ApplyOption
}

func NewTencent(option *ApplyOption) Applicant {
	return &tencent{
		option: option,
	}
}

func (t *tencent) Apply() (*Certificate, error) {

	access := &domain.TencentAccess{}
	json.Unmarshal([]byte(t.option.Access), access)

	os.Setenv("TENCENTCLOUD_SECRET_ID", access.SecretId)
	os.Setenv("TENCENTCLOUD_SECRET_KEY", access.SecretKey)
	dnsProvider, err := tencentcloud.NewDNSProvider()
	if err != nil {
		return nil, err
	}

	return apply(t.option, dnsProvider)
}
