package Types

import "math"

type MeshVector struct {
	X float64
	Y float64
	Z float64
}

func (a MeshVector) Add(b MeshVector) MeshVector {
	return MeshVector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a MeshVector) Sub(b MeshVector) MeshVector {
	return MeshVector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a MeshVector) Cross(b MeshVector) MeshVector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return MeshVector{x, y, z}
}

func (a MeshVector) Normalize() MeshVector {
	r := 1 / math.Sqrt(a.X*a.X+a.Y*a.Y+a.Z*a.Z)
	return MeshVector{a.X * r, a.Y * r, a.Z * r}
}

func (obj MeshVector) ToVertex(normal *MeshVector) *Vertex {
	return &Vertex{
		Position: obj,
		Normal:   normal,
	}
}

func (a MeshVector) Min(b *MeshVector) MeshVector {
	return MeshVector{
		math.Min(a.X, b.X),
		math.Min(a.Y, b.Y),
		math.Min(a.Z, b.Z),
	}
}

func (a MeshVector) Max(b *MeshVector) MeshVector {
	return MeshVector{
		math.Max(a.X, b.X),
		math.Max(a.Y, b.Y),
		math.Max(a.Z, b.Z),
	}
}

func (a MeshVector) Mult(b MeshVector) MeshVector {
	return MeshVector{
		a.X * b.X,
		a.Y * b.Y,
		a.Z * b.Z,
	}
}

func (a MeshVector) Div(b MeshVector) MeshVector {
	return MeshVector{
		a.X / b.X,
		a.Y / b.Y,
		a.Z / b.Z,
	}
}

type MeshMatrix struct {
	X00, X01, X02, X03 float64
	X10, X11, X12, X13 float64
	X20, X21, X22, X23 float64
	X30, X31, X32, X33 float64
}

type Vertex struct {
	Position MeshVector
	Normal   *MeshVector
}

type Triangle struct {
	V0 *Vertex
	V1 *Vertex
	V2 *Vertex
}

func (t *Triangle) Normal() MeshVector {
	e1 := t.V1.Position.Sub(t.V0.Position)
	e2 := t.V2.Position.Sub(t.V0.Position)
	return e1.Cross(e2).Normalize()
}

type Mesh struct {
	Triangles []*Triangle
}

func (obj *Mesh) AddTriangle(triangle *Triangle) {
	obj.Triangles = append(obj.Triangles, triangle)
}
