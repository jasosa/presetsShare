package service

import (
	"github.com/jasosa/PresetsManagement/pkg/datastore"
)

// PresetsService represents the available operations in the Presets service
type PresetsService interface {
	ListPresets() []string
	//UploadPreset(preset preset.Info) error
	//ShowPreset(presetName string) preset.Info
}

type presets struct {
	ds datastore.PresetsDataStore
}

// NewPresetsService returns a new PresetsService
func NewPresetsService(ds datastore.PresetsDataStore) PresetsService {
	return &presets{
		ds: ds,
	}
}

func (p *presets) ListPresets() []string {
	return p.ds.GetPresetNames()
}
