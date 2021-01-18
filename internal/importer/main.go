package importer

import (
	"net/http"

	"github.com/kadyrov-dev/observation-zones-importer/internal/client"
	"github.com/kadyrov-dev/observation-zones-importer/internal/coordinatesmatrix"
	"github.com/kadyrov-dev/observation-zones-importer/internal/types"
)

type Importer struct {
	accessToken string
	httpClient  *http.Client
}

func NewImporter(accessToken string, httpClient *http.Client) *Importer {
	return &Importer{accessToken: accessToken, httpClient: httpClient}
}

func (i *Importer) Import() (types.Zones, error) {
	var result types.Zones

	var (
		disabled    []types.Zone
		forbidden   []types.Zone
		paidParking []types.Zone
	)

	client := client.NewClient(i.accessToken, i.httpClient)

	matrix := coordinatesmatrix.Build()

	for _, rectangle := range matrix {
		zones, err := client.GetZones(rectangle)

		if err != nil {
			return result, err
		}

		for _, zone := range zones {
			if zone.IsProcessable() {
				switch t := zone.Type(); t {
				case "forbidden":
					forbidden = append(forbidden, zone)
				case "disabled":
					disabled = append(disabled, zone)
				case "paid_parking":
					paidParking = append(paidParking, zone)
				}
			}
		}
	}

	result = types.Zones{
		Disabled:    disabled,
		Forbidden:   forbidden,
		PaidParking: paidParking,
	}

	return result, nil
}
