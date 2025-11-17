package Types

type Model struct {
	Name            string
	Length          float32
	Width           float32
	Height          float32
	PrimitiveType   string // enum
	File            *string
	SVGOffsetX      float32
	SVGOffsetY      float32
	SVGSideOffsetX  float32
	SVGSideOffsetY  float32
	SVGFrontOffsetX float32
	SVGFrontOffsetY float32
	Mesh            *Mesh
}

func (obj *Model) CreateReferencePointer() {
	refPointers.Models[obj.Name] = obj
}
