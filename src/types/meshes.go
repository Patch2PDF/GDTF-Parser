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

func Identity() MeshMatrix {
	return MeshMatrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (a MeshMatrix) Mul(b MeshMatrix) MeshMatrix {
	m := MeshMatrix{}
	m.X00 = a.X00*b.X00 + a.X01*b.X10 + a.X02*b.X20 + a.X03*b.X30
	m.X10 = a.X10*b.X00 + a.X11*b.X10 + a.X12*b.X20 + a.X13*b.X30
	m.X20 = a.X20*b.X00 + a.X21*b.X10 + a.X22*b.X20 + a.X23*b.X30
	m.X30 = a.X30*b.X00 + a.X31*b.X10 + a.X32*b.X20 + a.X33*b.X30
	m.X01 = a.X00*b.X01 + a.X01*b.X11 + a.X02*b.X21 + a.X03*b.X31
	m.X11 = a.X10*b.X01 + a.X11*b.X11 + a.X12*b.X21 + a.X13*b.X31
	m.X21 = a.X20*b.X01 + a.X21*b.X11 + a.X22*b.X21 + a.X23*b.X31
	m.X31 = a.X30*b.X01 + a.X31*b.X11 + a.X32*b.X21 + a.X33*b.X31
	m.X02 = a.X00*b.X02 + a.X01*b.X12 + a.X02*b.X22 + a.X03*b.X32
	m.X12 = a.X10*b.X02 + a.X11*b.X12 + a.X12*b.X22 + a.X13*b.X32
	m.X22 = a.X20*b.X02 + a.X21*b.X12 + a.X22*b.X22 + a.X23*b.X32
	m.X32 = a.X30*b.X02 + a.X31*b.X12 + a.X32*b.X22 + a.X33*b.X32
	m.X03 = a.X00*b.X03 + a.X01*b.X13 + a.X02*b.X23 + a.X03*b.X33
	m.X13 = a.X10*b.X03 + a.X11*b.X13 + a.X12*b.X23 + a.X13*b.X33
	m.X23 = a.X20*b.X03 + a.X21*b.X13 + a.X22*b.X23 + a.X23*b.X33
	m.X33 = a.X30*b.X03 + a.X31*b.X13 + a.X32*b.X23 + a.X33*b.X33
	return m
}

func (a MeshMatrix) MulPosition(b MeshVector) MeshVector {
	x := a.X00*b.X + a.X01*b.Y + a.X02*b.Z + a.X03
	y := a.X10*b.X + a.X11*b.Y + a.X12*b.Z + a.X13
	z := a.X20*b.X + a.X21*b.Y + a.X22*b.Z + a.X23
	return MeshVector{x, y, z}
}

type Vertex struct {
	Position MeshVector
	Normal   *MeshVector
}

func (obj *Vertex) Copy() Vertex {
	normalCopy := *obj.Normal
	return Vertex{
		Position: obj.Position,
		Normal:   &normalCopy,
	}
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

func (obj *Triangle) Copy() Triangle {
	V0 := obj.V0.Copy()
	V1 := obj.V1.Copy()
	V2 := obj.V2.Copy()
	return Triangle{V0: &V0, V1: &V1, V2: &V2}
}

type Mesh struct {
	Triangles []*Triangle
}

func (obj *Mesh) AddTriangle(triangle *Triangle) {
	obj.Triangles = append(obj.Triangles, triangle)
}

func (obj *Mesh) Copy() Mesh {
	triangles := make([]*Triangle, len(obj.Triangles))
	for index, triangle := range obj.Triangles {
		temp := triangle.Copy()
		triangles[index] = &temp
	}
	return Mesh{Triangles: triangles}
}

func (obj *Mesh) Add(mesh *Mesh) *Mesh {
	obj.Triangles = append(obj.Triangles, mesh.Triangles...)
	return obj
}

func (obj *Mesh) RotateAndTranslate(translationMatrix MeshMatrix) {
	for _, triangle := range obj.Triangles {
		triangle.V0.Position = translationMatrix.MulPosition(triangle.V0.Position)
		triangle.V1.Position = translationMatrix.MulPosition(triangle.V1.Position)
		triangle.V2.Position = translationMatrix.MulPosition(triangle.V2.Position)
	}
}
