package main

import (
	"fmt"
)

type Player struct {
	// where the player is located
	position *BoundingBox

	// indicates whether the player jumped or not -- prevents double-jump
	hasJumped bool
}

func initPlayer() {
	PLAYER_HEIGHT := 3
	PLAYER_WIDTH  := 3

	return &Player{
		position: &BoundingBox{bottom:GROUND_LEVEL.top, 
			top:GROUND_LEVEL.top + PLAYER_HEIGHT, 
			left:SCORE_BOX.left, 
			right:SCORE_BOX.left + PLAYER_WIDTH},
		hasJumped: false
	}
}

func movePlayer() {

}

func checkDead() {

}