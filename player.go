package main

import (
)

var PLAYER_HEIGHT = 3
var PLAYER_WIDTH  = 3

type Player struct {
	// where the player is located
	position *BoundingBox

	// indicates whether the player jumped or not -- prevents double-jump
	hasJumped bool

	// used to indicate what portion of the jump is being executed
	jumpPoint int
}

/*****************************************************************************/
/* Creates new player and returns a pointer to it                            */
/*****************************************************************************/
func initPlayer() *Player {
	return &Player{
		position: &BoundingBox{bottom:GROUND_LEVEL.top, 
			top:GROUND_LEVEL.top + PLAYER_HEIGHT, 
			left:SCORE_BOX.left, 
			right:SCORE_BOX.left + PLAYER_WIDTH},
		hasJumped: false,
		jumpPoint: 0}
}

/*****************************************************************************/
/* Given the player and background on which being played, draws according to */
/* the player's current position.                                            */
/*****************************************************************************/
func drawPlayer(background *Background, player *Player) *Background {
	return insertOnBoard(background, player.position, "player")
}

/*****************************************************************************/
/* Given a player, checks whether it is in "jump status" and changes position*/
/* accordingly if so. Returns the new player w/ updated position             */
/*****************************************************************************/
func jumpPlayer(player *Player) *Player {
	DOWNWARD_PHASE := []int{1, 10}
	UPWARD_PHASE   := []int{11, 20}
	
	curJump := player.jumpPoint
	if curJump > 0 {
		if curJump >= DOWNWARD_PHASE[0] && curJump < DOWNWARD_PHASE[1] {
			player.position.bottom -= 1
		} else if curJump >= UPWARD_PHASE[0] && curJump < UPWARD_PHASE[1] {
			player.position.bottom += 1
		}
		player.position.top = player.position.bottom + PLAYER_HEIGHT
		player.jumpPoint -= 1
	} else {
		player.hasJumped = false
	}
	return player
}

/*****************************************************************************/
/* Given background and player, checks whether player has died (w/ enemy)    */
/*****************************************************************************/
func checkDead(background *Background, player *Player) bool {
	for y := player.position.bottom; y < player.position.top; y++ {
		for x := player.position.left; x < player.position.right; x++ {
			if background.board[y][x] == ENEMY {
				return true
			}	
		}	
	}
	return false
}