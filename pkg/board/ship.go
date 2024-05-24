package board

type ShipField struct {
	BoardField
	name string
	size int
}

func NewCarrier() *ShipField {
	return &ShipField{name: "carrier", size: 5}
}

func NewBattleship() *ShipField {
	return &ShipField{name: "battleship", size: 4}
}

func NewCruiser() *ShipField {
	return &ShipField{name: "cruiser", size: 3}
}

func NewSubmarine() *ShipField {
	return &ShipField{name: "submarine", size: 3}
}

func NewDestroyer() *ShipField {
	return &ShipField{name: "destroyer", size: 2}
}
