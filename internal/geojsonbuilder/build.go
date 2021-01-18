package geojsonbuilder

import (
	"github.com/kadyrov-dev/observation-zones-importer/internal/types"
)

func Build(zones []types.Zone) types.FeatureCollection {
	features := make([]types.Feature, 0)

	for _, z := range zones {
		features = append(features, buildFeature(z))
	}

	return types.FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
}

func buildFeature(z types.Zone) types.Feature {
	return types.Feature{
		Type: "Feature",
		Properties: types.Properties{
			ZoneID:     z.ID,
			ZoneTypeID: z.TypeID,
			Address:    z.Address,
		},
		Geometry: types.Polygon{
			Type:        "Polygon",
			Coordinates: z.BuildCoordinates(),
		},
	}
}
