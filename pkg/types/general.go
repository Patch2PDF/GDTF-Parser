package Types

import "github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"

// TODO: DMXValue Type
// see gdtf https://www.gdtf.eu/gdtf/file-spec/file-format-definition/#attrtype-dmxvalue
type DMXValue = string

type NodeReference[T any] struct {
	String string
	Ptr    *T
}

type DMXAddress struct {
	Address  int16
	Universe int
}

// format xyY (X,Y,Y2)
type ColorCIE struct {
	X  float32
	Y  float32
	Y2 float32
}

// The transformation matrix consists 4 x 4 floats.
// Stored in a row-major order.
// For example, each row of the matrix is stored as a 4- component vector.
// The mathematical definition of the matrix is in a column-major order.
// For example, the matrix rotation is stored in the first three columns,
// and the translation is stored in the 4th column.
// The metric system consists of the Right- handed Cartesian Coordinates XYZ:
//
//	X – from left (-X) to right (+X),
//	Y – from the outside of the monitor (-Y) to the inside of the monitor (+Y),
//	Z – from bottom (-Z) to top (+Z). 0,0,0 – center base.
type Matrix [4][4]float64

func (obj Matrix) toMeshMatrix() MeshTypes.Matrix {
	return MeshTypes.Matrix{
		X00: obj[0][0], X01: obj[0][1], X02: obj[0][2], X03: obj[0][3],
		X10: obj[1][0], X11: obj[1][1], X12: obj[1][2], X13: obj[1][3],
		X20: obj[2][0], X21: obj[2][1], X22: obj[2][2], X23: obj[2][3],
		X30: obj[3][0], X31: obj[3][1], X32: obj[3][2], X33: obj[3][3],
	}
}

// Rotation matrix, consist of 3*3 floats.
// Stored as row-major matrix, i.e. each row of the matrix is stored as a 3-component vector.
// Mathematical definition of the matrix is column-major,
// i.e. the matrix rotation is stored in the three columns.
// Metric system, right-handed Cartesian coordinates XYZ:
//
//	X – from left (-X) to right (+X),
//	Y – from the outside of the monitor (-Y) to the inside of the monitor (+Y),
//	Z – from the bottom (-Z) to the top (+Z).
type Rotation [3][3]float32

type FileReference = string // without extension and without path

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Hex = int
