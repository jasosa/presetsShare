package service

import (
	"testing"

	"github.com/jasosa/PresetsManagement/pkg/preset"
	"github.com/pkg/errors"

	"github.com/jasosa/PresetsManagement/pkg/datastore/memory"
	"github.com/jasosa/PresetsManagement/pkg/datastore/mock"
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

func TestUploadPresetSuccesfully(t *testing.T) {
	ds := memory.NewDataStore()
	ps := NewPresetsService(ds)
	presetInfo := preset.Info{Name: "Preset C", Amp: preset.AmpModels.Fender65UsTw}
	err := ps.UploadPreset(presetInfo)
	if err != nil {
		t.Fatalf("Non error was expected but got %q", err.Error())
	}

	numberOfPresets := len(ps.ListPresets())
	if numberOfPresets != 1 {
		t.Errorf("Expected a list of '%d' presets but got '%d'", 1, numberOfPresets)
	}
}

func TestUploadPresetError(t *testing.T) {
	mds := mock.NewFailingDataStore()
	ps := NewPresetsService(mds)
	presetInfo := preset.Info{Name: "Preset C", Amp: preset.AmpModels.Fender65UsTw}
	err := ps.UploadPreset(presetInfo)
	if err == nil {
		t.Fatalf("Error was expected but got %v", err)
	}

	if err != nil {
		switch err.(type) {
		case *ErrUploadingPreset:
			// everything ok
		default:
			t.Errorf("Expected error was '%q' but got '%v'", "ErrUploadingPreset", errors.Cause(err))
		}
	}

}
