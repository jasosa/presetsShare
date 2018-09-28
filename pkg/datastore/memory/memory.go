package memory

import (
	"github.com/jasosa/PresetsManagement/pkg/datastore"
	"github.com/jasosa/PresetsManagement/pkg/preset"
)

type memoryDataStore struct {
	presets map[string]preset.Info
}

//NewDataStore returns a new PresetsDatastore
func NewDataStore() datastore.PresetsDataStore {
	m := &memoryDataStore{
		presets: make(map[string]preset.Info)}
	return m
}

// SavePreset save preset information into a memory database.
// If a preset with the given name already exists information of the preset is updated.
func (m *memoryDataStore) SavePreset(presetInfo preset.Info) error {
	m.presets[presetInfo.Name] = presetInfo
	return nil
}

// LoadPreset loads the preset information corresponding to the given presetName.
// If the preset is not found a ErrPresetNotFound is returned.
func (m *memoryDataStore) LoadPreset(presetName string) (preset.Info, error) {
	pi, ok := m.presets[presetName]
	if ok {
		return pi, nil
	}

	return preset.Info{}, datastore.ErrPresetNotFound
}

func (m *memoryDataStore) GetPresetNames() []string {
	presetNames := make([]string, 0, len(m.presets))
	for k := range m.presets {
		presetNames = append(presetNames, k)
	}

	return presetNames
}
