package authorizer

import (
	"fmt"
	"net/http"
	"net/url"
)

func (a *Authorizer) fetchLocation(cookie string, state []byte, challenge string) (string, error) {
	callbackURL := parkingAuthBaseURL +
		"/connect/authorize/callback?" +
		fmt.Sprintf(returnURL, clientID, string(state), challenge)

	req, err := http.NewRequest("GET", callbackURL, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", cookie)

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		return "", fmt.Errorf("expected 302 status code, got %d", resp.StatusCode)
	}

	location := resp.Header["Location"]

	return location[0], nil
}

func (a *Authorizer) parseCodeFromLocation(location string) (string, error) {
	u, err := url.Parse(location)

	if err != nil {
		return "", err
	}

	m, _ := url.ParseQuery(u.RawQuery)
	code := m["code"]

	return code[0], nil
}
