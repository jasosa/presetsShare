package main

import (
	"fmt"

	"github.com/jasosa/PresetsManagement/pkg/preset"

	"github.com/jasosa/PresetsManagement/pkg/datastore/memory"
	"github.com/jasosa/PresetsManagement/pkg/service"
)

func main() {
	ds := memory.NewDataStore()
	ds.SavePreset(preset.Info{Name: "PresetA", Amp: preset.AmpModels.Fender65UsTw})
	ds.SavePreset(preset.Info{Name: "PresetB", Amp: preset.AmpModels.Fender65UsDx})

	s := service.NewPresetsService(ds)
	for _, name := range s.ListPresets() {
		fmt.Println(fmt.Sprintf("Preset: %s", name))
	}
}
