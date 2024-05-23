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
