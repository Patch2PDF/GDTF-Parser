package XMLTypes

import (
	"encoding/xml"
	"strings"
)

// TODO: DMXValue Type
// see gdtf https://www.gdtf.eu/gdtf/file-spec/file-format-definition/#attrtype-dmxvalue
type XMLDMXValue = string

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

type ColorCIE = string

type Matrix = string

type FileReference = string // without extension and without path
