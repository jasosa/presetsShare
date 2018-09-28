package preset

//AmpModels gets a list of all available amp models
// with their default settings
var AmpModels = getAmpRegistry()

var ar ampRegistry

func getAmpRegistry() *ampRegistry {
	if ar == (ampRegistry{}) {
		ar = ampRegistry{
			Fender65UsDx: Amplifier{Model: "Fender 65 Deluxe"},
			Fender65UsTw: Amplifier{Model: "Fender 65 Twin Reverb"},
		}
	}

	return &ar
}

//Amplifier represents and amplifier model with its settings
type Amplifier struct {
	Model string
}

type ampRegistry struct {
	Fender65UsDx Amplifier
	Fender65UsTw Amplifier
}
