package main

import (
	"fmt"
	// "log"
	"math/rand"
	"time"
	// gc "github.com/rthornton128/goncurses"
)

func main() {
	MAKE_RANGE := 100
	MAKE_ENEMY := 0
	MAKE_CLOUD := 1

	/* scr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End() */

	bg   := initBackground()
	play := initPlayer()
		
	rand.Seed(time.Now().UnixNano())
	for {
		seedDecision := rand.Intn(MAKE_RANGE)
		if seedDecision == MAKE_ENEMY {
			bg = insertEnemy(bg)	
		} else if seedDecision == MAKE_CLOUD {
			bg = insertCloud(bg)	
		} 

		bg = moveBackground(bg)
		if checkDead(bg, play) {
			fmt.Println("Game over!! You scored: ", bg.score)
			return
		}

		/* if play.hasJumped && gc.GetChar() == 65 {
			play.hasJumped = true
			play.jumpPoint   = 10
		} */

		play = jumpPlayer(play)
		bg = drawPlayer(bg, play)
		render(bg)
	}
}