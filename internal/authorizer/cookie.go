package authorizer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (a *Authorizer) fetchCookie(login, password string, state []byte, challenge string) (string, error) {
	authenticateURL := parkingAuthBaseURL + "/api/authenticate"

	type exp struct {
		UserName  string `json:"userName"`
		Password  string `json:"password"`
		ReturnURL string `json:"returnUrl"`
	}

	body := &exp{
		UserName:  login,
		Password:  password,
		ReturnURL: "/connect/authorize/callback?" + fmt.Sprintf(returnURL, clientID, string(state), challenge),
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(body)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", authenticateURL, buf)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("expected 200 status code, got %d", resp.StatusCode)
	}

	cookie := resp.Header["Set-Cookie"]

	if len(cookie) == 0 {
		return "", errors.New("empty cookie")
	}

	c := strings.Split(cookie[0], ";")[0]
	c += "; " + strings.Split(cookie[1], ";")[0]

	return c, nil
}
