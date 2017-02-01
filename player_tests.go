package main

import(
	"fmt"	
)

/*****************************************************************************/
/* Draws the user at a fixed x coordinate on the screen (w/ respect to size) */
/* A successful output shows a red block for the user on screen              */
/*****************************************************************************/
func test_draw_player() {
	bg   := initBackground()
	play := initPlayer()
	bg = drawPlayer(bg, play)
	render(bg)
}

/*****************************************************************************/
/* Should move the player when the up/down arrows are pressed and no others. */
/* A successful output will show player jumping up or crouching down resp.   */
/*****************************************************************************/
func test_move_player() {
	bg   := initBackground()
	play := initPlayer()
	play.jumpPoint = 30

	TEST_STEPS := 30
	for i := 0; i < TEST_STEPS; i++ {
		insertOnBoard(bg, play.position, "sky") 
		jumpPlayer(play)
		bg = drawPlayer(bg, play)
		render(bg)
	}
}

/*****************************************************************************/
/* Checks to make sure that, when make contact with an enemy, the player's   */
/* run ends. Successful output displays "Game Over" when user=enemy position */
/*****************************************************************************/
func test_kill_player() {
	bg   := initBackground()
	bg = insertEnemy(bg)
	bg = insertEnemy(bg)
	bg = insertEnemy(bg)
	play := initPlayer()
		
	TEST_STEPS := 125
	for i := 0; i < TEST_STEPS; i++ {
		bg = moveBackground(bg)
		if checkDead(bg, play) {
			fmt.Println("Successfully died!")
			return
		}
		bg = drawPlayer(bg, play)
		render(bg)
	}
}

/*****************************************************************************/
/* Runs all the tests written above -- used only in production compiles      */
/*****************************************************************************/
func main() {
	// test_draw_player()
	// test_move_player()
	test_kill_player()
}