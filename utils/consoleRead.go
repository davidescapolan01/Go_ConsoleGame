package utils

import (
	"asciigame/goGame/gameentity"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

func ReadUserInput(timeout time.Duration) {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for !gameentity.Kill {
		if char, _, err := keyboard.GetKey(); err == nil {
			if gameentity.InputMutex.TryLock() {
				gameentity.Inputs = append(gameentity.Inputs, string(char))
				gameentity.InputMutex.Unlock()
				time.Sleep(timeout - 1*time.Millisecond)
			}
		}

		time.Sleep(1 * time.Millisecond)
	}
}
