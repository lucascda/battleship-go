package board

type BoardField struct {
	isWater   bool
	isShip    bool
	isHit     bool
	fieldChar string
}

func NewBoardField() *BoardField {
	return &BoardField{isWater: true, isShip: false, isHit: false, fieldChar: ""}
}

func (b *BoardField) ChangeFieldChar(value string) {
	b.fieldChar = value
}

func (b *BoardField) ToggleisWater() {
	b.isWater = !b.isWater
}

func (b *BoardField) ToggleisShip() {
	b.isShip = !b.isShip
}

func (b *BoardField) ToggleisHit() {
	b.isHit = !b.isHit
}
