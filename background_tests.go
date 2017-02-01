package main

import (
	"fmt"
)

/*****************************************************************************/
/* Should produce a plain background, i.e. with nothing populating the scene */
/* A successful output is a green/blue scene for the grass and sky           */
/*****************************************************************************/
func test_plain_background() {
	fmt.Println("Testing the plain background...")
	bg := initBackground()
	render(bg)
}

/*****************************************************************************/
/* Should produce a populated background, i.e. with random clouds and trees. */
/* A successful output will have these and them changing each run            */
/*****************************************************************************/
func test_populated_background() {
	fmt.Println("Testing the clouded background...")
	bg := initBackground()
	bg =  insertCloud(bg)
	render(bg)
	bg =  insertCloud(bg)
	render(bg)
}

/*****************************************************************************/
/* Make it so that the background moves (with a potential variable speed).   */
/* A successful output will make the user feel as though he/she is moving    */
/*****************************************************************************/
func test_move_background() {
	fmt.Println("Testing the moving background...")
	bg := initBackground()
	bg =  insertCloud(bg)
	
	TEST_STEPS := 10
	for i := 0; i < TEST_STEPS; i++ {
		bg =  moveBackground(bg)
		render(bg)	
	}
}

/*****************************************************************************/
/* Display a score that goes up in correspondence to the current speed of    */
/* run. A successful output will display this score in the top-right (fixed) */
/*****************************************************************************/
func test_display_score() {
	return
}

/*****************************************************************************/
/* Display the same scene when the terminal has been resized. A successful   */
/* output produces the same ratio/feel after a resizing is completed         */
/*****************************************************************************/
func test_resize() {
	return
}

/*****************************************************************************/
/* Runs all the tests written above -- used only in production compiles      */
/*****************************************************************************/
func main() {
	test_plain_background()
	test_populated_background()
	test_move_background()
	// test_display_score()
	// test_resize()
}