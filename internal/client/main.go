package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kadyrov-dev/observation-zones-importer/internal/types"
)

type Client struct {
	accessToken string
	httpClient  *http.Client
}

func NewClient(accessToken string, httpClient *http.Client) *Client {
	return &Client{accessToken: accessToken, httpClient: httpClient}
}

func (c *Client) GetZones(coord types.Coordinates) ([]types.Zone, error) {
	var zones []types.Zone

	zonesURL := "https://parking.traffic-view.com/Api/Segment/GetAllZones?" +
		"bbox=%f%%2C%f%%2C%f%%2C%f&showSigns=false&page=1&start=0&limit=9999999&" +
		"group=zoneName&dir=ASC&" +
		"sort=%%5B%%7B%%22property%%22%%3A%%22ZoneId%%22%%2C%%22direction%%22%%3A%%22ASC%%22%%7D%%5D"

	req, err := http.NewRequest("GET", fmt.Sprintf(zonesURL, coord.LeftX, coord.LeftY, coord.RightX, coord.RightY), nil)

	if err != nil {
		return zones, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return zones, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return zones, fmt.Errorf("expected 200 status code, got %d", resp.StatusCode)
	}

	type responseStruct struct {
		Response []types.Zone
	}

	var result responseStruct

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return zones, err
	}

	return result.Response, nil
}
