package main

import (
	"fmt"
	"log"
	gc "github.com/rthornton128/goncurses"
)

type Keyboard struct {
	upPushed   bool
	downPushed bool
	scr        *gc.Window
}

func initKeys() *Keyboard {
	gcScr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
		return nil
	}

	return &Keyboard{
		upPushed:   false,
		downPushed: false,
		scr: gcScr}
}

func closeKeys(keys *Keyboard) {
	gc.End()
}

func readKeys(board *Keyboard) *Keyboard {
	key := board.scr.GetChar()
	fmt.Println(key)
	return board
}