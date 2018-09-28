package service

import (
	"testing"

	"github.com/jasosa/PresetsManagement/pkg/preset"
	"github.com/pkg/errors"

	"github.com/jasosa/PresetsManagement/pkg/datastore/memory"
)

func TestListPresets(t *testing.T) {
	// Arrange
	ds := memory.NewDataStore()
	ps := NewPresetsService(ds)
	ds.SavePreset(preset.Info{Name: "Preset A"})
	ds.SavePreset(preset.Info{Name: "Preset B"})

	// Act
	presetNames := ps.ListPresets()

	// Assert
	expectedNumberOfPresetNames := 2
	if len(presetNames) != expectedNumberOfPresetNames {
		t.Errorf("Expected %d preset names but got %d", expectedNumberOfPresetNames, len(presetNames))
	}
}

func TestLoadExistingPreset(t *testing.T) {
	// Arrange
	ds := memory.NewDataStore()
	ps := NewPresetsService(ds)
	ds.SavePreset(preset.Info{Name: "Preset A"})
	ds.SavePreset(preset.Info{Name: "Preset B"})

	// Act
	pi, _ := ps.ShowPreset("Preset A")

	// Assert
	if pi.Name != "Preset A" {
		t.Errorf("Expected preset name was %q but got %q", "Preset A", pi.Name)
	}
}

func TestLoadNonExistentPreset(t *testing.T) {
	ds := memory.NewDataStore()
	ps := NewPresetsService(ds)
	ds.SavePreset(preset.Info{Name: "Preset A"})
	ds.SavePreset(preset.Info{Name: "Preset B"})

	// Act
	_, err := ps.ShowPreset("Preset C")

	// Assert
	if err != nil {
		switch err.(type) {
		case *ErrPresetNotFound:
			// everything ok
		default:
			t.Errorf("Expected error was '%q' but got '%v'", "ErrPresetNotFound", errors.Cause(err))
		}
	}
}
