package XMLTypes

type Revision struct {
	Text       string  `xml:",attr"`
	Date       XMLTime `xml:",attr"`
	UserID     uint    `xml:",attr"`
	ModifiedBy string  `xml:",attr"`
}
