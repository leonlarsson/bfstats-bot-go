package shapes

// BaseData represents the base data for all image creations.
type BaseData struct {
	Identifier string `json:"identifier"`
	Username   string `json:"username"`
	Avatar     string `json:"avatar"`
	Platform   int    `json:"platform"`
	TimePlayed string `json:"timePlayed"`
	Meta       Meta   `json:"meta"`
}

// Meta represents the metadata for the image.
type Meta struct {
	Game    string `json:"game"`
	Segment string `json:"segment"`
}

// Slot represents a single canvas slot with a name, value and extra information.
type Slot struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Extra string `json:"extra"`
}

// GenericRegularData represents the data for a regular format image.
type GenericRegularData struct {
	BaseData BaseData
	Slots    GenericRegularSlots
}

// GenericRegularSlots represents the canvas slots for the regular format image.
type GenericRegularSlots struct {
	L1,
	L2,
	L3,
	L4,
	L5,
	L6,
	R1,
	R2,
	R3,
	R4 Slot
}

// GenericGridData represents the data for a grid format image.
type GenericGridData struct {
	BaseData BaseData
	Slots    []Slot
}
