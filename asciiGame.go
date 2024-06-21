package main

import (
	"time"

	"asciigame/goGame/gameengine"
	"asciigame/goGame/gamefield"
)

func main() {
	var engine gameengine.Engine

	f := gamefield.Init(31, 31)

	frameTime := time.Duration(300 * time.Millisecond)

	engine.InitEngine(frameTime, 2, 0, 3, f)

	engine.Run()
}
