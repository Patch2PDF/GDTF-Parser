package Type

import "time"

type Revision struct {
	Text       string
	Date       time.Time
	UserID     uint
	ModifiedBy string
}
