package Types

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

type Model struct {
	Name            string
	Length          float32 // in meter
	Width           float32 // in meter
	Height          float32 // in meter
	PrimitiveType   string  // enum
	File            *string
	SVGOffsetX      float32
	SVGOffsetY      float32
	SVGSideOffsetX  float32
	SVGSideOffsetY  float32
	SVGFrontOffsetX float32
	SVGFrontOffsetY float32
	Mesh            *MeshTypes.Mesh
}

func (obj *Model) CreateReferencePointer() {
	refPointers.Models[obj.Name] = obj
}
