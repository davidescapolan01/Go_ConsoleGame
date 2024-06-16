package gameengine

import (
	"fmt"
	"time"

	"asciigame/goGame/gameentity"
	"asciigame/goGame/gamefield"
)

type Engine struct {
	frameEnemyMovement  int
	framePalyerMovement int
	startEnemyNumnebr   int

	frameTime     time.Duration
	nextFrameTime time.Time

	field *gamefield.Field

	Enemies []gameentity.Entity
	Player  gameentity.Entity
}

func (
	e *Engine,
) InitEngine(
	fTime time.Duration,
	fEnemyMove int,
	fPlayerMov int,
	eNumb int,
	f *gamefield.Field,
) {
	e.frameTime = fTime
	e.nextFrameTime = time.Now().Add(e.frameTime)
	e.frameEnemyMovement = fEnemyMove
	e.framePalyerMovement = fPlayerMov
	e.startEnemyNumnebr = eNumb
	e.field = f

	e.Player = gameentity.Entity{}
	e.Player.Init((e.field.Width-1)/2, e.field.Height-2, e.framePalyerMovement, gameentity.Player)

	distanceFromEnemy := (e.field.Width - e.startEnemyNumnebr) / (e.startEnemyNumnebr + 1)
	xPositionEnemy := distanceFromEnemy

	for i := 0; i < e.startEnemyNumnebr; i++ {
		enemy := gameentity.Entity{}
		enemy.Init(xPositionEnemy, 1, e.frameEnemyMovement, gameentity.Enemy)
		e.Enemies = append(e.Enemies, enemy)
		xPositionEnemy += distanceFromEnemy
	}

}

func (
	e *Engine,
) Run() {
	e.renderField(false)
	exit := false

	for !exit {
		if time.Now().After(e.nextFrameTime) {
			e.nextFrameTime = e.nextFrameTime.Add(e.frameTime)

			for i := len(e.Enemies) - 1; i >= 0; i-- {
				if e.Enemies[i].IsDead() {
					e.Enemies = e.Enemies[:len(e.Enemies)-1]
				}
			}

			for i := len(e.Enemies) - 1; i >= 0; i-- {
				if e.Enemies[i].MoveForeward(1, e.field.Height) {
					e.Enemies[i].Kill()
					e.Player.Kill()
				}
			}

			e.checkPlayersKill()

			//readUserInput
			//movePlayer
			//che3ckPlayersKill

			exit = e.Player.IsDead()
			e.renderField(exit)
		}
	}
}

func (e *Engine) checkPlayersKill() {
	xPlayer, yPlayer := e.Player.GetPosition()

	for i, enemy := range e.Enemies {
		xEnemy, yEnemy := enemy.GetPosition()

		if xEnemy == xPlayer && yEnemy == yPlayer {
			e.Enemies[i].Kill()
		}
	}
}

func (e *Engine) renderField(exit bool) {
	var arrayEnemyReference []*gameentity.Entity

	for _, v := range e.Enemies {
		arrayEnemyReference = append(arrayEnemyReference, &v)
	}

	e.field.UpdateField(append(arrayEnemyReference, &e.Player))
	e.field.GeneratePrintableField(exit)
	fmt.Printf(e.field.PrintableField)
}
