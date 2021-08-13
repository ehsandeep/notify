package teams

import (
	"fmt"
	"strings"

	"github.com/containrrr/shoutrrr"
	"github.com/projectdiscovery/notify/pkg/utils"
)

type Provider struct {
	Teams []*Options `yaml:"teams,omitempty"`
}

type Options struct {
	ID              string `yaml:"id,omitempty"`
	TeamsWebHookURL string `yaml:"teams_webhook_url,omitempty"`
	TeamsFormat     string `yaml:"teams_format,omitempty"`
}

func New(options []*Options, ids []string) (*Provider, error) {
	provider := &Provider{}

	for _, o := range options {
		if len(ids) == 0 || utils.Contains(ids, o.ID) {
			provider.Teams = append(provider.Teams, o)
		}
	}

	return provider, nil
}

func (p *Provider) Send(message, CliFormat string) error {

	for _, pr := range p.Teams {
		msg := utils.FormatMessage(message, utils.SelectFormat(CliFormat, pr.TeamsFormat))

		teamsTokens := strings.TrimPrefix(pr.TeamsWebHookURL, "https://outlook.office.com/webhook/")
		teamsTokens = strings.ReplaceAll(teamsTokens, "IncomingWebhook/", "")
		url := fmt.Sprintf("teams://%s", teamsTokens)
		err := shoutrrr.Send(url, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
