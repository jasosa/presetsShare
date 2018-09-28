package service

import (
	"testing"

	"github.com/jasosa/PresetsManagement/pkg/preset"

	"github.com/jasosa/PresetsManagement/pkg/datastore/memory"
)

func TestListPresets(t *testing.T) {
	// Arrange
	ds := memory.NewDataStore()
	ps := NewPresetsService(ds)
	ds.SavePreset(preset.Info{Name: "Preset A"})
	ds.SavePreset(preset.Info{Name: "Preset B"})
	presetNames := ps.ListPresets()
	expectedNumberOfPresetNames := 2
	if len(presetNames) != expectedNumberOfPresetNames {
		t.Errorf("Expected %d preset names but got %d", expectedNumberOfPresetNames, len(presetNames))
	}
}
