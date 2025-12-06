package Types

import (
	"fmt"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

type GDTF struct {
	DataVersion string
	FixtureType FixtureType
}

func (obj *GDTF) CreateReferencePointer(refPointers *ReferencePointers) {
	obj.FixtureType.CreateReferencePointer(refPointers)
}

func (obj *GDTF) ResolveReference(refPointers *ReferencePointers) {
	obj.FixtureType.ResolveReference(refPointers)
}

// Assemble a mesh based on entire geometry
func (obj *GDTF) BuildMesh(dmxMode string) (*MeshTypes.Mesh, error) {
	mode := obj.FixtureType.DMXModes[dmxMode]
	if mode == nil {
		return nil, fmt.Errorf("unknown DMX Mode '%s' in Fixture %s", dmxMode, obj.FixtureType.Name)
	}
	return obj.FixtureType.DMXModes[dmxMode].Geometry.Ptr.Ptr.(MeshGenerator).GenerateMesh(MeshTypes.IdentityMatrix()), nil
}
