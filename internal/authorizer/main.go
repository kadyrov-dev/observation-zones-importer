package authorizer

import (
	"fmt"
	"net/http"

	"github.com/kadyrov-dev/observation-zones-importer/internal/secretsgenerator"
)

const clientID = "monitoring_old"
const parkingAuthBaseURL = "https://parking-auth.traffic-view.com"
const originURL = "https://parking.traffic-view.com"
const returnURL = "client_id=%s&" +
	"redirect_uri=https%%3A%%2F%%2Fparking.traffic-view.com&response_type=code&" +
	"scope=openid%%20profile%%20email%%20web_api%%20offline_access%%20roles&" +
	"state=%s&" +
	"code_challenge=%s&" +
	"code_challenge_method=S256"

type Authorizer struct {
	httpClient *http.Client
}

func NewAuthorizer(httpClient *http.Client) *Authorizer {
	return &Authorizer{httpClient: httpClient}
}

func (a *Authorizer) GetAccessToken(login, password string) (string, error) {
	state := secretsgenerator.GenerateState()
	verifier := secretsgenerator.GenerateVerifier(state)
	challenge := secretsgenerator.GenerateChallenge(verifier)

	cookie, err := a.fetchCookie(login, password, state, challenge)

	if err != nil {
		return "", fmt.Errorf("fetchCookie(): %v", err)
	}

	location, err := a.fetchLocation(cookie, state, challenge)

	if err != nil {
		return "", fmt.Errorf("fetchLocation(): %v", err)
	}

	code, err := a.parseCodeFromLocation(location)

	if err != nil {
		return "", fmt.Errorf("parseCodeFromLocation(): %v", err)
	}

	token, err := a.getAccessToken(location, code, verifier)

	if err != nil {
		return "", fmt.Errorf("getAccessToken(): %v", err)
	}

	return token, nil
}
