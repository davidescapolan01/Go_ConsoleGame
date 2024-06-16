package gameentity

type EntityType int

const (
	Player = iota
	Enemy
)

var EntityName = map[EntityType]string{
	Player: "player",
	Enemy:  "enemy",
}

func (
	et EntityType,
) String() string {
	return EntityName[et]
}

type Entity struct {
	Icon            string
	x               int
	y               int
	moveFrameTime   int
	life            int
	consecutiveMove int
}

func (
	e *Entity,
) Init(
	posX int,
	posY int,
	tFrame int,
	et EntityType,
) {
	e.x = posX
	e.y = posY
	e.moveFrameTime = tFrame
	e.consecutiveMove = tFrame

	switch et {
	case Enemy:
		e.life = 1
		e.Icon = "V"
		e.moveFrameTime = tFrame
		e.consecutiveMove = tFrame
	case Player:
		e.life = 3
		e.Icon = "A"
		e.moveFrameTime = 0
		e.consecutiveMove = 0
	}
}

func (
	e *Entity,
) MoveForeward(
	unit int,
	maxHeight int,
) bool {
	if e.consecutiveMove < e.moveFrameTime || e.moveFrameTime == 0 {
		if (e.y+unit) < maxHeight &&
			(e.y+unit) >= 0 {
			e.y += unit
			e.consecutiveMove++
		} else if (e.y+unit) == maxHeight && e.moveFrameTime != 0 {
			return true
		}
	} else {
		e.consecutiveMove = 0
	}
	return false
}

func (
	e *Entity,
) MoveLeft(
	unit int,
	maxWidth int,
) {
	if e.consecutiveMove < e.moveFrameTime || e.moveFrameTime == 0 {
		if (e.x+unit) < maxWidth &&
			(e.x+unit) >= 0 {
			e.x += unit
			e.consecutiveMove++
		}
	} else {
		e.consecutiveMove = 0
	}
}

func (
	e *Entity,
) GetPosition() (
	int,
	int,
) {
	return e.x, e.y
}

func (
	e *Entity,
) Kill() {
	e.life--
}

func (
	e *Entity,
) IsDead() bool {
	return e.life == 0
}
