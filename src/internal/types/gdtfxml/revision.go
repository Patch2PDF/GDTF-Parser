package XMLTypes

import xmldatetime "github.com/datainq/xml-date-time"

type Revision struct {
	Text       string                 `xml:",attr"`
	Date       xmldatetime.CustomTime `xml:"date,attr"`
	UserID     uint                   `xml:",attr"`
	ModifiedBy string                 `xml:",attr"`
}
