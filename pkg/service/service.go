package service

import (
	"github.com/jasosa/PresetsManagement/pkg/datastore"
	"github.com/jasosa/PresetsManagement/pkg/preset"
)

// PresetsService represents the available operations in the Presets service
type PresetsService interface {
	ListPresets() []string
	//UploadPreset(preset preset.Info) error
	ShowPreset(presetName string) (preset.Info, error)
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

func (p *presets) ShowPreset(presetName string) (preset.Info, error) {
	pi, err := p.ds.LoadPreset(presetName)
	if err != nil {
		if err == datastore.ErrPresetNotFound {
			return pi, newErrPresetNotFound(err.Error())
		}

		return pi, newErrUnknown(err.Error())
	}

	return pi, err
}
