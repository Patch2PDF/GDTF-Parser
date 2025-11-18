package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"

type Revision struct {
	Text       string  `xml:",attr"`
	Date       XMLTime `xml:",attr"`
	UserID     uint    `xml:",attr"`
	ModifiedBy string  `xml:",attr"`
}

func (revision Revision) Parse() Types.Revision {
	return Types.Revision{
		Text:       revision.Text,
		Date:       revision.Date.Time,
		UserID:     revision.UserID,
		ModifiedBy: revision.ModifiedBy,
	}
}
