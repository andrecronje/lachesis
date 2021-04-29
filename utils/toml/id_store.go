package toml

import (
	"fmt"
	//"github.com/naoina/toml"
)

// IDStore is implementation for migration id storage in toml file value
type IDStore struct {
	idChain []string
	data    *Helper
}

// NewIDStore return new toml.IDStore
func NewIDStore(d *Helper, idChain []string) *IDStore {
	return &IDStore{
		idChain: idChain,
		data:    d,
	}
}

// GetID return current saved migration ID from toml data
func (p *IDStore) GetID() string {
	v, err := p.data.GetParamString("Version", "")
	if err != nil {
		return ""
	}

	return p.human2id(v)
}

// SetID save migration ID in toml data
func (p *IDStore) SetID(id string) {
	v := p.id2human(id)
	_, ok := p.data.GetTable().Fields["Version"]

	var err error
	if !ok {
		err = p.data.AddParam("Version", "", v)
	} else {
		err = p.data.SetParam("Version", "", v)
	}
	if err != nil {
		panic(err)
	}
}

func (p *IDStore) id2human(id string) string {
	for i, x := range p.idChain {
		if x != id {
			continue
		}
		return fmt.Sprintf("v.%d.0", i+1)
	}
	panic("id is not from idChain")
}

func (p *IDStore) human2id(str string) string {
	var i int
	_, err := fmt.Sscanf(str, "v.%d.0", &i)
	if err == nil && i > 0 && i <= len(p.idChain) {
		return p.idChain[i-1]
	}

	return str
}
