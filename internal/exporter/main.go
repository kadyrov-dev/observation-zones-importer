package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kadyrov-dev/observation-zones-importer/internal/geojsonbuilder"
	"github.com/kadyrov-dev/observation-zones-importer/internal/types"
)

type Exporter struct {
	outputDir string
}

func NewExporter(outputDir string) *Exporter {
	return &Exporter{outputDir: outputDir}
}

func (e *Exporter) Export(zones types.Zones) error {
	disabled := geojsonbuilder.Build(zones.Disabled)

	if err := e.saveZone("disabled", disabled); err != nil {
		return err
	}

	forbidden := geojsonbuilder.Build(zones.Forbidden)

	if err := e.saveZone("forbidden", forbidden); err != nil {
		return err
	}

	paidParking := geojsonbuilder.Build(zones.PaidParking)

	if err := e.saveZone("paid_parking", paidParking); err != nil {
		return err
	}

	return nil
}

func (e *Exporter) saveZone(zoneType string, zones types.FeatureCollection) error {
	path := fmt.Sprintf("%s/observation_zones_%s.json", e.outputDir, zoneType)

	content, err := json.Marshal(zones)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, content, 0600); err != nil {
		return err
	}

	return nil
}
