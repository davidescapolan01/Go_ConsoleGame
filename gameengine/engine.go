package gameengine

import (
	"fmt"
	"strings"
	"time"

	"asciigame/goGame/gameentity"
	"asciigame/goGame/gamefield"
	"asciigame/goGame/utils"
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

	go utils.ReadUserInput(10 * time.Millisecond)

	for !exit {
		if time.Now().After(e.nextFrameTime) {
			e.nextFrameTime = e.nextFrameTime.Add(e.frameTime)

			//checkKilledEnemy
			for i := len(e.Enemies) - 1; i >= 0; i-- {
				if e.Enemies[i].IsDead() {
					e.Enemies = append(e.Enemies[:i], e.Enemies[i+1:]...)
				}
			}

			//checkEnemyWin
			for i := len(e.Enemies) - 1; i >= 0; i-- {
				if e.Enemies[i].MoveForeward(1, e.field.Height) {
					e.Enemies[i].Kill()
					e.Player.Kill()
				}
			}

			e.checkPlayersKill()

			//readUserInput
			inputRead := false
			var inputs []string
			for !inputRead {
				if gameentity.InputMutex.TryLock() {
					inputRead = true

					if len(gameentity.Inputs) > 0 {
						inputs = gameentity.Inputs
						gameentity.Inputs = make([]string, 0)
					}

					gameentity.InputMutex.Unlock()
				}
				time.Sleep(1 * time.Millisecond)
			}

			//parseInput
			var movement [4]int
			for _, input := range inputs {
				switch strings.ToLower(input) {
				case "h": //left
					movement[0] += 1
				case "j": //down
					movement[1] += 1
				case "k": //up
					movement[2] += 1
				case "l": //right
					movement[3] += 1
				}
			}

			//checkMostUsedMovement
			mostUsed := 0
			anyMovement := false
			for i := 0; i < len(movement); i++ {
				anyMovement = anyMovement || movement[i] > 0
				if mostUsed != i && movement[i] > movement[mostUsed] {
					mostUsed = i
				}
			}

			//movePlayer
			if anyMovement {
				switch mostUsed {
				case 0: //left
					e.Player.MoveLeft(-1, e.field.Width)
				case 1:
					e.Player.MoveForeward(1, e.field.Height)
				case 2:
					e.Player.MoveForeward(-1, e.field.Height)
				case 3:
					e.Player.MoveLeft(1, e.field.Width)
				}
			}

			e.checkPlayersKill()

			exit = e.Player.IsDead()
			e.renderField(exit)
		}
	}
}

func (
	e *Engine,
) checkPlayersKill() {
	xPlayer, yPlayer := e.Player.GetPosition()

	for i, enemy := range e.Enemies {
		xEnemy, yEnemy := enemy.GetPosition()

		if xEnemy == xPlayer && yEnemy == yPlayer {
			e.Enemies[i].Kill()
		}
	}
}

func (
	e *Engine,
) renderField(
	exit bool,
) {
	var arrayEnemyReference []*gameentity.Entity

	for _, v := range e.Enemies {
		arrayEnemyReference = append(arrayEnemyReference, &v)
	}

	e.field.UpdateField(append(arrayEnemyReference, &e.Player))
	e.field.GeneratePrintableField(exit)
	gameentity.Kill = exit
	utils.CallClear()
	fmt.Print(e.field.PrintableField)
}
