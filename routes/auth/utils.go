package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/archproj/slackoverflow/config"
)

// Unmarshal User Slack Authentication response body
type VerificationResponse struct {
	Code string `json:"code"`
}

const (
	// Easy assembly of URL, permanent endpoint
	baseURL = `https://www.slack.com/oauth/authorize?`
)

var (
	// Scopes to run app required at time of authentication
	scopes = [...]string{
		"channels:read",
		"groups:read",
		"incoming-webhook",
		"commands",
	}
)

// Assembles secerts to generate OAuth entrypoint URL
func GenerateOAuthURL(cfg *config.Variables) (string, error) {
	// TODO: error if params are not valid
	params := url.Values{}

	params.Set("client_id", cfg.SlackClientID)

	str := ""
	for _, s := range scopes {
		s += `+`
		str += s
	}

	fmt.Println(str)

	params.Set("scope", url.PathEscape(str))

	fmt.Println(str)

	return baseURL + params.Encode(), nil
}

func ParseOAuthVerCode(req *http.Request) (*string, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var v VerificationResponse

	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	return &v.Code, nil
}
