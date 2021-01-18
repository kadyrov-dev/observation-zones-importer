package geojsonbuilder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kadyrov-dev/observation-zones-importer/internal/geojsonbuilder"
	"github.com/kadyrov-dev/observation-zones-importer/internal/types"
)

func Test_Build_returnsGeoJSON(t *testing.T) {
	zone := types.NewZone(
		42,
		"ул. Генерала Иванова",
		0,
		false,
		[]types.Point{
			{
				X: 56.01053301310651,
				Y: 37.16082706530175,
			},
			{
				X: 56.01038381620671,
				Y: 37.16117038805566,
			},
		},
	)

	zones := Build([]types.Zone{*zone})

	coordinates := make([][][]float64, 1)
	coordinates[0] = [][]float64{
		{37.16082706530175, 56.01053301310651},
		{37.16117038805566, 56.01038381620671},
	}

	expected := types.FeatureCollection{
		Type: "FeatureCollection",
		Features: []types.Feature{
			{
				Type: "Feature",
				Properties: types.Properties{
					ZoneID:     42,
					ZoneTypeID: 0,
					Address:    "ул. Генерала Иванова",
				},
				Geometry: types.Polygon{
					Type:        "Polygon",
					Coordinates: coordinates,
				},
			},
		},
	}

	assert.EqualValues(t, zones, expected)
}
