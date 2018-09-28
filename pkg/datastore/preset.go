package datastore

import "github.com/jasosa/PresetsManagement/pkg/preset"

//PresetsDataStore represents a data store for presets
type PresetsDataStore interface {
	// SavePreset save preset information.
	// If a preset with the given name already exists information of the preset is updated.
	SavePreset(presetInfo preset.Info) error
	// LoadPreset loads the preset information corresponding to the given presetName.
	// If the preset is not found an ErrPresetNotFound is returned.
	LoadPreset(presetName string) (preset.Info, error)
	// GetPresetNames returns a list with the name of all presets
	GetPresetNames() []string
}
