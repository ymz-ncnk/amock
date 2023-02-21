package mock

import (
	amock_core "github.com/ymz-ncnk/amock/core"
)

func NewPersistor() Persistor {
	return Persistor{amock_core.New("Persistor")}
}

// persistor.Persistor
type Persistor struct {
	*amock_core.Mock
}

func (persistor Persistor) RegisterPersist(
	fn func(name string, data []byte, path string) error) Persistor {
	persistor.Register("Persist", fn)
	return persistor
}

func (persistor Persistor) Persist(name string, data []byte, path string) (
	err error) {
	vals, err := persistor.Call("Persist", name, data, path)
	if err != nil {
		return
	}
	err, _ = vals[0].(error)
	return
}
