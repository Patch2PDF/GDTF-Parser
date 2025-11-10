package Types

type Wheel struct {
	Name       string
	WheelSlots []WheelSlot
}

// TODO:
type WheelSlot struct {
	Name          string
	Color         string
	Filter        string // TODO:
	MediaFileName string
}
