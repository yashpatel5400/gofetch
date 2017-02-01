package main

import (
	"fmt"
	"os"
    "os/exec"
	"math/rand"
	"time"
)

func main() {
	// sets up variables for the game run -- constants, background, player
	MAKE_RANGE := 25
	MAKE_ENEMY := 0
	MAKE_CLOUD := 1

	bg   := initBackground()
	play := initPlayer()
	rand.Seed(time.Now().UnixNano())

	// continuously reads from stdin asynchronously
    go func() {
    	// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		for {
			var b []byte = make([]byte, 1)
            os.Stdin.Read(b)
		    if !play.hasJumped && b[0] == 65 {	
            	play.hasJumped = true
				play.jumpPoint = 20
            }
        }
    }()

    // Infinite game loop -- updates the screen and reads user input
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
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			return
		}

		play = jumpPlayer(play)
		bg = drawPlayer(bg, play)
		render(bg)
		time.Sleep(100 * time.Millisecond)
    }
}