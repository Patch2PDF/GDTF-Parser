package XMLTypes

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

// TODO: DMXValue Type
// see gdtf https://www.gdtf.eu/gdtf/file-spec/file-format-definition/#attrtype-dmxvalue
type XMLDMXValue = string

type DMXAddress struct {
	Address  int16
	Universe int
}

func (b *DMXAddress) UnmarshalXMLAttr(attr xml.Attr) error {
	if strings.Contains(attr.Value, ".") {
		frags := strings.Split(attr.Value, ".")
		value, err := strconv.ParseInt(frags[0], 10, 0)
		if err != nil {
			return err
		}
		b.Universe = int(value)

		value, err = strconv.ParseInt(frags[1], 10, 16)
		if err != nil {
			return err
		}
		b.Address = int16(value)
	} else {
		absolute, err := strconv.ParseInt(attr.Value, 10, 64)
		if err != nil {
			return err
		}
		b.Address = int16(absolute % 512)
		b.Universe = int(absolute / 512)
	}
	return nil
}

type XMLNodeReference = string

// Custom type for Yes/No -> bool conversion
type YesNoBool bool

// Implement xml.Unmarshaler
func (b *YesNoBool) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "yes", "true", "1":
		*b = true
	case "no", "false", "0":
		*b = false
	default:
		*b = false
	}
	return nil
}

// format xyY (X,Y,Y2)
type ColorCIE struct {
	X  float32
	Y  float32
	Y2 float32
}

func (object ColorCIE) Parse() Types.ColorCIE {
	return Types.ColorCIE(object)
}

func (dest *ColorCIE) UnmarshalXMLAttr(attr xml.Attr) error {
	frags := strings.Split(attr.Value, ",")
	if len(frags) != 3 {
		return fmt.Errorf("invalid structure for ColorCIE")
	}
	value, err := strconv.ParseFloat(frags[0], 32)
	if err != nil {
		return err
	}
	dest.X = float32(value)
	value, err = strconv.ParseFloat(frags[1], 32)
	if err != nil {
		return err
	}
	dest.Y = float32(value)
	value, err = strconv.ParseFloat(frags[2], 32)
	if err != nil {
		return err
	}
	dest.Y2 = float32(value)
	return nil
}

type Matrix [4][4]float64

// The transformation matrix consists 4 x 4 floats.
// Stored in a row-major order.
// For example, each row of the matrix is stored as a 4- component vector.
// The mathematical definition of the matrix is in a column-major order.
// For example, the matrix rotation is stored in the first three columns,
// and the translation is stored in the 4th column.
// The metric system consists of the Right- handed Cartesian Coordinates XYZ:
//   X – from left (-X) to right (+X),
//   Y – from the outside of the monitor (-Y) to the inside of the monitor (+Y),
//   Z – from bottom (-Z) to top (+Z). 0,0,0 – center base.

func (dest *Matrix) UnmarshalXMLAttr(attr xml.Attr) error {
	rows := strings.Split(strings.Trim(attr.Value, "{}"), "}{")
	if len(rows) != 4 {
		return fmt.Errorf("invalid structure for Matrix")
	}
	for index, row := range rows {
		columns := strings.Split(row, ",")
		if len(columns) != 4 {
			return fmt.Errorf("invalid structure for Matrix")
		}
		for column_index, column_value := range columns {
			value, err := strconv.ParseFloat(column_value, 64)
			if err != nil {
				return err
			}
			dest[index][column_index] = value
		}
	}
	return nil
}

type Rotation [3][3]float32

// Rotation matrix, consist of 3*3 floats.
// Stored as row-major matrix, i.e. each row of the matrix is stored as a 3-component vector.
// Mathematical definition of the matrix is column-major,
// i.e. the matrix rotation is stored in the three columns.
// Metric system, right-handed Cartesian coordinates XYZ:
//   X – from left (-X) to right (+X),
//   Y – from the outside of the monitor (-Y) to the inside of the monitor (+Y),
//   Z – from the bottom (-Z) to the top (+Z).

func (dest *Rotation) UnmarshalXMLAttr(attr xml.Attr) error {
	rows := strings.Split(strings.Trim(attr.Value, "{}"), "}{")
	if len(rows) != 3 {
		return fmt.Errorf("invalid structure for Rotation")
	}
	for index, row := range rows {
		columns := strings.Split(row, ",")
		if len(columns) != 3 {
			return fmt.Errorf("invalid structure for Rotation")
		}
		for column_index, column_value := range columns {
			value, err := strconv.ParseFloat(column_value, 32)
			if err != nil {
				return err
			}
			dest[index][column_index] = float32(value)
		}
	}
	return nil
}

type FileReference = string // without extension and without path

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func (dest *Vector3) UnmarshalXMLAttr(attr xml.Attr) error {
	frags := strings.Split(attr.Value, ",")
	if len(frags) != 3 {
		return fmt.Errorf("invalid structure for Vector3")
	}
	value, err := strconv.ParseFloat(frags[0], 32)
	if err != nil {
		return err
	}
	dest.X = float32(value)
	value, err = strconv.ParseFloat(frags[1], 32)
	if err != nil {
		return err
	}
	dest.Y = float32(value)
	value, err = strconv.ParseFloat(frags[2], 32)
	if err != nil {
		return err
	}
	dest.Z = float32(value)
	return nil
}

type Hex int

func (h *Hex) UnmarshalXMLAttr(attr xml.Attr) error {
	v, err := strconv.ParseInt(attr.Value, 0, 0)
	if err != nil {
		return err
	}
	*h = Hex(v)
	return nil
}

type XMLTime struct {
	time.Time
}

func (b *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	value, err := time.Parse("2006-01-02T15:04:05", attr.Value)
	if err != nil {
		return err
	}
	b.Time = value
	return nil
}

type IntList []int

func (h *IntList) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil // if empty, dont do anything
	}
	values := strings.Split(attr.Value, ",")
	for _, value := range values {
		v, err := strconv.ParseInt(value, 0, 0)
		if err != nil {
			return err
		}
		*h = append(*h, int(v))
	}
	return nil
}

type ConvertToDestinationStruct[T any] interface {
	Parse() T
}

type ConvertToDestinationMapStruct[T any] interface {
	ConvertToDestinationStruct[T]
	ParseKey() string
}

func ParseList[Source ConvertToDestinationStruct[Destination], Destination any](source *[]Source) []*Destination {
	var destination []*Destination = make([]*Destination, len(*source))
	for index, element := range *source {
		parsedElement := element.Parse()
		destination[index] = &parsedElement
	}
	return destination
}

func ParseMap[Source ConvertToDestinationMapStruct[Destination], Destination any](source *[]Source) map[string]*Destination {
	destination := make(map[string]*Destination)
	for _, element := range *source {
		parsedElement := element.Parse()
		destination[element.ParseKey()] = &parsedElement
	}
	return destination
}
