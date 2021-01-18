package authorizer

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (a *Authorizer) getAccessToken(location, code, verifier string) (string, error) {
	tokenURL := parkingAuthBaseURL + "/connect/token"

	form := url.Values{
		"client_id":     {clientID},
		"code":          {code},
		"code_verifier": {verifier},
		"redirect_uri":  {originURL},
		"grant_type":    {"authorization_code"},
	}

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(form.Encode()))

	if err != nil {
		return "", err
	}

	req.Header.Set("Referer", location)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if token, ok := result["access_token"].(string); ok {
		return token, nil
	}

	return "", errors.New("field access_token was not found in the result")
}
