package Types

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
type Matrix [4][4]float32

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
