package main

import (
	"github.com/jasosa/PresetsManagement/server"
)

func main() {
	/* ds := memory.NewDataStore()
	ds.SavePreset(preset.Info{Name: "PresetA", Amp: preset.AmpModels.Fender65UsTw})
	ds.SavePreset(preset.Info{Name: "PresetB", Amp: preset.AmpModels.Fender65UsDx})

	s := service.NewPresetsService(ds)
	for _, name := range s.ListPresets() {
		fmt.Println(fmt.Sprintf("Preset: %s", name))
	} */

	s, _ := server.NewServer()
	s.Start()
}
