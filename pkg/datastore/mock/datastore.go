package mock

import (
	"errors"

	"github.com/jasosa/PresetsManagement/pkg/datastore"
	"github.com/jasosa/PresetsManagement/pkg/preset"
)

type failingDataStore struct {
}

// NewFailingDataStore returns a new instance a of mock datastore
func NewFailingDataStore() datastore.PresetsDataStore {
	return &failingDataStore{}
}

func (f *failingDataStore) SavePreset(presetInfo preset.Info) error {
	return errors.New("Error saving preset")
}

func (f *failingDataStore) LoadPreset(presetName string) (preset.Info, error) {
	return preset.Info{}, errors.New("Preset cannot be loaded")
}

func (f *failingDataStore) GetPresetNames() []string {
	return []string{}
}
