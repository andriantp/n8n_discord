package webhook

import (
	"fmt"

	"github.com/imroc/req/v3"
)

type Setting struct {
	Url   string
	Name  string
	Value string
}
type repo struct {
	setting Setting
}

type RepositoryI interface {
	POSTAskMe(payload map[string]interface{}) error
}

func NewRepo(setting Setting) RepositoryI {
	return repo{
		setting: setting,
	}
}

func (r repo) POSTAskMe(payload map[string]interface{}) error {
	url := fmt.Sprintf("%s/ask-me/onmessage", r.setting.Url)

	client := req.C()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader(r.setting.Name, r.setting.Value).
		SetBody(payload).
		Post(url)

	if err != nil {
		return fmt.Errorf("❌ Failed to send request:%w", err)
	}

	if !resp.IsSuccessState() {
		return fmt.Errorf("status code [%d][%s]", resp.StatusCode, resp.Status)
	}

	return nil
}
