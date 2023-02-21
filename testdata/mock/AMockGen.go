package mock

import (
	amock_core "github.com/ymz-ncnk/amock/core"
	"github.com/ymz-ncnk/amockgen"
)

func NewAMockGen() AMockGen {
	return AMockGen{amock_core.New("AMockGen")}
}

// amockgen.AMockGen
type AMockGen struct {
	*amock_core.Mock
}

func (aMockGen AMockGen) RegisterGenerate(
	fn func(iDesc amockgen.MockImplDesc) ([]byte, error)) AMockGen {
	aMockGen.Register("Generate", fn)
	return aMockGen
}

func (aMockGen AMockGen) Generate(iDesc amockgen.MockImplDesc) (data []byte,
	err error) {
	vals, err := aMockGen.Call("Generate", iDesc)
	if err != nil {
		return
	}
	data = vals[0].([]byte)
	err, _ = vals[1].(error)
	return
}
