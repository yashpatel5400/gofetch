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
	MAKE_RANGE := 100
	MAKE_ENEMY := 0
	MAKE_CLOUD := 1

	bg   := initBackground()
	play := initPlayer()
	rand.Seed(time.Now().UnixNano())

	// continuously reads from stdin asynchronously
    ch := make(chan []byte)
    go func(ch chan []byte) {
    	// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		// restore the echoing state when exiting
		defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		for {
			var b []byte = make([]byte, 1)
            os.Stdin.Read(b)
		    ch <- b
        }
        close(ch)
    }(ch)

    // Infinite game loop -- updates the screen and reads user input
stdinloop:
    for {
        select {
        // Has input: change user state to reflect the jump
        case stdin, ok := <-ch:
            if !ok {
                break stdinloop
            } else {
            	newChar := stdin[0]
            	if !play.hasJumped && newChar == 65 {	
            		fmt.Println("Jumped!")
            		play.hasJumped = true
					play.jumpPoint = 10
            	}
            }

        // No input: just render and update the scene
        case <-time.After(250 * time.Millisecond):
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

			play = jumpPlayer(play)
			bg = drawPlayer(bg, play)
			render(bg)
        }
    }
}