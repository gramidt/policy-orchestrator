package admin

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hexa-org/policy-mapper/providers/openpolicyagent"
	"github.com/hexa-org/policy-mapper/sdk"
)

type opaProvider struct {
}

func (p opaProvider) detect(provider string) bool {
	return provider == "open_policy_agent" || provider == sdk.ProviderTypeOpa
}

func convertToName(c *openpolicyagent.Credentials) string {
	if c.GCP != nil {
		return c.GCP.BucketName
	}

	if c.AWS != nil {
		return c.AWS.BucketName
	}

	if c.GITHUB != nil {
		return c.GITHUB.Repo
	}

	return base64.StdEncoding.EncodeToString([]byte(c.BundleUrl))
}

func (p opaProvider) name(key []byte) (string, error) {

	var cred openpolicyagent.Credentials
	err := json.Unmarshal(key, &cred)
	if err != nil {
		return "", errors.New(fmt.Sprintf("unable to read key file: %s", err.Error()))
	}

	projectID := convertToName(&cred)
	return fmt.Sprintf("%s:open-policy-agent", projectID), nil
}
