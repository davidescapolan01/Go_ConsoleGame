package gameentity

import (
	"sync"
)

var InputMutex sync.Mutex
var Inputs []string
var Kill bool
