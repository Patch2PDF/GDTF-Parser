package gdtfxml_test

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func parsingTest[T any](t *testing.T, xmlData string, domain string, want T) {
	var result T
	err := xml.Unmarshal([]byte(xmlData), &result)
	if err != nil {
		t.Errorf("%s: XML Unmarshal threw error: %s", domain, err)
	} else {
		if !cmp.Equal(result, want) {
			t.Errorf("%s: XML Unmarshaling Output does not match", domain)
			t.Errorf("%s", cmp.Diff(result, want))
		}
	}
}

func floatPtr(v float32) *float32   { return &v }
func float64Ptr(v float64) *float64 { return &v }

func intPtr(v int) *int       { return &v }
func strPtr(v string) *string { return &v }
