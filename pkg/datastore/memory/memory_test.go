package memory

import (
	"testing"

	"github.com/jasosa/PresetsManagement/pkg/datastore"

	"github.com/jasosa/PresetsManagement/pkg/preset"
)

func TestSavePreset(t *testing.T) {
	// Arrange
	ds := NewDataStore()

	// Act
	ds.SavePreset(preset.Info{Name: "Preset 1"})

	// Assert
	ms := ds.(*memoryDataStore)
	_, ok := ms.presets["Preset 1"]
	if !ok {
		t.Errorf("Preset not saved!")
	}
}

func TestLoadPreset(t *testing.T) {

	// Arrange
	ds := NewDataStore()
	presetName := "Preset A"
	presetInfo := preset.Info{Name: presetName, Amp: preset.AmpModels.Fender65UsDx}
	ds.SavePreset(presetInfo)

	// Act
	loadedPreset, err := ds.LoadPreset(presetName)

	// Assert
	if err != nil {
		t.Errorf("Expected preset was not found")
	}

	if loadedPreset.Name != presetName {
		t.Errorf("Expected preset name was %s but got %s", presetName, loadedPreset.Name)
	}

	if loadedPreset.Amp.Model != preset.AmpModels.Fender65UsDx.Model {
		t.Errorf("Expected preset amplifier model was %s but got %s", preset.AmpModels.Fender65UsDx.Model, loadedPreset.Amp.Model)
	}
}

func TestLoadNotSavedPresetGivesError(t *testing.T) {
	// Arrange
	ds := NewDataStore()
	presetName := "Preset A"
	preset := preset.Info{Name: presetName}
	ds.SavePreset(preset)

	// Act
	_, err := ds.LoadPreset("Another preset name")

	// Assert
	if err != datastore.ErrPresetNotFound {
		t.Errorf("Expected error was %v but got %v", datastore.ErrPresetNotFound, err)
	}
}

func TestGetAllPresetNames(t *testing.T) {
	// Arrange
	ds := NewDataStore()
	presetInfoA := preset.Info{Name: "Preset A"}
	presetInfoB := preset.Info{Name: "Preset B"}
	presetInfoC := preset.Info{Name: "Preset C"}
	ds.SavePreset(presetInfoA)
	ds.SavePreset(presetInfoB)
	ds.SavePreset(presetInfoC)
	names := ds.GetPresetNames()
	expectedNumberOfNames := 3
	if len(names) != expectedNumberOfNames {
		t.Errorf("Expected %d preset names but got %d", expectedNumberOfNames, len(names))
	}
}
