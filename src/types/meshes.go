package Types

type MeshVector struct {
	X float64
	Y float64
	Z float64
}

func (obj *MeshVector) ToVertex(normal *MeshVector) *Vertex {
	return &Vertex{
		X:      obj.X,
		Y:      obj.Y,
		Z:      obj.Z,
		Normal: normal,
	}
}

type Vertex struct {
	X      float64
	Y      float64
	Z      float64
	Normal *MeshVector
}

type Triangle struct {
	V0 *Vertex
	V1 *Vertex
	V2 *Vertex
}

type Mesh struct {
	Triangles []*Triangle
}

func (obj *Mesh) AddTriangle(triangle *Triangle) {
	obj.Triangles = append(obj.Triangles, triangle)
}
