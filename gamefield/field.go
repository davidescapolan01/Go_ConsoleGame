package gamefield

import (
	"asciigame/goGame/gameentity"
)

type Field struct {
	Width          int
	Height         int
	PrintableField string
	FieldArray     [][]*gameentity.Entity
}

func Init(width, height int) *Field {
	if width <= 0 {
		width = 31
	}

	if height <= 0 {
		height = 31
	}

	f := Field{Width: width, Height: height}
	var fA [][]*gameentity.Entity
	for i := 0; i < height; i++ {
		fA = append(fA, make([]*gameentity.Entity, width))
	}

	f.FieldArray = fA

	return &f
}

func (f *Field) UpdateField(
	e []*gameentity.Entity,
) {
	for i := 0; i < f.Height; i++ {
		for j := 0; j < f.Width; j++ {
			f.FieldArray[i][j] = nil
		}
	}

	for _, v := range e {
		x, y := v.GetPosition()
		f.FieldArray[y][x] = v
	}

	//for i := 0; i < len(e); i++ {
	//	x, y := e[i].GetPosition()
	//	f.FieldArray[y][x] = e[i]
	//}
}

func (f *Field) gameOver() string {
	var gameOverLine string

	gameOverString := "G A M E   O V E R"
	for i := 0; i < (f.Width-len(gameOverString))/2; i++ {
		gameOverLine += " "
	}

	gameOverLine += gameOverString

	for i := 0; i < (f.Width-len(gameOverString))/2; i++ {

		gameOverLine += " "
	}

	return gameOverLine
}

func (f *Field) GeneratePrintableField(gameOver bool) {
	f.PrintableField = ""

	for i := 0; i < len(f.FieldArray); i++ {
		f.PrintableField += "*"

		if i == 0 {
			for j := 0; j < len(f.FieldArray[i]); j++ {
				f.PrintableField += "*"
			}

			f.PrintableField += "*"
			f.PrintableField += "\n"
			f.PrintableField += "*"
		}

		if gameOver && i == len(f.FieldArray)/2 {
			f.PrintableField += f.gameOver()
		} else {
			for j := 0; j < len(f.FieldArray[i]); j++ {
				if f.FieldArray[i][j] == nil {
					f.PrintableField += " "
				} else {
					f.PrintableField += f.FieldArray[i][j].Icon
				}
			}

		}

		if i == len(f.FieldArray)-1 {
			f.PrintableField += "*"
			f.PrintableField += "\n"
			f.PrintableField += "*"

			for j := 0; j < len(f.FieldArray[i]); j++ {
				f.PrintableField += "*"
			}

		}

		f.PrintableField += "*"
		f.PrintableField += "\n"
	}
}
