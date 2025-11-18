package Types

import "fmt"

type GDTF struct {
	DataVersion string
	FixtureType FixtureType
}

func (obj *GDTF) CreateReferencePointer() {
	obj.FixtureType.CreateReferencePointer()
}

func (obj *GDTF) ResolveReference() {
	obj.FixtureType.ResolveReference()
}

// Assemble a mesh based on entire geometry
func (obj *GDTF) BuildMesh(dmxMode string) (*Mesh, error) {
	mode := obj.FixtureType.DMXModes[dmxMode]
	if mode == nil {
		return nil, fmt.Errorf("unknown DMX Mode '%s' in Fixture %s", dmxMode, obj.FixtureType.Name)
	}
	return obj.FixtureType.DMXModes[dmxMode].Geometry.Ptr.Ptr.(MeshGenerator).GenerateMesh(Identity()), nil
}
